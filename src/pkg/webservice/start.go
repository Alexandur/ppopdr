package webservice

import (
  "goweb"
  "http"
  "webservice/controller"
)

func init() {
	// User
	goweb.MapRest("/user", new(controller.UserController))
    
    // Articles
	goweb.MapRest("/article", new(controller.ArticleController))
	
	// use the default Go handler
	http.Handle("/", goweb.DefaultHttpHandler)

// Wat test commentaar om de push te testen. Whoop whoop
}






































//import (
//	"twister/gae"
//	"twister/web"
//	"io"
//	"os"
//	"json"
//)
//
//type User struct {
//    Voornaam string
//    Achternaam string
//}
//
//func serveError(req *web.Request, status int, reason os.Error, header web.Header) {
//	header.Set(web.HeaderContentType, "text/plain; charset=utf-8")
//	w := req.Responder.Respond(status, header)
//	io.WriteString(w, web.StatusText(status))
//	if reason != nil {
//		io.WriteString(w, "\n")
//		io.WriteString(w, reason.String())
//	}
//}
//
//func handlePatrickHeaumeau(r *web.Request) {
//	w := r.Respond(web.StatusOK, web.HeaderContentType, "text/plain; charset=\"utf-8\"")
//    io.WriteString(w, "Patrick is " + r.URLParam["title"])
//}
//
//func handleUsers(r *web.Request) {
////	w := r.Respond(web.StatusOK, web.HeaderContentType, "text/plain; charset=\"utf-8\"")
//	user := User{"Patrick", "Kompier"}
//	b, _ := json.Marshal(user)
//	w := r.Respond(web.StatusOK, web.HeaderNameBytes(b))
//
//    io.WriteString(w, "JSON: ")
//}
//
//func init() {
//	const titleParam = "<title:[a-zA-Z0-9]+>"
//	gae.Handle("/",
//		web.SetErrorHandler(serveError,
//				web.NewRouter().
//					Register("/patrick/"+titleParam, "GET", handlePatrickHeaumeau).
//					Register("/getusers", "GET", handleUsers)))
//}





















