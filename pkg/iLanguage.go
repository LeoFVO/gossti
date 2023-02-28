package gossti

import (
	"fmt"
	"net/http"
)

// ILanguage is the interface for all languages that will be used to detect the language
type ILanguage interface {
	Detect(req *http.Request, queryString QueryString)
	
	GetConfidence() float64
	IncrementConfidence()

	GetName() string
	GetEngine() []IEngine
}

// Language is the struct that will be used to define a programming language
type Language struct {
	confidence	float64
	name				string
	Engines     []IEngine
}

// Increment the confidence of the language identification
func (l *Language) IncrementConfidence() {
	l.confidence++
}

// Returns the confidence of the language identification
func (l *Language) GetConfidence() float64 {
	return l.confidence / float64(len(l.Engines)) * 100
}

// Returns the list of engines that will be used to detect the language
func (l *Language) GetEngine() []IEngine {
	return l.Engines
}

// Set the list of engines that will be used to detect the language
func (l *Language) SetEngine(engines []IEngine) {
	l.Engines = engines
}

// Returns the name of the language
func (l *Language) GetName() string {
	return l.name
}

// Set the name of the language
func (l *Language) Detect(req *http.Request, queryString QueryString) {
	fmt.Printf("Detecting language with %s\n", l.GetName())

	for _, engine := range l.GetEngine() {
		if engine.Scan(req, queryString) {
			l.IncrementConfidence()
		}
	}

	fmt.Printf("Confidence for %s is %.2f %%\n\n",l.GetName(), l.GetConfidence())

	if l.GetConfidence() == 100.00 {
		fmt.Printf("Language %s is detected\n\n", l.GetName())
	}
}