/**
 * Input: Wordpress API => GO WebService
 * Hierin staat alle data die wij willen ontvangen van de Wordpress API, 
 * de rest van de gegevens wordt genegeerd
 *
 * !! Variablen MOETEN met hoofdletter anders pakt JSON decoding het niet  
 *
 * @author A. Glansbeek en P. Kompier
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
	
	// Hoofd structuur van categorieen
	type WordpressCategory struct {
		Status string
		Count int
		Posts []Post
	}
	
	// Hoofd structuur van hoofd categorieen
	type WordpressHeadCategory struct {
		Status string
		Count int
		Categories []Category
	}
	
	// Hoofd structuur van response van een CREATE
	type WordpressCommentResponse struct {
		Status string
	}
	
	// Structuur van post
	type Post struct {
		Id int
		Title string
		Content string
		Date string
		Status string
		Author Author
		Comment_count int
		Categories []Category
		Comments []Comment
	}
	
	// Comment
	type Comment struct {
		Id int
		Name string
		Date string
		Content string
		Author Author
	}
	
	// Author
	type Author struct {
		First_name string
		Last_name string
	}
	
	// Categorieen
	type Category struct {
		Id int
		Title string
		Parent int
	}