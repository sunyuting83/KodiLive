package controller

import (
	"errors"
	"strings"
)

// GetIndex get index data
func GetIndex(url string) (d string, e error) {
	var (
		config ConfigFile
		err    error
	)
	config, err = GetConfig()
	if err != nil {
		return "", errors.New("cant find config file")
	}
	d = MakeData(config.Menu, url, false)
	if len(config.Tags) > 0 {
		t := MakeData(config.Tags, url, true)
		d = strings.Join([]string{"", t}, d)
	}
	return d, nil
}

// MakeData make data fun
func MakeData(list []string, url string, t bool) (data string) {
	var (
		i        string
		n        string
		d        string
		u        string
		channels string = "/tags/"
	)
	if t == false {
		data = "#EXTM3U"
		channels = "/channels/"
	}
	for _, item := range list {
		i = strings.Join([]string{"", item}, "#EXTINF:-1,")
		n = strings.Join([]string{"", ".m3u"}, item)
		d = strings.Join([]string{"", n}, channels)
		u = strings.Join([]string{"\r", d}, url)
		data += strings.Join([]string{"\r", u}, i)
	}
	return
}
