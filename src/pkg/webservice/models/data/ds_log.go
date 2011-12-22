package data

import (
	"appengine"
    "appengine/datastore"
	"os"
	"http"
	"time"
)

type PageView struct {
	// Pagina die aangeroepen wordt
    Page  string
    // Tijd van aanroepen    
    Date    datastore.Time
    // TODO: wie hem heeft aangeroepen de pagina. Wordt nu 1 vaste waarde
    ViewedBy string
}

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