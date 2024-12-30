package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	cl "github.com/jufianto/comic-info-scraper/services"

	"gopkg.in/yaml.v3"
)

func StoreToYaml(result []cl.InfoComic) (dataB []byte, err error) {

	fileName := fmt.Sprintf("result-scrape-%s.yaml", time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.Create("results/" + fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s", err)
	}

	infoYaml := cl.FileInfoComic{
		Total:     len(result),
		Comic:     result,
		Timestamp: time.Now(),
	}

	data, err := yaml.Marshal(infoYaml)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data %s", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to write data %s", err)
	}

	log.Printf("success store data to file %s\n", fileName)
	return data, nil
}

func ConvertToJSON(results []cl.InfoComic) (data map[string]interface{}, err error) {

	databyte, err := json.Marshal(results)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data %s", err)
	}

	return map[string]interface{}{
		"total":     len(results),
		"timestamp": time.Now(),
		"results":   string(databyte),
	}, nil
}
