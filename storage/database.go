package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//Struct to interpret CSV tables to
type Table struct {
	Headers []string
	Entries [][]string
}

//Print the table in a table format
func (t Table) String() string {
	var temp strings.Builder
	ncols := len(t.Headers)
	nrows := len(t.Entries)
	//Entry Spacing code
	//find highest numspaces for each col
	numspaces := make(map[string]int) //This method may not be the most resource conscious
	//get baseline col numspaces
	for i := 0; i < ncols; i++ {
		numspaces[t.Headers[i]] = len(t.Headers[i])
	}
	//get highest numspaces per col
	for i := 0; i < ncols; i++ {
		for j := 0; j < nrows; j++ {
			if len(t.Entries[j][i]) > numspaces[t.Headers[i]] {
				numspaces[t.Headers[i]] = len(t.Entries[j][i])
			}
		}
	}

	//Display headers
	for i := 0; i < ncols; i++ {
		temp.WriteString("|")
		for n := 0; n <= (numspaces[t.Headers[i]]-len(t.Headers[i]))/2; n++ {
			temp.WriteString(" ")
		}
		temp.WriteString(fmt.Sprintf("%s", t.Headers[i]))
		for n := 0; n <= (numspaces[t.Headers[i]]-len(t.Headers[i]))/2; n++ {
			temp.WriteString(" ")
		}
		temp.WriteString("|")
	}
	temp.WriteString("\n")
	//Display entries
	for j := 0; j < nrows; j++ {
		for i := 0; i < ncols; i++ {
			temp.WriteString("|")
			for n := 0; n <= (numspaces[t.Headers[i]]-len(t.Entries[j][i]))/2; n++ {
				temp.WriteString(" ")
			}
			temp.WriteString(fmt.Sprintf("%s", t.Entries[j][i]))
			for n := 0; n <= (numspaces[t.Headers[i]]-len(t.Entries[j][i]))/2; n++ {
				temp.WriteString(" ")
			}
			//for oddnumbers add an extra trailing space
			if (numspaces[t.Headers[i]]-len(t.Entries[j][i]))%2 != 0 {
				temp.WriteString(" ")
			}
			temp.WriteString("|")
		}
		temp.WriteString("\n")
	}
	return temp.String()
}

//Struct representing how the data is held
type DB struct {
	Name   string   //the name of the table is also the directory where entries will be stored
	Tables []string //file names of the tables
}

//Delete the nth row
func (database *DB) DelRow(name string, n int) error {
	table, err := database.ReadTable(name)
	if err != nil {
		return err
	}
	table.Entries = append(table.Entries[:n], table.Entries[n+1:]...)

	path := fmt.Sprintf("%s/%s.csv", database.Name, name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	temp := [][]string{}
	temp = append(temp, table.Headers)
	temp = append(temp, table.Entries...)
	err = w.WriteAll(temp)
	if err != nil {
		return err
	}

	return nil
}

//Add a row to the nth position
func (database *DB) AddRow(name string, entry []string, n int) error {
	//Ensure table ncols == entry ncols
	table, err := database.ReadTable(name)
	if err != nil {
		return err
	}
	if len(table.Headers) != len(entry) {
		return errors.New("new entry has different number of cols")
	}

	//make table object
	temp := make([][]string, len(table.Entries[n:]))
	copy(temp, table.Entries[n:])
	table.Entries = append(table.Entries[:n], entry)
	table.Entries = append(table.Entries, temp...)

	//write data to file
	path := fmt.Sprintf("%s/%s.csv", database.Name, name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	temp = [][]string{}
	temp = append(temp, table.Headers)
	temp = append(temp, table.Entries...)
	err = w.WriteAll(temp)
	if err != nil {
		return err
	}

	return nil
}

//Adds a table (.csv File) to a given database
func (database *DB) MakeTable(name string, headers []string) error {
	//return error if table already exists
	if strings.Contains(strings.Join(database.Tables, " "), name) {
		return errors.New(fmt.Sprintf("table %s already exists", name))
	}
	//make table
	path := fmt.Sprintf("%s/%s.csv", database.Name, name)
	temp := strings.Join(headers, ",")
	err := ioutil.WriteFile(path, []byte(temp), os.ModePerm)
	if err != nil {
		return err
	}
	database.Tables = append(database.Tables, name)
	return nil
}

//Deletes a table from the DB
func (database *DB) DelTable(name string) error {
	//Strings package instead of a loop to try and keep it fast on large DB's
	tables := strings.Join(database.Tables, " ")
	//Check if table exists
	if !strings.Contains(tables, name) {
		return errors.New(fmt.Sprintf("table %s does not exist", name))
	}
	tables = strings.Replace(tables, " "+name, "", 1)
	database.Tables = strings.Split(tables, " ")
	err := os.Remove(fmt.Sprintf("%s/%s.csv", database.Name, name))
	if err != nil {
		return err
	}
	return nil
}

//Reads in the table / csv
func (database *DB) ReadTable(name string) (Table, error) {
	path := fmt.Sprintf("%s/%s.csv", database.Name, name)
	file, err := os.Open(path)
	if err != nil {
		return Table{}, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	var temp Table
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return Table{}, err
		} else if temp.Headers == nil {
			temp.Headers = record
		} else {
			temp.Entries = append(temp.Entries, record)
		}
	}
	return temp, nil
}
