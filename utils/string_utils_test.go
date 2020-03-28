package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteEmpty(t *testing.T) {
	values := []string{"This", "", "is", "", "a", "", "slice", "", "with", "", "empty", "", "values"}
	expectedResult := []string{"This", "is", "a", "slice", "with", "empty", "values"}

	result := DeleteEmpty(values)

	assert.Equal(t, result, expectedResult)
}
