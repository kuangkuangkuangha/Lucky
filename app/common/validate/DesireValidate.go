package validate

import "lucky/app/helper"

var DesireValidate helper.Validator

func init() {
	rules := map[string]string{
		"desire":         "required",
		"id":             "required",
		"wishman_name":   "required",
		"wishman_wechat": "required",
		"wishman_tel":    "required",
		"wishman_qq":     "required",
		"message":        "required",
		"desire_id":      "required",
		"light_name":     "required",
	}

	scenes := map[string][]string{
		"add":     {"desire", "wishman_name", "wishman_qq"},
		"light":   {"desire_id", "light_name"},
		"achieve": {"id"},
		"byid":    {"id"},
		"cancel":  {"id", "message"},
		"getUser": {""},
	}
	DesireValidate.Rules = rules
	DesireValidate.Scenes = scenes
}
