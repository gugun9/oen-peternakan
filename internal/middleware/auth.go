package middleware

import (
    "fmt"
    "os"
    "strings"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

// AuthRequired adalah middleware yang mengecek apakah user login
func AuthRequired(c *fiber.Ctx) error {
    // 1. Ambil Header Authorization
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Token tidak ditemukan. Silakan login.",
        })
    }

    // 2. Format token harus: "Bearer <token>"
    tokenString := strings.Split(authHeader, " ")[1]
    if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Format token salah.",
        })
    }

    // 3. Parsing Token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validasi algoritma signing
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Metode signing token salah: %v", token.Header["alg"])
        }

        secret := os.Getenv("JWT_SECRET")
        if secret == "" {
            secret = "rahasia_farm_project_123" // GANTI INI DI PRODUCTION!
        }
        return []byte(secret), nil
    })

    if err != nil || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Token tidak valid atau kadaluarsa.",
        })
    }

    // 4. Ambil klaim data user dari token (misal: user_id)
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Gagal membaca data user dari token.",
        })
    }

    // 5. Simpan data user di context Fiber (untuk dipakai controller jika perlu)
    c.Locals("user_id", claims["user_id"])
    c.Locals("role", claims["role"])

    return c.Next()
}

// IsAdmin adalah middleware tambahan untuk membatasi akses hanya Admin
func IsAdmin(c *fiber.Ctx) error {
    role := c.Locals("role")
    if role != "admin" {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "message": "Akses ditolak. Hanya Admin yang diperbolehkan.",
        })
    }
    return c.Next()
}