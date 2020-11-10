package main

import (
  "sync"
  "encoding/gob"
  "os"
  "github.com/golang/glog"
)

// Количество и стоимость товаров на складах

type WarehouseGoods struct {
  Warehouse_ID       string  
  Goods_ID           int64  
  Quantity           int
  Cost               int    
}

type WarehouseGoodsLite struct {
  Quantity           int
  Cost               int    
}

var muPrWH   sync.RWMutex
var memPrWH = make(map[int64]map[string]WarehouseGoodsLite)

func WarehouseGoodsAppend(info *WarehouseGoods) {
  muPrWH.Lock()
  if _, ok := memPrWH[info.Goods_ID]; !ok {
    memPrWH[info.Goods_ID] = make(map[string]WarehouseGoodsLite)
  }
  var whl WarehouseGoodsLite
  whl.Quantity = info.Quantity
  whl.Cost = info.Cost
  memPrWH[info.Goods_ID][info.Warehouse_ID] = whl
  muPrWH.Unlock()
}

func RlGetGoodsInWarehouse(goods_id int64, warehouse_id string) (*WarehouseGoodsLite) {
  muPrWH.RLock()
  item, ok := memPrWH[goods_id]
  muPrWH.RUnlock()
  if ok {
    res, ok2 := item[warehouse_id]
    if ok2 {
      return &res
    }
  }
  
  return &WarehouseGoodsLite{Quantity: 0, Cost: 0}
}

/*
var memPrWH = make(map[int64]map[string]WarehouseGoods)

func WarehouseGoodsAppend(info *WarehouseGoods) {
  muPrWH.Lock()
  if _, ok := memPrWH[info.Goods_ID]; !ok {
    memPrWH[info.Goods_ID] = make(map[string]WarehouseGoods)
  }
  memPrWH[info.Goods_ID][info.Warehouse_ID] = *info
  muPrWH.Unlock()
}

func RlGetGoodsInWarehouse(goods_id int64, warehouse_id string) (*WarehouseGoods) {
  muPrWH.RLock()
  item, ok := memPrWH[goods_id]
  muPrWH.RUnlock()
  if ok {
    res, ok2 := item[warehouse_id]
    if ok2 {
      return &res
    }
  }
  
  return &WarehouseGoods{Warehouse_ID: warehouse_id, Goods_ID: goods_id, Quantity: 0, Cost: 0}
}
*/

func WriteFileQuantityInWarehouse(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, _ := os.Create(filename)
  defer file.Close()
  encoder := gob.NewEncoder(file)
  encoder.Encode(memPrWH)
}

func LoadFileQuantityInWarehouse(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, err := os.Open(filename)
  if err !=nil {
    glog.Errorf("ERR: Load(%s): %v", filename, err)
    return
  }
  defer file.Close()
  
  decoder := gob.NewDecoder(file)
  err = decoder.Decode(&memPrWH)
  if err != nil {
    glog.Errorf("ERR: Decoder(%s): %v", filename, err)
    return
  }
}
