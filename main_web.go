package main

import (
  "strconv"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)
func webGoods4Sale(w http.ResponseWriter, r *http.Request)  {
  params := mux.Vars(r)
  
  
  goods_id_str, okp1 := params["goods_id"]
  if !okp1 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  goods_id, err1 := strconv.ParseInt(goods_id_str, 10, 64)
  if err1 != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  
  region_id_str, okp2 := params["region_id"]
  if !okp2 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  region_id, err2 := strconv.ParseInt(region_id_str, 10, 64)
  if err2 != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  
  res := GetGoods4Sale(goods_id, region_id)
  
  jsonRes, _ := json.Marshal(res)
  w.Write(jsonRes)
}
