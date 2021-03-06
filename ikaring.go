/*
Package ikaring provides http client Api for SplatNet; web service for Splatoon by Nintendo.
*/
package ikaring

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
)

// IkaClient is a http client for SplatNet.
// it includes http.Client.
type IkaClient struct {
	hc      *http.Client
	BaseURL *url.URL    // Splatnet Domain URL
	AuthURL *url.URL    // Nintendo Network OAUTH URL
	Logger  *log.Logger // Logger
}

const (
	splatoonCookieName = "_wag_session"

	splatoonDomainURL = "https://splatoon.nintendo.net/"
	nintendoOauthURL  = "https://id.nintendo.net/oauth/authorize"

	loginPageURL  = "users/auth/nintendo"
	scheduleAPI   = "schedule.json"
	rankingAPI    = "ranking.json"
	friendListAPI = "friend_list/index.json"
)

// CreateClient generates ikaClient, http client object for Splatnet.
// It provides a http client with empty cookiejar.
func CreateClient() (*IkaClient, error) {
	return newClient(splatoonDomainURL, nintendoOauthURL, nil)
}

// newCleint generates ikaClient, http client object for Splatnet.
// this is inner implement for CreateClient() and used for tests.
func newClient(baseURL, authURL string, logger *log.Logger) (*IkaClient, error) {
	// cookie
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cookie jar")
	}
	// http client
	hc := &http.Client{Jar: jar}
	client := &IkaClient{
		hc: hc,
	}
	// base URL
	client.BaseURL, err = url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", baseURL)
	}
	// auth URL
	client.AuthURL, err = url.ParseRequestURI(authURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", authURL)
	}
	// logger
	client.Logger = logger
	if logger == nil {
		var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
		client.Logger = discardLogger
	} else {
		client.Logger = logger
	}
	return client, nil
}

// SetSession sets session cookie to receiver IkaClient.
func (c *IkaClient) SetSession(session string) {
	c.hc.Jar.SetCookies(c.BaseURL, []*http.Cookie{
		&http.Cookie{
			Secure:   true,
			HttpOnly: true,
			Name:     splatoonCookieName,
			Value:    session,
		}})
}

// Login sends http request to authorize Nintendo Network.
// it require NNID and password and return session cookie.
func (c *IkaClient) Login(ctx context.Context, name string, password string) (string, error) {
	req, err := c.newRequest(ctx, "GET", loginPageURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	query, err := getOauthQuery(resp, name, password)
	if err != nil {
		return "", err
	}

	req, err = c.newAuthRequest(ctx, strings.NewReader(query.Encode()))
	if err != nil {
		return "", err
	}
	resp, err = c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	session := getSessionFromCookie(splatoonCookieName, resp.Cookies())

	return session, nil
}

// Authorized judges wheather the client authorized
// It checks cookies for session that used for authorization
func (c *IkaClient) Authorized() bool {
	uri := c.BaseURL
	session := getSessionFromCookie(splatoonCookieName, c.hc.Jar.Cookies(uri))
	return len(session) != 0
}

// GetStageInfo get Stage Info from SplatNet.
// this API send GET request and parse stage schedules from JSON.
func (c *IkaClient) GetStageInfo(ctx context.Context) (*StageInfo, error) {
	req, err := c.newRequest(ctx, "GET", scheduleAPI, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)

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

// GetRanking get Ranking of Friends from SplatNet.
// this API send GET request and parse ranking from JSON.
func (c *IkaClient) GetRanking(ctx context.Context) (*RankingInfo, error) {
	req, err := c.newRequest(ctx, "GET", rankingAPI, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)

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

	return decodeJSONRanking(body)
}

// GetFriendList get Friend List form SplatNet.
// this API send GET request and parse friend online status from JSON
func (c *IkaClient) GetFriendList(ctx context.Context) ([]Friend, error) {
	req, err := c.newRequest(ctx, "GET", friendListAPI, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.hc.Do(req)

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
	return decodeJSONFriendList(body)
}

// GetWeaponMap get Weapon Set from SplatNet.
// this API send GET request and parse weapon map by scraping HTML
func (c *IkaClient) GetWeaponMap(ctx context.Context) (map[string]string, error) {
	req, err := c.newRequest(ctx, "GET", "", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	weapons := map[string]string{}
	doc.Find("#user_intention_weapon").Children().Each(func(_ int, s *goquery.Selection) {
		key, ok := s.Attr("value")
		if ok {
			weapons[key] = s.Text()
		}
	})
	return weapons, nil
}

func (c *IkaClient) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	if len(spath) > 0 {
		u.Path = path.Join(c.BaseURL.Path, spath)
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return req, nil
}

func (c *IkaClient) newAuthRequest(ctx context.Context, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", c.AuthURL.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return req, nil
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
