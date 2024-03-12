/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	bcserv01 "github.com/vinceyoumans/ms/mock/bcserv01"

	"github.com/spf13/cobra"
)

// bcserv01Cmd represents the bcserv01 command
var bcserv01Cmd = &cobra.Command{
	Use:   "bcserv01",
	Short: "Sends a mock service message",
	Long:  `A mock service `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bcserv01 called")
		bcserv01.BCServ01("http://localhost:8080/v1/status")
	},
}

func init() {
	rootCmd.AddCommand(bcserv01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bcserv01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bcserv01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
