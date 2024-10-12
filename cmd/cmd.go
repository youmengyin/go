package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func tip() {
	usageStr := `欢迎使用 ` + (`go-admin v1.0.0`) + ` 可以使用 ` + (`-h`) + ` 查看命令`
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

var startCmd = &cobra.Command{
	Use:          "hh",
	Short:        "打印 'Hello World'",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing required argument")
		}
		fmt.Println("Starting the application with argument:", args[0])
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	// rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := startCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
