/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const url = "http://localhost:3001/api/v1/rates"

var (
	pair string
)

// rootCmd represents the base command when called without any subcommands
var rateCmd = &cobra.Command{
	Use:   "api/v1/rates",
	Short: "Comand for get one pair rates",
	Long: `This is command, which used for get request to local server. It has flag pari, which must contain one pair. 
		   In response you get map of pair and price, and you can see print wiht this price.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := getRate(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rateCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rateCmd.Flags().StringVar(&pair, "pair", "", "Pair for rate")

	if err := rateCmd.MarkFlagRequired("pair"); err != nil {
		fmt.Println(err)
	}
}

func getRate() (float32, error) {
	var client = http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("pairs", pair)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	var mapResp map[string]float32

	if err := json.NewDecoder(resp.Body).Decode(&mapResp); err != nil {
		return 0, err
	}

	value, ok := mapResp[pair]
	if !ok {
		return 0, errors.New("Pair not Found")
	}

	return value, nil
}
