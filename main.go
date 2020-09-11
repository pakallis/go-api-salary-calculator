package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pakallis/go-api-salary-calculator/taxcalculator"
)

type SalaryRequest struct {
	salary    float32
	insurance float32
	kids      int
}

func main() {
	r := gin.Default()
	r.GET("/salary-calculation", func(c *gin.Context) {
		salary, err := strconv.ParseFloat(c.Query("salary"), 32)
		insurance, err := strconv.ParseFloat(c.Query("insurance"), 32)
		kids, err := strconv.ParseInt(c.Query("kids"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return

		}
		res := taxcalculator.GrossForNetSalary(float32(salary), float32(insurance), int(kids), 1, 1)
		c.JSON(http.StatusOK, gin.H{
			"salary": res,
		})
	})
	r.Run()
}
