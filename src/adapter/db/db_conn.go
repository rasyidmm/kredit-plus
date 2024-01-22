package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"kredit-plus/internal/config"
	dbConfig "kredit-plus/internal/config/db"
	"kredit-plus/src/adapter/repository/entity"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	cfg := config.GetConfig()
	DB, err = connect(cfg.Database.Kreditplus)
	if err != nil {
		panic(err)
	}
	err = MigrateSchema(DB)
	if err == nil {
		Seed(DB)
	}

}
func connect(config dbConfig.Database) (*gorm.DB, error) {
	var (
		dbConn *gorm.DB
		err    error
	)
	user := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.Dbname

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	currentWaitTime := 2
	trialCount := 0

	for dbConn == nil && trialCount < 5 {
		trialCount++
		dbConn, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			fmt.Println("unable connecting to DB.")
			if trialCount == 5 {
				return nil, err
			}
			fmt.Println("retrying in", currentWaitTime, "seconds...")
			time.Sleep(time.Duration(currentWaitTime) * time.Second)
			currentWaitTime = currentWaitTime * 2
			dbConn = nil
		}
	}
	conn, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(7)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(1 * time.Hour)

	return dbConn, err
}

var tables = []interface{}{
	&entity.UsersEntity{},
	&entity.MasterTenorEntity{},
	&entity.LoansEntity{},
	&entity.BillingsEntity{},
}

func MigrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}
