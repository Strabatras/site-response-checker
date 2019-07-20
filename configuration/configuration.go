package configuration

// Конфигурация приложения
type ConfigurationInterface interface {
	GetPreferences()	ConfigurationPreferencesInterface;
}

// Конфигурация приложения секция общих настроек
type ConfigurationPreferencesInterface interface {
	GetBasePath() 		string;
	IsProduction()		bool;
	IsDevelopment()		bool;
	IsTesting()			bool;

}

func Get( configurationType string ) ConfigurationInterface {

	var configuration ConfigurationInterface;
	if configurationType == "json" {
		configuration = Json();
	}
	return configuration;
}
