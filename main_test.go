package main

import (
	"ManticoreDB/misc"
	"ManticoreDB/storage"
	"testing"
)

//test the DB package
func Test(t *testing.T) {
	var err error
	misc.Conf, err = misc.GetConfig("config.yml")
	if err != nil {
		t.Error(err)
	}

	vers := misc.GetVersion()
	if vers != "0.2" {
		t.Error("Incorrect version returned")
	}

	var sc storage.Schema

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
