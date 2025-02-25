package main

import (
	"github.com/rodrinoblega/prop-filter/src/adapters/filters_provider"
	"github.com/rodrinoblega/prop-filter/src/adapters/readers"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/rodrinoblega/prop-filter/src/use_cases"
	"log"
)

func main() {
	propertyReader := readers.NewJSONPropertyReader("./properties.json")

	filterProvider := filters_provider.NewArgsFilterProvider()

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
