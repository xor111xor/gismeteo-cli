package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/xor111xor/gismeteo-cli/internal/requests"
	"os"
)

func main() {
	url := "https://xml.meteoservice.ru/export/gismeteo/point/9615.xml"

	weather, err := requests.GetWeather(url)
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
