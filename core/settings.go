package core

import (
	"encoding/json"
	"os"
)

type Settings struct {
	SettingsPath string `json:"SettingsPath"`
	ItemPath     string `json:"ItemPath"`
}

func NewSettings() (*Settings, error) {
	file, err := os.Open("settings.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Create a new settings file
			file, err = os.Create("settings.json")
			if err != nil {
				return nil, err
			}
			defer file.Close()

			// Initialize default settings
			defaultSettings := &Settings{
				SettingsPath: "./settings.json",
				ItemPath:     "./items.json",
			}

			// Encode and write default settings to the file
			encoder := json.NewEncoder(file)
			err = encoder.Encode(defaultSettings)
			if err != nil {
				return nil, err
			}

			return defaultSettings, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	settings := Settings{}
	err = decoder.Decode(&settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
