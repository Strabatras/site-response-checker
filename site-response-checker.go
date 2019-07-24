package main

import (
	"./configuration"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
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
	Line				string		;
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


func preferences() configuration.ConfigurationPreferencesInterface {
	return CONFIGURATION.GetPreferences();
}


func worker( lines chan DataLine, waitGroup *sync.WaitGroup ) {
	defer waitGroup.Done()
	for {
		j, more := <-lines
		if more {
			rand.Seed(time.Now().UnixNano())
			min := 100
			max := 5000
			rnd := rand.Intn(max-min) + min

			duration := time.Duration(rnd) * time.Millisecond
			time.Sleep(duration)
			fmt.Println("received job => ", j )
		} else {
			fmt.Println("received all jobs => " )
			//done <- true
			return
		}
	}

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
	for j := 1; j <= 30; j++ {
		lines <- DataLine { Id: j };
		fmt.Println("sent line", j );
	}
	close( lines );
	fmt.Println("закрыли канал jobs" )

	waitGroup.Wait()

	logging("======= STOP  Site Response Checker =======");
}
