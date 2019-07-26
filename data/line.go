package data

import (
	"../helpers"
	//"fmt"
	"../request"
	"fmt"
)
// Данные - строка
type Line struct {
	id 					int				;
	cells				[]string		;
	urltMain			request.Url		;
	urlFast				[]request.Url	;
}

func ( l *Line ) SetId( id int ) {
	l.id = id;
}

func ( l Line) GetId() int {
	return l.id;
}


func ( l *Line ) SetCells( cells []string ) {
	l.cells = cells;
}

func ( l Line ) GetCells() []string {
	return l.cells;
}

// Подготовка данных строки для дальнейшей обработки
func ( l Line ) Prepare() {
	pattern := `(http)|(https)://\w+\.\w{2,}`
	for _, cell := range l.GetCells() {
		if ( helpers.Matched( pattern, cell ) ) {
			fmt.Println( cell , l.GetId());
			fmt.Println( "" );
		}

		fmt.Println( l.GetId() );
	}
}