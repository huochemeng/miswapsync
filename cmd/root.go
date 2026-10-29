package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "ycb",
	Short: "ycb server template short",
	Long:  "ycb server template long long long, very huge",
	// 注释掉Run，root只是当做一个根命令的容器
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("ycb server run..., huge")
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root init")
	cobra.OnInitialize(initConfig)
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&cfgFile, "config", "c", "./config/config_import.toml", "config file (default is $HOME/.config_import.toml)")
}

func initConfig() {
	fmt.Println("root initConfig")
	if cfgFile != "" {
		//从flag中获取配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		//主目录 /Users/$home$
		dir, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// 在主目录下搜索文件名为config_import的配置文件
		viper.AddConfigPath(dir)
		viper.SetConfigName("config_import")
	}
	viper.AutomaticEnv()
	viper.SetConfigType("toml")
	viper.SetEnvPrefix("MiSwap")
	//设置替换器，在环境变量中大多数不支持.的形式，这里viper自动做替换。如配置文件中database.host，就会去环境变量中找MiSwap_DATABASE_HOST
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	//真正读取找到的配置文件
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
}
