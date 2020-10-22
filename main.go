package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// Employee is struct which contain employee details field
type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

func main() {

	csvFile, err := os.Open("sample.csv") // read mode only

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	reader, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, rec := range reader {
		emp.ID = rec[0]
		emp.EmployeeName = rec[1]
		emp.EmployeeSalary = rec[2]
		emp.EmployeeAge = rec[3]
		emp.ProfileImage = rec[4]

		employees = append(employees, emp)
	}

	// convert to json
	jsonData, err := json.MarshalIndent(employees, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// print json data
	fmt.Println(string(jsonData))

	// create json file
	jsonFile, err := os.Create("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
