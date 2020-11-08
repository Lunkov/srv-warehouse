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

var memShop    = make(map[string]Shop)
var memShopReg = make(map[int64][]string) // index region_id -> shops
var muShop   sync.RWMutex

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

