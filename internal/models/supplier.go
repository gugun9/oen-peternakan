package models

import (
    "time"
)

// Supplier merepresentasikan tabel suppliers
type Supplier struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    Name          string    `json:"name" gorm:"type:varchar(100);not null"`
    ContactPerson string    `json:"contact_person" gorm:"type:varchar(50)"`
    Phone         string    `json:"phone" gorm:"type:varchar(20)"`
    Address       string    `json:"address" gorm:"type:text"`
    CreatedAt     time.Time `json:"created_at"`
}