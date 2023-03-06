package gossti

import (
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type QueryString struct {
	Name string
	Value string
}

type Payload struct {
	Name string
	Value string
	Expected string
}

func (s *QueryString) Check(req *http.Request, payload Payload) bool {
	log.Tracef("   Checking if the payload %s is vulnerable", payload.Name)
	client := &http.Client{}

	q := req.URL.Query()
	q.Del(s.Name) // Remove the query string from the request to avoid false positives
	q.Add(s.Name, strings.Replace(s.Value, "SSTI", payload.Value, -1))
	req.URL.RawQuery = q.Encode()

	log.Debugf("   Sending the request to %s", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Tracef("   Response status code: %d", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Tracef("   Crawling the response body")
	return strings.Contains(string(body), payload.Expected)
}