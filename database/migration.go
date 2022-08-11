package database

import (
	"fmt"
	"waysbucks/models"
	"waysbucks/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println(("Migration Success"))
}
