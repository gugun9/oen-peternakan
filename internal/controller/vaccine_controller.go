package controllers

import (
    "oen-peternakan/internal/config"
    "oen-peternakan/internal/models"
    "github.com/gofiber/fiber/v2"
)

// GetVaccines mengambil semua data vaksin
func GetVaccines(c *fiber.Ctx) error {
    var vaccines []models.Vaccine
    // Preload Supplier untuk mengambil detail supplier juga
    config.DB.Preload("Supplier").Find(&vaccines)
    return c.JSON(vaccines)
}

// CreateVaccine menambah vaksin baru
func CreateVaccine(c *fiber.Ctx) error {
    var vaccine models.Vaccine
    if err := c.BodyParser(&vaccine); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }

    // Validasi Stok Minus (Double check di backend)
    if vaccine.Stock < 0 {
        return c.Status(400).JSON(fiber.Map{"message": "Stok tidak boleh negatif"})
    }

    config.DB.Create(&vaccine)
    return c.JSON(vaccine)
}

// UpdateVaccine mengedit data vaksin
func UpdateVaccine(c *fiber.Ctx) error {
    id := c.Params("id")
    var vaccine models.Vaccine

    if err := config.DB.First(&vaccine, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Data tidak ditemukan"})
    }

    var updateData models.Vaccine
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }

    // Validasi Stok
    if updateData.Stock < 0 {
        return c.Status(400).JSON(fiber.Map{"message": "Stok tidak boleh negatif"})
    }

    config.DB.Model(&vaccine).Updates(updateData)
    return c.JSON(fiber.Map{"message": "Berhasil update vaksin"})
}

// DeleteVaccine menghapus vaksin
func DeleteVaccine(c *fiber.Ctx) error {
    id := c.Params("id")
    config.DB.Delete(&models.Vaccine{}, id)
    return c.JSON(fiber.Map{"message": "Berhasil menghapus vaksin"})
}