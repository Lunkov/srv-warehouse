package main

import (
  "sync"
)

// Количество и стоимость товаров на складах

type WarehouseGoods struct {
  Warehouse_ID       string  
  Goods_ID           int64  
  Quantity           int
  Cost               int    
}

var memPrWH = make(map[int64]map[string]WarehouseGoods)
var muPrWH   sync.RWMutex

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
