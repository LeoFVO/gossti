package gossti

import (
	"fmt"
	"net/http"
)

// IEngine is the interface for all engines that will be used to detect the language
// Engine are defined as libraries/framework that are used to render templates in a language
type IEngine interface {
	Scan(req *http.Request, queryString QueryString) bool

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
	e.confidence++
}

// Returns the confidence of the engine identification
func (e *Engine) GetConfidence() float64 {
	return e.confidence / float64(len(e.payloads)) * 100
}

// Try each well known payload to see if the engine is vulnerable 
// and deduce the confidence
// Returns true if the confidence is above 50%
func (e *Engine) Scan(req *http.Request, queryString QueryString) bool {
	fmt.Printf("-> Trying to detect %s engine.\n", e.name)

	for _, payload := range e.payloads {
		if queryString.Check(req, payload) {
			fmt.Printf("Payload %s is injectable\n", payload.Name)
			e.IncrementConfidence()
		}
	}

	fmt.Printf("   Confidence for %s is %.2f %%\n\n", e.name, e.GetConfidence())
	return e.GetConfidence() >= 50.00
}