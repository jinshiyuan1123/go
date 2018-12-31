package main

import (
    "fmt"
    "net/http"
    "log"
)

func responseHello(w http.ResponseWriter, r *http.Request){
    //将字符串发送给客户端
    fmt.Fprintf(w, "Hello Golang!")
}

func main(){
    //设置要解析的URL路由
    http.HandleFunc("/", responseHello)
    //设置监听的端口，开始监听
    errInfo := http.ListenAndServe(":8080", nil) 
    if errInfo != nil {
        log.Fatal("ListenAndServe: ", errInfo)
    }
}
