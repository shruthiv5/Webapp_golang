/*********************************************************************
Objective:Using a webserver with negroni,
1.To create a database
2.Write Operations into db
3.Read Operations from db
4.Json format

**********************************************************************/

package main

import ( "fmt"
	       "net/http"
	       "github.com/codegangsta/negroni"
         "github.com/julienschmidt/httprouter"
	       "database/sql"
	       _ "github.com/mattn/go-sqlite3"
	       "encoding/json"
	)


/**********************************
 Structure for json object
**********************************/
type Info struct{

	Table string     `json:"Table,omitempty"`
	Slno int       	 `json:"Slno"`
	Name string	     `json:"Name"`
	Duration string  `json:"Duration"`
}

/***************************************************************************************
				HTTP POST Handler
****************************************************************************************/

func CreateResource(res http.ResponseWriter,req * http.Request,p httprouter.Params){


/*** converting json to struct ***/

	 var clidata Info
        err:=json.NewDecoder(req.Body).Decode(&clidata)
        if err!=nil{
                panic(err)
        }
        fmt.Println(clidata)

	custom:=p.ByName("id")

	dd:=NewDB("dvr.db")
	Insert_Data(dd,custom,clidata)



}
/*************************************************************************************
				HTTP GET Handler
*************************************************************************************/

func FetchResource(res http.ResponseWriter, req *http.Request,p httprouter.Params){

	custom:=p.ByName("id")
	dd:=NewDB("dvr.db")
	Read_Data(dd,custom,res)

}

/**********************************************************************************
			      MAIN FUNC
***********************************************************************************/

func main(){

/*** Negroni Init ***/

	n:=negroni.Classic()

/*** Sqlite3 Init ***/

	dvrdb:=NewDB("dvr.db")
	CreateTable("Recordings",dvrdb)
	CreateTable("Scheduled",dvrdb)

/*** HTTP Init ***/

	r:=httprouter.New()
	r.POST("/insert/:id",CreateResource)
	r.GET("/fetch/:id",FetchResource)

/*** Server ON ***/

	n.UseHandler(r)
	n.Run(":8080")
}

/***********************************************************
		Create a database
************************************************************/

func NewDB(dbname string) *sql.DB{

	db,err:=sql.Open("sqlite3",dbname)
	if err!=nil{
		panic(" Unable to create data base")
	}
	return db
}

/***********************************************************
		Create tables
***********************************************************/

func CreateTable(tablename string,db *sql.DB) {

	cstring:="create table "+tablename+"(slno integer,name varchar(30),duration integer)"
	db.Exec(cstring)
}
/************************************************************
		Insert in dvr Database
*************************************************************/

func Insert_Data(db *sql.DB, lt string,cd Info ){

	qstring:="insert into "+lt+"(slno,name,duration) values(?,?,?)"
	wd,_:=db.Prepare(qstring)
	wd.Exec(cd.Slno,cd.Name,cd.Duration)

}
/**************************************************************
		To Print json data
**************************************************************/

func ConvertJson(db *sql.DB,lt string,res http.ResponseWriter){

	var rd Info
	qstring:="select slno,name,duration from "+lt
		r,_:=db.Query(qstring)
 		for r.Next(){
                         err:=r.Scan(&rd.Slno,&rd.Name,&rd.Duration)
                        if err!=nil{
                                panic("cannot read from data base")
                        }
                }
		err:=json.NewEncoder(res).Encode(rd)
                if err!=nil{
                        panic(err)
                }
}

/***************************************************************
		To Read from a specific table in dvr.db
***************************************************************/

func Read_Data(db *sql.DB,lt string,res http.ResponseWriter){

	if lt=="Recordings"{
		ConvertJson(db,lt,res)
	}else if lt=="Scheduled"{
		ConvertJson(db,lt,res)
	}
}
