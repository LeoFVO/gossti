package gossti

type php struct {
    Language
}

func newPhp() ILanguage {
		return &php{
				Language: Language{
						name: "php",
						confidence: 0.0,
						Engines: []IEngine{
							newTwig(),
							newSmarty(),
						},
				},
		}
}

type twig struct {
	Engine
}

func newTwig() IEngine {
	return &twig{ 
		Engine: Engine{
			name: "Twig",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Twig_1", Value: "{{21389+219839}}", Expected: "241228"},
				{ Name: "Twig_2", Value: "{{21389*90}}", Expected: "1925010"},
				{ Name: "Twig_3", Value: "{{21389*'90'}}", Expected: "1925010"},
				{ Name: "Twig_4", Value: "${21389*90}", Expected: "${21389*90}"},
			},
		},
	}
}

type smarty struct {
	Engine
}

func newSmarty() IEngine {
	return &smarty{
		Engine: Engine{
			name: "Smarty",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Smarty_1", Value: "{$21389+219839}", Expected: "241228"},
				{ Name: "Smarty_2", Value: "{$21389*90}", Expected: "1925010"},
				{ Name: "Smarty_3", Value: "{$21389*'90'}", Expected: "1925010"},
				{ Name: "Smarty_4", Value: "${21389*90}", Expected: "${21389*90}"},
			},
		},
	}
}
