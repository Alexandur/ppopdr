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
	
	// Search
	goweb.MapRest("/search", new(controller.SearchController))
	
	// Documentation wadl
	goweb.MapFunc("/documentation/wadl", func(c *goweb.Context) {
		var docTemplate = template.Must(template.ParseFile("webservice/views/wadl.wadl"))
	
		if err := docTemplate.Execute(c.ResponseWriter, nil); err != nil {
	        c.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
	    }
	})
	
	// Documentation
	goweb.MapFunc("/documentation", func(c *goweb.Context) {
		var docTemplate = template.Must(template.ParseFile("webservice/views/documentation.html"))
	
		if err := docTemplate.Execute(c.ResponseWriter, nil); err != nil {
	        c.RespondWithErrorMessage(err.String(), http.StatusInternalServerError)
	    }
	})
	
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








