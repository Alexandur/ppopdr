/**
 * Implementeer hier elke REST methode (+restMany methode) voor comments (tips)
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

type CommentController struct{}

/**
 * Creert een nieuwe comment (tip)
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
func (cr *CommentController) Create(cx *goweb.Context) {
	reqValues := cx.GetReqData()

	params := map[string]string {
		"post_id":reqValues.Get("post_id"),
		"name":reqValues.Get("name"),
		"email":reqValues.Get("email"),
		"content":reqValues.Get("content"),
	}
	
	var m structs.WordpressCommentResponse
	
	models.CallWordpressApi(cx, &m, "submit_comment", params)
	
	cx.RespondWithData(structs.TipResponse{m.Status})
}

/**
 * Onderstaande functies worden nog niet gebruikt. Nog niet nodig gehad.
 */
func (cr *CommentController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CommentController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CommentController) Read(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CommentController) ReadMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CommentController) Update(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *CommentController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}





