package repository

import (
	"errors"

	"github.com/aparnasukesh/shoezone/pkg/db"
	"github.com/aparnasukesh/shoezone/pkg/domain"
	"gorm.io/gorm"
)

func CreateUser(userData *domain.User) error {
	err := db.DB.Create(&userData)
	if err != nil {
		return err.Error
	}
	return nil
}
func FindUserByEmail(userData *domain.User) (*domain.User, error) {
	dbData := &domain.User{}
	result := db.DB.Where("email = ?", userData.Email).First(dbData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return dbData, nil
}

func DeleteUserByEmail(userData *domain.User) error {
	result := db.DB.Where("email = ? ", userData.Email).Delete(userData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUsers() (*[]domain.User, error) {
	userData := []domain.User{}
	res := db.DB.Table("users").Select("id,created_at,updated_at,deleted_at,username,email,phone,isverified,isadmin,dateofbirth,gender").Find(&userData)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userData, nil
}
func GetUserByID(id int) (*domain.User, error) {
	userData := domain.User{}
	res := db.DB.Table("users").Select("id,created_at,updated_at,deleted_at,username,email,phone,isverified,isadmin,dateofbirth,gender").Where("id= ?", id).First(&userData)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userData, nil

}

func BlockUser(id int) error {
	userData := domain.User{}
	res := db.DB.First(&userData, id)
	if res.Error != nil {
		return res.Error
	}
	if userData.Isverified != false {
		userData.Isverified = false
		err := db.DB.Save(&userData)
		if err.Error != nil {
			return err.Error
		}
	} else {
		return errors.New("User Already Blocked")
	}
	return nil

}

func UnblockUser(id int) error {
	userData := domain.User{}
	res := db.DB.First(&userData, id)
	if res.Error != nil {
		return res.Error
	}
	if userData.Isverified == false {
		userData.Isverified = true
		err := db.DB.Save(&userData)
		if err != nil {
			return err.Error
		}
	} else {
		return errors.New("User Already Unblocked")
	}
	return nil

}

func AddAddress(userAdd *domain.Address, id int) error {
	userAdd.UserID = id
	if err := db.DB.Create(&userAdd).Error; err != nil {
		return err
	}
	return nil
}

func EditUserProfile(updateuser domain.UserProfileUpdate, id int) (*domain.User, error) {
	db.DB.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)

	result := db.DB.Model(&domain.User{}).Where("id = ?", id).Updates(updateuser)
	if result.Error != nil {
		return nil, result.Error
	}
	updatedUser := domain.User{}
	if err := db.DB.Where("id=?", id).First(&updatedUser).Error; err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func ProfileDetails(id int) (*domain.UserProfileUpdate, error) {
	userDetails := domain.UserProfileUpdate{}
	if err := db.DB.Model(&domain.User{}).Where("id=?", id).Preload("DefaultAddress").First(&userDetails).Error; err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func ViewAddress(id int) ([]domain.Address, error) {
	userAdd := []domain.Address{}
	if err := db.DB.Where("user_id=?", id).Find(&userAdd).Error; err != nil {
		return nil, err
	}
	return userAdd, nil
}

func CheckValidAddress(userId, addressId int) error {
	address := domain.Address{}
	if err := db.DB.Where("id=? AND user_id=?", addressId, userId).First(&address).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByEmailResetPassword(email string) (*domain.User, error) {
	user := domain.User{}
	if err := db.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateOtp(email, otp string) error {
	user := &domain.User{}
	if err := db.DB.Model(user).Where("email =?", email).Update("otp", otp).Error; err != nil {
		return err
	}
	return nil
}

func ResetPassword(email, password string) error {
	user := &domain.User{}
	if err := db.DB.Model(&user).Where("email=?", email).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
