package main

import (
  "fmt"
  "math/rand"
  "unsafe"
  
  "github.com/golang/glog"
)

// Заполняем тестовыми данными
func fillData4Tests() {
  var i, cWH, cSP, cPR, cPRWH, cPRSP, cRegions int64

  cRegions =   200 // Количество регионов
  cWH   =    10000 // Количество складов
  cSP   =     5000 // Количество магазинов
  cPR   =   200000 // Количество товаров
  cPRWH =  cPR * cWH /100 // Количество товаров на складах
  cPRSP =  cPR * cSP /100  // Количество товаров в магазинах

  glog.Infof("LOG: Regions:        %d", RegionCount())
  glog.Infof("LOG: Regions: sizeof=%d", unsafe.Sizeof(Region{}))
  glog.Infof("LOG: Warehouses:     %d", WarehouseCount())
  glog.Infof("LOG: Warehouses: sizeof=%d", unsafe.Sizeof(Warehouse{}))
  glog.Infof("LOG: Shops:          %d", ShopCount())
  glog.Infof("LOG: Goods:          %d", GoodsCount())

  glog.Infof("LOG: Goods in WH:     %d", cPRWH)
  glog.Infof("LOG: Goods in WH: sizeof=%d", unsafe.Sizeof(WarehouseGoods{}))
  glog.Infof("LOG: Goods in WHL:sizeof=%d", unsafe.Sizeof(WarehouseGoodsLite{}))
  glog.Infof("LOG: Goods in Shops:  %d", cPRSP)
  glog.Infof("LOG: Goods in Shops: sizeof=%d", unsafe.Sizeof(ShopGoods{}))
  
  for i = 0; i <= cRegions; i++ {
    RegionAppend(&Region{ID: i, Name: fmt.Sprintf("Name_WH_%d", i) })
  }
  
  for i = 0; i <= cWH; i++ {
    WarehouseAppend(&Warehouse{CODE: fmt.Sprintf("a%d", i), Name: fmt.Sprintf("Name_WH_%d", i), Region_ID: rand.Int63n(RegionCount()) })
  }
  
  for i = 0; i <= cSP; i++ {
    ShopAppend(&Shop{CODE: fmt.Sprintf("s%d", i), Name: fmt.Sprintf("SHOP_%d", i), Region_ID: rand.Int63n(RegionCount()) })
  }

  for i = 0; i <= cPR; i++ {
    GoodsAppend(&Goods{ID: i, Name: fmt.Sprintf("Product_%d", i) })
  }
  //  Случайно раскидываем товары по магазинам
  for i = 0; i <= cPRWH; i++ {
    WarehouseGoodsAppend(&WarehouseGoods{Warehouse_ID: fmt.Sprintf("a%d", rand.Int63n(WarehouseCount())), Goods_ID: rand.Int63n(GoodsCount()), Cost: 1000+rand.Intn(10000), Quantity: rand.Intn(100)})
  }
  // Случайно раскидываем товары по складам
  for i = 0; i <= cPRSP; i++ {
    ShopGoodsAppend(&ShopGoods{Shop_ID: fmt.Sprintf("s%d", rand.Intn(ShopCount())), Goods_ID: rand.Int63n(GoodsCount()), Cost: 1000+rand.Intn(10000), Quantity: rand.Intn(100)})
  }
  
  glog.Infof("LOG: Regions:        %d", RegionCount())
  glog.Infof("LOG: Warehouses:     %d", WarehouseCount())
  glog.Infof("LOG: Shops:          %d", ShopCount())
  glog.Infof("LOG: Goods:          %d", GoodsCount())

  glog.Infof("LOG: Goods in WH:     %d", cPRWH)
  glog.Infof("LOG: Goods in Shops:  %d", cPRSP)
  
}
