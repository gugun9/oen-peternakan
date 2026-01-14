package controllers

import (
    "oen-peternakan/internal/config"
    "oen-peternakan/internal/models"
    "github.com/gofiber/fiber/v2"
)

func GetSchedules(c *fiber.Ctx) error {
    var schedules []models.Schedule
    // Preload Vaccine agar user tahu nama vaksinnya, bukan cuma ID
    config.DB.Preload("Vaccine").Find(&schedules)
    return c.JSON(schedules)
}

func CreateSchedule(c *fiber.Ctx) error {
    var schedule models.Schedule
    if err := c.BodyParser(&schedule); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }
    config.DB.Create(&schedule)
    return c.JSON(schedule)
}

func UpdateSchedule(c *fiber.Ctx) error {
    id := c.Params("id")
    var schedule models.Schedule
    if err := config.DB.First(&schedule, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Data tidak ditemukan"})
    }

    var updateData models.Schedule
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }

    config.DB.Model(&schedule).Updates(updateData)
    return c.JSON(fiber.Map{"message": "Berhasil update jadwal"})
}

func DeleteSchedule(c *fiber.Ctx) error {
    id := c.Params("id")
    config.DB.Delete(&models.Schedule{}, id)
    return c.JSON(fiber.Map{"message": "Berhasil menghapus jadwal"})
}