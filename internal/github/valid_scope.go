package github

import "slices"

func IsValidScope(scope string, scopes []string) bool {
	if slices.Contains(scopes, scope) {
		return true
	}
	return false
}

func IsValidCommand(command string, rules map[string]map[string][]string, scope string) bool {
	if _, ok := rules[scope][command]; ok {
		return true
	}
	return false
}

func IsValidFlag(flag string, flags []string) bool {
	if slices.Contains(flags, flag) {
		return true
	}
	return false
}
