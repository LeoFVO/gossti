package gossti

type python struct {
    Language
}

func newPython() ILanguage {
		return &python{
				Language: Language{
						name: "python",
						confidence: 0.0,
						Engines: []IEngine{newMako(), newJinja2(), newTornado()},
				},
		}
}

type mako struct {
	Engine
}

func newMako() IEngine {
	return &mako{ 
		Engine: Engine{
			name: "Mako",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Mako_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Mako_2", Value: "${21389*90}", Expected: "1925010"},
			},
		},
	}
}

type jinja2 struct {
	Engine
}

func newJinja2() IEngine {
	return &jinja2{
		Engine: Engine{
			name: "Jinja2",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Jinja2_1", Value: "{{20*'7'}}", Expected: "77777777777777777777"},
				{ Name: "Jinja2_2", Value: "${21389*90}", Expected: "${21389*90}"},
			},
		},
	}
}

type tornado struct {
	Engine
}

func newTornado() IEngine {
	return &tornado{
		Engine: Engine{
			name: "Tornado",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Tornado_1", Value: "{{20*'7'}}", Expected: "77777777777777777777"},
				{ Name: "Tornado_2", Value: "{{21389*90}}", Expected: "1925010"},
				{ Name: "Tornado_3", Value: "${21389*90}", Expected: "${21389*90}"},
			},
		},
	}
}