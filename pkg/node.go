package gossti

type node struct {
    Language
}

func newNode() ILanguage {
		return &node{
				Language: Language{
						name: "NodeJS",
						confidence: 0.0,
						Engines: []IEngine{newPug()},
				},
		}
}

type pug struct {
	Engine
}
func newPug() IEngine {
	return &pug{ 
		Engine: Engine{
			name: "pug",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "pug_1", Value: "#{21389*90}", Expected: "1925010"},
			},
		},
	}
}