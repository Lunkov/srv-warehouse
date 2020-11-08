package main

import (
  "sync"
)

// Количество и стоимость товаров в магазинах

type ShopGoods struct {
  Shop_ID       string  
  Goods_ID      int64  
  Quantity      int
  Cost          int    
}

var memPrSP = make(map[int64]map[string]ShopGoods)
var muPrSP   sync.RWMutex

func ShopGoodsAppend(info *ShopGoods) {
  muPrSP.Lock()
  if _, ok := memPrSP[info.Goods_ID]; !ok {
    memPrSP[info.Goods_ID] = make(map[string]ShopGoods)
  }
  memPrSP[info.Goods_ID][info.Shop_ID] = *info
  muPrSP.Unlock()
}

func RlGetGoodsInShop(goods_id int64, shop_id string) (*ShopGoods) {
  muPrSP.RLock()
  item, ok := memPrSP[goods_id]
  muPrSP.RUnlock()
  if ok {
    res, ok2 := item[shop_id]
    if ok2 {
      return &res
    }
  }
  
  return &ShopGoods{Shop_ID: shop_id, Goods_ID: goods_id, Quantity: 0, Cost: 0}
}
