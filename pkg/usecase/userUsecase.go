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
		return errors.New("Invalid Email address")
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

func EditUserProfile(user domain.UserProfileUpdate, id int) (*domain.UserProfileUpdate, error) {
	var password string
	if user.Password != "" {
		password = user.Password

		user.Password = util.HashPassword(user.Password)
	}

	updatedUser, err := repository.EditUserProfile(user, id)
	if err != nil {
		return nil, err
	}
	user, err = util.BuildUserProfileUpdate(*updatedUser, password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func ProfileDetails(id int) (*domain.ProfileDetails, error) {
	userDetails, err := repository.ProfileDetails(id)
	profileDetails := util.BuildProfileDetails(userDetails)
	if err != nil {
		return nil, err
	}
	return profileDetails, nil
}

func ViewAddress(id int) ([]domain.Address, error) {
	userAdd, err := repository.ViewAddress(id)
	if err != nil {
		return nil, err
	}
	return userAdd, nil
}

func ForgotPassword(email string) error {
	user, err := repository.FindUserByEmailResetPassword(email)
	if err != nil {
		return err
	}
	otp := util.Otpgeneration(user.Email)
	err = repository.UpdateOtp(email, otp)
	if err != nil {
		return err
	}
	return nil
}

func ResetPassword(data domain.ResetPassword, email string) error {
	user, err := repository.FindUserByEmailResetPassword(email)
	if err != nil {
		return err
	}
	password := util.HashPassword(data.NewPassword)
	if email == user.Email && data.OTP == user.Otp {
		if password == user.Password {
			return errors.New("Try another password")
		}
		err := repository.ResetPassword(email, password)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid otp")
	}
	return nil
}
