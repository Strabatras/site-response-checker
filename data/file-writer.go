package data;

import (
	"encoding/csv"
	"log"
	"sync"
)

type FileWriter struct {
	mx   sync.Mutex;
	writer csv.Writer
}

func (fw *FileWriter) SetWriter( writer *csv.Writer ){
	fw.writer = *writer;
}

func (fw *FileWriter) WriteLine( line []string) {
	fw.mx.Lock();
	defer fw.mx.Unlock();
	err := fw.writer.Write(line);
	if err != nil {
		log.Fatal(err);
	}
}
