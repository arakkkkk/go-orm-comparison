package config

import (
	"github.com/go-sql-driver/mysql"
)

func DSN() string {
	mc := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3305",
		DBName:               "sample_db",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return mc.FormatDSN()
}
