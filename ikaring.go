package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	simplejson "github.com/bitly/go-simplejson"
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
	TimeBegin string     `json:"datetime_begin"`
	TimeEnd   string     `json:"datetime_end"`
	Stages    regulation `json:"stages"`
	GachiRule string     `json:"gachi_rule"`
}

type stageInfo struct {
	Festival  bool       `json:"festival"`
	Schedules []schedule `json:"schedule"`
}

const (
	splatoonClientID = "12af3d0a3a1f441eb900411bb50a835a"
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
		case "cliend_id":
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
	query.Add("client_id", splatoonClientID)
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

func (c *ikaClient) login(name string, password string) error {
	query, err := getOauthQuery(splatoonOauthURL, name, password)
	if err != nil {
		return err
	}

	resp, err := (*http.Client)(c).PostForm(nintendoOauthURL, query)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil
}

func (c *ikaClient) getStageInfo() (*stageInfo, error) {

	resp, err := (*http.Client)(c).Get(splatoonScheduleAPI)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return decodeJSONSchedule(resp.Body)
}

func main() {
	_, err := createClient()
	if err != nil {
		fmt.Println(err)
	}
}
