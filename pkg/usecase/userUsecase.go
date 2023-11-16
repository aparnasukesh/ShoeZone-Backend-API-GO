package usecase

import (
	"errors"
	"fmt"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/repository"
	"github.com/aparnasukesh/shoezone/pkg/util"
)

func RegisterUser(userData *domain.User) error {

	validateErr := util.ValidateUser(*userData)
	if validateErr != nil {
		return validateErr
	}

	_, err := repository.FindUserByEmail(userData)
	if err != nil {
		otp := util.Otpgeneration(userData.Email)
		userData.Otp = otp
		pass := util.HashPassword(userData.Password)
		userData.Password = pass
		err := repository.CreateUser(userData)
		return err
	}
	return errors.New("User already exist")
}

func RegisterValidate(userData *domain.User) error {
	enterdOtp := userData.Otp
	res, err := repository.FindUserByEmail(userData)
	if err != nil {
		return err
	}

	fmt.Println("this is : ", enterdOtp, "thi is db otp : ", res.Otp)

	if userData.Email == res.Email && enterdOtp == res.Otp {
		return nil
	}

	err = repository.DeleteUserByEmail(userData)
	if err != nil {
		return err
	}

	return errors.New("Invalid otp")
}
