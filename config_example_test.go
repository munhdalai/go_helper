package go_helper_test

import (
	"fmt"

	go_helper "github.com/munhdalai/go_helper"
)

type AppConfig struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
}

func ExampleLoadConfig() {
	// config.yaml файл дараах агуулгатай гэж үзвэл:
	//   port: 8080
	//   host: "localhost"
	//   database: "mydb"

	config, err := go_helper.LoadConfig[AppConfig]("config", "yaml", "./")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}

	fmt.Println(config.Port)
	fmt.Println(config.Host)
}
