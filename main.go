package main

import (
	"github.com/rodrinoblega/prop-filter/src/adapters/cli"
	"github.com/rodrinoblega/prop-filter/src/adapters/filters_provider"
	"github.com/rodrinoblega/prop-filter/src/adapters/readers"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/rodrinoblega/prop-filter/src/use_cases"
	"log"
)

func main() {
	args := cli.ParseFlags()

	propertyReader, err := readers.NewJSONPropertyReader(args)
	if err != nil {
		log.Fatalf("error with input: %v", err.Error())
	}

	filterProvider := filters_provider.NewArgsFilterProvider(args)

	propertyFinder := use_cases.NewPropertyFinder(propertyReader, filterProvider)

	properties := propertyFinder.Execute()

	log.Printf("Loaded %d properties after filtering.\n", len(properties))

	printProperties(properties)
}

func printProperties(properties []entities.Property) {
	log.Println("Filtered Properties:")
	for i, prop := range properties {
		log.Printf("%d. %+v\n", i+1, prop)
	}
}
