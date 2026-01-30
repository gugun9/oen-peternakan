package main

import (
    "oen-peternakan/internal/config"        
    "oen-peternakan/internal/controller"   
    "oen-peternakan/internal/middleware"    
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // 1. Inisialisasi Database
    config.InitDB()

    // 2. Inisialisasi Fiber App
    app := fiber.New(fiber.Config{
        Prefork:       false,
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "Farm Management API",
    })

    // 3. Middleware
    app.Use(cors.New())
    app.Use(logger.New())

    // 4. Setup Static Files (HTML Frontend)
    app.Static("/", "./static")

    // 5. Setup API Routes
    api := app.Group("/api")

    // --- PUBLIC ROUTES (Tanpa Login) ---
    api.Post("/register", controllers.Register)
    api.Post("/login", controllers.Login)

    // --- PROTECTED ROUTES (Butuh Auth) ---
    // Dashboard
    api.Get("/dashboard", middleware.AuthRequired, controllers.GetDashboardStats)
    api.Get("/dashboard/chart", middleware.AuthRequired, controllers.GetDashboardChart)
    api.Get("/export/excel", middleware.AuthRequired, controllers.ExportExcel)
    api.Get("/export/pdf", middleware.AuthRequired, controllers.ExportPDF)

    // Vaccines
    api.Get("/vaccines", middleware.AuthRequired, controllers.GetVaccines)
    api.Post("/vaccines", middleware.AuthRequired, controllers.CreateVaccine)
    api.Put("/vaccines/:id", middleware.AuthRequired, controllers.UpdateVaccine)
    api.Delete("/vaccines/:id", middleware.AuthRequired, controllers.DeleteVaccine)

    // Supplier
    api.Get("/suppliers", middleware.AuthRequired, controllers.GetSuppliers)
    api.Post("/suppliers", middleware.AuthRequired, controllers.CreateSupplier)
    api.Put("/suppliers/:id", middleware.AuthRequired, controllers.UpdateSupplier)
    api.Delete("/suppliers/:id", middleware.AuthRequired, controllers.DeleteSupplier)

    // Purchases
    api.Get("/purchases", middleware.AuthRequired, controllers.GetPurchases)
    api.Post("/purchases", middleware.AuthRequired, controllers.CreatePurchase)
    api.Delete("/purchases/:id", middleware.AuthRequired, controllers.DeletePurchase)

    // Schedules
    api.Get("/schedules", middleware.AuthRequired, controllers.GetSchedules)
    api.Post("/schedules", middleware.AuthRequired, controllers.CreateSchedule)
    api.Put("/schedules/:id", middleware.AuthRequired, controllers.UpdateSchedule)
    api.Delete("/schedules/:id", middleware.AuthRequired, controllers.DeleteSchedule)

    // 6. Menjalankan Server
    log.Println("Server berjalan di http://localhost:3000")
    err := app.Listen(":3000")
    if err != nil {
        log.Fatal("Gagal menjalankan server:", err)
    }
}