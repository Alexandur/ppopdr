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
    "appengine"
    "appengine/datastore"
    "http"
    "template"
	"webservice/models/data"
)

var logTemplate = template.Must(template.ParseFile("webservice/views/log.html"))

type LogController struct{}

/*
 * Ziet er raar uit, maar dit is "de" manier in appengine
 * Bron: stackoverflow.com/questions/1062540/how-to-delete-all-datastore-in-google-app-engine (python)
 * The best approach is the remote API method as suggested by Nick, he's an App Engine engineer from Google
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-21
 */
func (cr *LogController) DeleteMany(cx *goweb.Context) {
	c := appengine.NewContext(cx.GetRequest())
	pageviews := make([]data.PageView, 0)
	q := datastore.NewQuery("PageView")
	
	keys, err := q.GetAll(c, &pageviews);
	
	if err != nil {
    	cx.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
        return
    }
    
	datastore.DeleteMulti(c, keys)
}

/**
 * Haal laatste 50 logs op en toon deze met log.html
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-21
 */
func (cr *LogController) ReadMany(cx *goweb.Context) {
	c := appengine.NewContext(cx.GetRequest())
    q := datastore.NewQuery("PageView").Order("-Date").Limit(50)
    pageviews := make([]data.PageView, 0, 50)
    
    if _, err := q.GetAll(c, &pageviews); err != nil {
    	cx.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
        return
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








