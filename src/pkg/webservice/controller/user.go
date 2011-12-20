/**
 * Implementeer hier elke REST methode (+restMany methode) voor gebruiker
 *
 * @author A. Glansbeek
 * @version 1.0
 * @date 2011-12-18
 */
 
package controller

import (
	"goweb"
	"http"
	"webservice/structs"
)

type UserController struct{}

func (cr *UserController) Create(cx *goweb.Context) {
	cx.RespondWithData(structs.User{"1", "Mat", 28})
}

func (cr *UserController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithOK()
}

func (cr *UserController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *UserController) Read(id string, cx *goweb.Context) {
	if id == "1" {
		cx.RespondWithData(structs.User{id, "Mat", 28})
	} else if id == "2" {
		cx.RespondWithData(structs.User{id, "Laurie", 27})
	} else {
		cx.RespondWithNotFound()
	}
}

func (cr *UserController) ReadMany(cx *goweb.Context) {
	cx.RespondWithData([]structs.User{structs.User{"1", "Mat", 28}, structs.User{"2", "Laurie", 27}})
}

func (cr *UserController) Update(id string, cx *goweb.Context) {
	cx.RespondWithData(structs.User{id, "Mat", 28})
}

func (cr *UserController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithData(structs.User{"1", "Mat", 28})
}




