package main

import (
)

// Склады

type Warehouse struct {
  CODE         string    `db:"code"           json:"code"            yaml:"code"`
  Region_ID    int64  
  Name         string    `db:"name"           json:"name"            yaml:"name"`
  Description  string    `db:"description"    json:"description"     yaml:"description"`
}

var memWH = make(map[string]Warehouse)

func WarehouseCount() int64 {
  return int64(len(memWH))
}

func WarehouseAppend(info *Warehouse) {
  memWH[info.CODE] = *info
}

func GetWarehouseByCode(code string) (*Warehouse) {
  item, ok := memWH[code]
  if ok {
    return &item
  }
  return nil
}
