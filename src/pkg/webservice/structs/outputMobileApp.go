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
	Id int
	Title string
	Author string
	Date string
	Content string
	TipsCount int
	Category string
	Tips []Tips
}

type Tips struct {
	Name string
	Date string
	Content string
	Author string
}

// Tips toevoeging status response
type TipResponse struct {
	Status string
}

// Categorieen
type HeadCategory struct {
	Id int
	Title string
}