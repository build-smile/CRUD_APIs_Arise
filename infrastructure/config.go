package infrastructure

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	CFG = Config{}
)

func InitConfig() {

	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&CFG)
	log.Println("Config loaded successfully", CFG.App.Name, "on port", CFG.App.Port)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	CFG.Jwt.SecretKey = []byte(CFG.Jwt.SecretKeyString)

}

type Config struct {
	App App      `mapstructure:"app"`
	DB  DBConfig `mapstructure:"db"`
	Jwt Jwt      `mapstructure:"jwt"`
}
type App struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}
type DBConfig struct {
	MyDb Database `mapstructure:"mydb"`
}
type Database struct {
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}
type Jwt struct {
	SecretKeyString string `mapstructure:"secret-key"`
	SecretKey       []byte
}
