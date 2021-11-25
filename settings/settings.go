package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml") //指定配置文件名称（不带后缀)
	//viper.SetConfigType("yaml")   //指定配置文件类型
	viper.AddConfigPath(".")   //指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig() //读取配置文件
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper配置文件读取发生致命性错误%s\n", err)
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Sprintf("viper.Unmarshal 失败%v\n", err)
	}
	viper.WatchConfig() //当config发生改变，自动热加载
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改....")
		if err := viper.Unmarshal(Conf); err != nil { //重新加载到Conf文件中
			fmt.Sprintf("viper.Unmarshal 失败%v\n", err)
		}
	})
	return
}
