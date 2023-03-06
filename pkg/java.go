package gossti

type java struct {
    Language
}

func newJava() ILanguage {
		return &java{
				Language: Language{
						name: "Java",
						confidence: 0.0,
						Engines: []IEngine{
							newExpressionLanguage(),
							newFreemarker(),
							newGroovy(),
							newGeneric(),
							newJinjava(),
							newSpring(),
							newVelocity(),
							newThymeleaf(),
							newPebble(),
							newHubl(),
						},
				},
		}
}

type generic struct {
	Engine
}

func newGeneric() IEngine {
	return &generic{
		Engine: Engine{
			name: "Generic",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Generic_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Generic_2", Value: "${21389*90}", Expected: "1925010"},
				{ Name: "Generic_3", Value: "${{21389*90}}", Expected: "1925010"},
				{ Name: "Generic_4", Value: "${{21389*90}}", Expected: "1925010"},
			},
		},
	}
}

type freemarker struct {
	Engine
}

func newFreemarker() IEngine {
	return &freemarker{
		Engine: Engine{
			name: "Freemarker",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Freemarker_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Freemarker_2", Value: "${21389*90}", Expected: "1925010"},
				{ Name: "Freemarker_3", Value: "{{21389*90}}", Expected: "{{21389*90}}"},
				{ Name: "Freemarker_legacy", Value: "#{{21389*90}}", Expected: "1925010"},
				{ Name: "Freemarker_alternative", Value: "[=21389*90]", Expected: "1925010"},
			},
		},
	}
}

type velocity struct {
	Engine
}

func newVelocity() IEngine {
	return &velocity{
		Engine: Engine{
			name: "Velocity",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Velocity_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Velocity_2", Value: "${21389*90}", Expected: "1925010"},
				{ Name: "Velocity_3", Value: "#set($x=21389*90)${x}", Expected: "1925010"},
			},
		},
	}
}

type groovy struct {
	Engine
}

func newGroovy() IEngine {
	return &groovy{
		Engine: Engine{
			name: "Groovy",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Groovy_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Groovy_2", Value: "${21389*90}", Expected: "1925010"},
			},
		},
	}
}

type thymeleaf struct {
	Engine
}

func newThymeleaf() IEngine {
	return &thymeleaf{
		Engine: Engine{
			name: "Thymeleaf",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Thymeleaf_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Thymeleaf_2", Value: "${21389*90}", Expected: "1925010"},
				{ Name: "Thymeleaf_3", Value: "[[${21389*90}]]", Expected: "1925010"},
				{ Name: "Thymeleaf_4", Value: "[(${21389*90})]", Expected: "1925010"},
			},
		},
	}
}

type spring struct {
	Engine
}

func newSpring() IEngine {
	return &spring{
		Engine: Engine{
			name: "Spring",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Spring_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "Spring_2", Value: "${21389*90}", Expected: "1925010"},
				{ Name: "Spring_3", Value: "#{21389+219839}", Expected: "241228"},
				{ Name: "Spring_4", Value: "#{21389*90}", Expected: "1925010"},
				{ Name: "Spring_5", Value: "*{21389+219839}", Expected: "241228"},
				{ Name: "Spring_6", Value: "*{21389*90}", Expected: "1925010"},
				{ Name: "Spring_7", Value: "@{21389+219839}", Expected: "241228"},
				{ Name: "Spring_8", Value: "@{21389*90}", Expected: "1925010"},
				{ Name: "Spring_9", Value: "~{21389+219839}", Expected: "241228"},
				{ Name: "Spring_10", Value: "~{21389*90}", Expected: "1925010"},
			},
		},
	}
}

type pebble struct {
	Engine
}

func newPebble() IEngine {
	return &pebble{
		Engine: Engine{
			name: "Pebble",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Pebble_1", Value: "{{ someString.toUPPERCASE() }}", Expected: "SOMESTRING"},
				{ Name: "Pebble_2", Value: "{% 21389*90 %}", Expected: "1925010"},
			},
		},
	}
}

type jinjava struct {
	Engine
}

func newJinjava() IEngine {
	return &jinjava{
		Engine: Engine{
			name: "Jinjava",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Jinjava_1", Value: "{{ request }}", Expected: ".context.TemplateContextRequest@"},
				{ Name: "Jinjava_2", Value: "{{'someString'.toUpperCase()}} ", Expected: "SOMESTRING"},
			},
		},
	}
}

type hubl struct {
	Engine
}

func newHubl() IEngine {
	return &hubl{
		Engine: Engine{
			name: "HubL",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "HubL_1", Value: "{{'someString'.toUpperCase()}}", Expected: "SOMESTRING"},
				{ Name: "HubL_2", Value: "{{'someString'.concat('b')}}", Expected: "someStringb"},
				{ Name: "HubL_3", Value: "{{'a'.getClass()}}", Expected: "java.lang.String"},
				{ Name: "HubL_4", Value: "{{request.getClass()}}", Expected: "com.hubspot.content.hubl.context.TemplateContextRequest"},
				{ Name: "HubL_5", Value: "{{request.getClass().getDeclaredMethods()[0]}}", Expected: "public boolean com.hubspot.content.hubl.context.TemplateContextRequest.isDebug()"},
			},
		},
	}
}

type expressionLanguage struct {
	Engine
}

func newExpressionLanguage() IEngine {
	return &expressionLanguage{
		Engine: Engine{
			name: "ExpressionLanguage",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "ExpressionLanguage_1", Value: "${21389+219839}", Expected: "241228"},
				{ Name: "ExpressionLanguage_2", Value: "${21389*90}", Expected: "1925010"},
				{ Name: "ExpressionLanguage_3", Value: "${\"someString\"}", Expected: "\"someString\""},
				{ Name: "ExpressionLanguage_4", Value: "#{21389+219839}", Expected: "241228"},
				{ Name: "ExpressionLanguage_5", Value: "#{21389*90}", Expected: "1925010"},
			},
		},
	}
}