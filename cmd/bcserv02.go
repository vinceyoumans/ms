/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	bcser02 "github.com/vinceyoumans/ms/mock/bcserv02"
)

// bcserv02Cmd represents the bcserv02 command
var bcserv02Cmd = &cobra.Command{
	Use:   "bcserv02",
	Short: "Mock Multiple Service BROADCAST Service",
	Long: `Mock Multiple Broadcast Service 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bcserv02 called")

		mockJsonFile, _ := cmd.Flags().GetString("MockJSONFile")
		url, _ := cmd.Flags().GetString("serviceURL")
		bcser02.BC02Serv(mockJsonFile, url)
	},
}

func init() {
	rootCmd.AddCommand(bcserv02Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bcserv02Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bcserv02Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//bcserv02Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bcserv02Cmd.Flags().String("MockJSONFile", "./mockJSON/BCServ02.json", "Override location of a MOCK JSON file")
	bcserv02Cmd.Flags().String("serviceURL", "localhost:8080/", "service location to BroadCast")

}
