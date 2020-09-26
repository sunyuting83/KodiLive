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
		u        string
		channels string = "/tags/"
	)
	if t == false {
		data = "#EXTM3U\n"
		channels = "/channels/"
	}
	for _, item := range list {
		i = strings.Join([]string{"#EXTINF:-1", item}, ",")
		u = strings.Join([]string{url, channels, item, ".m3u", "\n"}, "")
		data += strings.Join([]string{i, "\n", u}, "")
	}
	return
}
