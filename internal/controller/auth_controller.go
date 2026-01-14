package controllers

import (
    "oen-peternakan/internal/config"
    "oen-peternakan/internal/models"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

// Register menambahkan user baru
func Register(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Gagal membaca data"})
    }

    // Cek username sudah ada atau belum
    var existingUser models.User
    config.DB.Where("username = ?", user.Username).First(&existingUser)
    if existingUser.ID != 0 {
        return c.Status(400).JSON(fiber.Map{"message": "Username sudah terdaftar"})
    }

    // Hashing password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 10)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Gagal mengamankan password"})
    }
    user.PasswordHash = string(hashedPassword)

    // Simpan ke DB
    config.DB.Create(&user)
    return c.JSON(fiber.Map{"message": "Registrasi berhasil"})
}

// Login memvalidasi user
func Login(c *fiber.Ctx) error {
    var input models.User
    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Gagal membaca data"})
    }

    var user models.User
    // Cari user berdasarkan username
    config.DB.Where("username = ?", input.Username).First(&user)

    if user.ID == 0 {
        return c.Status(401).JSON(fiber.Map{"message": "Username atau password salah"})
    }

    // Cek password hash
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.PasswordHash)); err != nil {
        return c.Status(401).JSON(fiber.Map{"message": "Username atau password salah"})
    }

    // (Di tahap selanjutnya, Anda bisa generate JWT Token di sini)
    // Untuk sekarang, kembalikan data user
    return c.JSON(fiber.Map{
        "message": "Login berhasil",
        "user":    user,
    })
}