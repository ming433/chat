package global

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Gorm struct {
	SkipDefaultTx bool
	TablePrefix   string
	SingularTable bool
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Enable   bool
	Gorm     Gorm
}

var (
	DB           *gorm.DB
	GlobalConfig Config
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.Unmarshal(&GlobalConfig)
	fmt.Println(GlobalConfig)
}
