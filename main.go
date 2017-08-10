package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tokoku_backend/config"
	"github.com/tokoku_backend/repository"
	"github.com/tokoku_backend/services"
)

var configPath = flag.String("c", "./config.json", "path to config file")
var migrationPath = flag.String("m", "./migrations/postgres", "path to migration directory")

func main() {
	flag.Parse()

	if err := config.LoadConfig(*configPath); err != nil {
		fmt.Printf("Error: %s loading configuration file: %s\n", *configPath, err.Error())
		os.Exit(1)
	}
	args := flag.Args()

	if len(args) > 0 {
		switch args[0] {
		case "migration":
			executeMigration(args)
		case "server":
			connectToDB()
			CreateRouter()
			//services.RetrieveAllWalmartItems()
			// aList, err := repository.Mgr.GetListOfWalmartCategory()
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// fmt.Println(aList)
		case "email":
			services.SendEmail()
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			os.Exit(1)
		}
	} else {
		flag.Usage()
	}
}

func connectToDB() {
	if err := repository.InitDBConnection(); err != nil {
		fmt.Println("error connecting to db: ", err.Error())
		os.Exit(1)
	}
}
