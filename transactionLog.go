package main

import (
	"fmt"
	"strings"
	"time"
)

type transaction struct {
	ID        int
	Type      string
	Time      time.Time
	TableName string
	Data      []interface{}
}

// Transaction constructor
func CreateTransaction(ID int, Type string, TableName string, Data ...interface{}) transaction {
	var t transaction
	t.ID = ID
	t.Type = Type
	t.Time = time.Now()
	t.TableName = TableName
	t.Data = Data

	return t
}

func (t transaction) String() string {
	var temp strings.Builder
	temp.WriteString(fmt.Sprintf("ID: %d\n", t.ID))
	temp.WriteString(fmt.Sprintf("Type: %s\n", t.Type))
	temp.WriteString(fmt.Sprintf("Time: %s\n", t.Time.String()))
	temp.WriteString(fmt.Sprintf("Table: %s\n", t.TableName))
	temp.WriteString(fmt.Sprintf("Data: %s\n", t.Data))
	return temp.String()
}
