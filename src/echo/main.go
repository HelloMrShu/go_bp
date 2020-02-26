package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "text/template"
)

type Test struct {
    Id    int32
    Name  string
}

type Conn struct {
    Username string
    Password string
    Host string
    Port int
    Db string
}

func initDB() *gorm.DB {
    cn := Conn {
        "root",
        "root",
        "127.0.0.1",
        3306,
        "test",
    }

    connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cn.Username,cn.Password, cn.Host, cn.Port, cn.Db )
    db, err := gorm.Open("mysql", connArgs)
    if err != nil {
        log.Fatal(err)
    }
    db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    return db
}

func index(w http.ResponseWriter, r *http.Request) {
    db := initDB()
    defer db.Close()

    var data []Test

    db.Find(&data)

    wd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    t, err := template.ParseFiles(wd + "/src/echo/views/index.gtpl")
    if err!=nil {
        fmt.Fprintln(w, err)
        return
    }
    t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", index)                 //模板列表路由
    err := http.ListenAndServe(":8081", nil)    //监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}