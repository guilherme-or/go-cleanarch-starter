// Makefile entity command entry point
// This file is used to create an application entity
package main

import (
	"github.com/guilherme-or/go-cleanarch-starter/cmd/scaffold"
	"github.com/guilherme-or/go-cleanarch-starter/pkg/util"
)

func main() {
	entityName := util.Ask("What's the entity name? (e.g. User, Post, Comment)")
	util.Normalize(&entityName)

	wantTests := util.AskBool("Create a file for %s tests?", entityName)
	wantRepository := util.AskBool("Create a file for %s repository?", entityName)
	wantDTOs := util.AskBool("Create a file for %s DTOs?", entityName)

	scaffold.Entity(entityName)

	if wantTests {
		scaffold.EntityTest(entityName)
	}
	if wantRepository {
		scaffold.Repository(entityName)
	}
	if wantDTOs {
		scaffold.DTO(entityName)
	}

	util.Say("\nAll set for %s entity files!", entityName)
}
