/**
 * Implementeer hier elke REST methode (+restMany methode) voor categorieen
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
package controller

import (
	"goweb"
	"http"
    "webservice/structs"
    "webservice/models"
)

type CategoryController struct{}

/**
 * Retourneer hoofd categorieen
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
func (cr *CategoryController) ReadMany(cx *goweb.Context) {
	var m structs.WordpressHeadCategory
	
	models.CallWordpressApi(cx, &m, "get_category_index", nil)

    var headCategories []structs.HeadCategory
    for _, category := range m.Categories {
    	// Alleen parents 0 (dat is hoofd category)
     	if category.Parent == 0 {
     		headCategories = append(headCategories, structs.HeadCategory{category.Id, category.Title})
     	}
	}
    
    cx.RespondWithData(headCategories)
}

/**
 * Onderstaande functies worden nog niet gebruikt. Nog niet nodig gehad.
 */
func (cr *CategoryController) Read(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}
 
func (cr *CategoryController) Create(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CategoryController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CategoryController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CategoryController) Update(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CategoryController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}














