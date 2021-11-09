package validate

import (
	"lucky/app/helper"
)

var UserValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id":       "required",
		"idcard_number": "required",
		"nick":          "required",
		"password":      "required",
		"school":        "required",
		"major":         "required",
		"contact":       "required",
		"email":         "required|email",
	}

	scenes := map[string][]string{
		"login": {"idcard_number", "password"},
		"email": {"email"},
	}
	UserValidate.Rules = rules
	UserValidate.Scenes = scenes
}
