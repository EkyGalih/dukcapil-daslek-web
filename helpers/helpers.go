package helpers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(c *gin.Context, tmpl string, data interface{}) error {
	t := template.New("").Funcs(template.FuncMap{})

	// Parsing partial templates dari folder partials
	partialTemplates := filepath.Join("templates", "layouts", "*.html")
	t, err := t.ParseGlob(partialTemplates)
	if err != nil {
		log.Printf("Error parsing partial templates: %v", err)
		c.String(http.StatusInternalServerError, "Template error: %v", err)
		return err
	}

	// Parsing main template (file .html) di folder yang sesuai
	tmplPath := filepath.Join("templates", tmpl)
	t, err = t.ParseFiles(tmplPath, filepath.Join("templates", "layouts", "app.html"))
	if err != nil {
		log.Printf("Error parsing main template: %v", err)
		c.String(http.StatusInternalServerError, "Template error: %v", err)
		return err
	}

	// render the template
	err = t.ExecuteTemplate(c.Writer, "app.html", data)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		c.String(http.StatusInternalServerError, "Template error: %v", err)
		return err
	}

	return nil
}
