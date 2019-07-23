package main

import (
	"./configuration"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
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

func task(t chan int)  {

	rand.Seed(time.Now().UnixNano());
	min := 100;
	max := 500;
	rnd := rand.Intn(max - min) + min;

	duration := time.Duration( rnd ) * time.Millisecond;
	time.Sleep(duration);

	t<- 1
	close(t);
}

func worker(c chan int) {
	var worker = true;

	//t := make(chan int);

	var i int = 0;
	for ( worker ){

		if ( i > 10 ){
			worker = false;
		}

		//task(t);

		c <- i * i;

		i++;
	}
/*
	for {
		val, ok := <-t
		if ok == false {
			fmt.Println(val, ok, "<-- W loop close!")
			break;
		} else {
			fmt.Println(val, ok)
		}
	}

*/	close(c);
}

func main()  {

	defer loggingClose();

	logging("======= START Site Response Checker =======");

	c := make(chan int)

	go worker(c) // start goroutine

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop close!")
			break;
		} else {
			fmt.Println(val, ok)
		}
	}

	logging("======= STOP  Site Response Checker =======");
}
