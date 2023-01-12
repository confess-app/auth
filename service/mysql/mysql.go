package mysql

import (
	"auth/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TableName = "user"
)

var Client *gorm.DB

func Init() {
	conn := os.Getenv("MYSQL_DATABASE_URI")
	var err error
	Client, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = Client.AutoMigrate(&model.User{})

	if err != nil {
		panic(err)
	}
}

func QueryUserByUsername(username string) (*model.User, error) {
	var user model.User
	Client.Table(TableName).Where("username = ?", username).Scan(&user)
	return &user, nil
}

func QueryUserByEmail(email string) (*model.User, error) {
	var user model.User
	Client.Table(TableName).Where("email = ?", email).Scan(&user)
	return &user, nil
}

func Save(user *model.User) error {
	Client.Table(TableName).Save(user)
	return nil
}
