package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/franela/goreq"
	"github.com/garetht/fullstory-api/query"
	"github.com/garetht/fullstory-api/session"
	"github.com/garetht/fullstory-api/user"
)

type API struct {
	key string
}

type Config struct {
	Query   *query.FsQuery
	Options map[string]string
}

const (
	fsBase      = "https://www.fullstory.com"
	fsLogin     = "/loginrequest"
	fsCookieKey = "fs_session"
	fsCsrfKey   = "csrftoken"
)

func (a *API) Init(username string, password string) {
	params := map[string]interface{}{
		"Username": username,
		"Password": password,
	}

	resp := a.post(fsLogin, params, url.Values{})
	cookies := resp.Cookies()

	for _, c := range cookies {
		if c.Name == fsCookieKey {
			a.key = c.Value
			break
		}
	}
}

func (a *API) UserQuery(c Config) (users user.FsUsers) {
	qp := url.Values{}
	qp.Set("limit", "25")
	qp.Set("start", "0")
	qp.Set("tz", "America/New_York")

	if c.Options != nil {
		for k, v := range c.Options {
			qp.Set(k, v)
		}
	}

	resp := a.post(user.Endpoint, c.Query, qp)

	users = user.FsUsers{}
	resp.Body.FromJsonTo(&users)

	return
}

func (a *API) SessionQuery(c Config) (sessions session.FsSessions) {
	if c.Options["userId"] == "" {
		panic("A userId must be provided to make a SessionQuery.")
	}

	qp := url.Values{}
	qp.Set("limit", "100")
	qp.Set("tz", "America/New_York")

	for k, v := range c.Options {
		qp.Set(k, v)
	}

	resp := a.post(session.Endpoint, c.Query, qp)

	sessions = session.FsSessions{}
	resp.Body.FromJsonTo(&sessions)

	return
}

func (a *API) post(endpoint string, data interface{}, queryParams url.Values) (resp *goreq.Response) {
	url := strings.Join([]string{fsBase, endpoint}, "")
	req := goreq.Request{
		Uri:         url,
		Method:      "POST",
		Body:        data,
		Accept:      "application/json",
		ContentType: "application/json",
		ShowDebug:   true,
		QueryString: queryParams,
	}

	// If the key exists, send it.
	if a.key != "" {
		req = req.WithCookie(&http.Cookie{
			Name:  fsCookieKey,
			Value: a.key,
		}).WithHeader(fsCsrfKey, a.key)
	}

	resp, err := req.Do()

	if err != nil {
		fmt.Printf("%v : error encountered during Fullstory API client POST")
	}

	return
}
