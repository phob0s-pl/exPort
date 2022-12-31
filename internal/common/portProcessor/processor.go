package portProcessor

import (
	"encoding/json"
	"fmt"
	"github.com/phob0s-pl/exPort/domain"
	"io"
	"os"
)

type Port struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type PortCallback func(port *domain.Port)

func ProcessJSONFile(fileName string, onPort PortCallback) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return processJSONStream(file, onPort)
}

func processJSONStream(reader io.Reader, onPort PortCallback) error {
	decoder := json.NewDecoder(reader)

	// opening bracket
	if _, err := decoder.Token(); err != nil {
		return err
	}

	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return err
		}

		keyValue, ok := key.(string)
		if !ok {
			return fmt.Errorf("expected key to be string value")
		}

		var port Port

		err = decoder.Decode(&port)
		if err != nil {
			return err
		}

		onPort(&domain.Port{
			Key:         keyValue,
			Name:        port.Name,
			Coordinates: port.Coordinates,
			City:        port.City,
			Province:    port.Province,
			Country:     port.Country,
			Alias:       port.Alias,
			Regions:     port.Regions,
			Timezone:    port.Timezone,
			Unlocs:      port.Unlocs,
			Code:        port.Code,
		})
	}

	return nil
}
