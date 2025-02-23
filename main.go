package main

import (
	"fmt"
	"github.com/rodrinoblega/prop-filter/src/adapters/filters_provider.go"
	"github.com/rodrinoblega/prop-filter/src/adapters/repositories"
	"github.com/rodrinoblega/prop-filter/src/use_cases"
	"log"
)

func main() {

	repo, err := repositories.NewJSONPropertyRepository("./properties.json")
	if err != nil {
		log.Fatalf("Error reading properties: %v", err)
	}

	filterProvider := filters_provider_go.NewArgsFilterProvider()

	propertyFinder := use_cases.NewPropertyFinder(repo, filterProvider)

	properties, err := propertyFinder.Execute()
	if err != nil {
		log.Fatalf("Error executing property finder: %v", err)
	}

	fmt.Printf("Loaded %d properties\n", len(properties))
}
