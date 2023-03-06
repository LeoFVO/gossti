package gossti

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// IEngine is the interface for all engines that will be used to detect the language
// Engine are defined as libraries/framework that are used to render templates in a language
type IEngine interface {
	Scan(req *http.Request, queryString QueryString) bool

	GetName() string
	GetConfidence() float64
	IncrementConfidence()
}

// Engine is the struct that will be used to define an engine
type Engine struct {
	name string
	confidence	float64
	payloads []Payload
}

// Increment the confidence of the engine identification
func (e *Engine) IncrementConfidence() {
	log.Tracef("Incrementing confidence for engine %s", e.name)
	e.confidence++
}

// Returns the confidence of the engine identification
func (e *Engine) GetConfidence() float64 {
	log.Tracef("Getting confidence for engine %s", e.name)
	return e.confidence / float64(len(e.payloads)) * 100
}

// Returns the name of the engine
func (e *Engine) GetName() string {
	return e.name
}

// Try each well known payload to see if the engine is vulnerable 
// and deduce the confidence
// Returns true if the confidence is above 50%
func (e *Engine) Scan(req *http.Request, queryString QueryString) bool {
	log.Infof("-> Trying to detect %s engine.\n", e.name)

	for _, payload := range e.payloads {
		log.WithFields(log.Fields{
			"injected": payload.Value,
			"expected": payload.Expected,
		}).Debugf("   Trying payload %s\n", payload.Name)
		if queryString.Check(req, payload) {
			log.Warnf("   Payload %s is injectable\n", payload.Name)
			e.IncrementConfidence()
		}
	}

	log.Infof("   Confidence for %s is %.2f %%\n\n", e.name, e.GetConfidence())
	return e.GetConfidence() >= 50.00
}