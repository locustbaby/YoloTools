package utils

import "gopkg.in/yaml.v2"

// ParseYAML parses the content of a YAML file and returns a Go map.
func ParseYAML(content []byte) (map[string]interface{}, error) {
	valuesMap := make(map[string]interface{})
	err := yaml.Unmarshal(content, &valuesMap)
	if err != nil {
		return nil, err
	}
	return valuesMap, nil
}
