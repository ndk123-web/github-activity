package models

func Rules() map[string]map[string][]string {
	rules := map[string]map[string][]string{
		"user": {
			"pushes": {"--limit"},
			"pulls":  {"--limit"},
			"issues": {"--state"},
		},
		"repo": {
			"info": {},
		},
	}

	return rules
}

func Scopes() []string {
	scopes := []string{
		"user",
		"repo",
	}

	return scopes
}
