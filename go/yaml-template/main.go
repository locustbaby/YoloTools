package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	// 读取 values.yaml 文件
	values, err := os.ReadFile("values.yaml")
	if err != nil {
		fmt.Println("Error reading values.yaml:", err)
		return
	}

	// 解析 values.yaml 文件，将其内容解析为 Go map
	valuesMap := make(map[string]interface{})
	err = yaml.Unmarshal(values, &valuesMap)
	if err != nil {
		fmt.Println("Error parsing values.yaml:", err)
		return
	}

	// 遍历 templates 目录下的所有文件
	templatesDir := "templates"
	err = filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 忽略目录
		if info.IsDir() {
			return nil
		}

		// 读取模板文件
		templateContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// 创建模板对象并解析模板
		tmpl, err := template.New("template").Parse(string(templateContent))
		if err != nil {
			return err
		}

		// 渲染模板并输出到标准输出
		err = tmpl.Execute(os.Stdout, valuesMap)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error processing templates:", err)
		return
	}
}
