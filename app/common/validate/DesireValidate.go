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
	}

	scenes := map[string][]string{
		"add":     {"desire", "wishman_name", "wishman_qq"},
		"light":   {"id"},
		"achieve": {"id"},
		"byid":    {"id"},
		"cancel":  {"id", "message"},
		"getUser": {""},
	}
	DesireValidate.Rules = rules
	DesireValidate.Scenes = scenes
}
