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

func UserLogin(userData *domain.User) (error, *domain.User) {
	res, err := repository.FindUserByEmail(userData)
	if err != nil {
		return err, nil
	}

	if !res.Isverified {
		return errors.New("Invalid User"), nil
	}

	if res.Isadmin == false {
		isVerified := util.VerifyPassword(userData.Password, res.Password)
		if userData.Email == res.Email && isVerified == true {
			fmt.Println(userData.Password, res.Password)

			return nil, res
		}
	}

	if res.Isadmin == true {
		if userData.Email == res.Email && userData.Password == res.Password {
			return nil, res
		}
	}
	return errors.New("Incorrect Password"), nil

}

func GetUsers() ([]domain.User, error) {
	var users []domain.User
	res, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range *res {
		if !user.Isadmin {
			users = append(users, user)
		}
	}
	return users, nil
}

func GetUserByID(userId int) (*domain.User, error) {
	var user domain.User
	res, err := repository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	if !res.Isadmin {
		user = *res
	}
	return &user, nil
}

func BlockUser(id int) error {
	err := repository.BlockUser(id)
	if err != nil {
		return err
	}
	return nil
}

func UnblockUser(id int) error {
	err := repository.UnblockUser(id)
	if err != nil {
		return err
	}
	return nil
}

// User- Address-----------------------------------------------------------------------------------------------

func AddAddress(userAdd *domain.Address, id int) error {
	err := repository.AddAddress(userAdd, id)
	if err != nil {
		return err
	}
	return nil
}

func EditUserProfile(user domain.UserProfileUpdate, id int) error {
	user.Password = util.HashPassword(user.Password)
	err := repository.EditUserProfile(user, id)
	if err != nil {
		return err
	}

	return nil
}

func ProfileDetails(id int) (*domain.UserProfileUpdate, error) {
	userDetails, err := repository.ProfileDetails(id)
	if err != nil {
		return nil, err
	}
	return userDetails, nil
}
