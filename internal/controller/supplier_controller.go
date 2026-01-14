package controllers

import (
    "oen-peternakan/internal/config"
    "oen-peternakan/internal/models"
    "github.com/gofiber/fiber/v2"
)

func GetSuppliers(c *fiber.Ctx) error {
    var suppliers []models.Supplier
    config.DB.Find(&suppliers)
    return c.JSON(suppliers)
}

func CreateSupplier(c *fiber.Ctx) error {
    var supplier models.Supplier
    if err := c.BodyParser(&supplier); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }
    config.DB.Create(&supplier)
    return c.JSON(supplier)
}

func UpdateSupplier(c *fiber.Ctx) error {
    id := c.Params("id")
    var supplier models.Supplier
    if err := config.DB.First(&supplier, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Data tidak ditemukan"})
    }

    var updateData models.Supplier
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }

    config.DB.Model(&supplier).Updates(updateData)
    return c.JSON(fiber.Map{"message": "Berhasil update supplier"})
}

func DeleteSupplier(c *fiber.Ctx) error {
    id := c.Params("id")
    config.DB.Delete(&models.Supplier{}, id)
    return c.JSON(fiber.Map{"message": "Berhasil menghapus supplier"})
}