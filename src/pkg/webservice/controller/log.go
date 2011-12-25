/**
 * Implementeer hier elke REST methode (+restMany methode) voor log
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-20
 */
package controller

import (
	"goweb"
    "http"
    "template"
	"webservice/models/data"
)

var logTemplate = template.Must(template.ParseFile("webservice/views/log.html"))

type LogController struct{}

/*
 * Verwijder alle log entries
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-21
 */
func (cr *LogController) DeleteMany(cx *goweb.Context) {
	err := data.DeleteAllLogs(cx.GetRequest())
	
	if err != nil {
    	cx.RespondWithErrorMessage("Verwijderen is niet gelukt", http.StatusInternalServerError)
    }
}

/**
 * Haal laatste 50 logs op en toon deze met log.html
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-21
 */
func (cr *LogController) ReadMany(cx *goweb.Context) {
	err, pageviews := data.GetLogs(cx.GetRequest())
	
    if err != nil {
    	cx.RespondWithErrorMessage("Kan de lijst met logs entries niet tonen", http.StatusInternalServerError)
    }
     
    if err := logTemplate.Execute(cx.GetResponseWriter(), pageviews); err != nil {
        cx.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
    }
}

/**
 * Onderstaande functies worden nog niet gebruikt. Nog niet nodig gehad.
 */
func (cr *LogController) Create(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *LogController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *LogController) Read(input string, cx *goweb.Context) {
	if input == "bbb" {
		var bbbTemplate = template.Must(template.ParseFile("webservice/views/bbb.html"))
	
		if err := bbbTemplate.Execute(cx.GetResponseWriter(), nil); err != nil {
	        cx.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
	    }
	} 
}

func (cr *LogController) Update(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *LogController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}








