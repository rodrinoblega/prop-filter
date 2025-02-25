package use_cases

import (
	"fmt"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"sync"
)

const numWorkers = 4

type PropertyReader interface {
	FindProperties(chan<- entities.Property, chan error)
}

type FilterProvider interface {
	GetFilters() *entities.Filters
}

type PropertyFinder struct {
	propertyReader PropertyReader
	filterProvider FilterProvider
}

func NewPropertyFinder(propertyReader PropertyReader, filterProvider FilterProvider) *PropertyFinder {
	return &PropertyFinder{propertyReader: propertyReader, filterProvider: filterProvider}
}

func (pf *PropertyFinder) Execute() []entities.Property {
	var filteredProperties []entities.Property
	resultChan := make(chan entities.Property, 100)
	propertiesChan := make(chan entities.Property, 100)
	errorChan := make(chan error, 10)

	go pf.propertyReader.FindProperties(propertiesChan, errorChan)

	go pf.handleErrors(errorChan)

	filters := pf.filterProvider.GetFilters()

	pf.processProperties(filters, propertiesChan, resultChan)

	close(resultChan)

	filteredProperties = pf.collectResults(resultChan)

	return filteredProperties
}

func (pf *PropertyFinder) handleErrors(errorChan chan error) {
	for err := range errorChan {
		fmt.Printf("Error: %v\n", err)
	}
}

func (pf *PropertyFinder) processProperties(filters *entities.Filters, propertiesChan chan entities.Property, resultChan chan entities.Property) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for property := range propertiesChan {
				if filters.ApplyFilters(property) {
					resultChan <- property
				}
			}
		}(i)
	}

	wg.Wait()
}

func (pf *PropertyFinder) collectResults(resultChan chan entities.Property) []entities.Property {
	var filteredProperties []entities.Property
	for prop := range resultChan {
		filteredProperties = append(filteredProperties, prop)
	}
	return filteredProperties
}
