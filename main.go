package main

import (
  "flag"
  "net/http"
  "github.com/gorilla/mux"
  
  "github.com/golang/glog"
)

func main() {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  flag.Set("v", "9")
  flag.Parse()
  
  fillData4Tests()
  //LoadAll()
  SaveAll()
  
  router := mux.NewRouter()
  
  router.HandleFunc("/product4sale/{goods_id}/{region_id}", webGoods4Sale)
  
  glog.Infof("LOG: Start service")
  err := http.ListenAndServe(":3000", router)
  if err != nil {
    glog.Errorf("ERR: HTTP server: %s", err)
  }
}

