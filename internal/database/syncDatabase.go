package database

import "github.com/gasuhwbab/goAuth/internal/model"

func SyncDatabase() {
	DB.AutoMigrate(&model.User{})
}
