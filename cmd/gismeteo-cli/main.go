package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/xor111xor/gismeteo-cli/internal/configs"
	"github.com/xor111xor/gismeteo-cli/internal/requests"
	"os"
)

func main() {

	// Configure flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Cli tool for get weather \n")
		flag.PrintDefaults()
	}
	flagVersion := flag.Bool("v", false, "show version")
	flagConfig := flag.String("config", "$HOME/.config/gismeteo-cli.json", "set configuration file")
	flag.Parse()

	var AppVersion string

	if *flagVersion {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	config := fmt.Sprint(os.ExpandEnv(*flagConfig))

	// Get data
	myConf := new(configs.Conf)
	myConf.GetUrl(config, nil)

	weather, err := requests.GetWeather(myConf.Url)
	if err != nil {
		fmt.Println(err)
	}

	var data [][]string
	forecast := weather.REPORT.TOWN.FORECAST
	for i := range forecast {
		data = append(data, weather.CreateString(i))
	}

	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"DATE", "TEMP", "WILD", "VERBOSE"})

	for _, v := range data {
		table.Append(v)
	}

	table.SetBorder(true)
	table.SetRowLine(true)
	table.Render()
}
