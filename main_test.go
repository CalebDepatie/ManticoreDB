package main

import (
	"testing"
)

//test the DB package
func Test(t *testing.T) {
	var err error
	Conf, err = GetConfig("config.yml")
	if err != nil {
		t.Error(err)
	}

	var sc Schema

	err = sc.Makedb("bar")
	if err != nil {
		t.Error(err)
	}

	err = sc.Databases[0].MakeTable("foo", []string{"name", "job"})
	if err != nil {
		t.Error(err)
	}
	err = sc.Databases[0].AddRow("foo", []string{"John", "Software Dev"}, 0)
	if err != nil {
		t.Error(err)
	}
	err = sc.Databases[0].DelRow("foo", 0)
	if err != nil {
		t.Error(err)
	}
	err = sc.Databases[0].DelTable("foo")
	if err != nil {
		t.Error(err)
	}

	err = sc.Deldb("bar")
	if err != nil {
		t.Error(err)
	}
}
