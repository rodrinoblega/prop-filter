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

type PropertyFinderInputs struct {
	PropertyReader PropertyReader
	FilterProvider FilterProvider
}

type PropertyFinder struct {
	pfi PropertyFinderInputs
}

func NewPropertyFinder(pfi PropertyFinderInputs) *PropertyFinder {
	return &PropertyFinder{pfi: pfi}
}

func (pf *PropertyFinder) Execute() ([]entities.Property, error) {
	var wg sync.WaitGroup
	var filteredProperties []entities.Property
	resultChan := make(chan entities.Property, 100)
	propertiesChan := make(chan entities.Property, 100)
	errorChan := make(chan error, 10)

	go pf.pfi.PropertyReader.FindProperties(propertiesChan, errorChan)

	go pf.handleErrors(errorChan)

	filters, err := pf.pfi.FilterProvider.GetFilters()
	if err != nil {
		return nil, err
	}

	pf.processProperties(&wg, filters, propertiesChan, resultChan)

	go pf.waitAndCloseResultChan(&wg, resultChan)

	filteredProperties = pf.collectResults(resultChan)

	return filteredProperties, nil
}

func (pf *PropertyFinder) handleErrors(errorChan chan error) {
	for err := range errorChan {
		fmt.Printf("Error: %v\n", err)
	}
}

func (pf *PropertyFinder) processProperties(wg *sync.WaitGroup, filters *entities.Filters, propertiesChan chan entities.Property, resultChan chan entities.Property) {
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
