package model

import (
	"gin-monitor/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)



// Config 配置参数
type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// NewDB 创建DB实例
func NewDB(c *Config) (*gorm.DB, error) {
	db, err := gorm.Open(c.DBType, c.DSN)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	err = db.DB().Ping()
	db.SingularTable(true)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	if dbType := viper.GetString("gorm.db_type"); dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		&Schedule{},
		&Task{},
		).Error
}

func InitMysql() error {
	db ,err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	return AutoMigrate(db)
}

func getDB() (*gorm.DB, error) {
	dns := config.MySQLDSN()
	return NewDB(&Config{
		Debug:        viper.GetBool("gorm.debug"),
		DBType:       viper.GetString("gorm.db_type"),
		DSN:          dns,
		MaxLifetime:  viper.GetInt("gorm.max_lifetime"),
		MaxOpenConns: viper.GetInt("gorm.max_open_conns"),
		MaxIdleConns: viper.GetInt("gorm.max_idle_conns"),
	})
}