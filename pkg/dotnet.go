package gossti

type dotNet struct {
    Language
}

func newdDotNet() ILanguage {
		return &dotNet{
				Language: Language{
						name: ".Net",
						confidence: 0.0,
						Engines: []IEngine{newRazor(), newAsp()},
				},
		}
}

type Razor struct {
	Engine
}

func newRazor() IEngine {
	return &Razor{ 
		Engine: Engine{
			name: "Razor",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "Razor_1", Value: "@(21389+219839)", Expected: "43372"},
			},
		},
	}
}

type asp struct {
	Engine
}

func newAsp() IEngine {
	return &asp{ 
		Engine: Engine{
			name: "ASP",
			confidence: 0.0,
			payloads: []Payload{
				{ Name: "ASP_1", Value: "<%= 21389+219839 %>", Expected: "43372"},
				{ Name: "ASP_2", Value: "<%= \"uAnPoLhekuV6mH5Lp2re\"+\"k6GrhGWq3xwItdTkt3uB\" %>", Expected: "uAnPoLhekuV6mH5Lp2rek6GrhGWq3xwItdTkt3uB"},
			},
		},
	}
}
