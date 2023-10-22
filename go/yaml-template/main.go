package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	// Define command-line flags
	valuesFile := flag.String("values", "", "values file")
	templateFile := flag.String("template", "", "template file or directory path")
	outputDir := flag.String("output", "", "output directory path")

	// Parse command-line flags
	flag.Parse()

	// Check if required flags are provided
	if *templateFile == "" {
		flag.PrintDefaults()
		return
	}

	// Print flag values
	fmt.Println("Template file path:", *templateFile)
	fmt.Println("Output directory path:", *outputDir)

	// Read values.yaml file
	values, err := os.ReadFile(*valuesFile)
	if err != nil {
		fmt.Println("Error reading values.yaml:", err)
		return
	}

	// Parse values.yaml file and convert its content to a Go map
	valuesMap := make(map[string]interface{})
	err = yaml.Unmarshal(values, &valuesMap)
	if err != nil {
		fmt.Println("Error parsing values.yaml:", err)
		return
	}

	// Create the output directory
	if *outputDir != "" {
		err = os.MkdirAll(*outputDir, 0755)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	// Traverse all files in the templates directory
	templatesDir := *templateFile
	err = filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Ignore directories
		if info.IsDir() {
			return nil
		}

		// Read the template file
		templateContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Create and parse the template object
		tmpl, err := template.New("template").Parse(string(templateContent))
		if err != nil {
			return err
		}

		// Render the template
		var outputContent strings.Builder
		err = tmpl.Execute(&outputContent, valuesMap)
		if err != nil {
			return err
		}

		// Print to standard output
		fmt.Println(outputContent.String())

		// Write to file
		if *outputDir != "" {
			outputFileName := filepath.Join(*outputDir, info.Name())
			outputFile, err := os.Create(outputFileName)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			_, err = outputFile.WriteString(outputContent.String())
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error processing templates:", err)
		return
	}
}
