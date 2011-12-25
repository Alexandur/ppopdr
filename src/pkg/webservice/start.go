package webservice

import (
  "goweb"
  "http"
  "webservice/controller"
  "template"
)

func init() {
    // Articles
	goweb.MapRest("/article/category/{category_id}", new(controller.ArticleController))
	
	// Articles
	goweb.MapRest("/article", new(controller.ArticleController))
	
	// Comments
	goweb.MapRest("/comment", new(controller.CommentController))
	
	// Categories 
	goweb.MapRest("/category", new(controller.CategoryController))
	
	// Logs
	goweb.MapRest("/log", new(controller.LogController))
	
	// Index
	goweb.MapFunc("/", func(c *goweb.Context) {
		var indexTemplate = template.Must(template.ParseFile("webservice/views/index.html"))
	
		if err := indexTemplate.Execute(c.ResponseWriter, nil); err != nil {
	        c.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
	    }
	})
	
	// use the default Go handler
	http.Handle("/", goweb.DefaultHttpHandler)
}








