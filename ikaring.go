package ikaring

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/bitly/go-simplejson"
)

const (
	splatoonCookieName = "_wag_session"
	splatoonDomainURL  = "https://splatoon.nintendo.net/"

	splatoonOauthURL = "https://splatoon.nintendo.net/users/auth/nintendo"
	nintendoOauthURL = "https://id.nintendo.net/oauth/authorize"

	splatoonScheduleAPI = "https://splatoon.nintendo.net/schedule.json"
)

func CreateClient() (*ikaClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &ikaClient{}
	client.Jar = jar
	return client, nil
}

func (c *ikaClient) SetSession(session string) {
	uri, _ := url.Parse(splatoonDomainURL)
	c.Jar.SetCookies(uri, []*http.Cookie{
		&http.Cookie{
			Secure:   true,
			HttpOnly: true,
			Name:     splatoonCookieName,
			Value:    session,
		}})
}

func (c *ikaClient) Login(name string, password string) (string, error) {
	query, err := getOauthQuery(splatoonOauthURL, name, password)
	if err != nil {
		return "", err
	}

	resp, err := c.PostForm(nintendoOauthURL, query)
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

func (c *ikaClient) GetStageInfo() (*stageInfo, error) {

	resp, err := c.Get(splatoonScheduleAPI)

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
