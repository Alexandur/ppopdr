/**
 * Input: Wordpress API => GO WebService
 * Hierin staat alle data die wij willen ontvangen van de Wordpress API, 
 * de rest van de gegevens wordt genegeerd
 *
 * !! Variablen MOETEN met hoofdletter anders pakt JSON decoding het niet  
 *
 * @author A. Glansbeek
 * @version 1.0
 * @date 2011-12-18
 */
 
package structs

/* POSTS */
	// Hoofd structuur van meerdere posts (artikelen)
	type WordpressPostMany struct {
		Status string
		Count int
		Posts []Post
	}
	
	// Hoofd structuur van 1 post (artikel)
	type WordpressPost struct {
		Status string
		Count int
		Post Post
	}
	
	// Structuur van post
	type Post struct {
		Id int
		Title string
		Date string
		Author struct {
			First_name string
			Last_name string
		}
		Comment_count int
	}
	
/* COMMENTS */
	// Hoofd structuur van response van een CREATE
	type WordpressCommentResponse struct {
		Status string
	}

/* USERS */