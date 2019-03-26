package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type RegisterReq struct {
	Username       string `validate:"gt=0"`
	PasswordNew    string `validate:"gt=0"`
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	Email          string `validate:"email"`
}

var req = RegisterReq{
	Username:       "Xargin",
	PasswordNew:    "ohno",
	PasswordRepeat: "ohn",
	Email:          "alex@abc.com",
}

func main() {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err)
	}
}
