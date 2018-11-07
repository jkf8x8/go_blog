package ctls
import(
	// "fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
)

func Login(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	// fmt.Fprint(w , "login")

	temp,_ := template.ParseFiles("./views/login.html","./views/layout/nav.html","./views/layout/footer.html")
	temp.ExecuteTemplate(w ,"login.html","欢迎登陆")

}