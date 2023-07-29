package models

import (
	"encoding/xml"
	"fmt"
)

type Mmweather struct {
	XMLName xml.Name `xml:"MMWEATHER"`
	Text    string   `xml:",chardata"`
	REPORT  struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		TOWN struct {
			Text      string `xml:",chardata"`
			Index     string `xml:"index,attr"`
			Sname     string `xml:"sname,attr"`
			Latitude  string `xml:"latitude,attr"`
			Longitude string `xml:"longitude,attr"`
			FORECAST  []struct {
				Text      string `xml:",chardata"`
				Day       string `xml:"day,attr"`
				Month     string `xml:"month,attr"`
				Year      string `xml:"year,attr"`
				Hour      string `xml:"hour,attr"`
				Tod       string `xml:"tod,attr"`
				Predict   string `xml:"predict,attr"`
				Weekday   string `xml:"weekday,attr"`
				PHENOMENA struct {
					Text          string `xml:",chardata"`
					Cloudiness    string `xml:"cloudiness,attr"`
					Precipitation string `xml:"precipitation,attr"`
					Rpower        string `xml:"rpower,attr"`
					Spower        string `xml:"spower,attr"`
				} `xml:"PHENOMENA"`
				PRESSURE struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"PRESSURE"`
				TEMPERATURE struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"TEMPERATURE"`
				WIND struct {
					Text      string `xml:",chardata"`
					Min       string `xml:"min,attr"`
					Max       string `xml:"max,attr"`
					Direction string `xml:"direction,attr"`
				} `xml:"WIND"`
				RELWET struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"RELWET"`
				HEAT struct {
					Text string `xml:",chardata"`
					Min  string `xml:"min,attr"`
					Max  string `xml:"max,attr"`
				} `xml:"HEAT"`
			} `xml:"FORECAST"`
		} `xml:"TOWN"`
	} `xml:"REPORT"`
}

func (t Mmweather) CreateString(i int) []string {
	var array []string

	// Time
	day := t.REPORT.TOWN.FORECAST[i].Day
	month := t.REPORT.TOWN.FORECAST[i].Month
	hour := t.REPORT.TOWN.FORECAST[i].Hour
	timeColumn := fmt.Sprintf("%v/%v %v:00", day, month, hour)
	array = append(array, timeColumn)

	// Temp
	tempMin := t.REPORT.TOWN.FORECAST[i].TEMPERATURE.Min
	tempMax := t.REPORT.TOWN.FORECAST[i].TEMPERATURE.Max
	tempColumn := fmt.Sprintf("%v-%v C", tempMin, tempMax)
	array = append(array, tempColumn)

	// Wind
	wildMin := t.REPORT.TOWN.FORECAST[i].WIND.Min
	wildMax := t.REPORT.TOWN.FORECAST[i].WIND.Max
	wildDir := t.REPORT.TOWN.FORECAST[i].WIND.Direction
	switch wildDir {
	case "1":
		wildDir = "N"
	case "2":
		wildDir = "NE"
	case "3":
		wildDir = "E"
	case "4":
		wildDir = "SE"
	case "5":
		wildDir = "S"
	case "6":
		wildDir = "SW"
	case "7":
		wildDir = "W"
	case "8":
		wildDir = "NW"
	}

	wildColumn := fmt.Sprintf("%v-%v m/s %v", wildMin, wildMax, wildDir)
	array = append(array, wildColumn)

	// Verbose
	verCloud := t.REPORT.TOWN.FORECAST[i].PHENOMENA.Cloudiness
	switch verCloud {
	case "-1":
		verCloud = "Fog"
	case "0":
		verCloud = "Clear"
	case "1":
		verCloud = "Low cloudy"
	case "2":
		verCloud = "Cloudy"
	case "3":
		verCloud = "Overcast"
	}

	verPrecip := t.REPORT.TOWN.FORECAST[i].PHENOMENA.Precipitation
	switch verPrecip {
	case "3":
		verPrecip = "Mixed"
	case "4":
		verPrecip = "Rain"
	case "5":
		verPrecip = "Showers"
	case "6", "7":
		verPrecip = "Snow"
	case "8":
		verPrecip = "Thunderstorm"
	case "9":
		verPrecip = ""
	case "10":
		verPrecip = "No precipitation"
	}

	verRpower := t.REPORT.TOWN.FORECAST[i].PHENOMENA.Rpower
	switch verRpower {
	case "0":
		verRpower = "Possible precipitation"
	case "1":
		verRpower = "Fallout"
	}

	verWetMin := t.REPORT.TOWN.FORECAST[i].RELWET.Min
	verWetMax := t.REPORT.TOWN.FORECAST[i].RELWET.Max
	verColumn := fmt.Sprintf("%v, %v, %v, %v-%v%% rh", verCloud, verPrecip, verRpower, verWetMin, verWetMax)

	array = append(array, verColumn)
	return array
}
