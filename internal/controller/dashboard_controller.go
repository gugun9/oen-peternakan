package controllers

import (
    "oen-peternakan/internal/config" 
    "oen-peternakan/internal/models" 
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// GetDashboardStats mengambil ringkasan data untuk dashboard
func GetDashboardStats(c *fiber.Ctx) error {
    var vaccineCount int64
    var supplierCount int64
    var scheduleCount int64
    var lowStockVaccines []models.Vaccine

    // 1. Hitung total vaksin
    config.DB.Model(&models.Vaccine{}).Count(&vaccineCount)

    // 2. Hitung stok menipis (< 2)
    config.DB.Where("stock < ?", 2).Find(&lowStockVaccines)

    // 3. Hitung total supplier
    config.DB.Model(&models.Supplier{}).Count(&supplierCount)

    // 4. Hitung jadwal aktif
    config.DB.Model(&models.Schedule{}).Count(&scheduleCount)

    return c.JSON(fiber.Map{
        "total_vaksin":    vaccineCount,
        "total_supplier":  supplierCount,
        "total_jadwal":    scheduleCount,
        "low_stock_count": len(lowStockVaccines),
        "low_stock_items": lowStockVaccines,
    })
}

// GetDashboardChart mengambil data khusus untuk grafik (Chart.js)
func GetDashboardChart(c *fiber.Ctx) error {
    var vaccines []models.Vaccine

    // Ambil semua data vaksin
    config.DB.Find(&vaccines)

    // Siapkan data untuk Chart.js
    var labels []string
    var dataStock []int

    for _, v := range vaccines {
        labels = append(labels, v.Name)
        dataStock = append(dataStock, v.Stock)
    }

    return c.JSON(fiber.Map{
        "labels": labels,
        "data":   dataStock,
    })
}

// ExportExcel mengexport data vaksin menjadi file .csv (bisa dibuka di Excel)
func ExportExcel(c *fiber.Ctx) error {
    // Ambil data vaksin beserta supplier
    var vaccines []models.Vaccine
    config.DB.Preload("Supplier").Find(&vaccines)

    // Set header agar browser menganggap ini file Excel/CSV
    c.Set("Content-Type", "text/csv")
    c.Set("Content-Disposition", "attachment; filename=laporan_vaksin.csv")

    // Tulis Header CSV
    c.WriteString("Nama Vaksin,Stok,Nama Supplier\n")

    // Tulis baris data
    for _, v := range vaccines {
        supplierName := v.Supplier.Name
        if supplierName == "" {
            supplierName = "Tidak ada supplier"
        }
        c.WriteString(fmt.Sprintf("%s,%d,%s\n", v.Name, v.Stock, supplierName))
    }

    return nil
}

// ExportPDF mengexport data (Simulasi)
func ExportPDF(c *fiber.Ctx) error {
    c.Set("Content-Type", "text/plain")
    c.Set("Content-Disposition", "attachment; filename=laporan.txt")

    content := "LAPORAN STOK VAKSIN\n====================\n\n" +
        "Silakan gunakan fitur 'Print' di menu Laporan Web untuk menyimpan PDF.\n" +
        "Backend ini hanya mendukung Export Excel (.csv) secara langsung."

    c.WriteString(content)
    return nil
}