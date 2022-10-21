package demo

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var visitors atomic.Int32

// var lock sync.Mutex
var regex = regexp.MustCompile(`^\w*$`)

func HandleHi(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.RemoteAddr)
	//log.Println("handleHi hit")
	if !regex.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	//var currVisitor int
	//lock.Lock()
	//visitors++
	//currVisitor = visitors
	//lock.Unlock()
	currVisitor := int(visitors.Add(1))
	// w.Header().Set("Content-Type", "text/html; charset=utf-8") net/http will do this for us? TODO where?

	//w.Write( // underlying 2.1GB
	//	[]byte( // 174MB
	//		fmt.Sprintf( // underlying 206MB
	//			"<h1 style='color: %s'>Welcome!</h1>You are visitor number %d !",
	//			r.FormValue("color"),
	//			currVisitor,
	//		),
	//	),
	//)
	_, err := fmt.Fprintf( // 2.27GB
		w, // io writer
		"<h1 style='color: %s'>Welcome!</h1>You are visitor number %d !",
		r.FormValue("color"), // 13.50MB
		currVisitor,
	)
	if err != nil {
		log.Println(err)
	}
}
