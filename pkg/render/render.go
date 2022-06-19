package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/fastscripts/bookings/pkg/config"
	"github.com/fastscripts/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTEmplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}
func RenderTemplate(w http.ResponseWriter, tmpl string, templeteData *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		// get the template cache from the app config
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[tmpl]

	if !ok {
		fmt.Println("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	templeteData = AddDefaultData(templeteData)

	_ = template.Execute(buf, templeteData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}

}

// Crate Template Cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
