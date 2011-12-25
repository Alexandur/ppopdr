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
    "webservice/structs"
    "webservice/models"
)

type ArticleController struct{}


/**
 * GET artikel met id [1..n], retourneer een artikel
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
func (cr *ArticleController) Read(id string, cx *goweb.Context) {
	params := map[string]string {"post_id":id}
	var m structs.WordpressPost
	
	models.CallWordpressApi(cx, &m, "get_post", params)

	if m.Post.Id == 0 {
		var str []string
		str = append(str, "Het opgevraagde id bestaat niet")

		cx.Respond(nil, 202, str, cx)
	} else {
		cx.RespondWithData(createArticle(m.Post))
	}
}

/**
 * Retourneer alle artikelen
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
func (cr *ArticleController) ReadMany(cx *goweb.Context) {
	
	var m structs.WordpressPostMany
	
	// Als category_id is, pak dan alleen article van die category
	if cx.PathParams["category_id"] != "" {
		params := map[string]string {
			"category_id":cx.PathParams["category_id"],
		}

	 	models.CallWordpressApi(cx, &m, "get_category_posts", params)
	} else {
		models.CallWordpressApi(cx, &m, "get_recent_posts", nil)
	}
	
	var mapper  = make(map[string][]structs.Article)

    for _, post := range m.Posts {
    	
		// Nieuw struct artikel
		article := createArticle(post)
		
		// Zet het artikel in de goede category 
		_, present := mapper[article.Category]
		
	 	if present {
			mapper[article.Category] = append(mapper[article.Category], article)
	 	} else {
	 		var articles []structs.Article
 			articles = append(articles, article)
 			mapper[article.Category] = articles
		}
	}
			
    cx.RespondWithData(mapper)
}

/**
 * Creer article 
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-25
 */
func createArticle (post structs.Post) (structs.Article){
	// Tips
	var tips []structs.Tips
	for _, tip := range post.Comments {
    	tips = append(tips, structs.Tips{
    		tip.Name,
    		tip.Date,
    		tip.Content,
    		tip.Author.First_name + " " + tip.Author.Last_name,
    	})
	}

	var categoryName = ""
	if len(post.Categories) < 1 {
		categoryName = "Onbekend"
	} else {
		categoryName = post.Categories[0].Title
	}
	
	// Nieuw struct artikel
	article := structs.Article{
		post.Id, 
		post.Title, 
		post.Author.First_name + " " + post.Author.Last_name, 
		post.Date, 
		post.Content, 
		post.Comment_count,
		categoryName,
		tips,
	}
	
	return article
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



















