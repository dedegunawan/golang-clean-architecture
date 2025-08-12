package app

import (
	"time"

	"github.com/dedegunawan/golang-clean-architecture/config"
	"github.com/dedegunawan/golang-clean-architecture/internal/domain/user"
	"github.com/dedegunawan/golang-clean-architecture/internal/repository/mysql"
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
	gormLg := gormLogger.New(nil, gormLogger.Config{
		SlowThreshold: time.Second,
		LogLevel:      gormLogger.Warn,
	})

	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{Logger: gormLg})
	if err != nil { return nil, err }

	// router + middleware
	r := server.New(lg)

	// module User
	userRepo := mysql.NewUserRepository(db)
	userSvc := user.NewService(userRepo, lg)
	userHttp.RegisterRoutes(r.Engine, userSvc)

	return &App{Router: r, DB: db}, nil
}
