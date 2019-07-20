package main

import (
	"./configuration"
	"fmt"
	"log"
	"os"
)

var (
	CONFIGURATION	configuration.ConfigurationInterface	;
	PATH_SEPARATOR	string 									= string( os.PathSeparator );
	LOG_FILE		*os.File								;
)


func init(){

	CONFIGURATION  = configuration.Get( "json" );

	if ( CONFIGURATION.GetPreferences().IsProduction() ) {
		loggingInit();
	}

}

// Инициализация логирования
func loggingInit() {
	LOG_FILE, err := os.OpenFile(  preferences().GetBasePath() + PATH_SEPARATOR + "application.log" , os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666);
	if err != nil {
		log.Fatalf("Error opening log file: %v", err);
	}
	log.SetOutput( LOG_FILE );
}

// Закрываем ссылку на файл логирования
func loggingClose() {
	if LOG_FILE != nil {
		LOG_FILE.Close();
	}
}

// Логирование
func logging( message string ){

	var preferences = preferences();
	if ( preferences.IsProduction() ) {
		log.Println( message );
	}
	if ( preferences.IsDevelopment() ){
		fmt.Println( message );
	}
}


func preferences() configuration.ConfigurationPreferencesInterface {
	return CONFIGURATION.GetPreferences();
}

func main()  {

	defer loggingClose();

	logging("======= START Site Response Checker =======");

	logging("======= STOP  Site Response Checker =======");
}
