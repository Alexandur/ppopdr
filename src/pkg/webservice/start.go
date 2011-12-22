package webservice

import (
  "goweb"
  "http"
  "webservice/controller"
)

func init() {
    // Articles
	goweb.MapRest("/article", new(controller.ArticleController))
	
	// Comments
	goweb.MapRest("/comment", new(controller.CommentController))
	
	// Logs
	goweb.MapRest("/log", new(controller.LogController))
	
	// use the default Go handler
	http.Handle("/", goweb.DefaultHttpHandler)
}








