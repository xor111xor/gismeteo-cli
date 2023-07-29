package requests

import (
	"encoding/xml"
	"fmt"
	"github.com/xor111xor/gismeteo-cli/internal/models"
	"io"
	"net/http"
)

func GetWeather(url string) (models.Mmweather, error) {
	var weather models.Mmweather
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return weather, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return weather, err
	}
	err = xml.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println(err)
		return weather, err
	}
	return weather, nil
}
