package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//test github/gin
//import (
//	"fmt"
//	"net/http"
//	"strings"
//	"github.com/gin-gonic/gin"
//	"time"
//)
//
//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//
//	fmt.Println(r.Form)
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello astaxie")
//}
//
//func test1(c *gin.Context) {
//	id:=c.Query("id")
//	name:=c.DefaultQuery("name","shuai")
//	no:=c.DefaultQuery("no","qwer")
//
//	c.JSON(200,gin.H{
//		"status":"ok",
//		"message":fmt.Sprintf("id is %s,name is %s,no is %s",id,name,no),
//		"nick":name,
//		"method":c.Request.Method,
//	})
//	//c.String(http.StatusOK,"id is %s,name is %s,no is %s",id,name,no)
//}
//
//func test2(c *gin.Context) {
//	name:=c.Param("name")
//	c.String(http.StatusOK, "my name is %s",name)
//}
//
//func test3(c *gin.Context) {
//	c.String(http.StatusOK, "test3")
//}
//
//func main() {
//	/*http.HandleFunc("/",sayhelloName)
//	err:=http.ListenAndServe(":3131",nil)
//	if err!=nil {
//		log.Fatal("ListenAndServer: ",err)
//	}*/
//
//	router := gin.Default()
//
//	router.GET("/", test1)
//	router.GET("/home:name", test2)
//	router.GET("/qian", test3)
//
//	s := http.Server{
//		Addr:           ": ",
//		Handler:        router,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	fmt.Println(s.ListenAndServe())
//}

// 使用go自带http库实现简单的web服务器
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//func echo(wr http.ResponseWriter, r *http.Request) {
//	msg, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		wr.Write([]byte("echo error"))
//		return
//	}
//	fmt.Println(msg)
//	fmt.Println(r.Method)
//
//	writeLen, err := wr.Write(msg)
//	if err != nil || writeLen != len(msg) {
//		log.Println(err, "write len:", writeLen)
//	}
//}
//
//func main() {
//	http.HandleFunc("/", echo)
//
//	err := http.ListenAndServe("192.168.1.97:2121", nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// 使用middleware剥离非业务逻辑

var logger = log.New(os.Stdout, "", 0)

func hello(wr http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	wr.Write([]byte("hello"))
}

// timeMiddleware 函数使用接口型函数，用函数实现接口。
func timeMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)

		logger.Println(timeElapsed)
	})
}

func main() {
	// 以接口方式注册
	//http.Handle("/", timeMiddleware(hello))

	// 以函数方式注册
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("192.168.1.96:2121", nil)
	if err != nil {
		log.Fatal(err)
	}
}
