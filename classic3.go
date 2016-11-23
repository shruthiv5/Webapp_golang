/* authenticating before doin a post */
package main

import ( "fmt"
          "net/http"
        "github.com/codegangsta/negroni"
	 "github.com/julienschmidt/httprouter"
        )

func MessageRend(w http.ResponseWriter,r * http.Request,p httprouter.Params){

        fmt.Fprintf(w,"L&T!")
}


func ReplyFromSer(w http.ResponseWriter,r * http.Request, p httprouter.Params){

        fmt.Fprintf(w,"Stars,black hole,earth,sun,moon")
}


func main(){


	n:=negroni.Classic()

	check:=0
	m:=func(res http.ResponseWriter,req *http.Request,next http.HandlerFunc){

	if check==0{
	if req.URL.Query().Get("password")== "secret123" {
        	//fmt.Fprint(res,"Before...")
        	next(res,req)
        }else{
                http.Error(res,"Not Authorised",401)

        }
	check+=1
	}else{
        	next(res,req)
	        //fmt.Fprint(res,"...After")
       	     }
	}
	

        n.Use(negroni.HandlerFunc(m))


	
	r:=httprouter.New()
	r.GET("/",MessageRend)
	r.POST("/posts",ReplyFromSer)

	n.UseHandler(r)
	n.Run(":8080")
}


