package dbc

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseOptions 数据库配置
type DatabaseOptions struct {
	Dialact               string
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

// DSN 通过 DatabaseOptions 返回 DSN
func (opts *DatabaseOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")
}

func NewDatabase(opts *DatabaseOptions) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbIns, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns 设置到数据库的最大打开连接数
	dbIns.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime 设置连接可重用的最长时间
	dbIns.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns 设置空闲连接池的最大连接数
	dbIns.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
