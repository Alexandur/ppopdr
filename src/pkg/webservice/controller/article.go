/**
 * Implementeer hier elke REST methode (+restMany methode) voor artikelen
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
package controller

import (
	"goweb"
	"http"
    "strconv"
    "webservice/structs"
    "webservice/models"
)

type ArticleController struct{}

// GET artikel met id [1..n]
func (cr *ArticleController) Read(id string, cx *goweb.Context) {
	params := map[string]string {"post_id":id}
	var m structs.WordpressPost
	
	models.CallWordpressApi(cx, &m, "get_post", params)

	if m.Post.Id == 0 {
		var str []string
		str = append(str, "Het opgevraagde id bestaat niet")

		cx.Respond(nil, 202, str, cx)
	} else {
		cx.RespondWithData(structs.Article{strconv.Itoa(m.Post.Id), m.Post.Title, m.Post.Author.First_name + " " + m.Post.Author.Last_name, m.Post.Date, m.Post.Comment_count})
	}
}

// GET alle artikelen
func (cr *ArticleController) ReadMany(cx *goweb.Context) {
	var m structs.WordpressPostMany
	
	models.CallWordpressApi(cx, &m, "get_recent_posts", nil)

    var articles []structs.Article
    for _, post := range m.Posts {
    	articles = append(articles, structs.Article{strconv.Itoa(post.Id), post.Title, post.Author.First_name + " " + post.Author.Last_name, post.Date, post.Comment_count})
	}
    
    cx.RespondWithData(articles)
}

/**
 * Onderstaande functies worden nog niet gebruikt. Nog niet nodig gehad.
 */
func (cr *ArticleController) Create(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *ArticleController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *ArticleController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *ArticleController) Update(id string, cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

func (cr *ArticleController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}














