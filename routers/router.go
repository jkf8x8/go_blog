package route

import (
	"net/http"
	"go_blog/controllers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// type PreHandle http.Handler

// func (ph PreHandle) ServerHTTP( w http.ResponseWriter , r *http.Request){
// 	// if(r.RequestURI =="/")
// 	fmt.Println("preHandle")
// }


func init(){
	router := httprouter.New()

	// sql
	// db,openErr :=sql.Open("mysql","root:admins@blog?charset=utf8")
	// if openErr !=nil{
	// 	fmt.Println(" ok ")
	// }
	// router 
	http.Handle("/static/", http.StripPrefix("/static/" ,http.FileServer(http.Dir("/public"))))
	router.GET("/",ctls.Home)
	router.GET("/login",ctls.Login)
	router.GET("/register",ctls.Register)
	router.POST("/register",ctls.RegisterPost)

	// fmt.Println(http.Dir("/public"))

	




	server:=http.Server{
		Addr:":8080",
		Handler:router,
	}

	server.ListenAndServe()
	fmt.Println("listen port 8080")
}

func Run(){
	
}