package main

import (
  "sync"
  "encoding/gob"
  "os"
  "unsafe"
  "github.com/golang/glog"
)

// Товары

type Goods struct {
  ID           int64     `db:"id"             json:"id"              yaml:"id"`
  Name         string    `db:"name"           json:"name"            yaml:"name"`
  Description  string    `db:"description"    json:"description"     yaml:"description"`
}

var memG map[int64]Goods
var muG   sync.RWMutex

func GoodsInit(max int64) {
  memG = make(map[int64]Goods, max)
  glog.Infof("LOG: Goods: max          = %d", max)
  glog.Infof("LOG: Goods: sizeof(item) = %d", unsafe.Sizeof(Goods{}))
  glog.Infof("LOG: Goods: sizeof(map)  = %d", unsafe.Sizeof(memG))
}

func GoodsCount() int64 {
  return int64(len(memG))
}

func GoodsAppend(info *Goods) {
  muG.Lock()
  memG[info.ID] = *info
  muG.Unlock()
}

func GetGoodsByID(id int64) (*Goods) {
  muG.RLock()
  item, ok := memG[id]
  muG.RUnlock()
  if ok {
    return &item
  }
  return nil
}

func WriteFileGoods(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, _ := os.Create(filename)
  defer file.Close()
  encoder := gob.NewEncoder(file)
  encoder.Encode(memG)
}

func LoadFileGoods(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, err := os.Open(filename)
  if err !=nil {
    glog.Errorf("ERR: Load(%s): %v", filename, err)
    return
  }
  defer file.Close()
  
  decoder := gob.NewDecoder(file)
  err = decoder.Decode(&memG)
  if err != nil {
    glog.Errorf("ERR: Decoder(%s): %v", filename, err)
    return
  }
}
