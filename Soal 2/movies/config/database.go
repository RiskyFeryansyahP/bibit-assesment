package config

import (
	"fmt"
	"time"

	"github.com/RiskyFeryansyahP/bibit-movies/ent"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq" // postgres driver
)

// Database ...
type Database struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

// ConnectDB ...
func (mc *MapConfig) ConnectDB() (*ent.Client, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", mc.Database.User, mc.Database.Pass, mc.Database.Host, mc.Database.Port, mc.Database.Name, mc.Database.SSLMode)

	driver, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	db := driver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return ent.NewClient(ent.Driver(driver)), nil
}
