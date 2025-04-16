/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get a desired Gopher",
	Long: `This command will call a Github repository in order to return the desired Gopher`,
	Run: func(cmd *cobra.Command, args []string) {
		// default name
		var gopherName = "dr-who.png"

		// if there are arguments and the first (arg[0]) is not empty,
		// use arg[0] as the gopherName
		if len(args) >= 1 && args[0] !=""{
			gopherName = args[0]
		}

		// build URL
		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

		fmt.Println("Trying to get '" + gopherName + "' Gopher...")

		// get the data, check if error
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		// defer (await) the response
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Create the file
			out, err := os.Create(gopherName + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			// Write the body to file
			_, err = io.Copy(out, response.Body)
			if err!= nil{
				fmt.Println(err)
			}

			fmt.Println("Perfect! Saved Gopher in " + out.Name() +"!")

		} else {
			fmt.Println("Error fetching gopher... :(")
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
