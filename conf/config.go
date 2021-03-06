package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/ericyaoxr/mcube/logger/zap"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CIPHER_TEXT_PREFIX = "@ciphered@"
)

func newConfig() *Config {
	return &Config{
		App:   newDefaultAPP(),
		Log:   newDefaultLog(),
		MySQL: newDefaultMySQL(),
	}
}

type Config struct {
	App   *app   `toml:"app"`
	Log   *log   `toml:"log"`
	MySQL *mySQL `toml:"mysql"`
}

type app struct {
	Name       string `toml:"name" env:"APP_NAME"`
	HttpHost   string `toml:"http_host" env:"APP_HTTP_HOST"`
	HttpPort   string `toml:"http_port" env:"APP_HTTP_PORT"`
	GRPCHost   string `toml:"grpc_host" env:"APP_GRPC_HOST"`
	GRPCPort   string `toml:"grpc_port" env:"APP_GRPC_PORT"`
	EncryptKey string `toml:"encrypt_key" env:"APP_ENCRYPT_KEY"`
}

func (a *app) HTTPAddr() string {
	return fmt.Sprintf("%s:%s", a.HttpHost, a.HttpPort)
}

func (a *app) GRPCAddr() string {
	return fmt.Sprintf("%s:%s", a.GRPCHost, a.GRPCPort)
}

func newDefaultAPP() *app {
	return &app{
		Name:       "cmdb",
		HttpHost:   "0.0.0.0",
		HttpPort:   "8060",
		GRPCHost:   "0.0.0.0",
		GRPCPort:   "18060",
		EncryptKey: "defualt app encrypt key",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

// newDefaultLog todo
func newDefaultLog() *log {
	return &log{
		Level:  "debug",
		Format: "text",
		To:     "stdout",
	}
}

// MySQL todo
type mySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`

	lock sync.Mutex
}

var (
	db *sql.DB
)

func (m *mySQL) GetDB() (*sql.DB, error) {
	// ???????????????????????????
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func (m *mySQL) getDBConn() (*sql.DB, error) {
	zap.L().Infof("connect to mysql: %s:%s", m.Host, m.Port)
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

// newDefaultMySQL todo
func newDefaultMySQL() *mySQL {
	return &mySQL{
		Database: "cmdb",
		Host:     "127.0.0.1",
		Port:     "3306",
		// MaxOpenConn: 200,
		// MaxIdleConn: 50,
		// MaxLifeTime: 1800,
		// MaxIdleTime: 600,
		MaxOpenConn: 600,
		MaxIdleConn: 200,
		MaxLifeTime: 3600,
		MaxIdleTime: 1200,
	}
}
