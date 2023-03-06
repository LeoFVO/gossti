package gossti

import (
	"net/http"

	log "github.com/sirupsen/logrus"
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
	log.Tracef("Incrementing confidence for language %s", l.name)
	l.confidence++
}

// Returns the confidence of the language identification
func (l *Language) GetConfidence() float64 {
	log.Tracef("Getting confidence for language %s", l.name)
	return l.confidence / float64(len(l.Engines)) * 100
}

// Returns the list of engines that will be used to detect the language
func (l *Language) GetEngine() []IEngine {
	log.Tracef("Getting engines for language %s", l.name)
	return l.Engines
}

// Set the list of engines that will be used to detect the language
func (l *Language) SetEngine(engines []IEngine) {
	log.Tracef("Setting engines for language %s", l.name)
	l.Engines = engines
}

// Returns the name of the language
func (l *Language) GetName() string {
	return l.name
}

// Set the name of the language
func (l *Language) Detect(req *http.Request, queryString QueryString) {
	log.Infof("Starting detection for %s language.\n", l.GetName())

	for _, engine := range l.GetEngine() {
		log.Debug("   Trying engine ", engine.GetName())
		if engine.Scan(req, queryString) {
			l.IncrementConfidence()
		}
	}

	log.Printf("Confidence for %s is %.2f %%\n",l.GetName(), l.GetConfidence())

	if l.GetConfidence() == 100.00 {
		log.Warnf("Successfully detected %s language\n\n", l.GetName())
	}

	log.Printf("--------------------------------------------------\n\n")
}