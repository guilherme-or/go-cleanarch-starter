package scaffold

import "github.com/guilherme-or/go-cleanarch-starter/pkg/util"

func template(msg, entityName string, tplType util.TemplateType) {
	util.Say(msg)
	filename := util.CreateTemplate(tplType, entityName)
	util.Say(filename)
	util.Say(" OK\n")
}

func Entity(entityName string) {
	template("\nCreating entity: ", entityName, util.EntityTempl)
}

func EntityTest(entityName string) {
	template("\nCreating tests: ", entityName, util.EntityTestTempl)
}

func DTO(entityName string) {
	template("\nCreating DTO: ", entityName, util.DTOTempl)
}

func Repository(entityName string) {
	template("\nCreating repository: ", entityName, util.RepositoryTempl)
}
