package repository

import(
	//"../model"
	//"github.com/khalid/apiWithGO/model"
	"github.com/khalidalhabibie/APIWITHGO/model"
	"gorm.io/gorm"
)


type UserRepository interface{
	AddUser(model.User) (model.User, error)
	GetUser(int)(model.User,error)
	GetByEmail(string)(model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
	GetProductTransaction(int) ([]model.Transaction, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.connection.First(&user, id).Error
}

func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}

func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.connection.Find(&users).Error
}



func (db *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Model(&user).Updates(&user).Error
}

func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetProductTransaction(userID int) (orders []model.Transaction, err error) {
	return orders, db.connection.Where("user_id = ?", userID).Set("gorm:auto_preload", true).Find(&orders).Error
}