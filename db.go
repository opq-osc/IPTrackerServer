package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type VistorItem struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Timestamp int64  `json:"timestamp"`
}

type Vistor struct {
	gorm.Model
	Key        string
	VistorItem VistorItem `gorm:"embedded"`
}

var DB *gorm.DB

const dbName = "vistors.db"

func initDB(verbose bool) (err error) {

	var level logger.LogLevel
	if verbose {
		level = logger.Info
	} else {
		level = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: level,
			Colorful: true,
		},
	)

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %s", err)
	}

	db.AutoMigrate(&Vistor{})

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	return
}

// InsertItem 插入一条请求
func InsertItem(key, ip, userAgent string) {
	DB.Create(
		&Vistor{
			Key: key,
			VistorItem: VistorItem{
				IP:        ip,
				UserAgent: userAgent,
				Timestamp: time.Now().Unix(),
			},
		},
	)
}

// GetItems 返回key对应的所有请求
func GetItems(key string) []VistorItem {

	var vistors []Vistor
	DB.Where("key = ?", key).Find(&vistors)

	var items []VistorItem

	for _, vistor := range vistors {
		items = append(items, vistor.VistorItem)
	}

	return items
}

// CheckKey key已存在返回false， 不存在返回true
func CheckKey(key string) bool {
	var vistors []Vistor
	DB.Limit(1).Where("key = ?", key).Find(&vistors)
	return len(vistors) == 0
}
