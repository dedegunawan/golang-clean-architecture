// internal/app/app.go
package app

import (
	"github.com/dedegunawan/golang-clean-architecture/internal/server/middleware"
	"github.com/dedegunawan/golang-clean-architecture/internal/transport/http/auth"
	"github.com/dedegunawan/golang-clean-architecture/pkg/jwtmanager"
	"log"
	"os"
	"time"

	"github.com/dedegunawan/golang-clean-architecture/config"
	"github.com/dedegunawan/golang-clean-architecture/internal/domain/user"
	repository "github.com/dedegunawan/golang-clean-architecture/internal/repository/mysql"
	"github.com/dedegunawan/golang-clean-architecture/internal/server"
	userHttp "github.com/dedegunawan/golang-clean-architecture/internal/transport/http/user"
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type App struct {
	Router *server.Engine
	DB     *gorm.DB
}

func New(cfg config.Config, lg *logger.Logger) (*App, error) {
	// ✅ Pakai stdout (atau ganti os.Stdout -> io.Discard kalau mau senyap total)
	stdlog := log.New(os.Stdout, "[gorm] ", log.LstdFlags)

	gormLg := gormLogger.New(stdlog, gormLogger.Config{
		SlowThreshold: time.Second,
		LogLevel:      gormLogger.Warn, // atur sesuai kebutuhan
	})

	// db
	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{Logger: gormLg})
	if err != nil {
		return nil, err
	}

	// load jwt library
	jwt := jwtmanager.New(cfg.JWTSecret, cfg.JWTIssuer, time.Duration(cfg.JWTExpiresMinutes)*time.Minute)

	r := server.New(lg)

	// repository
	userRepo := repository.NewUserRepository(db)

	// service / use case
	userSvc := user.NewService(userRepo, lg)

	// handler
	authHandler := auth.NewAuthHandler(userSvc)
	userHandler := userHttp.NewUserHandler(userSvc)

	// Container: middleware (jadikan gin.HandlerFunc di sini)
	m := &server.Middleware{
		AuthJWT:   middleware.AuthJWT(jwt),
		RequestID: middleware.RequestID(),
		Logger:    middleware.ZapLogger(lg),
		Recovery:  middleware.Recovery(lg),
		// unimplemented cors
		CORS: middleware.DefaultMiddleware(),
		Rate: middleware.DefaultMiddleware(300),
	}

	// Daftarkan semua route terpusat
	handlers := &server.Handlers{
		Auth: authHandler,
		User: userHandler,
	}

	server.RegisterRoutes(r.Engine, handlers, m)

	return &App{Router: r, DB: db}, nil
}
