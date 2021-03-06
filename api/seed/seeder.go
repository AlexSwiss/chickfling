package seed

import (
	"log"

	"github.com/AlexSwiss/chickfling/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Firstname: "Alexander",
		Lastname:  "Ikeh",
		Email:     "aleks@gmail.com",
		Phone:     "08182432388",
		Bio:       "An introvert, software Developer",
		Password:  "password",
	},
	models.User{
		Firstname: "Jon",
		Lastname:  "Snow",
		Email:     "snow@gmail.com",
		Phone:     "27002983999",
		Bio:       "An extrovert, software Developer",
		Password:  "password",
	},
}

// Load data
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}
