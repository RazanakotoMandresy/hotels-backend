package service

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func authValidator(mail string, params interface{}) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return err	
	}
	if isMail := govalidator.IsEmail(mail); !isMail {
		return fmt.Errorf("%v is not an valid mail", mail)
	}
	return nil
}
