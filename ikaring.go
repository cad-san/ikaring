package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bgentry/speakeasy"
	"github.com/bitly/go-simplejson"
)

type ikaClient http.Client

type stage struct {
	Name  string `json:"name"`
	Image string `json:"asset_path"`
}

type regulation struct {
	Regular []stage
	Gachi   []stage
}

type schedule struct {
	TimeBegin time.Time  `json:"datetime_begin"`
	TimeEnd   time.Time  `json:"datetime_end"`
	Stages    regulation `json:"stages"`
	GachiRule string     `json:"gachi_rule"`
}

type stageInfo struct {
	Festival  bool       `json:"festival"`
	Schedules []schedule `json:"schedule"`
}

const (
	splatoonCookieName = "_wag_session"
	splatoonDomainURL  = "https://splatoon.nintendo.net/"

	splatoonOauthURL = "https://splatoon.nintendo.net/users/auth/nintendo"
	nintendoOauthURL = "https://id.nintendo.net/oauth/authorize"

	splatoonScheduleAPI = "https://splatoon.nintendo.net/schedule/index.json"
)

func checkJSONError(data []byte) error {
	js, err := simplejson.NewJson(data)
	if err != nil {
		return err
	}

	if info := js.Get("error").MustString(); len(info) != 0 {
		return errors.New(info)
	}
	return nil
}

func decodeJSONSchedule(data []byte) (*stageInfo, error) {
	info := &stageInfo{}
	if err := json.Unmarshal(data, info); err != nil {
		return nil, err
	}
	return info, nil
}

func getSessionFromCookie(resp *http.Response) string {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "_wag_session" {
			return cookie.Value
		}
	}
	return ""
}

func getOauthQuery(oarthURL string, id string, password string) (url.Values, error) {
	doc, err := goquery.NewDocument(oarthURL)

	if err != nil {
		return nil, err
	}

	query := url.Values{}
	doc.Find("input").Each(func(_ int, s *goquery.Selection) {
		name, ok := s.Attr("name")
		if !ok {
			return
		}

		// parse from docment
		switch name {
		case "client_id":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		case "state":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		case "redirect_uri":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		case "response_type":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		}
	})

	// fixed value
	query.Add("nintendo_authenticate", "")
	query.Add("nintendo_authorize", "")
	query.Add("scope", "")
	query.Add("lang", "ja-JP")

	// user info
	query.Add("username", id)
	query.Add("password", password)

	return query, nil
}

func createClient() (*ikaClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Jar: jar}
	return (*ikaClient)(client), nil
}

func (c *ikaClient) setSession(session string) {
	uri, _ := url.Parse(splatoonDomainURL)
	(*http.Client)(c).Jar.SetCookies(uri, []*http.Cookie{
		&http.Cookie{
			Secure:   true,
			HttpOnly: true,
			Name:     splatoonCookieName,
			Value:    session,
		}})
}

func (c *ikaClient) login(name string, password string) (string, error) {
	query, err := getOauthQuery(splatoonOauthURL, name, password)
	if err != nil {
		return "", err
	}

	resp, err := (*http.Client)(c).PostForm(nintendoOauthURL, query)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	session := getSessionFromCookie(resp)

	return session, nil
}

func (c *ikaClient) getStageInfo() (*stageInfo, error) {

	resp, err := (*http.Client)(c).Get(splatoonScheduleAPI)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = checkJSONError(body); err != nil {
		return nil, err
	}

	return decodeJSONSchedule(body)
}

func getCacheFile() (string, error) {
	me, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(me.HomeDir, ".ikaring.session"), nil
}

func readSession(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buff)), nil
}

func writeSession(path string, session string) error {
	return ioutil.WriteFile(path, []byte(session), 600)
}

func getAccount(r io.Reader) (string, string, error) {
	scanner := bufio.NewScanner(r)
	for {
		fmt.Print("User: ")
		if scanner.Scan() {
			break
		}
	}
	username := scanner.Text()
	password, err := speakeasy.Ask("Password: ")
	return username, password, err
}

func (s schedule) String() string {
	timefmt := "01/02 15:04:05"
	str := fmt.Sprintf("%s - %s\n",
		s.TimeBegin.Format(timefmt), s.TimeEnd.Format(timefmt))

	str += "レギュラーマッチ\n"
	for i, stage := range s.Stages.Regular {
		if i == 0 {
			str += "\t"
		} else {
			str += ", "
		}
		str += stage.Name
	}
	str += "\n"

	str += fmt.Sprintf("ガチマッチ (%s)\n", s.GachiRule)
	for i, stage := range s.Stages.Gachi {
		if i == 0 {
			str += "\t"
		} else {
			str += ", "
		}
		str += stage.Name
	}
	str += "\n"
	return str
}

func main() {
	client, err := createClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	path, err := getCacheFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := readSession(path)
	if err == nil && len(session) > 0 {
		client.setSession(session)
	} else {
		username, password, err := getAccount(os.Stdin)
		session, err = client.login(username, password)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if len(session) <= 0 {
		fmt.Println("ログインできませんでした")
		return
	}

	writeSession(path, session)

	info, err := client.getStageInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range info.Schedules {
		fmt.Printf("%v\n", s)
	}
}
