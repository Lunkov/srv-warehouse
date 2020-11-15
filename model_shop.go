package main

import (
  "sync"
  "encoding/gob"
  "os"
  "unsafe"
  "github.com/golang/glog"
)

// Магазины

type Shop struct {
  CODE         string    `db:"code"           json:"code"            yaml:"code"`
  Region_ID    int64  
  Name         string    `db:"name"           json:"name"            yaml:"name"`
  Description  string    `db:"description"    json:"description"     yaml:"description"`
}

var memShop map[string]Shop
var memShopReg = make(map[int64][]string, 100) // index region_id -> shops
var muShop   sync.RWMutex

func ShopInit(max int64) {
  memShop = make(map[string]Shop, max)
  glog.Infof("LOG: Shops: max          = %d", max)
  glog.Infof("LOG: Shops: sizeof(item) = %d", unsafe.Sizeof(Shop{}))
  glog.Infof("LOG: Shops: sizeof(map)  = %d", unsafe.Sizeof(memShop))
}

func ShopCount() int {
  return len(memShop)
}

func ShopAppend(info *Shop) {
  muShop.Lock()
  memShop[info.CODE] = *info
  if _, ok := memShopReg[info.Region_ID]; !ok {
    memShopReg[info.Region_ID] = make([]string, 1)
  }
  memShopReg[info.Region_ID] = append(memShopReg[info.Region_ID], info.CODE)
  muShop.Unlock()
}

func GetShopByID(code string) (*Shop) {
  muShop.RLock()
  item, ok := memShop[code]
  muShop.RUnlock()
  if ok {
    return &item
  }
  return nil
}

func GetShopByRegionID(region_id int64) ([]string) {
  muShop.RLock()
  items, ok := memShopReg[region_id]
  muShop.RUnlock()
  if ok {
    return items
  }
  return make([]string, 0)
}

func WriteFileShops(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, _ := os.Create(filename)
  defer file.Close()
  encoder := gob.NewEncoder(file)
  encoder.Encode(memShop)

  fileIndex, _ := os.Create(filename+".index")
  defer fileIndex.Close()
  encoderIndex := gob.NewEncoder(fileIndex)
  encoderIndex.Encode(memShopReg)
}


func LoadFileShops(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, err := os.Open(filename)
  if err !=nil {
    glog.Errorf("ERR: Load(%s): %v", filename, err)
    return
  }
  defer file.Close()
  
  decoder := gob.NewDecoder(file)
  err = decoder.Decode(&memShop)
  if err != nil {
    glog.Errorf("ERR: Decoder(%s): %v", filename, err)
    return
  }
  
  fileI, errI := os.Open(filename+".index")
  if errI !=nil {
    glog.Errorf("ERR: Load(%s): %v", filename+".index", errI)
    return
  }
  defer file.Close()
  
  decoderI := gob.NewDecoder(fileI)
  errI = decoderI.Decode(&memShopReg)
  if errI != nil {
    glog.Errorf("ERR: Decoder(%s): %v", filename+".index", errI)
    return
  }
}

