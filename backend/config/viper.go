package config

import (
	"errors"
	"fmt"
	
	"github.com/spf13/viper"
	
)


type EnvVars struct {
	Database DatabaseConfigurations
	Port string
	AppUrl string
	AppName string
}

type DatabaseConfigurations struct {
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	Databaseip		string
}

func LoadConfig() (config EnvVars, err error) {

	viper.SetConfigName("config")	

	//viper.AddConfigPath("/etc/app_config/bookscan") //for non docker
	viper.AddConfigPath("./")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		
		return EnvVars{}, err
	} 


	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return EnvVars{}, errors.New("unable to decode into struct")
	} else {
		fmt.Println("read into struct was OK")
		fmt.Println(config.Database.DatabaseName)
		fmt.Println(config.Database.DatabaseUser)
		fmt.Println(config.Database.DatabasePassword)
		fmt.Println(config.Database.Databaseip)
		fmt.Println(config.AppUrl)
		fmt.Println(config.AppName)
	}

	return
}
