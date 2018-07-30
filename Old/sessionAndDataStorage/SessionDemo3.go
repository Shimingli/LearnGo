package main

import (
	"net/http"
	"github.com/astaxie/beego/session"
	"html/template"
)

func main() {
	http.HandleFunc("/login",count)
	http.ListenAndServe(":9092",nil)
}
var g *session.Manager

func count(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	sess,_ := g.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}