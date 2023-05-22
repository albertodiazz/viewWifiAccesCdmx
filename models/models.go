package models

type ConfigFile struct {
	Server struct {
		URLCDMX   string `yaml:"urlCDMX"`
		CLASSNAME string `yaml:"className"`
	}
}
