package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func JsonToMap(data []byte) {
	var employee []Employee

	var employeeMap []map[string]interface{}

	err := json.Unmarshal([]byte(data), &employeeMap)

	if err != nil {
		panic(err)
	}

	for _, personData := range employeeMap {

		// convert map to array of Person struct
		var p Employee
		p.Name = fmt.Sprintf("%s", personData["Name"])
		p.Age, _ = strconv.Atoi(fmt.Sprintf("%v", personData["Age"]))
		p.Job = fmt.Sprintf("%s", personData["Job"])
		employee = append(employee, p)

	}
	fmt.Println(employee)
}

func main() {

	// read data from CSV file
	csvFile, err := os.Open("./data.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Employee
	var allRecords []Employee

	for _, each := range csvData {
		oneRecord.Name = each[0]
		oneRecord.Age, _ = strconv.Atoi(each[1]) // need to cast integer to string
		oneRecord.Job = each[2]
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//convert JSON to Map Data Structure and print it.
	JsonToMap(jsondata)

	// now write to JSON file
	jsonFile, err := os.Create("./data.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}
