package taxcalculator

import (
	"math"
	"testing"
)

const epsilon = 10e-5

func isClose(n1 float32, n2 float32) bool {
	return math.Abs(float64(n1-n2)) < epsilon
}

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
	type TestCase struct {
		salary    float32
		insurance float32
		kids      int
		expected  SalaryParts
	}

	testCases := []TestCase{
		TestCase{
			salary:    1000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 1000,
				initialTax:    90,
				eisfora:       0,
				reduction:     810,
				totalTax:      -720,
				netSalary:     1720,
			},
		},
		TestCase{
			salary:    10000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 10000,
				initialTax:    900,
				eisfora:       0,
				reduction:     810,
				totalTax:      90,
				netSalary:     9910,
			},
		},
		TestCase{
			salary:    20000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 20000,
				initialTax:    3100,
				eisfora:       176,
				reduction:     650,
				totalTax:      2626,
				netSalary:     17374,
			},
		},
		TestCase{
			salary:    40000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 40000,
				initialTax:    9500,
				eisfora:       1326,
				reduction:     250,
				totalTax:      10576,
				netSalary:     29424,
			},
		},
		TestCase{
			salary:    1000,
			insurance: 0,
			kids:      3,
			expected: SalaryParts{
				taxableSalary: 1000,
				initialTax:    90,
				eisfora:       0,
				reduction:     1120,
				totalTax:      -1030,
				netSalary:     2030,
			},
		},
		TestCase{
			salary:    10000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 10000,
				initialTax:    900,
				eisfora:       0,
				reduction:     810,
				totalTax:      90,
				netSalary:     9910,
			},
		},
		TestCase{
			salary:    30000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 30000,
				initialTax:    5900,
				eisfora:       676,
				reduction:     450,
				totalTax:      6126,
				netSalary:     23874,
			},
		},
		TestCase{
			salary:    40000,
			insurance: 0,
			kids:      1,
			expected: SalaryParts{
				taxableSalary: 40000,
				initialTax:    9500,
				eisfora:       1326,
				reduction:     250,
				totalTax:      10576,
				netSalary:     29424,
			},
		},
		TestCase{
			salary:    80000,
			insurance: 0,
			kids:      0,
			expected: SalaryParts{
				taxableSalary: 80000,
				initialTax:    27100,
				eisfora:       4551,
				reduction:     0,
				totalTax:      31651,
				netSalary:     48349,
			},
		},
		TestCase{
			salary:    100000,
			insurance: 0,
			kids:      0,
			expected: SalaryParts{
				taxableSalary: 100000,
				initialTax:    35900,
				eisfora:       6351,
				reduction:     0,
				totalTax:      42251,
				netSalary:     57749,
			},
		},
		TestCase{
			salary:    100000,
			insurance: 2000,
			kids:      0,
			expected: SalaryParts{
				taxableSalary: 98000,
				initialTax:    35020,
				eisfora:       6171,
				reduction:     0,
				totalTax:      41191,
				netSalary:     58809,
			},
		},
	}

	for _, te := range testCases {
		r := getSalaryParts(te.salary, te.insurance, te.kids)
		if !isClose(r.taxableSalary, te.expected.taxableSalary) {
			t.Errorf("getSalaryParts().taxableSalary = %f, expected = %f", r.taxableSalary, te.expected.taxableSalary)
		}
		if !isClose(r.initialTax, te.expected.initialTax) {
			t.Errorf("getSalaryParts().initialTax = %f, expected = %f", r.initialTax, te.expected.initialTax)
		}
		if !isClose(r.eisfora, te.expected.eisfora) {
			t.Errorf("getSalaryParts().eisfora = %f, expected = %f", r.eisfora, te.expected.eisfora)
		}
		if !isClose(r.reduction, te.expected.reduction) {
			t.Errorf("getSalaryParts().reduction = %f, expected = %f", r.reduction, te.expected.reduction)
		}
		if !isClose(r.totalTax, te.expected.totalTax) {
			t.Errorf("getSalaryParts().totalTax = %f, expected = %f", r.totalTax, te.expected.totalTax)
		}
		if !isClose(r.netSalary, te.expected.netSalary) {
			t.Errorf("getSalaryParts().netSalary = %f, expected = %f", r.netSalary, te.expected.netSalary)
		}
	}
}

func TestGrossForNetSalary(t *testing.T) {
	type TestCase struct {
		netSalary float32
		insurance float32
		kids      int
		expected  float32
	}

	testCases := []TestCase{
		TestCase{
			netSalary: 10000,
			insurance: 0,
			kids:      0,
			expected:  10157,
		},
		TestCase{
			netSalary: 20000,
			insurance: 0,
			kids:      0,
			expected:  24090,
		},
		TestCase{
			netSalary: 30000,
			insurance: 0,
			kids:      0,
			expected:  41308,
		},
		TestCase{
			netSalary: 50000,
			insurance: 0,
			kids:      0,
			expected:  83511,
		},
		TestCase{
			netSalary: 100000,
			insurance: 0,
			kids:      0,
			expected:  189894,
		},
	}
	for _, te := range testCases {
		res := GrossForNetSalary(te.netSalary, te.insurance, te.kids, 1, 1)
		if !isClose(te.expected, res) {
			t.Errorf("GrossForNetSalary() = %f, expected = %f", res, te.expected)
		}
	}
}
