//Initial inspiration taken from Ivy (https://github.com/jameycribbs/ivy)
package main

import (
	"fmt"
	"log"
)

var sc Schema

func init() {

	fmt.Println("----- Starting ManticoreDB -----")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Loading Configuration ...")
	var err error
	Conf, err = GetConfig("./config.yml")
	if err != nil {
		log.Panicf("Error Loading Config: %s", err)
	}
	fmt.Printf("Database Count: %d\n", len(Conf.Databases))
	for i := 0; i < len(Conf.Databases); i++ {
		err = sc.Makedb(Conf.Databases[i])
		if err != nil {
			log.Panicf("Error Loading Database: %s", err)
		}
		fmt.Printf("%s Table Count: %d\n", sc.Databases[i].Name, len(sc.Databases[i].Tables))
		err = sc.Databases[i].LoadLog()
		if err != nil {
			//Being unable to load the transaction log doesn't lead to a panic
			log.Printf("Error Loading Transaction Log: %s", err)
		}
	}
	fmt.Println("----- Running ManticoreDB -----")
}

func main() {

	err := sc.Databases[0].AddRow("people", []string{"Jane", "Doe", "Toronto"}, 0)
	if err != nil {
		fmt.Println(err)
	}

	err = sc.Databases[0].DelRow("people", 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sc.Databases[0].TLog)

	err = terminateSession()
	if err != nil {
		fmt.Printf("Error Terminating Session: %s", err)
	}
}

//Properly shuts down the system, storing all unsaved changes first
func terminateSession() error {
	Conf.UpdateConfigDBs(sc.DBNames())
	err := Conf.SaveConfig("./config.yml")
	if err != nil {
		return err
	}
	for _, db := range sc.Databases {
		err = db.SaveLog()
		if err != nil {
			return err
		}
	}
	return nil
}
