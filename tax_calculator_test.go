package tax_calculator

import (
	"math"
	"testing"
)

func TestCalculatePercentageLevels(t *testing.T) {
	taxLevels := []float32{10000.0, 20000.0, 30000, 40000, 100000000}
	taxPercentages := []float32{0.09, 0.22, 0.28, 0.36, 0.44}

	var want float32
	var amount float32

	want = 90.0
	amount = 1000.0
	if res := calculatePercentageLevels(amount, taxLevels, taxPercentages); res != want {
		t.Errorf("calculatePercentageLevels() = %f, want = %f", res, want)
	}

	want = 0.09
	amount = 1
	if res := calculatePercentageLevels(amount, taxLevels, taxPercentages); res != want {
		t.Errorf("calculatePercentageLevels() = %f, want = %f", res, want)
	}

	want = 900.0
	amount = 10000.0
	var epsilon float64 = 10e-5
	if res := calculatePercentageLevels(amount, taxLevels, taxPercentages); math.Abs(float64(res-want)) > epsilon {
		t.Errorf("calculatePercentageLevels() = %f, want = %f", res, want)
	}

	want = 2000.0
	amount = 15000.0
	if res := calculatePercentageLevels(amount, taxLevels, taxPercentages); math.Abs(float64(res-want)) > epsilon {
		t.Errorf("calculatePercentageLevels() = %f, want = %f", res, want)
	}

	want = 3100.0
	amount = 20000.0
	if res := calculatePercentageLevels(amount, taxLevels, taxPercentages); math.Abs(float64(res-want)) > epsilon {
		t.Errorf("calculatePercentageLevels() = %f, want = %f", res, want)
	}
}

func TestGetReduction(t *testing.T) {
	type TestCase struct {
		amount   float32
		kids     int
		expected float32
	}
	testCases := []TestCase{
		TestCase{
			amount:   1000,
			kids:     0,
			expected: 777,
		},
		TestCase{
			amount:   1000,
			kids:     1,
			expected: 810,
		},
		TestCase{
			amount:   1000,
			kids:     2,
			expected: 900,
		},
		TestCase{
			amount:   1000,
			kids:     4,
			expected: 1340,
		},
		TestCase{
			amount:   100,
			kids:     4,
			expected: 1340,
		},
		TestCase{
			amount:   500,
			kids:     0,
			expected: 777,
		},
		TestCase{
			amount:   800,
			kids:     3,
			expected: 1120,
		},
	}
	for _, testCase := range testCases {
		if res := getReduction(testCase.amount, testCase.kids); res != testCase.expected {
			t.Errorf("getReduction() = %f, expected = %f", res, testCase.expected)
		}
	}
}

func TestGetSalaryParts(t *testing.T) {
	var expected float32 = 176.0
	res := getSalaryParts(20000, 0, 0)
	if res.eisfora != expected {
		t.Errorf("getSalaryParts() = %f, expected = %f", res.eisfora, expected)
	}
}
