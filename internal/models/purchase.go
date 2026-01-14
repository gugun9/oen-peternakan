package models

import (
    "time"
)

// Purchase merepresentasikan tabel purchases
type Purchase struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    
    // Relasi ke Vaccine (Belongs To)
    Vaccine     Vaccine   `json:"vaccine" gorm:"foreignKey:VaccineID"`
    VaccineID   uint      `json:"vaccine_id"`
    
    Quantity    int       `json:"quantity" gorm:"not null"`
    TotalPrice  float64   `json:"total_price" gorm:"type:decimal(15,2)"`
    Description string    `json:"description" gorm:"type:text"`
    
    PurchaseDate time.Time `json:"purchase_date"`
    
    // Relasi ke User (Audit: siapa yang input)
    CreatedBy   uint      `json:"created_by"`
    CreatedAt   time.Time `json:"created_at"`
}