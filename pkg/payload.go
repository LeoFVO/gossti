package gossti

import (
	"io"
	"net/http"
	"strings"
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
	client := &http.Client{}

	q := req.URL.Query()
	q.Add(s.Name, strings.Replace(s.Value, "SSTI", payload.Value, -1))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return strings.Contains(string(body), payload.Expected)
}