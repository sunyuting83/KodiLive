package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// ConfigFile Config File
type ConfigFile struct {
	Menu []string `json:"menu"`
	Tags []string `json:"tags"`
}

// GetConfig get config fun
func GetConfig() (c ConfigFile, err error) {
	var p string = "config/config.json"
	data, erra := ioutil.ReadFile(p)
	if erra != nil {
		// fmt.Println("\x1B[31m错误：配置文件路径错误\x1B[0m")
		return ConfigFile{}, errors.New("cant find config file")
	}
	config := ConfigTojson(data)
	// fmt.Println(config)
	return config, nil
}

// ConfigTojson fun
func ConfigTojson(s []byte) (p ConfigFile) {
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		// fmt.Println(err.Error())
		return
	}
	return
}
