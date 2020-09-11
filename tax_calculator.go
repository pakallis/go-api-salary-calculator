package tax_calculator

import "math"

const infinity float32 = 100000000

var taxLevels = []float32{10000, 20000, 30000, 40000, infinity}
var taxPercentages = []float32{0.09, 0.22, 0.28, 0.36, 0.44}

var eisforaLevels = []float32{12000, 20000, 30000, 40000, 65000, 220000, infinity}
var eisforaPercentages = []float32{0, 0.022, 0.05, 0.065, 0.075, 0.09, 0.1}
var reductionPerKids = []float32{777, 810, 900, 1120, 1340}

const reductionPerExtraKid = 220

func calculatePercentageLevels(amount float32, levelTables []float32, percentageTables []float32) float32 {
	var total float32 = 0
	var previousLevel float32
	for i := 0; i < len(levelTables); i++ {
		currentLevel := levelTables[i]
		if i == 0 {
			previousLevel = 0
		} else {
			previousLevel = levelTables[i-1]
		}
		currentPercentage := percentageTables[i]
		if amount < currentLevel {
			total += (amount - previousLevel) * currentPercentage
			break
		} else {
			total += (currentLevel - previousLevel) * currentPercentage
		}
	}
	return total
}

func getReduction(salary float32, kids int) float32 {
	maxDefinedKidsInTable := len(reductionPerKids) - 1
	var reduction float32
	if kids > maxDefinedKidsInTable {
		// max reduction + 220 per extra kid
		reduction = reductionPerKids[maxDefinedKidsInTable] + (float32(kids)-float32(maxDefinedKidsInTable))*220.0
	} else {
		reduction = reductionPerKids[kids] // specific reduction
	}
	if salary > 12000 && kids < 5 { // no decrease in reduction for salaries under 12000 or more than 5 kids
		reduction -= ((salary - 12000) / 1000 * 20) // decrease reduction for 20 per 1000 extra in salary above 12000
	}
	return float32(math.Max(float64(reduction), 0))
}

type SalaryParts struct {
	taxableSalary float32
	initialTax    float32
	eisfora       float32
	reduction     float32
	totalTax      float32
	netSalary     float32
}

func getSalaryParts(salary float32, insurance float32, kids int) SalaryParts {
	taxableSalary := salary - insurance
	initialTax := calculatePercentageLevels(taxableSalary, taxLevels, taxPercentages)
	eisfora := calculatePercentageLevels(taxableSalary, eisforaLevels, eisforaPercentages)
	reduction := getReduction(taxableSalary, kids)
	totalTax := float32(math.Round(float64(initialTax-reduction+eisfora)*100.0) / 100.0)
	netSalary := salary - totalTax
	return SalaryParts{
		taxableSalary: taxableSalary,
		initialTax:    initialTax,
		eisfora:       eisfora,
		reduction:     reduction,
		totalTax:      totalTax,
		netSalary:     netSalary,
	}
}
