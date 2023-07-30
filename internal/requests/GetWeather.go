package requests

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/xor111xor/gismeteo-cli/internal/models"
	"io"
	"net/http"
	"time"
)

func GetWeather(url string) (models.Mmweather, error) {
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()

	var weather models.Mmweather

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
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
