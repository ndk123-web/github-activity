package models

func Rules() map[string]map[string][]string {
	rules := map[string]map[string][]string{
		"user": {
			"pushes": {"--limit"},
			"pulls":  {"--limit", "--state"},
			"issues": {"--state"},
		},
		"repo": {
			"info": {},
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
