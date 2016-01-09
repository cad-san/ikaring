package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type ikaClient http.Client

const (
	splatoonClientID = "12af3d0a3a1f441eb900411bb50a835a"
	splatoonOauthURL = "https://splatoon.nintendo.net/users/auth/nintendo"
	nintendoOauthURL = "https://id.nintendo.net/oauth/authorize"
)

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

func main() {
}
