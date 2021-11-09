package validate

import "lucky/app/helper"

var MessageValidate helper.Validator

func init() {
	rules := map[string]string{
		"desire_id": "required",
		"message":   "required",
	}

	scenes := map[string][]string{
		"leave": {"desire_id", "message"},
		"get":   {"desire_id"},
	}
	MessageValidate.Rules = rules
	MessageValidate.Scenes = scenes
}
