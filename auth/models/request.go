package models

import (
	"errors"
	"github/com/protocol10/deep-dive-auth-lab/auth/validations"
)

const (
	PasswordDoesNotMatchError = "passwords do not match"
)

type UserRegisterRequest struct {
	EmailID         string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (req *UserRegisterRequest) Validate() error {
	if req.Password != req.ConfirmPassword {
		return errors.New(PasswordDoesNotMatchError)
	}

	isValid, err := validations.IsPasswordLengthValid(req.Password)
	if err != nil {
		return err
	}

	isValid, err = validations.IsPasswordComplexityValid(req.Password)
	if err != nil {
		return err
	}
	if isValid {
		return nil
	}

	return nil
}
