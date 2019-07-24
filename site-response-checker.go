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


func worker( jobs chan int, done chan bool ) {

	for {
		j, more := <-jobs
		if more {
/*			rand.Seed(time.Now().UnixNano())
			min := 100
			max := 500
			rnd := rand.Intn(max-min) + min

			duration := time.Duration(rnd) * time.Millisecond
			time.Sleep(duration)*/
			fmt.Println("received job => ", j )
		} else {
			fmt.Println("received all jobs => " )
			done <- true
			return
		}
	}

}

func main()  {

	defer loggingClose();

	logging("======= START Site Response Checker =======");

	jobs := make(chan int, 5)
	done := make(chan bool)

	for k := 0; k < 5; k++ {
		go worker(jobs, done )
	}

	// читаем файл строки посылаем в канал
	for j := 1; j <= 30; j++ {
		jobs <- j
		fmt.Println("sent line", j)
	}
	close(jobs)
	fmt.Println("закрыли канал jobs")

	<-done

	logging("======= STOP  Site Response Checker =======");
}
