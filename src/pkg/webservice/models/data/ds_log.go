/**
 * Database logica voor log datastore
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-22
 */
package data

import (
	"appengine"
    "appengine/datastore"
	"os"
	"http"
	"time"
)

/* Hoe de datastore eruit ziet */
type PageView struct {
	// Pagina die aangeroepen wordt
    Page  string
    // Tijd van aanroepen    
    Date    datastore.Time
    // TODO: wie hem heeft aangeroepen de pagina. Wordt nu 1 vaste waarde
    ViewedBy string
}

/**
 * Insert in datastore log
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-22
 * @param string de url 
 * @param string de naam van wie url heeft aangeroepen
 * @param *http.Request 
 * @return os.Error errors
 */
func SaveLog(url string, viewedby string, r *http.Request) os.Error {
    c := appengine.NewContext(r)

    pv := PageView{   	
    	Page: url,	
        ViewedBy: viewedby,        
        Date:    datastore.SecondsToTime(time.Seconds()),
    }
    
    // Record toevoegen aan de datastore
    _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "PageView", nil), &pv)
    
    // Wanneer er een error optreedt deze terug geven
    return err  
}

/**
 * Retourneer alle logs
 *
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-22
 * @param *http.Request 
 * @return os.Error errors
 * @return interface struct met pageviews
 */
func GetLogs(r *http.Request) (os.Error, interface{}) {
    c := appengine.NewContext(r)
    q := datastore.NewQuery("PageView").Order("-Date").Limit(50)
    pageviews := make([]PageView, 0, 50)
    
    if _, err := q.GetAll(c, &pageviews); err != nil {
        return err, nil
    }
    	
    return nil, pageviews
}

/**
 * Delete alle logs
 *
 * Ziet er raar uit, maar dit is "de" manier in appengine
 * Bron: stackoverflow.com/questions/1062540/how-to-delete-all-datastore-in-google-app-engine (python)
 * "The best approach is the remote API method as suggested by Nick, he's an App Engine engineer from Google"
 * 
 * @author A. Glansbeek en P. Kompier
 * @version 1.0
 * @date 2011-12-22
 * @param *http.Request
 * @return os.Error errors
 */
func DeleteAllLogs(r *http.Request) os.Error {
	c := appengine.NewContext(r)
	pageviews := make([]PageView, 0)
	q := datastore.NewQuery("PageView")
	
	keys, err := q.GetAll(c, &pageviews);
	
	if err != nil {
        return err
    }
    
	err = datastore.DeleteMulti(c, keys)
	
	if err != nil {
        return err
    }
    
	return nil
}