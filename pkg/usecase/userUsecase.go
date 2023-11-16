package usecase

import (
	"errors"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/repository"
	"github.com/aparnasukesh/shoezone/pkg/util"
)

func RegisterUser(userData *domain.User) error {

	validateErr := util.ValidateUser(*userData)
	if validateErr != nil {
		return validateErr
	}

	res, err := repository.FindUserByEmail(userData)

	if err != nil && res == nil {
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

	if userData.Email == res.Email && enterdOtp == res.Otp {
		return nil
	}

	err = repository.DeleteUserByEmail(userData)
	if err != nil {
		return err
	}

	return errors.New("Invalid otp")
}

func UserLogin(userData *domain.User) error {
	res, err := repository.FindUserByEmail(userData)
	if err != nil {
		return err
	}
	isVerified := util.VerifyPassword(userData.Password, res.Password)
	if userData.Email == res.Email && isVerified == true {
		return nil
	}
	return errors.New("Incorrect Password")

}

func GetUsers() (*[]domain.User, error) {
	res, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}
	return res, nil

}
