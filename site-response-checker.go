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
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
);

var (
	WORKER_MAX          int    = 5;
	ATTEMPT_REQUEST_MAX int    = 3;
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
			if ( (len(split) > 0) && (split[0] != "") ) {
				var request interfaces.Request = &request.Request{};
				request.SetUrl( split[0] );
				request.SetHash(helpers.HashSHA1(request.GetUrl()));
				line.GetRequestList().SetRequest(request);
				line.GetRequestList().IncrementInWork();
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

func send(request interfaces.Request, netTransport http.RoundTripper, attempt int) (*http.Response, bool) {

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	response, ok := netClient.Get(request.GetUrl());
	defer response.Body.Close();

	if ok == nil {
		return response, true;
	}

	if ( ATTEMPT_REQUEST_MAX > attempt ) {
		fmt.Println( "attempt > ", attempt );
		return send(request, netTransport, attempt+1);
	}
	return response, false;
}

func sendRequest(request interfaces.Request) {
	response, ok := send(request, &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}, 0);
	request.SetFinished();
	if ok {
		request.SetStatusCode(response.StatusCode);
	}
}

func worker(lines chan interfaces.Line, inProgress interfaces.InProgress, waitGroupWorker *sync.WaitGroup, writer chan interfaces.Line, waitGroupWriter *sync.WaitGroup) {
	defer waitGroupWorker.Done();
	for {
		line, more := <-lines
		if more {
			prepareLine(line);
			// проверяем себя и если есть подобные в строке
			for _, relations := range line.GetRequestList().GetRelations() {
				request := line.GetRequestList().GetRequest(relations[0]);
				// если запрос не выполнялся ранее
				if (inProgress.ToObservation(request, line) == false) {
					sendRequest(request);
					inProgress.FromObservation(request, writer, waitGroupWriter);
					return;
				}
				if ( line.GetRequestList().GetInWork() == 0 ) {
					fmt.Println( "worker() => ToWriteLine" );
				}
			}
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

func NewInProgress() interfaces.InProgress {
	var inProgress interfaces.InProgress = &request.InProgress{};
	var checked interfaces.CheckedList = &data.CheckedList{};
	var observation interfaces.Observation = &data.Observation{};
	checked.Init();
	observation.Init();
	inProgress.SetCheckedList(checked);
	inProgress.SetObservation(observation);
	return inProgress;
}

func main() {
	fmt.Println("======= START Site Response Checker =======");
	var waitGroupWorker sync.WaitGroup
	waitGroupWorker.Add(WORKER_MAX);
	lines := make(chan interfaces.Line, WORKER_MAX);

	var waitGroupWriter  sync.WaitGroup;
	writer := make(chan interfaces.Line);

	inProgress := NewInProgress();
	for i := 0; i < WORKER_MAX; i++ {
		go worker(lines, inProgress, &waitGroupWorker, writer, &waitGroupWriter);
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
		lines <- NewLine(j, cells);
	}

	close(lines);
	waitGroupWorker.Wait();

	close(writer);
	waitGroupWriter.Wait();

	fmt.Println("======= STOP  Site Response Checker =======");
}
