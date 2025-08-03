package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "shopping-cart/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to DB")
    }

    db.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.CartItem{}, &models.Order{})
    DB = db
}
