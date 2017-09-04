package model

// Migrate will create/update models in database.
func Migrate() {
	DB.AutoMigrate(&Post{})
}
