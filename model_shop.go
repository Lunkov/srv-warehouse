package main

import (
  "sync"
)

// Магазины

type Shop struct {
  CODE         string    `db:"code"           json:"code"            yaml:"code"`
  Region_ID    int64  
  Name         string    `db:"name"           json:"name"            yaml:"name"`
  Description  string    `db:"description"    json:"description"     yaml:"description"`
}

var memShop = make(map[string]Shop)
var muShop   sync.RWMutex

func ShopCount() int {
  return len(memShop)
}

func ShopAppend(info *Shop) {
  muShop.Lock()
  memShop[info.CODE] = *info
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


