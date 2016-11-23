package main

import ( "fmt"
          "net/http"
        "github.com/codegangsta/negroni"
	 "github.com/julienschmidt/httprouter"
        )

func MessageRend(w http.ResponseWriter,r * http.Request,p httprouter.Params){

	custom:=p.ByName("id")
	switch(custom){
		case "g8"  :  fmt.Fprintf(w,"it's a g8 box")
		case "g6"  :  fmt.Fprintf(w,"it's a g6 box")
		case "g10" :  fmt.Fprintf(w,"it's a g10 box")
		default    :  fmt.Fprintf(w,"invalid set box version")
	}

}



func ReplyFromSer(w http.ResponseWriter,r * http.Request, p httprouter.Params){

        fmt.Fprintf(w,"Stars,earth,sun,moon")
}


func main(){


	n:=negroni.Classic()
	
	r:=httprouter.New()
	r.GET("/:id",MessageRend)
	r.POST("/posts",ReplyFromSer)

	n.UseHandler(r)
	n.Run(":8080")
}

