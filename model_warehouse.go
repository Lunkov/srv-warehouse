package main

import (
  "sync"
)

// Склады

type Warehouse struct {
  CODE         string    `db:"code"           json:"code"            yaml:"code"`
  Region_ID    int64  
  Name         string    `db:"name"           json:"name"            yaml:"name"`
  Description  string    `db:"description"    json:"description"     yaml:"description"`
}

var memWH = make(map[string]Warehouse)
var memWHReg = make(map[int64][]string) // index region_id -> warehouses
var muWH   sync.RWMutex

func WarehouseCount() int64 {
  return int64(len(memWH))
}

func WarehouseAppend(info *Warehouse) {
  muWH.Lock()
  memWH[info.CODE] = *info
  if _, ok := memWHReg[info.Region_ID]; !ok {
    memWHReg[info.Region_ID] = make([]string, 1)
  }
  memWHReg[info.Region_ID] = append(memWHReg[info.Region_ID], info.CODE)
  muWH.Unlock()
}

func GetWarehouseByCode(code string) (*Warehouse) {
  item, ok := memWH[code]
  if ok {
    return &item
  }
  return nil
}

func GetWarehousesByRegionID(region_id int64) ([]string) {
  muWH.RLock()
  items, ok := memWHReg[region_id]
  muWH.RUnlock()
  if ok {
    return items
  }
  return make([]string, 0)
}
