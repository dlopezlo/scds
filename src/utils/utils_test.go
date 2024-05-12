package utils

import "testing"

func TestGetUsernameValid(t *testing.T) {
	testCases := []struct {
		username string
	}{
		{"dlopez"},
		{"msolsona"},
		{"jroque"}, // too short
		{"cris.b"},
		{"xavi.gonzalez"},
		{"user77"},
		{"d-lopez"},
	}

	for _, tc := range testCases {
		t.Run("Valid username "+tc.username, func(t *testing.T) {
			result, err := GetUsername(tc.username)
			if result != tc.username {
				t.Errorf("Expected username %s. Got %s\n", tc.username, result)
			}
			if err != nil {
				t.Error("Valid usernames should not produce errors")
			}
		})
	}
}

func TestGetUsernameInvalid(t *testing.T) {
	testCases := []struct {
		username string
	}{
		{"0startswithdigit"},
		{"invalidChars.$%ˆ&@#"},
		{"a"}, // too short
		{"toolongusernametoolongusernametoolongusername"},
		{"_startswithunderscore"},
		{"-startswithhyphen"},
	}

	for _, tc := range testCases {
		t.Run("Invalid username "+tc.username, func(t *testing.T) {
			result, err := GetUsername(tc.username)
			if err == nil {
				t.Error("Invalid username should return an error")
			}
			if result != "" {
				t.Error("Invalid username should return an empty string")
			}
		})
	}
}
