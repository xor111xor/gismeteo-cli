package configs

import (
	"bytes"
	"os"
	"testing"
)

func TestGetUrl(t *testing.T) {
	in := bytes.NewBufferString("https://xml.meteoservice.ru/export/gismeteo/point/9615.xml")

	myConf := new(Conf)
	got := myConf.GetUrl("test_config", in)
	defer os.Remove("test_config")
	want := "https://xml.meteoservice.ru/export/gismeteo/point/9615.xml"

	if got != want {
		t.Errorf("Got %s not equal Want %s", got, want)
	}

}
