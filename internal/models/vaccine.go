package models

import (
    "time"
)

// Vaccine merepresentasikan tabel vaccines
type Vaccine struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"type:varchar(100);not null"`
    Stock     int       `json:"stock" gorm:"default:0"`
    
    // Relasi ke Supplier (Belongs To)
    Supplier   Supplier  `json:"supplier" gorm:"foreignKey:SupplierID"`
    SupplierID uint      `json:"supplier_id"`
    
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}