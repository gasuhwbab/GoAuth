package initializer

import "github.com/gasuhwbab/goAuth/internal/database"

func InitDb() {
	database.ConnectToDb()
	database.SyncDatabase()
}
