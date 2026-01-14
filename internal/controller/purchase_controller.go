package controllers

import (
    "oen-peternakan/internal/config"
    "oen-peternakan/internal/models"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func GetPurchases(c *fiber.Ctx) error {
    var purchases []models.Purchase
    // Kita preload Vaccine dan Supplier (di dalam Vaccine) agar lengkap
    config.DB.Preload("Vaccine.Supplier").Find(&purchases)
    return c.JSON(purchases)
}

func CreatePurchase(c *fiber.Ctx) error {
    var purchase models.Purchase
    if err := c.BodyParser(&purchase); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
    }

    // LOGIKA UPDATE STOK (ANTI-CURANG)
    err := config.DB.Transaction(func(tx *gorm.DB) error {
        // 1. Simpan Data Pembelian
        if err := tx.Create(&purchase).Error; err != nil {
            return err
        }

        // 2. Update Stok Vaksin
        if err := tx.Model(&models.Vaccine{}).
            Where("id = ?", purchase.VaccineID).
            Update("stock", gorm.Expr("stock + ?", purchase.Quantity)).Error; err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Gagal memproses pembelian, transaksi dibatalkan"})
    }

    return c.JSON(fiber.Map{"message": "Pembelian berhasil & Stok ditambahkan"})
}

func DeletePurchase(c *fiber.Ctx) error {
    id := c.Params("id")
    
    // Opsional: Saat hapus pembelian, kurangi stok kembali?
    config.DB.Delete(&models.Purchase{}, id)
    return c.JSON(fiber.Map{"message": "Berhasil menghapus data pembelian"})
}