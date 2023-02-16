package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cfgFilename viper 配置文件名
var cfgFilename = "app"

func NewApp() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "",
		Short:        "",
		Long:         "",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runApp()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 额外命令
	cmd.PersistentFlags().StringVarP(&cfgFilename, "config", "c", "app", "The path to the app configuration file. Empty string for no configuration file.")

	// 初始化viper配置
	cobra.OnInitialize(initConfig)

	return cmd
}

// runApp App 实际业务代码入口
func runApp() error {
	return nil
}
