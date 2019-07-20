package main

import (
	"./configuration"
	"log"
	"os"
)

var (
	BASE_PATH		string									;
	PATH_SEPARATOR	string 		= string( os.PathSeparator );
	LOG_FILE_EXT	string		= "log";
	LOG_FILE		string		= "application";
)


func init(){

	var configuration configuration.ConfigurationInterface  = configuration.Get( "json" );
	BASE_PATH 		= configuration.GetPreferences().GetBasePath();


}

func logFilePath() string {
	return BASE_PATH + PATH_SEPARATOR + LOG_FILE + "." + LOG_FILE_EXT;
}

func main()  {
	logFile, err := os.OpenFile(logFilePath() , os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666);
	if err != nil {
		log.Fatalf("Error opening log file: %v", err);
	}
	defer logFile.Close();
	log.SetOutput(logFile);

	log.Println("======= START Site Response Checker =======");

	log.Println("======= STOP  Site Response Checker =======");
}
