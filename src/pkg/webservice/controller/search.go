/**
 * Implementeer hier elke REST methode (+restMany methode) voor log
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2012-01-06
 */
package controller

import (
	"goweb"
    "http"
    "webservice/structs"
    "webservice/models"
)

type SearchController struct{}

/**
 * Zoek functie!
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2012-01-06
 */
func (cr *SearchController) Read(input string, cx *goweb.Context) {
	params := map[string]string {"search":input}
	var m structs.WordpressPostMany
	
	models.CallWordpressApi(cx, &m, "get_search_results", params)

	var articles []structs.Article

    for _, post := range m.Posts {
    	
		// Nieuw struct artikel
		article := createArticle(post)
		articles = append(articles, article)
	}
			
    cx.RespondWithData(articles)
}

/**
 * Onderstaande functies worden nog niet gebruikt. Nog niet nodig gehad.
 */
func (cr *SearchController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *SearchController) ReadMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *SearchController) Create(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *SearchController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *SearchController) Update(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *SearchController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}








