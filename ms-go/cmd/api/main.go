package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	appmatrix "github.com/alonso804/ms-go/internal/application"
	httpinfra "github.com/alonso804/ms-go/internal/infrastructure/http"
	"github.com/alonso804/ms-go/internal/infrastructure/implementations"
)

type structValidator struct {
	validate *validator.Validate
}

func NewStructValidator() *structValidator {
	return &structValidator{
		validate: validator.New(),
	}
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	_ = godotenv.Load()

	app := fiber.New(fiber.Config{
		StructValidator: NewStructValidator(),
	})

	logger := slog.Default()

	// == START Injection ==

	httpClient := &http.Client{}

	statisticsProvider := implementations.NewAPIProvider(
		os.Getenv("STATISTICS_API_URL"),
		httpClient,
		logger,
	)

	jwtProvider := implementations.NewJWTProvider(os.Getenv("JWT_SECRET"))

	service := appmatrix.NewService(
		statisticsProvider,
	)

	authService := appmatrix.NewAuthService(jwtProvider)

	handler := httpinfra.NewHandler(
		service,
	)

	authHandler := httpinfra.NewAuthHandler(
		authService,
	)

	httpinfra.RegisterRoutes(
		app,
		handler,
		authHandler,
		jwtProvider,
	)

	// == END Injection ==

	PORT := os.Getenv("PORT")

	log.Println("Server is running on port " + PORT)

	log.Fatal(app.Listen(":" + PORT))
}
