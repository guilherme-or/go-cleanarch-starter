package util

import (
	"os"
	"strings"
	"text/template"
)

type TemplateType uint8

const (
	EntityTempl TemplateType = iota
	EntityTestTempl
	RepositoryTempl
	DTOTempl
)

type templateData struct {
	EntityName string
}

func CreateTemplate(t TemplateType, entityName string) string {
	var tplName, fileName string
	lowerEntityName := strings.ToLower(entityName)

	switch t {
	case EntityTempl:
		// Create entity template
		tplName = "/pkg/util/template/entity"
		fileName = "/internal/domain/entity/" + lowerEntityName + ".go"
	case EntityTestTempl:
		// Create entity test template
		tplName = "/pkg/util/template/entity_test"
		fileName = "/internal/domain/entity/" + lowerEntityName + "_test.go"
	case RepositoryTempl:
		// Create repository template
		tplName = "/pkg/util/template/repository"
		fileName = "/internal/domain/repository/" + lowerEntityName + "_repository.go"
	case DTOTempl:
		// Create DTO template
		tplName = "/pkg/util/template/dto"
		fileName = "/internal/domain/dto/" + lowerEntityName + "_dto.go"
	}

	tpl := parseTemplate(BasePath + tplName)
	return writeTemplate(tpl, BasePath+fileName, &templateData{EntityName: entityName})
}

func parseTemplate(filenames ...string) *template.Template {
	tpl, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}
	return tpl
}

func writeTemplate(tpl *template.Template, filename string, data *templateData) string {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	return file.Name()
}
