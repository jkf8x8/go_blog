package ctls

import(
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
)

func Home(w http.ResponseWriter , r *http.Request,_ httprouter.Params){
	temp,_ := template.ParseFiles("./views/index.html","./views/layout/nav.html","./views/layout/footer.html")
	temp.Execute(w,"欢迎来到小站")

	fmt.Println("ok")
}
