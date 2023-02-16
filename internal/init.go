package internal

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func initConfig() {
	// 配置路径
	baseDir := "configs"
	path := baseDir + "/" + cfgFilename + ".yaml"

	// 读取匹配的环境变量
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "[config] config not found >>", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stdout, "[config] using config file >>", viper.ConfigFileUsed())
	}
}
