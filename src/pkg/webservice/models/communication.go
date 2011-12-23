/**
 * Logic van inkomende en uitgaande verbindingen
 *
 * @author A. Glansbeek en P. Kompier
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
	"url"
	"webservice/models/data"
)

// URL van Wordpress API
const URL_WORDPRESS = "http://sil.erwinvd.nl/sil/api/"

/**
 * Roep de wordpress API aan
 *
 * @author A. Glansbeek en P. Kompier
 * @params *goweb.Context Context van goweb
 * @params interface De struct waarin json opgeslagen wordt
 * @params string functie naam van API wordpress
 * @params map[string]string map waarin params staan
 */
func CallWordpressApi (cx *goweb.Context, v interface{}, callfunction string, params map[string]string) {	
	var contents []uint8
	
	url := createURL(callfunction, params)
    client := urlfetch.Client(appengine.NewContext(cx.GetRequest()))
    response, err := client.Get(url)
    data.SaveLog(url, "A.Glansbeek of P.Kompier", cx.GetRequest())
    
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

/**
 * Roep de wordpress API aan
 *
 * @author A. Glansbeek en P. Kompier
 * @params string functie naam van API wordpress
 * @params map[string]string map waarin params staan
 */
func createURL (callfunction string, params map[string]string) string {
	return URL_WORDPRESS + callfunction + "/" + parseParams(params)
}

/**
 * Maak van de params een mooie string en bruikbaar voor wordpress API
 * 
 * @author A. Glansbeek en P. Kompier
 * @params map[string]string map waarin params staan
 */
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
		
		s = s + prefix + key + "=" + url.QueryEscape(value)
		i++
	}
	
	return s
}

