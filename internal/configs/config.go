package configs

import (
	"log"

	"github.com/spf13/viper"
)

var config *Config

type option struct {
	configFolders []string
	configFile string
	configType string
}

func Init(opts ...Option) error {
	// opt := &option{
	// 	configFolders: getDefaultConfigFolder(),
	// 	configFile: getDefaultConfigFile(),
	// 	configType: getDefaultConfigType(),
	// }

	opt := new(option)
	opt.configFolders = getDefaultConfigFolder()
	opt.configFile = getDefaultConfigFile()
	opt.configType = getDefaultConfigType()

	for _, funcOpt := range opts{
		funcOpt(opt)
	}

	for _ , configFolder := range opt.configFolders {
		viper.AddConfigPath(configFolder)
	}

	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	config = new(Config)
	err := viper.ReadInConfig()
	
	if err != nil{
		return  err
	}

	return viper.Unmarshal(config)

}

type Option func(*option)

func getDefaultConfigFolder() []string {
	return []string{"./internal/configs/"}
}

func getDefaultConfigFile() string{
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) Option {
	return func(opt *option){
		log.Println("Running with config folder : " , configFolders)
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option){
		log.Println("Running with config file : " , configFile)
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option){
		log.Println("Running with config type : " , configType)
		opt.configType = configType
	}
}

func Get() *Config{
	if config == nil {
		return &Config{}
	}

	return config
}

