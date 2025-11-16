package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/yadav-shubh/base-backend/config"
	"go.uber.org/zap"
	"runtime"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func finalizer(db *sql.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	once.Do(func() {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.Get().Database.User,
			config.Get().Database.Password,
			config.Get().Database.Host,
			config.Get().Database.Port,
			config.Get().Database.DBName,
			config.Get().Database.SSLMode)

		var err error

		Log.Info("Connecting to db", zap.String("conStr", connStr))
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			Log.Error("Unable to connect to db", zap.Error(err))
			panic(err)
		} else {
			err := db.Ping()
			if err != nil {
				Log.Error("Unable to ping db", zap.Error(err))
				return
			} else {
				Log.Error("Successfully connected to db")
			}
		}

		// Set the maximum number of open connections
		db.SetMaxOpenConns(viper.GetInt("db.max_open_conns"))

		// Set the maximum number of idle connections
		db.SetMaxIdleConns(viper.GetInt("db.max_idle_conns"))

		// need to control by user, not the package
		runtime.SetFinalizer(db, finalizer)
	})

	return db
}
