package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    // 1. Load file .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: File .env tidak ditemukan, menggunakan environment variable sistem")
    }

    // 2. Ambil data dari environment
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    // 3. Buat Koneksi Database (METODE CONCATENATION / GABUNGAN)
    dsn := "host=" + host + 
           " port=" + port + 
           " user=" + user + 
           " password=" + password + 
           " dbname=" + dbname + 
           " sslmode=disable TimeZone=Asia/Jakarta"

    // Debug: Cetak di terminal supaya tahu isinya
    log.Println("Mencoba koneksi ke DB...")
    log.Println("DSN:", dsn)

    // 4. Buka Koneksi
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi ke Database:", err)
    }

    log.Println("Koneksi Database Berhasil!")
    return DB
}