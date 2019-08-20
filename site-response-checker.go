package main;

import (
	"./data"
	"./helpers"
	"./interfaces"
	"./request"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
);

var (
	WORKER_MAX                 = 5;
	PATH_SEPARATOR      string = string(os.PathSeparator);
	SEARCH_LINK_PATTERN string = `(http)|(https)://\w+\.\w{2,}`;
);

func csvFile() string {
	return "/Users/dmd/Documents/temp/site-response-checker" + PATH_SEPARATOR + "input" + PATH_SEPARATOR + "41423731.csv";
}

// Разбор строки с ссылками
// TODO Добавить в параметры конфигурации номер ячейки для ссылки и быстрых ссылок
//      Позволит не искать ссылки в ячейке csv.
func prepareLineRequests(key int, cell string, line interfaces.Line) {
	if (helpers.Matched(SEARCH_LINK_PATTERN, cell)) {
		links := strings.Split(cell, "||");
		for _, link := range links {
			split := strings.Split(link, "|");
			if (len(split) > 0) {
				var request interfaces.Request = &request.Request{};
				request.SetUrl(split[0]);
				request.SetHash(helpers.HashSHA1(request.GetUrl()));
				line.GetRequestList().SetRequest(request);
			}
		}

	}
}

// Подготовка данных строки для дальнейшей обработки
func prepareLine(line interfaces.Line) {
	for key, cell := range line.GetCells() {
		prepareLineRequests(key, cell, line);
	}
}

func worker(lines chan interfaces.Line, inProgress interfaces.InProgress, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for {
		line, more := <-lines
		if more {
			// 1) разобрать строку
			prepareLine(line);
			for _, relations := range line.GetRequestList().GetRelations() {
				//fmt.Println( "KEY =>", key , " relations => " , relations , " request => " , line.GetRequestList().GetRequest(relations[0]))
				if ( inProgress.ToObservation( line.GetRequestList().GetRequest(relations[0]), line ) == false ){
					fmt.Println( "inProgress.ToObservation( request, line ) == false " );
					fmt.Println( "Send GET request" );
				}
			}
			fmt.Println("=================")
			/*
			for _, request := range line.GetRequestList().GetRequests() {
				if ( inProgress.ToObservation( request, line ) == false ){
					fmt.Println( "inProgress.ToObservation( request, line ) == false " );
					fmt.Println( "Send GET request" );
				}

			}
			*/
			//if (len(line.GetRequestList().GetRequests()) > 0) {
			//		fmt.Println("line.GetRequestList()", line.GetRequestList().GetRequests())
			//}
			/*			if ( line.GetLink() != nil ) {
							//sendRequest( line.GetLink(), line, inProgress );
						}

						if ( line.GetFast() != nil ) {
							for _, request := range line.GetFast(){
								sendRequest( request, line, inProgress );
							}
						}*/

			//fmt.Println("line.GetLink()", line.GetLink(), line.GetFast() )

			/*
				2) отправить запрос
				3) получить данные
				4) разобрать данные
				5) Записать данные
			*/

		} else {
			return;
		}
	}

}

func NewRequestList() interfaces.RequestList {
	var requestList interfaces.RequestList = &data.RequestList{};
	requestList.Init();
	return requestList;
}

func NewLine(id int, cells []string) interfaces.Line {

	var line interfaces.Line = &data.Line{};
	line.SetId(id);
	line.SetCells(cells);
	line.SetRequestList(NewRequestList());
	return line;
}

func NewInProgress() interfaces.InProgress  {
	var inProgress interfaces.InProgress = &request.InProgress{};
	var checked interfaces.CheckedList = &data.CheckedList{};
	var observation interfaces.Observation = &data.Observation{};

	checked.Init();
	observation.Init();

	inProgress.SetCheckedList( checked );
	inProgress.SetObservation( observation );
	return inProgress;
}

func main() {
	fmt.Println("======= START Site Response Checker =======");
	var waitGroup sync.WaitGroup
	waitGroup.Add(WORKER_MAX);
	lines := make(chan interfaces.Line, WORKER_MAX);

	inProgress := NewInProgress();

	for i := 0; i < WORKER_MAX; i++ {
		go worker(lines, inProgress, &waitGroup);
	}

	csvFile, err := os.Open(csvFile());
	if err != nil {
		// TODO обработать фатал
		log.Fatalln("Couldn't open the csv file", err);
	}
	defer csvFile.Close();
	r := csv.NewReader(csvFile);
	r.Comma = ';';
	r.Comment = '#';
	for j := 1; ; j++ {
		cells, err := r.Read()
		if err == io.EOF {
			break;
		}
		if err != nil {
			fmt.Println(err);
		}
		data := NewLine(j, cells);
		lines <- data;
	}

	close(lines);
	waitGroup.Wait();
	fmt.Println("======= STOP  Site Response Checker =======");
}
