package use_cases

import (
	"github.com/rodrinoblega/prop-filter/src/entities"
)

type PropertyRepository interface {
	FindProperties() ([]entities.Property, error)
}

type FilterProvider interface {
	GetFilters() (*entities.Filters, error)
}

type PropertyFinder struct {
	repo           PropertyRepository
	filterProvider FilterProvider
}

func NewPropertyFinder(repo PropertyRepository, filterProvider FilterProvider) *PropertyFinder {
	return &PropertyFinder{repo: repo, filterProvider: filterProvider}
}

func (pf *PropertyFinder) Execute() ([]entities.Property, error) {
	properties, err := pf.repo.FindProperties()
	if err != nil {
		return []entities.Property{}, err
	}

	filters, err := pf.filterProvider.GetFilters()
	if err != nil {
		return []entities.Property{}, err
	}

	var filteredProperties []entities.Property
	for _, property := range properties {
		if filters.ApplyFilters(property) {
			filteredProperties = append(filteredProperties, property)
		}
	}

	return filteredProperties, nil
}
