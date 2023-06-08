package ssti

import (
	"fmt"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/multipart"
)

type Payload struct {
	ID string `yaml:"id"`
	Payload string `yaml:"payload"`
	Response struct {
		Expected string `yaml:"expected"`
		Invalidate string `yaml:"invalidate"`
		Error bool `yaml:"error"`
	} `yaml:"response"`
}

type Engine struct {
		Name string `yaml:"name"`
		Payloads []Payload `yaml:"payloads"`
}

type Language struct {
	Version string `yaml:"version"`
	Name string `yaml:"name"`
	Engines []Engine `yaml:"engines"`
}

func (l *Language) Detect(url string, cli *gentleman.Client, options Options) bool {
	log.Tracef("Detecting SSTI for language %s using plugin version %s", l.Name, l.Version)
	
	detected := false

	log.Infof("Starting detection for language %s", l.Name)
	for _, engine := range l.Engines {
		confidence, err := engine.Detect(url, cli, options)
		if err != nil {
			log.Errorf("Error using engine %s: %v", engine.Name, err)
			continue
		}

		log.Tracef("Confidence for engine %s: %f", engine.Name, confidence)

		if confidence == 1 {
			detected = true
			log.Warnf("Successfully detected SSTI using engine %s", engine.Name)
		} else if confidence > 0.65 {
			log.Warnf("Detected engine %s with confidence of %f", engine.Name, confidence)
		}
	}

	log.Infof("Finished detection for language %s", l.Name)

	return detected
}

func (l *Language) Exploit() {
	fmt.Println("Exploit not implemented yet")
}

func (e *Engine) Detect(url string, cli *gentleman.Client, options Options) (float32, error) {
	log.Tracef("Detecting SSTI for engine %s", e.Name)

	log.Debugf("Starting detection for engine %s", e.Name)
	var confidences int
	for _, payload := range e.Payloads {
		detected, err := payload.Detect(url, cli, options)
		if err != nil {
			log.Errorf("Error using payload %s: %v",payload.ID, err)
			continue
		}

		if detected {
			log.Infof("SSTI detected using payload %s", payload.ID)
			confidences += 1
		}
	}

	return float32(confidences/len(e.Payloads)), nil
}


func (p *Payload) Detect(u string, cli *gentleman.Client, options Options) (bool, error) {
	log.Debugf("Using payload %s", p.ID)

	u = strings.Replace(u, "SSTI", p.Payload, -1)

		request := cli.Request().Method(options.Method).URL(u)

	if options.Form != nil {
		if options.FormType == "multipart" {
			log.Tracef("Using form data: %v", options.Form)
			df := make(multipart.DataFields)
			for k, v := range options.Form {
				df[k] = multipart.Values{strings.Replace(v, "SSTI", p.Payload, -1)}
			}
			cli.Use(multipart.Fields(df))
		} else if options.FormType == "urlencoded" {
			data := url.Values{}
			for k, v := range options.Form {
				data.Set(k, strings.Replace(v, "SSTI", p.Payload, -1))
			}
			request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			log.Tracef("Sending data: %v", data.Encode())
			request.Body(strings.NewReader(data.Encode()))
		}
	}


	log.Tracef("Sending %s request to %s", options.Method, u)
	res, err := request.Send()
  if err != nil {
    log.Errorf("Request error: %s\n", err)
    return false, err
  }

	if p.Response.Error {
		if res.Error != nil {
			log.Infof("SSTI detected using payload %s", p.ID)
			return true, nil
		}
	}

  if !res.Ok {
    fmt.Printf("Invalid server response: %d\n", res.StatusCode)
    return false, nil
  }


	log.Debugf("Response: %s", res.String())
	if strings.Contains(res.String(), p.Response.Expected) && !strings.Contains(res.String(), p.Response.Invalidate) {
		log.Debugf("Response contains expected string %s", p.Response.Expected)
		return true, nil
	}

	return false,nil
}
