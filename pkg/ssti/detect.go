package ssti

import (
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
)

type Options struct {
	Method string
	UserAgent string
	Cookies []string
	Timeout time.Duration
	Form map[string]string
	FormType string
}

func Detect(url string, options Options) error {
	log.Debug("Starting SSTI detection")

	cli := configureHttpClient(options)
	log.Debugf("Sending %s request to %s", options.Method, url)

	languages := getAvailableLanguages()
	for _, language := range languages {		
		l, err := getLanguage(language)
		if err != nil {
			log.Errorf("Error loading language plugin: %v", err)
			continue
		}

		l.Detect(url, cli, options)
	}

	return nil
}

func configureHttpClient(options Options) *gentleman.Client {
	log.Debugf("Configuring http client with options: %+v", options)
	cli := gentleman.New()

	log.Tracef("Set user agent to %s", options.UserAgent)
	cli.SetHeader("User-Agent", options.UserAgent)
	cli.SetHeader("Accept", "*/*")

	if options.Cookies != nil && len(options.Cookies) > 0 {
		cookies := []*http.Cookie{}
		for _, cookie := range options.Cookies {
			cookies = append(cookies, &http.Cookie{Name: strings.Split(cookie, "=")[0], Value: strings.Split(cookie, "=")[1]})
		}
		log.Tracef("Set cookies to: %v", strings.Join(options.Cookies, "; "))
		cli.AddCookies(cookies)
	}

	if options.Timeout != 0 {
		log.Tracef("Set timeout to %+v", options.Timeout)
  	cli.Use(timeout.Request(options.Timeout))
	}

	return cli
}
