package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie")
}

func test1(c *gin.Context) {
	id:=c.Query("id")
	name:=c.DefaultQuery("name","shuai")
	no:=c.DefaultQuery("no","qwer")

	c.JSON(200,gin.H{
		"status":"ok",
		"message":fmt.Sprintf("id is %s,name is %s,no is %s",id,name,no),
		"nick":name,
		"method":c.Request.Method,
	})
	//c.String(http.StatusOK,"id is %s,name is %s,no is %s",id,name,no)
}

func test2(c *gin.Context) {
	name:=c.Param("name")
	c.String(http.StatusOK, "my name is %s",name)
}

func test3(c *gin.Context) {
	c.String(http.StatusOK, "test3")
}

func main() {
	/*http.HandleFunc("/",sayhelloName)
	err:=http.ListenAndServe(":3131",nil)
	if err!=nil {
		log.Fatal("ListenAndServer: ",err)
	}*/

	router := gin.Default()

	router.GET("/", test1)
	router.GET("/home:name", test2)
	router.GET("/qian", test3)

	s := http.Server{
		Addr:           ": ",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(s.ListenAndServe())
}
