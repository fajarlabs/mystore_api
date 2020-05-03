package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "gopkg.in/ini.v1"
    "os"
    "github.com/jmoiron/sqlx"
    "github.com/restapi_go/body_request"
    "github.com/restapi_go/response_body"
    "github.com/restapi_go/entity_model"
)

/* ========== VARIABLE INITIALIZE =============== */

var host string
var port int
var user string
var password string
var dbname string
var port_app string

/* ========== REQUEST INITIALIZE =============== */

func initialize(host *string, port *int, user *string, password *string, dbname *string, port_app *string){
    // init setup database configuration ini
    cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        // shutdown apps
        os.Exit(1)
    }

    *host = cfg.Section("database").Key("hostname").String()
    *user = cfg.Section("database").Key("username").String()
    *password = cfg.Section("database").Key("password").String()
    *port = cfg.Section("database").Key("port").MustInt()
    *dbname = cfg.Section("database").Key("database").String()
    *port_app = cfg.Section("application").Key("port_app").String()
}

/* ========== CONTROLLER =============== */

func indexPage(w http.ResponseWriter, r *http.Request){
    responseCommands := response_body.ResponseCommands{"ok","Rest API is ready to use"}
    w.Header().Set("Content-Type", "application/json")
    jsonInBytes, _ := json.Marshal(responseCommands)
    w.Write(jsonInBytes)
}

func loginAuth(w http.ResponseWriter, r *http.Request){
    var p body_request.BodyRequestLogin    
    switch r.Method {
    case http.MethodPost:
        // Parser vesselfinder binnary
        // Try to decode the request body into the struct. If there is an error,
        // respond to the client with the error message and a 400 status code.
        err := json.NewDecoder(r.Body).Decode(&p)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Println("Endpoint Hit: LoginAuth")

        pg_con_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
        db, err := sqlx.Connect("postgres", pg_con_string)
        if err != nil {
            log.Fatalln(err)
        }

        userData := entity_model.App_user{}
        err = db.Get(&userData, "SELECT id, username, email, first_name, last_name, password, counter, status, remark, change_password, phone, photo, is_delete, who_delete_it FROM app_user WHERE username = $1 AND password = $2 ", p.Username, p.Password)
        
        if err == nil {
            // set response json
            responseLogin := response_body.ResponseLogin{"success",userData}
            w.Header().Set("Content-Type", "application/json")
            jsonInBytes, _ := json.Marshal(responseLogin)
            w.Write(jsonInBytes)
        } else {
            fmt.Println(err)
            // Give an error message.
            responseCommands := response_body.ResponseCommands{"failed","no data!"}
            fmt.Println(userData)
            w.Header().Set("Content-Type", "application/json")
            jsonInBytes, _ := json.Marshal(responseCommands)
            w.Write(jsonInBytes)
        }
    default:
        // Give an error message.
        responseCommands := response_body.ResponseCommands{"failed","Url route not found!"}
        w.Header().Set("Content-Type", "application/json")
        jsonInBytes, _ := json.Marshal(responseCommands)
        w.Write(jsonInBytes)
    }
}

func orderList(w http.ResponseWriter, r *http.Request){
}

func processList(w http.ResponseWriter, r *http.Request){
}

func deliveredList(w http.ResponseWriter, r *http.Request){
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)

    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", indexPage)
    myRouter.HandleFunc("/loginAuth", loginAuth)
    myRouter.HandleFunc("/getOrderList", orderList)
    myRouter.HandleFunc("/getProcessList", processList)
    myRouter.HandleFunc("/getDeliveredList", deliveredList)

    // port application
    log.Fatal(http.ListenAndServe(port_app, myRouter))
}

/* ========== MAIN PROGRAM =============== */

func main() {
    initialize(&host, &port, &user, &password, &dbname, &port_app )
    handleRequests()
}