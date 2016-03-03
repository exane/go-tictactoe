package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "log"
  "github.com/gorilla/websocket"
)

func serverSetup() {
  router := mux.NewRouter().StrictSlash(true)

  /*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.HandleFunc("/set/:index", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request from /init!")
    fmt.Fprintf(w, "Request from /init!")
  })*/

  //router.Handle("/", http.FileServer(http.Dir("./www/")))
  //router.HandleFunc("/set/{fieldID}", setField)
  //router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www")))

  //router.Handle("/www", http.FileServer(http.Dir("./static/")))

  //http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./src/tictactoe/static/"))))
  //router.HandleFunc("/", index)
  router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./src/tictactoe/www/"))))
  router.HandleFunc("/ws", wsHandler)

  http.ListenAndServe(":65123", router)

  webSocketSetup()

}

func webSocketSetup() {

}

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  _ = conn
  if err != nil {
    log.Println(err)
    return
  }
  log.Println("new Connection!")

}

func setField(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fieldID := vars["fieldID"]
  _ = fieldID
}
func index(w http.ResponseWriter, r *http.Request) {
  log.Println("yoo")
}



