/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	mock_rgserv "github.com/vinceyoumans/ms/mock/regserv"
)

// regservCmd represents the regserv command
var regservCmd = &cobra.Command{
	Use:   "regserv",
	Short: "One time register of services",
	Long:  `Mock register of services manually`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("regserv called")
		mock_rgserv.RegisterService00()
	},
}

func init() {
	rootCmd.AddCommand(regservCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regservCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regservCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
