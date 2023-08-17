package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfic holds the applicatino config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
