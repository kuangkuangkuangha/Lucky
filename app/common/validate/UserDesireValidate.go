package validate

import "lucky/app/helper"

var UserDesireValidate helper.Validator

func init() {
	rules := map[string]string{
		"id":        "required",
		"user_id":   "required",
		"desire_id": "required",
	}

	scenes := map[string][]string{
		"getUser": {"user_id"},
	}
	UserDesireValidate.Rules = rules
	UserDesireValidate.Scenes = scenes
}
