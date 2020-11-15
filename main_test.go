package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "testing"
  "github.com/stretchr/testify/assert"
)

func BenchmarkWrite(b *testing.B) {
  fillData4Tests()
  //LoadAll()
  b.ResetTimer()
  SaveAll()
}

func BenchmarkRead(b *testing.B) {
  LoadAll()
}

func BenchmarkWHSerial(b *testing.B) {
  //fillData4Tests()
  LoadAll()
  
  //assert.Equal(b, &Warehouse{CODE:"a4", Name:"Name_WH_4", Description:""},   GetWarehouseByCode("a4"))
  //assert.Equal(b, &Warehouse{CODE:"a23", Name:"Name_WH_23", Description:""}, GetWarehouseByCode("a23"))

  //assert.Equal(b, &Shop{CODE:"s1",   Name:"SHOP_1",   Description:""}, GetShopByID("s1"))
  //assert.Equal(b, &Shop{CODE:"s543", Name:"SHOP_543", Description:""}, GetShopByID("s543"))

  assert.Equal(b, &Goods{ID:1,      Name:"Product_1", Description:""}, GetGoodsByID(1))
  assert.Equal(b, &Goods{ID:543,    Name:"Product_543", Description:""}, GetGoodsByID(543))
  id := GoodsCount() - 2
  assert.Equal(b, &Goods{ID:id, Name:fmt.Sprintf("Product_%d", id), Description:""}, GetGoodsByID(id))
  
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    goods_id := rand.Int63n(GoodsCount())
    p := GetGoods4Sale(goods_id, rand.Int63n(RegionCount()))
    
    assert.Equal(b, GetGoodsByID(goods_id), &p.Prod)
    //assert.Equal(b, Goods4Sale{Prod:Goods{ID:i, Name:fmt.Sprintf("Product_%d", i), Description:""}, Shops:[]Shop(nil), Warehouses:[]Warehouse(nil)}, p)
  }

}


func BenchmarkWHParallel(b *testing.B) {
  //fillData4Tests()
  LoadAll()
  
  assert.Equal(b, &Goods{ID:1,      Name:"Product_1", Description:""}, GetGoodsByID(1))
  assert.Equal(b, &Goods{ID:543,    Name:"Product_543", Description:""}, GetGoodsByID(543))
  id := GoodsCount() - 2
  assert.Equal(b, &Goods{ID:id, Name:fmt.Sprintf("Product_%d", id), Description:""}, GetGoodsByID(id))
  
  b.ResetTimer()
  for i := 1; i <= 8; i *= 2 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.SetParallelism(i)
      b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
          goods_id := rand.Int63n(GoodsCount())
          p := GetGoods4Sale(goods_id, rand.Int63n(RegionCount()))
          
          assert.Equal(b, GetGoodsByID(goods_id), &p.Prod)
        }
      })
    })
  }
}
