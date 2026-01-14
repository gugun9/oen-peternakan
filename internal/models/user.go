package models

import (
    "time"
    // HAPUS BARIS INI -> "gorm.io/gorm" 
)

// User merepresentasikan tabel users
type User struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    FullName     string    `json:"full_name" gorm:"type:varchar(100);not null"`
    Username     string    `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
    PasswordHash string    `json:"-" gorm:"type:varchar(255);not null"`
    Role         string    `json:"role" gorm:"type:varchar(20);default:'peternak'"`
    CreatedAt    time.Time `json:"created_at"`
}