package models

func Rules() map[string]map[string][]string {
	rules := map[string]map[string][]string{
		"user": {
			"pushes":  {"--limit"},
			"pulls":   {"--limit", "--state"},
			"issues":  {"--limit", "--state"},
			"watches": {"--limit"},
			"summary": {"--limit"},
		},
		"repo": {
			"info": {"--limit"},
		},
		"set": {
			"token": {},
		},
		"get": {
			"token": {},
		},
	}

	return rules
}

func Scopes() []string {
	scopes := []string{
		"user",
		"repo",
		"set",
		"get",
	}

	return scopes
}
