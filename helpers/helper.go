package helpers

import "regexp"

// Поиск совпадений
func Matched( pattern string, text string ) bool {
	matched, _ := regexp.Match( pattern, []byte( text ) );
	if matched {
		return true;
	}
	return false;
}
