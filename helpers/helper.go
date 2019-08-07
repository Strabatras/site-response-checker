package helpers

import (
	"crypto/sha1"
	"encoding/hex"
	"regexp"
)

// Поиск совпадений
func Matched( pattern string, text string ) bool {
	matched, _ := regexp.Match( pattern, []byte( text ) );
	if matched {
		return true;
	}
	return false;
}

// Хэш строки sha1
func Hash( str string ) string  {
	sha := sha1.New();
	sha.Write( []byte( str ) );
	return hex.EncodeToString( sha.Sum( nil ) );
}