package main

import (
	"./configuration"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
)

var (
	CONFIGURATION	configuration.ConfigurationInterface	;
	PATH_SEPARATOR	string 									= string( os.PathSeparator );
	LOG_FILE		*os.File								;
	WORKER_MAX		int										= 5;
)

// Результат запроса
type Request struct {
	Url 		string	;
	StatusCode	int		;
}

// Данные строки
type DataLine struct {
	Id 					int			;
	Line				[]string	;
	RequestMain			Request		;
	RequestAdditionals	[]Request	;
}

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

func matched( pattern string, text string ) bool {
	matched, _ := regexp.Match( pattern, []byte( text ) );
	if matched {
		return true;
	}
	return false;
}

func prepareDataLine ( dataLine *DataLine ) {
	if ( dataLine.Id == 40 ) {
		pattern := `(http)|(https)://\w+\.\w{2,}`
		for _, value := range dataLine.Line {
			if ( matched( pattern, value ) ) {
				fmt.Println( value );
				fmt.Println( "" );
			}
		}
	}
}

func worker( lines chan DataLine, waitGroup *sync.WaitGroup ) {
	defer waitGroup.Done()
	for {
		line, more := <-lines
		if more {
			/*
				1) разобрать строку
				2) отправить запрос
				3) получить данные
				4) разобрать данные
			 */

			prepareDataLine( &line );
			//fmt.Println("DataLine", line );
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

	lines := make( chan DataLine, WORKER_MAX );

	for i := 0; i < WORKER_MAX; i++ {
		go worker( lines, &waitGroup );
	}

	// читаем файл строки посылаем в канал
	//for j := 1; j <= 30; j++ {
	//	lines <- DataLine { Id: j };
	//	fmt.Println("sent line", j );
	//}

	csvfile, err := os.Open(preferences().GetBasePath() + PATH_SEPARATOR + "input" + PATH_SEPARATOR + "41423731.csv" );
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close();

	r := csv.NewReader( csvfile );
	r.Comma = ';';
	r.Comment = '#';
	for j :=1; ; j++ {
		records, err := r.Read()
		if err == io.EOF {
			break;
		}
		if err != nil {
			logging( errorToLogging( err ) );
		}
		lines <- DataLine{ Id : j, Line: records };
	}

		close( lines );
	fmt.Println("закрыли канал" )

	waitGroup.Wait();

	logging("======= STOP  Site Response Checker =======");
}
