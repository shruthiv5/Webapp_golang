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

        fmt.Fprintf(w,"Stars,black hole,earth,sun,moon")
}


func main(){


	n:=negroni.Classic()

	m:=func(res http.ResponseWriter,req *http.Request,next http.HandlerFunc){
        fmt.Fprint(res,"Before...")
        next(res,req)
        fmt.Fprint(res,"...After")
        }

        n.Use(negroni.HandlerFunc(m))


	
	r:=httprouter.New()
	r.GET("/",MessageRend)
	r.POST("/posts",ReplyFromSer)

	n.UseHandler(r)
	n.Run(":8080")
}

