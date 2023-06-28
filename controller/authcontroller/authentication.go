package authcontroller

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func SessionChecker(c *fiber.Ctx) error {
	envFile := godotenv.Load()

	if envFile != nil {
		fmt.Print(envFile)
	}

	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "You are not logged in"})
	}

	jwt_secret := os.Getenv("API_SECRET")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(jwt_secret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid token claim"})
	}

	var account models.Account

	database.DB.First(&account, "account_id = ?", fmt.Sprint(claims["i"]))

	ourIntegerUser, err := strconv.Atoi(fmt.Sprint(claims["i"]))

	if account.AccountId != ourIntegerUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "the user belonging to this token no logger exists"})
	}

	return c.Next()
}

func Login(c *fiber.Ctx) error {
	envFile := godotenv.Load()

	if envFile != nil {
		fmt.Print(envFile)
	}

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please fill username and password",
		})
	}

	var account models.Account

	if err := database.DB.First(&account, "account_username = ? and account_password = ?", payload.Username, payload.Password).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := jwt.MapClaims{
		"i":   account.AccountId,
		"u":   account.AccountUsername,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // in hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	})
}

func Register(c *fiber.Ctx) error {
	var account models.Account

	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Save(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(account)
}

func DetailAccount(c *fiber.Ctx) error {
	payload := struct {
		UserId int `json:"account_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var account models.Account

	database.DB.Where("account_id = ?", payload.UserId).Preload("AccountAddress", "account_address_is_active = ?", 1).Find(&account)

	return c.Status(fiber.StatusOK).JSON(account)
}
