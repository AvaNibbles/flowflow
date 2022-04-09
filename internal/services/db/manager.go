package db

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/avanibbles/flowflow/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func New(db *config.DatabaseConfig, logger *zap.Logger) (*gorm.DB, error) {
	switch strings.ToLower(db.Type) {
	case "postgres":
		return newPostgres(db.Postgres, logger)
	default:
		return nil, errors.New(fmt.Sprintf("unknown db type: \"%s\"", db.Type))
	}
}

func newPostgres(db *config.PostgresDatabaseConfig, logger *zap.Logger) (*gorm.DB, error) {
	var user *url.Userinfo = nil
	if len(db.Username) > 0 {
		if len(db.Password) > 0 {
			user = url.UserPassword(db.Username, db.Password)
		} else {
			user = url.User(db.Username)
		}
	}

	dsn := url.URL{
		Scheme:   "postgres",
		User:     user,
		Host:     fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:     db.Database,
		RawQuery: (&url.Values{"sslmode": []string{db.SslMode}}).Encode(),
	}

	dbLogger := zapgorm2.New(logger.With(zap.String("component", "orm")))
	// dbLogger.LogLevel = gormlog.Info
	return gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: dbLogger,
	})
}
