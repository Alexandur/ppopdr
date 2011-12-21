/**
 * Implementeer hier elke REST methode (+restMany methode) voor artikelen
 *
 * @author A. Glansbeek
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

func (cr *ArticleController) Create(cx *goweb.Context) {
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

func (cr *ArticleController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithOK()
}

func (cr *ArticleController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}

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

func (cr *ArticleController) ReadMany(cx *goweb.Context) {
	var m structs.WordpressPostMany
	
	models.CallWordpressApi(cx, &m, "get_recent_posts", nil)

    var articles []structs.Article
    for _, post := range m.Posts {
    	articles = append(articles, structs.Article{strconv.Itoa(post.Id), post.Title, post.Author.First_name + " " + post.Author.Last_name, post.Date, post.Comment_count})
	}
    
    cx.RespondWithData(articles)
}

func (cr *ArticleController) Update(id string, cx *goweb.Context) {
	cx.RespondWithData(structs.Article{"1", "Titel1", "Lex", "20111219", 0})
}

func (cr *ArticleController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithData(structs.Article{"1", "Titel1", "Lex", "20111219", 0})
}














