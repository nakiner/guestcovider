package database

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	_ "os"
	"time"
)

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string `mapstructure:"database_name"`
	Secure       string
}

type Connection struct {
	Master  *gorm.DB
}

func Connect(ctx context.Context, master Config) (*Connection, error) {
	var res Connection
	var err error

	// connect to master
	res.Master, err = ConnectPool(ctx, master)
	if err != nil {
		return nil, errors.Wrap(err, "Master DB connect")
	}

	return &res, nil
}

func ConnectPool(ctx context.Context, db Config) (conn *gorm.DB, err error) {
	dsn := url.URL{
		User:     url.UserPassword(db.User, db.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:     db.DatabaseName,
		RawQuery: (&url.Values{"sslmode": []string{db.Secure}}).Encode(),
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	conn, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: dbLogger,
	})

	return conn, err
}

func (c *Connection) Close() {
	db, _ :=c.Master.DB()
	db.Close()
}

func GetMasterConn(ctx context.Context, conn *Connection) (*gorm.DB, error) {
	dbConn := conn.Master.WithContext(ctx)

	return dbConn, nil
}
