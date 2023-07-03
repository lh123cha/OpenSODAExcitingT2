/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show data from OpenDigger",
	Long:  `show data from api and print in screen`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取结果
		repoInfo = getResult(queryPara)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		

		// 检查是否有compare,这里只有一个subcommand
		if (strings.Contains(cmd.CommandPath(), "compare") ){
			// TODO qk: a more pretty output?
			fmt.Println(repoInfo)
			fmt.Println(repoInfoCompare)
		} else{
			fmt.Println(repoInfo)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}