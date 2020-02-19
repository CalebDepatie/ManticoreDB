//Test DBMS setup before go_test files are created

//Initial inspiration taken from Ivy (https://github.com/jameycribbs/ivy)
package main

import (
	"ManticoreDB/misc"
	"ManticoreDB/storage"
	"fmt"
	"log"
)

var sc storage.Schema

func init() {

	fmt.Println("----- Starting ManticoreDB -----")
	fmt.Printf("Version: %s\n", misc.GetVersion())
	fmt.Println("Loading Configuration ...")
	var err error
	misc.Conf, err = misc.GetConfig("./config.yml")
	if err != nil {
		log.Panicf("Error Loading Config: %s", err)
	}
	fmt.Printf("Database Count: %d\n", len(misc.Conf.Databases))
	for i := 0; i < len(misc.Conf.Databases); i++ {
		err = sc.Makedb(misc.Conf.Databases[i])
		if err != nil {
			log.Panicf("Error Loading Database: %s", err)
		}
		fmt.Printf("%s Table Count: %d\n", sc.Databases[i].Name, len(sc.Databases[i].Tables))
	}
	fmt.Println("----- Running ManticoreDB -----")
}

func main() {
	//initial table
	temp, err := sc.Databases[0].ReadTable("people")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)

	//delete row 2
	sc.Databases[0].DelRow("people", 1)
	temp, err = sc.Databases[0].ReadTable("people")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)

	//Add row 2 back
	sc.Databases[0].AddRow("people", []string{"Jack", "Daniels", "Tennessee"}, 1)
	temp, err = sc.Databases[0].ReadTable("people")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)

	err = terminateSession()
	if err != nil {
		fmt.Printf("Error Terminating Session: %s", err)
	}
}

//Properly shuts down the system, storing all unsaved changes first
func terminateSession() error {
	misc.Conf.UpdateConfigDBs(sc.DBNames())
	err := misc.Conf.SaveConfig("./config.yml")
	if err != nil {
		return err
	}
	return nil
}
