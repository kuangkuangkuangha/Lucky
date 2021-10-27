package helper

import "testing"

func Test(t *testing.T) {
	token := CreatToken("011111111")
	println(token)
	idcard_number, err := VerifyToken(token)
	if err != nil {
		println(err.Error())
	}
	println(idcard_number)
	idcard_number, err = VerifyToken("")
	if err != nil {
		println(err.Error())
	}
	println(idcard_number)
}
