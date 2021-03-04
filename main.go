package main

import (
	"github.com/colombia9503/quasar-fire-operation/router"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

//se halla el punto media trilateración

func main() {
	//fmt.Println(location.GetLocation(100, 115.5, 142.7))
	//fmt.Println(messages.GetMessage(
	//	[]string{"este", "", "", "mensaje", ""},
	//	[]string{"", "es", "", "", "secreto"},
	//	[]string{"este", "", "un", "", ""},
	//))
	//app := App{}
	//app.InitializeRouters(
	//	os.Getenv("APP_DB_USERNAME"),
	//	os.Getenv("APP_DB_PASSWORD"),
	//	os.Getenv("APP_DB_NAME"))
	//
	//app.Run(":8081")

	var ar router.AppRouter
	ar.InitializeRouters()

	n := negroni.Classic()
	n.UseHandler(ar.Router)

	log.Println("listening on port :8081")
	http.ListenAndServe(":8080", n)
}
