package main

import (
	"github.com/guilherme-or/go-cleanarch-starter/cmd/scaffold"
	"github.com/guilherme-or/go-cleanarch-starter/pkg/util"
)

func main() {
	entityName := util.Ask("What's the entity name? (e.g. User, Post, Comment)")
	util.Normalize(&entityName)

	scaffold.DTO(entityName)

	util.Say("\nAll set for %s DTO file!", entityName)
}