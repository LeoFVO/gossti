package gossti

type ruby struct {
    Language
}

func newRuby() ILanguage {
		return &ruby{
				Language: Language{
						name: "Ruby",
						confidence: 0.0,
						Engines: []IEngine{newERB(), newSlim()},
				},
		}
}

type ERB struct {
	Engine
}

func newERB() IEngine {
	return &ERB{ 
		Engine: Engine{
			name: "ERB",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "ERB_1", Value: "<%= 21389+219839 %>", Expected: "241228"},
				{ Name: "ERB_2", Value: "{{21389*90}}", Expected: "{{21389*90}}"},
				{ Name: "ERB_3", Value: "${21389*90}", Expected: "${21389*90}"},
			},
		},
	}
}

type Slim struct {
	Engine
}

func newSlim() IEngine {
	return &Slim{ 
		Engine: Engine{
			name: "Slim",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Slim_1", Value: "= 21389+219839", Expected: "241228"},
				{ Name: "Slim_2", Value: "= 21389*90", Expected: "1925010"},
				{ Name: "Slim_3", Value: "= 21389*90", Expected: "1925010"},
				{ Name: "Slim_4", Value: "{ 21389 * 90 }", Expected: "1925010"},
			},
		},
	}
}