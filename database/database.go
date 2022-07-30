package database

var db *gorm.DB

func StartDatabase() {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=admin"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{}) {
		log.Fatal("Error: ", err)
	}

	db := database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func getDatabase() *gorm.DB {
	return db
}