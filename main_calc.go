package main

import (
)

type Goods4Sale struct {
  Prod         Goods          `json:"goods"`
  Quantity     int            `json:"quantity"`
  Shops        []Shop         `json:"shops"`
  Warehouses   []Warehouse    `json:"warehouses"`
}

// Основная функция: получить остатки по товару в зависимости от региона
func GetGoods4Sale(goods_id int64, region_id int64) Goods4Sale {
  var res Goods4Sale
  res.Quantity = 0
  p := GetGoodsByID(goods_id)
  if p == nil {
    return res
  }
  res.Prod = *p
  for _, shop_code := range GetShopByRegionID(region_id) {
    t := RlGetGoodsInShop(goods_id, shop_code)
    if t.Quantity > 0 {
      res.Quantity += t.Quantity
      res.Shops = append(res.Shops, (*GetShopByID(shop_code)))
    }
  }
  for _, warehouse_code := range GetWarehousesByRegionID(region_id) {
    t := RlGetGoodsInWarehouse(goods_id, warehouse_code)
    if t.Quantity > 0 {
      res.Quantity += t.Quantity
      res.Warehouses = append(res.Warehouses, (*GetWarehouseByCode(warehouse_code)))
    }
  }
  return res
}
