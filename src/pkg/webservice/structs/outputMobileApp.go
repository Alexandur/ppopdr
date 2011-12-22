/**
 * Output: GO WebService => Mobile App
 * Hierin staat de structuur van de data. 
 * 
 * !! Variablen MOETEN met hoofdletter anders pakt JSON encoding het niet  
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-18
 */
 
package structs

// Artikelen output
type Article struct {
	Id string
	Title string
	Author string
	Date string
	TipsCount int
}

// Gebruiker output
type User struct {
	Id   string
	Name string
	Age  int
}

// Tips toevoeging status response
type TipResponse struct {
	Status string
}