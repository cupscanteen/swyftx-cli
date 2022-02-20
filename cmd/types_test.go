package cmd

import (
	"reflect"
	"testing"
)

func Test_CalculateMetadata(t *testing.T) {
	t.Run("CalculateMetadata returns an empty struct when supplied totalRecords=0", func(t *testing.T) {
		totalRecords := 0
		pageSize := 20
		page := 1

		result := CalculateMetadata(totalRecords, page, pageSize)
		if !reflect.DeepEqual(result, Metadata{}) {
			FatalfFormatter(t, result, Metadata{})
		}
	})
	t.Run("CalculateMetadata returns a valid struct", func(t *testing.T) {

		totalRecords := 100
		pageSize := 20
		page := 1

		result := CalculateMetadata(totalRecords, page, pageSize)
		expected := Metadata{
			CurrentPage:  1,
			PageSize:     20,
			FirstPage:    1,
			LastPage:     5,
			TotalRecords: 100,
		}
		if !reflect.DeepEqual(result, expected) {
			FatalfFormatter(t, result, expected)
		}
	})
}
