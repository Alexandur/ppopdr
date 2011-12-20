/**
 * Logic van inkomende en uitgaande verbindingen
 *
 * @author A. Glansbeek
 * @version 1.0
 * @date 2011-12-18
 */
 
package models

import (
	"goweb"
	"appengine"
    "appengine/urlfetch"
    "io/ioutil"
	"json"
	"http"
)

const URL_WORDPRESS = "http://sil.erwinvd.nl/sil/api/"

type Communication struct{}

func (com *Communication) Bla() {

}

func Bla2() {

}

func CallWordpressApi (cx *goweb.Context, v interface{}, callfunction string, params map[string]string) {
	var contents []uint8

    client := urlfetch.Client(appengine.NewContext(cx.GetRequest()))
    response, err := client.Get(createURL(callfunction, params))
    
    // Error check
    if err != nil {
    	cx.RespondWithErrorMessage("Kan geen verbinding maken met de wordpress webservice", http.StatusBadRequest)
    	return
    } else {
    	// Aan het einde van deze functie sluit response
        defer response.Body.Close()
        
        // Maak van de content bytes
        contents, err = ioutil.ReadAll(response.Body)
        
        // Error check
        if err != nil {
        	cx.RespondWithErrorMessage("Kan response niet lezen", http.StatusBadRequest)
    		return
        }
    }

    err = json.Unmarshal(contents, v)
  
  	// Error check
  	if err != nil {
  		cx.RespondWithErrorMessage("Kan json niet parsen", http.StatusBadRequest)
    	return
    }
}

func createURL (callfunction string, params map[string]string) string {
	return URL_WORDPRESS + callfunction + "/" + parseParams(params)
}

func parseParams (params map[string]string) string {
	s := ""
	prefix := ""
	i := 0
	
	for key, value := range params {
		if (i == 0) {
			prefix = "?"
		} else {
			prefix = "&"
		}
		
		s = s + prefix + key + "=" + value
		i++
	}
	
	return s
}
