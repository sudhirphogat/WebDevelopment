package config

import (
	"html/template"
	"log"
)

//All application configurations are held here in AppConfig
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
}
