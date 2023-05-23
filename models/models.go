package models

type ConfigFile struct {
	Server struct {
		URLCDMX   string `yaml:"urlCDMX"`
		CLASSNAME string `yaml:"className"`
	}
	DataBase struct {
		IP   string `yaml:"ip"`
		PORT string `yaml:"port"`
		DB   string `yaml:"db"`
	}
}

type DatosWifi struct {
	Id       string
	Latitud  string
	Longitud string
	Colonia  string
	Alcaldia string
}
