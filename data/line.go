package data

import (
	"../helpers"
	"../request"
	"strings"
)

var(
	SEARCH_LINK_PATTERN		string	=	`(http)|(https)://\w+\.\w{2,}`;
)

// Данные - строка
type Line struct {
	id 					int				;
	cells				[]string		;
	link				request.Url		;
	fast				[]request.Url	;
}

func ( l *Line ) SetId( id int ) {
	l.id = id;
}

func ( l Line ) GetId() int {
	return l.id;
}


func ( l *Line ) SetCells( cells []string ) {
	l.cells = cells;
}

func ( l Line ) GetCells() []string {
	return l.cells;
}

func ( l *Line ) SetLink( link string ) {
	url := request.Url{};
	url.SetUrl( link );
	l.link = url;
}

func ( l Line ) GetLink() request.Url {
	return l.link;
}

func ( l *Line ) SetFast( link string ) {
	url := request.Url{};
	url.SetUrl( link );
	l.fast = append( l.fast, url);
}

func ( l Line ) GetFast() []request.Url {
	return l.fast;
}

// Разбор строки с ссылкой
func ( l *Line ) prepareLink ( link string, isLink bool )  {
	url := strings.Split( link , "|" );
	if ( len( url ) < 1 ) {
		return;
	}

	if ( isLink ) {
		l.SetLink( url[0] );
	} else {
		l.SetFast( url[0] );
	}
}

// Разбор строки с ссылками
// TODO Добавить в параметры конфигурации номер ячейки для ссылки и быстрых ссылок
//      Позволит не искать ссылки в ячейке csv.
func ( l *Line ) prepareLinks ( key int, cell string )  {
	if ( helpers.Matched( SEARCH_LINK_PATTERN, cell ) ){//&& l.GetId() == 40 ) {
		links := strings.Split( cell , "||" );
		isLink := ( ( len( links ) == 1 ) && ( l.GetLink().GetUrl() == "" ) );
		for _, link := range links {
			l.prepareLink( link, isLink );
		}
	}
}

// Подготовка данных строки для дальнейшей обработки
func ( l *Line ) Prepare() {
	for key, cell := range l.GetCells() {
		l.prepareLinks( key, cell );
	}
}