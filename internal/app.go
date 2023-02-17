package internal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valyevo/gosimple/internal/routes"
)

var (
	// cfgFilename viper 配置文件名
	cfgFilename = "app"
	// routeGroupName 应用路由组名
	routeGroupName = "app"
)

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
	if err := initGin(routes.ApiRoutes); err != nil {
		return err
	}

	return nil
}
