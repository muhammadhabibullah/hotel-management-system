package mysql

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"hotel-management-system/domains/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.Hotel{},
			&models.Room{},
			&models.Reservation{},
			&models.Stay{},
			&models.StayRoom{},
		); err != nil {
		panic(err)
	}

	return db
}
