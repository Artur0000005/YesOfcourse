package encoding

import (
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"

	"encoding/json"
	"os"
)

type JSONData struct {
	DockerCompose models.DockerCompose
	FileInput     string
	FileOutput    string
}

type YAMLData struct {
	DockerCompose models.DockerCompose
	FileInput     string
	FileOutput    string
}

type MyEncoder interface {
	Encoding() error
}

func (j *JSONData) Encoding() error {
	
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		return err
	}

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}

	yamFile, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}

defer yamFile.Close()

_, err = yamFile.Write(yamlData)
if err != nil {
	return err
}

	return nil
}

func (y *YAMLData) Encoding() error {
	
	yamFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamFile, &y.DockerCompose)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}