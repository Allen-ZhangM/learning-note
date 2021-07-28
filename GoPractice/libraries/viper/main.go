package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("GoPractice/libraries/viper/")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))

	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout: ", viper.GetDuration("server.timeout"))

	fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))
	fmt.Println("mysql port: ", viper.GetInt("mysql.port"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())

	// 绑定环境变量
	viper.AutomaticEnv()
	fmt.Println("GOPATH: ", viper.Get("GOPATH"))

	viper.WatchConfig()

	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	// 手动修改文件，如果是自动保存则需要手动再保存下才会生效
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))
}
