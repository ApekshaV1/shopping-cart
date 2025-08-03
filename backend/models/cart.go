package models

import "gorm.io/gorm"

type Cart struct {
    gorm.Model
    UserID uint      `json:"user_id"`
    Items  []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
    gorm.Model
    CartID uint `json:"cart_id"`
    ItemID uint `json:"item_id"`
}
