package main

import (
  "sync"
  "encoding/gob"
  "os"
  "unsafe"
  "github.com/golang/glog"
)

// Количество и стоимость товаров в магазинах

type ShopGoods struct {
  Shop_ID       string  
  Goods_ID      int64  
  Quantity      int
  Cost          int    
}

type ShopGoodsLite struct {
  Quantity      int
  Cost          int    
}

var muPrSP   sync.RWMutex

var memPrSP = make(map[int64]map[string]ShopGoodsLite, 100000)
var maxx_shops int64
var maxx_goods int64

func ShopGoodsInit(max_goods int64, max_shops int64) {
  maxx_shops = max_shops
  maxx_goods = max_goods
  memPrSP = make(map[int64]map[string]ShopGoodsLite, max_goods)
  
  glog.Infof("LOG: Goods In Shops: sizeof(item) = %d", unsafe.Sizeof(ShopGoodsLite{}))
  glog.Infof("LOG: Goods In Shops: sizeof(map)  = %d", unsafe.Sizeof(memPrSP))
}

func ShopGoodsAppend(info *ShopGoods) {
  muPrSP.Lock()
  if _, ok := memPrSP[info.Goods_ID]; !ok {
    memPrSP[info.Goods_ID] = make(map[string]ShopGoodsLite, maxx_goods)
  }
  var sgl ShopGoodsLite
  sgl.Quantity = info.Quantity
  sgl.Cost = info.Cost  
  memPrSP[info.Goods_ID][info.Shop_ID] = sgl
  muPrSP.Unlock()
}

func RlGetGoodsInShop(goods_id int64, shop_id string) (*ShopGoodsLite) {
  muPrSP.RLock()
  item, ok := memPrSP[goods_id]
  muPrSP.RUnlock()
  if ok {
    res, ok2 := item[shop_id]
    if ok2 {
      return &res
    }
  }
  
  return &ShopGoodsLite{Quantity: 0, Cost: 0}
}

func WriteFileQuantityInShops(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, _ := os.Create(filename)
  defer file.Close()
  encoder := gob.NewEncoder(file)
  encoder.Encode(memPrSP)
}

func LoadFileQuantityInShops(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, err := os.Open(filename)
  if err !=nil {
    glog.Errorf("ERR: Load(%s): %v", filename, err)
    return
  }
  defer file.Close()
  
  decoder := gob.NewDecoder(file)
  err = decoder.Decode(&memPrSP)
  if err != nil {
    glog.Errorf("ERR: Decoder(%s): %v", filename, err)
    return
  }
}
