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
				{ Name: "GoExpressionLanguage_1", Value: "{{printf \"%%s\" \"rqXEcdfab9H9NyaBkFJRkMqbGBAQS5sYY6kD53d6\"}}", Expected: "rqXEcdfab9H9NyaBkFJRkMqbGBAQS5sYY6kD53d6"},
				{ Name: "GoExpressionLanguage_2", Value: "{{html \"Ax8dXg86yiRQyr4EN4ndNxmkEFQsotEs6zM3iN4C\"}}", Expected: "Ax8dXg86yiRQyr4EN4ndNxmkEFQsotEs6zM3iN4C"},
				{ Name: "GoExpressionLanguage_3", Value: "{{js \"jFL4GnD8hmhXeEqioYnDQLgmSD6QQmnKFzdfXQK7\"}}", Expected: "jFL4GnD8hmhXeEqioYnDQLgmSD6QQmnKFzdfXQK7"},
			},
		},
	}
}