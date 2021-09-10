package config

import "html/template"

//All application configurations are held here in AppConfig
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
