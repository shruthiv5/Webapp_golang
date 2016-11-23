package main

import ( "fmt"
          "net/http"
        "github.com/codegangsta/negroni"
	 "github.com/julienschmidt/httprouter"
        )

func MessageRend(w http.ResponseWriter,r * http.Request,p httprouter.Params){

        fmt.Fprintf(w,"Hello Universe!")
}


func ReplyFromSer(w http.ResponseWriter,r * http.Request, p httprouter.Params){

        fmt.Fprintf(w,"Stars,earth,sun,moon")
}


func main(){


	n:=negroni.Classic()
	
	r:=httprouter.New()
	r.GET("/",MessageRend)
	r.POST("/posts",ReplyFromSer)
	r.GET("/custom",MessageRend)

	n.UseHandler(r)
	n.Run(":8080")
}

