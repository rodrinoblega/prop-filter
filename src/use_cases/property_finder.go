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
	GetFilters() (*entities.Filters, error)
}

type PropertyFinder struct {
	reader         PropertyReader
	filterProvider FilterProvider
	propertiesChan chan entities.Property
	errorChan      chan error
}

func NewPropertyFinder(propertiesChan chan entities.Property, errorChan chan error, reader PropertyReader, filterProvider FilterProvider) *PropertyFinder {
	return &PropertyFinder{propertiesChan: propertiesChan, errorChan: errorChan, reader: reader, filterProvider: filterProvider}
}

func (pf *PropertyFinder) Execute() ([]entities.Property, error) {
	var wg sync.WaitGroup
	var filteredProperties []entities.Property
	resultChan := make(chan entities.Property, 100)

	go pf.reader.FindProperties(pf.propertiesChan, pf.errorChan)

	go pf.handleErrors()

	filters, err := pf.filterProvider.GetFilters()
	if err != nil {
		return nil, err
	}

	pf.processProperties(&wg, filters, resultChan)

	go pf.waitAndCloseResultChan(&wg, resultChan)

	filteredProperties = pf.collectResults(resultChan)

	return filteredProperties, nil
}

func (pf *PropertyFinder) handleErrors() {
	for err := range pf.errorChan {
		fmt.Printf("Error: %v\n", err)
	}
}

func (pf *PropertyFinder) processProperties(wg *sync.WaitGroup, filters *entities.Filters, resultChan chan entities.Property) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for property := range pf.propertiesChan {
				if filters.ApplyFilters(property) {
					resultChan <- property
				}
			}
		}(i)
	}
}

func (pf *PropertyFinder) waitAndCloseResultChan(wg *sync.WaitGroup, resultChan chan entities.Property) {
	wg.Wait()
	close(resultChan)
}

func (pf *PropertyFinder) collectResults(resultChan chan entities.Property) []entities.Property {
	var filteredProperties []entities.Property
	for prop := range resultChan {
		filteredProperties = append(filteredProperties, prop)
	}
	return filteredProperties
}
