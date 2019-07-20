// Разбор параметров конфигурации приложения в формате JSON
package configuration;

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Название файла конфигурации
var file_name = "configuration.json";

// Описание структуры файла конфигурации
type ConfigurationJson struct {
	Preferences		ConfigurationPreferencesJson    `json:"Preferences"`
}

// Описание структуры блока 'Preferences' файла конфигурации
type ConfigurationPreferencesJson struct {
	BasePath		string		`json:"base_path"`;
	Environment		string  	`json:"environment"`;
}

// Возвращает содержимое блока 'Preferences' файла конфигурации
func ( json ConfigurationJson ) GetPreferences() ConfigurationPreferencesInterface {
	return json.Preferences;
}

func ( json ConfigurationPreferencesJson ) GetBasePath() string {
	return json.BasePath;
}

//	Рабочая
func ( json ConfigurationPreferencesJson ) IsProduction() bool {
	return ( json.Environment == "production" );
}
// Разработка
func ( json ConfigurationPreferencesJson ) IsDevelopment() bool {
	return ( json.Environment == "development" );
}
// Тестирование
func ( json ConfigurationPreferencesJson ) IsTesting() bool {
	return ( json.Environment == "testing" );
}

func Json() ConfigurationInterface {
	jsonConfigurationFile, err := os.Open( file_name );

	if err != nil {
		fmt.Println( "Failed to open json configuration file: " + file_name + "." );
	}
	defer jsonConfigurationFile.Close();

	var configuration ConfigurationJson;

	jsonFileByteValue, _ := ioutil.ReadAll( jsonConfigurationFile );

	json.Unmarshal( jsonFileByteValue, &configuration );

	return configuration;

}
