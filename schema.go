package main

import (
	"io/ioutil"
	"os"
	"strings"
)

//Struct containing pointers to DB's
type Schema struct {
	Databases []*DB
}

//Gets an array of the DB names
func (s Schema) DBNames() []string {
	var names []string
	for i := 0; i < len(s.Databases); i++ {
		names = append(names, s.Databases[i].Name)
	}
	return names
}

//Connects to the database entries and adds it to the Schema doesnt add it to the conf
func (s *Schema) Makedb(name string) error {
	var temp DB
	temp.Name = name
	MakeDir(name) //attempts to make the DB directory if it needs to
	//get the table names
	files, err := ioutil.ReadDir(name)
	if err != nil {
		return err
	}
	if len(files) != 0 {
		for i := 0; i < len(files); i++ {
			if !strings.Contains(files[i].Name(), "tlog.gob") {
				temp.Tables = append(temp.Tables, files[i].Name())
			}
		}
	}
	s.Databases = append(s.Databases, &temp)
	return nil
}

//Deletes a database from the Schema
func (s *Schema) Deldb(name string) error {
	//remove DB entry
	for i := 0; i < len(s.Databases); i++ {
		if s.Databases[i].Name == name {
			s.Databases = append(s.Databases[:i], s.Databases[i+1:]...)
		}
	}
	//remove DB directory
	err := os.RemoveAll(name)
	if err != nil {
		return err
	}
	return nil
}
