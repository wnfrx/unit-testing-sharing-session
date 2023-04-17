package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wnfrx/unit-testing-sharing-session/1-basic/utils"
)

func TestSum(t *testing.T) {
	var (
		a        int = 2
		b        int = 5
		expected int = 7
	)

	actual := utils.Sum(a, b)

	assert.Equal(t, expected, actual)
}

func TestMax_GreaterA(t *testing.T) {
	var (
		a        int = 10
		b        int = 5
		expected int = 10
	)

	actual := utils.Max(a, b)

	assert.Equal(t, expected, actual)
}

func TestMax_GreaterB(t *testing.T) {
	var (
		a        int = 5
		b        int = 10
		expected int = 10
	)

	actual := utils.Max(a, b)

	assert.Equal(t, expected, actual)
}

func TestMax(t *testing.T) {
	var testTable = []struct {
		name  string
		input struct {
			A int
			B int
		}
		expected int
	}{
		{
			name: "a < b",
			input: struct {
				A int
				B int
			}{A: 5, B: 10},
			expected: 10,
		},
		{
			name: "a < b",
			input: struct {
				A int
				B int
			}{A: 7, B: 3},
			expected: 7,
		},
		{
			name: "a = b",
			input: struct {
				A int
				B int
			}{A: 5, B: 5},
			expected: 5,
		},
	}

	for _, tc := range testTable {
		actual := utils.Max(tc.input.A, tc.input.B)

		assert.Equal(t, tc.expected, actual)
	}
}
