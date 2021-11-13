package common

import (
	"log"
	"os"
	"time"
)

const CodeExpries = -2
const CodeError = -1
const CodeSuccess = 0

const Login = true
const UnLogin = false

const WishHaveDelete = 3
const WishHaveRealize = 2
const WishHaveLight = 1
const WishNotLight = 0

const LightWish = 0
const CancelLight = 1
const DeleteWish = 2
const HaveAchieve = 3

var ChinaTime *time.Location

var RedirectURL string

func init() {
	RedirectURL = os.Getenv("redirect_url")
	var err error
	ChinaTime, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println(err)
	}
}
