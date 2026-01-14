package models

import (
    "time"
)

// Schedule merepresentasikan tabel schedules
type Schedule struct {
    ID                uint      `json:"id" gorm:"primaryKey"`
    LivestockIDOrName string    `json:"livestock_id_or_name" gorm:"type:varchar(100);not null"`
    
    // Relasi ke Vaccine (Belongs To)
    Vaccine           Vaccine   `json:"vaccine" gorm:"foreignKey:VaccineID"`
    VaccineID         uint      `json:"vaccine_id"`
    
    ScheduleDate      time.Time `json:"schedule_date" gorm:"type:date"`
    Status            string    `json:"status" gorm:"type:varchar(20);default:'Terjadwal'"`
    CreatedAt         time.Time `json:"created_at"`
}