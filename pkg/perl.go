package gossti

type perl struct {
    Language
}

func newPerl() ILanguage {
		return &perl{
				Language: Language{
						name: "Perl",
						confidence: 0.0,
						Engines: []IEngine{newMojolicious()},
				},
		}
}

type Mojolicious struct {
	Engine
}

func newMojolicious() IEngine {
	return &Mojolicious{ 
		Engine: Engine{
			name: "Mojolicious",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Mojolicious_1", Value: "<%= 21389+219839 %>", Expected: "241228"},
				{ Name: "Mojolicious_2", Value: "{{21389*90}}", Expected: "{{21389*90}}"},
				{ Name: "Mojolicious_3", Value: "${21389*90}", Expected: "${21389*90}"},
			},
		},
	}
}
