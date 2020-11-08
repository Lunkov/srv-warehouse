package main

import (
  "sync"
)

// Товары

type Goods struct {
  ID           int64     `db:"id"             json:"id"              yaml:"id"`
  Name         string    `db:"name"           json:"name"            yaml:"name"`
  Description  string    `db:"description"    json:"description"     yaml:"description"`
}

var memG = make(map[int64]Goods)
var muG   sync.RWMutex

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
