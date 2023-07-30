package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Conf struct {
	Url string
}

func (t *Conf) GetUrl(path string) string {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Configuration file %v does not exist.\n"+
			"Open https://meteodays.com/en/content/export.\n"+
			"And generate url for your location:\nurl: ", path)

		fmt.Scan(&t.Url)

		file, _ := json.Marshal(t)

		err = os.WriteFile(path, file, 0644)
		if err != nil {
			fmt.Println(err)
		}
		return t.Url
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.NewDecoder(file).Decode(&t); err != nil {
		fmt.Printf("JSON unmarshaling failed: %s", err)
	}
	return t.Url
}
