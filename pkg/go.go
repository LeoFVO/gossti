package gossti

type golang struct {
    Language
}

func newGo() ILanguage {
		return &golang{
				Language: Language{
						name: "GoLang",
						confidence: 0.0,
						Engines: []IEngine{
							newGoEL(),
						},
				},
		}
}

type golangEL struct {
	Engine
}

func newGoEL() IEngine {
	return &golangEL{
		Engine: Engine{
			name: "GoLang Expression Language",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "GoExpressionLanguage_1", Value: "{{printf \"%%s%%s\" \"rqXEcdfab9H9NyaBkFJRkMqbGBAQS5sYY6kD53d6\" \"Ax8dXg86yiRQyr4EN4ndNxmkEFQsotEs6zM3iN4C\"}}", Expected: "rqXEcdfab9H9NyaBkFJRkMqbGBAQS5sYY6kD53d6Ax8dXg86yiRQyr4EN4ndNxmkEFQsotEs6zM3iN4C"},
				{ Name: "GoExpressionLanguage_2", Value: "{{239892183 -}} < {{- 4521368712}}", Expected: "239892183>4521368712"},
			},
		},
	}
}