package tools

import (
	"io/ioutil"
	"encoding/json"
)

type config struct {
  SaveDB string `json:"save_db"`
  LogFile string `json:"log_file"`
	TemplateFolder string `json:"tpl_folder"`
}

var Config config

func LoadConfig(filename string) error {

  file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

  err = json.Unmarshal(file, &Config)

	if err != nil {
		return err
	}

  return nil

}
