package gossti

type node struct {
    Language
}

func newNode() ILanguage {
		return &node{
				Language: Language{
						name: "NodeJS",
						confidence: 0.0,
						Engines: []IEngine{newPug(), newJade(), newHandlebars(), newJSRender(), newNunjucks()},
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

type Jade struct {
	Engine
}

func newJade() IEngine {
	return &Jade{ 
		Engine: Engine{
			name: "jade",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "jade_1", Value: "#{21389*90}", Expected: "1925010"},
			},
		},
	}
}

type Handlebars struct {
	Engine
}

func newHandlebars() IEngine {
	return &Handlebars{ 
		Engine: Engine{
			name: "handlebars",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "handlebars_1", Value: "${7*7}", Expected: "${7*7}"},
			},
		},
	}
}

type JSRender struct {
	Engine
}

func newJSRender() IEngine {
	return &JSRender{ 
		Engine: Engine{
			name: "jsrender",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "jsrender_1", Value: "{{:21389*90}}", Expected: "1925010"},
			},
		},
	}
}

type Nunjucks struct {
	Engine
}

func newNunjucks() IEngine {
	return &Nunjucks{ 
		Engine: Engine{
			name: "nunjucks",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "nunjucks_1", Value: "{{21389*90}}", Expected: "1925010"},
				{ Name: "nunjucks_2", Value: " #{7*7}", Expected: " #{7*7}"},
			},
		},
	}
}
