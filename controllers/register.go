package ctls

import(
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
)

func Register(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	temp,_ := template.ParseFiles("./views/register.html","./views/layout/footer.html","./views/layout/nav.html")
	temp.ExecuteTemplate(w ,"register.html", "欢迎注册")
	// temp,_ := template.ParseFiles("./views/register.html","./views/layout/nav.html","./views/layout/footer.html")
	// temp.ExecuteTemplate(w ,"register.html","欢迎注册qaa")
}

func RegisterPost(w http.ResponseWriter , r *http.Request , _ httprouter.Params ){
	// r.ParseForm()
	uname := r.FormValue("uname")
	pwd := r.FormValue("pwd")
	fmt.Println(uname,pwd)
}