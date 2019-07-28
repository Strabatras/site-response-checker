package main

import (
	"./configuration"
	"./data"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var (
	CONFIGURATION	configuration.ConfigurationInterface	;
	PATH_SEPARATOR	string 									= string( os.PathSeparator );
	LOG_FILE		*os.File								;
	WORKER_MAX		int										= 5;
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

// Общие настройки
func preferences() configuration.ConfigurationPreferencesInterface {
	return CONFIGURATION.GetPreferences();
}

func worker( lines chan data.Line, waitGroup *sync.WaitGroup ) {
	defer waitGroup.Done()
	for {
		line, more := <-lines
		if more {
			// 1) разобрать строку
			line.Prepare();
			/*

				2) отправить запрос
				3) получить данные
				4) разобрать данные
				5) Записать данные
			 */

			//prepareDataLine( &line );
			//fmt.Println( line.GetFast() );
			//fmt.Println("received job => ", line.Id );

		} else {
			return;
		}
	}

}

func errorToLogging(  error error ) string {
	return error.Error();
}

func main()  {

	defer loggingClose();

	logging("======= START Site Response Checker =======");

	var waitGroup sync.WaitGroup
	waitGroup.Add( WORKER_MAX );

	lines := make( chan data.Line, WORKER_MAX );

	for i := 0; i < WORKER_MAX; i++ {
		go worker( lines, &waitGroup );
	}

	csvFile, err := os.Open(preferences().GetBasePath() + PATH_SEPARATOR + "input" + PATH_SEPARATOR + "41423731.csv" );
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvFile.Close();

	r := csv.NewReader( csvFile );
	r.Comma = ';';
	r.Comment = '#';
	for j :=1; ; j++ {
		cells, err := r.Read()
		if err == io.EOF {
			break;
		}
		if err != nil {
			logging( errorToLogging( err ) );
		}
		data := data.Line{};
		data.SetId( j );
		data.SetCells( cells );
		lines <- data;
	}

	close( lines );
	waitGroup.Wait();

	logging("======= STOP  Site Response Checker =======");
}
