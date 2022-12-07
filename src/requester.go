package src

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"time"
)

func (s SurrDB) BasecAuth() string {
	auth := s.User + ":" + s.Pass
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (s SurrDB) Requester(p string) (string, int) {
	var client *http.Client
	var req *http.Request
	var resp *http.Response
	var rbody []byte
	var err error

	body := strings.NewReader(p)
	req, err = http.NewRequest("POST", s.Schema+"://"+s.Host+"/sql", body)
	if err != nil {
		HandErrs(err)
	}
	req.Header.Add("Authorization", "Basic "+s.BasecAuth())
	req.Header.Add("NS", s.Namespace)
	req.Header.Add("DB", s.Database)
	req.Header.Add("Accept", "application/json")

	client = &http.Client{Timeout: s.Timeout * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		HandErrs(err)
	}

	rbody, err = io.ReadAll(resp.Body)
	if err != nil {
		HandErrs(err)
	}
	return string(rbody), resp.StatusCode
}
