package seeds

import (
	"log"
	"login-register/models"

	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
)

func hash(pw string) string {
	password, err := argon2id.CreateHash(pw, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return password
}

var users = []models.User{
	models.User{
		Username: "afrizal",
		Email:    "asd@gmail.com",
		Password: hash("admin"),
	},
	models.User{
		Username: "rizal",
		Email:    "ds@gmail.com",
		Password: hash("admin"),
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
