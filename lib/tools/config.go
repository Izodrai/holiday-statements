package tools

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
  SaveDB string `json:"save_db"`
  LogFile string `json:"log_file"`
	TemplateFolder string `json:"tpl_folder"`
}

func LoadConfig(filename string) (Config, error) {

  var config Config

  file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

  err = json.Unmarshal(file, &config)

	if err != nil {
		return Config{}, err
	}

  return config, nil

}
