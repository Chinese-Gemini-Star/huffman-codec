package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

func GetDB() (*gorm.DB, error) {
	if dB == nil {
		dsn := "huffman_codec:BANGkY5AphMNwRHL@tcp(127.0.0.1:3306)/huffman_codec?charset=utf8mb4&parseTime=True&loc=Local"
		if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
			dB = db
			return dB, nil
		} else {
			return nil, err
		}
	} else {
		return dB, nil
	}
}
