package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	yamlFile := "input.yaml"

	jsonFile := "output.json"

	yamlData, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v\n", err)
	}

	var yamlMap map[string]interface{}
	if err := yaml.Unmarshal(yamlData, &yamlMap); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v\n", err)
	}

	jsonData, err := json.MarshalIndent(yamlMap, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v\n", err)
	}

	if err := os.WriteFile(jsonFile, jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON file: %v\n", err)
	}

	fmt.Printf("Converted YAML to JSON successfully! Output file: %s\n", jsonFile)
}
