# Before tests
```
sudo apt-get install libcanberra-gtk-module
sudo apt-get install libcanberra-gtk-module libcanberra-gtk3-module
sudo apt-get install graphviz
```

# Tests
About tools: https://blog.golang.org/pprof
```
go test -bench=. -benchmem -benchtime=10s -cpuprofile=cpu.out -memprofile=mem.out
```

```
go tool pprof ./mem.out
go tool pprof ./cpu.out
```

```
top10
web mallocgc
```
# Tests results

Intel® Core™ i5-4210U CPU @ 1.70GHz × 4

* cRegions =  1000 // Количество регионов
* cWH   =    10000 // Количество складов
* cSP   =     5000 // Количество магазинов
* cPR   =   200000 // Количество товаров
* cPRWH =  5000000 // Количество товаров на складах
* cPRSP =  1000000 // Количество товаров в магазинах
```  
      flat  flat%   sum%        cum   cum%
 8212.20MB 76.55% 76.55%  8212.20MB 76.55%  _/home/sergey/src/GO/lib-warehouse.WarehouseGoodsAppend
 1640.41MB 15.29% 91.84%  1640.41MB 15.29%  _/home/sergey/src/GO/lib-warehouse.ShopGoodsAppend
  649.01MB  6.05% 97.89% 10663.17MB 99.40%  _/home/sergey/src/GO/lib-warehouse.fillData4Tests
     132MB  1.23% 99.12%      132MB  1.23%  fmt.Sprintf
```

## 300 regions
```
go test -bench=. -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial       	    3416	    346176 ns/op	     461 B/op	      10 allocs/op
BenchmarkWHSerial-2     	    3435	    545280 ns/op	     486 B/op	      10 allocs/op
BenchmarkWHSerial-4     	    3430	    370342 ns/op	     506 B/op	      10 allocs/op
BenchmarkWHSerial-8     	    2211	    469228 ns/op	     542 B/op	      11 allocs/op
BenchmarkWHParallel     	    3334	    461272 ns/op	     574 B/op	      11 allocs/op
```
## 1000 regions
```
go test -bench=. -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial       	    3138	    402735 ns/op	     448 B/op	      10 allocs/op
BenchmarkWHSerial-2     	    3075	    391101 ns/op	     453 B/op	      10 allocs/op
BenchmarkWHSerial-4     	    3067	    392282 ns/op	     460 B/op	      10 allocs/op
BenchmarkWHSerial-8     	    3020	    452692 ns/op	     467 B/op	      10 allocs/op
BenchmarkWHParallel     	    2947	    516536 ns/op	     476 B/op	      10 allocs/op
```
## 1000 regions + Mutex
```
go test -bench=. -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial       	    3513	    339993 ns/op	    1928 B/op	      40 allocs/op
BenchmarkWHSerial-2     	    3512	    344325 ns/op	    1935 B/op	      41 allocs/op
BenchmarkWHSerial-4     	    3516	    341486 ns/op	    1934 B/op	      41 allocs/op
BenchmarkWHSerial-8     	    3523	    526883 ns/op	    1935 B/op	      40 allocs/op
BenchmarkWHParallel     	    3423	    341600 ns/op	    1946 B/op	      41 allocs/op
```

## 300 regions + Mutex + Region Index
```
go test -bench=. -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial       	   26665	     66836 ns/op	   19913 B/op	     414 allocs/op
BenchmarkWHSerial-2     	   10000	    110450 ns/op	   34568 B/op	     712 allocs/op
BenchmarkWHSerial-4     	    7467	    168922 ns/op	   49552 B/op	    1009 allocs/op
BenchmarkWHSerial-8     	    6375	    203547 ns/op	   64669 B/op	    1304 allocs/op
BenchmarkWHParallel     	    4825	    274171 ns/op	   80082 B/op	    1598 allocs/op
```

## 1000 regions + Mutex + Region Index
```
time go test -bench=. -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial       	   64258	     25393 ns/op	    6446 B/op	     134 allocs/op
BenchmarkWHSerial-2     	   30985	     48734 ns/op	   12303 B/op	     255 allocs/op
BenchmarkWHSerial-4     	   20785	     63798 ns/op	   18284 B/op	     374 allocs/op
BenchmarkWHSerial-8     	   10000	    106853 ns/op	   22862 B/op	     463 allocs/op
BenchmarkWHParallel     	   65337	     24672 ns/op	    6446 B/op	     134 allocs/op
BenchmarkWHParallel-2   	   53988	     29460 ns/op	   12303 B/op	     255 allocs/op
BenchmarkWHParallel-4   	   46034	     28187 ns/op	   18286 B/op	     374 allocs/op
BenchmarkWHParallel-8   	   34252	     41041 ns/op	   24408 B/op	     493 allocs/op
```
## 1000 regions - Mutex + Region Index
```
go test -bench=. -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial       	  100926	     14739 ns/op	     480 B/op	      10 allocs/op
BenchmarkWHSerial-2     	   48650	     28738 ns/op	     642 B/op	      12 allocs/op
BenchmarkWHSerial-4     	   31650	     40957 ns/op	     985 B/op	      14 allocs/op
BenchmarkWHSerial-8     	   23517	     56914 ns/op	    1516 B/op	      16 allocs/op
BenchmarkWHParallel     	  102074	     14903 ns/op	     479 B/op	      10 allocs/op
BenchmarkWHParallel-2   	   84792	     15833 ns/op	     642 B/op	      12 allocs/op
BenchmarkWHParallel-4   	   83028	     15226 ns/op	     987 B/op	      14 allocs/op
BenchmarkWHParallel-8   	   48744	     20605 ns/op	    1513 B/op	      16 allocs/op
```
```
Showing nodes accounting for 22.18GB, 98.94% of 22.42GB total
Dropped 44 nodes (cum <= 0.11GB)
Showing top 10 nodes out of 12
      flat  flat%   sum%        cum   cum%
   14.51GB 64.73% 64.73%    14.51GB 64.73%  _/home/sergey/src/GO/srv-warehouse.WarehouseGoodsAppend
    3.26GB 14.56% 79.29%     3.26GB 14.56%  _/home/sergey/src/GO/srv-warehouse.ShopGoodsAppend
    1.94GB  8.66% 87.94%     1.94GB  8.66%  _/home/sergey/src/GO/srv-warehouse.RlGetGoodsInWarehouse
    1.39GB  6.21% 94.16%    19.32GB 86.16%  _/home/sergey/src/GO/srv-warehouse.fillData4Tests
    0.97GB  4.31% 98.47%     0.97GB  4.31%  _/home/sergey/src/GO/srv-warehouse.RlGetGoodsInShop
    0.08GB  0.37% 98.84%     3.01GB 13.41%  _/home/sergey/src/GO/srv-warehouse.GetGoods4Sale
    0.02GB 0.098% 98.94%     3.10GB 13.82%  _/home/sergey/src/GO/srv-warehouse.BenchmarkWHParallel.func1
         0     0% 98.94%    19.32GB 86.16%  _/home/sergey/src/GO/srv-warehouse.BenchmarkWHParallel
```
## Write & Read


```
        200 - Количество регионов
     10,000 - Количество складов
      5,000 - Количество магазинов
    200,000 - Количество товаров
 20,000,000 - Количество товаров на складах
 10,000,000 - Количество товаров в магазинах
```
  
Size of files:
```
131,485,953 cnt_shops.gob          - Товары в магазинах
264,228,381 cnt_warehouses.gob     - Товары на складах
  5,023,066 goods.gob              - Магазины
     3,778  regions.gob            - Регионы
   136,894  shops.gob              - Магазины
    29,746  shops.gob.index        - Индекс магазинов по регионам
   307,227  warehouses.gob         - Склады
    59,747  warehouses.gob.index   - Индек складов по регионам
```
### Serial
```
$go test -bench=BenchmarkWHSerial -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHSerial     	   64107	     17903 ns/op	    3000 B/op	     164 allocs/op
BenchmarkWHSerial-2   	   44845	     25972 ns/op	    2999 B/op	     164 allocs/op
BenchmarkWHSerial-4   	   45054	     26308 ns/op	    3002 B/op	     164 allocs/op
BenchmarkWHSerial-8   	   46059	     25989 ns/op	    3000 B/op	     164 allocs/op
```
### Parallel
```
$go test -bench=BenchmarkWHParallel -cpu 1,2,4,8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWHParallel/1   	   63542	     17169 ns/op	    2998 B/op	     164 allocs/op
BenchmarkWHParallel/1-2 	  104888	     11037 ns/op	    3000 B/op	     164 allocs/op
BenchmarkWHParallel/1-4 	  119359	      9735 ns/op	    2998 B/op	     164 allocs/op
BenchmarkWHParallel/1-8 	  122959	      9505 ns/op	    2999 B/op	     164 allocs/op
BenchmarkWHParallel/2   	   67255	     16981 ns/op	    3001 B/op	     164 allocs/op
BenchmarkWHParallel/2-2 	  101695	     10973 ns/op	    2998 B/op	     164 allocs/op
BenchmarkWHParallel/2-4 	  124570	      9380 ns/op	    3000 B/op	     164 allocs/op
BenchmarkWHParallel/2-8 	  118636	     11886 ns/op	    2998 B/op	     164 allocs/op
BenchmarkWHParallel/4   	   66622	     17146 ns/op	    3001 B/op	     164 allocs/op
BenchmarkWHParallel/4-2 	  104336	     11100 ns/op	    2998 B/op	     164 allocs/op
BenchmarkWHParallel/4-4 	  123817	     12366 ns/op	    2999 B/op	     164 allocs/op
BenchmarkWHParallel/4-8 	  100856	     14142 ns/op	    2999 B/op	     164 allocs/op
BenchmarkWHParallel/8   	   66759	     18963 ns/op	    2999 B/op	     164 allocs/op
BenchmarkWHParallel/8-2 	  107932	     11261 ns/op	    2998 B/op	     164 allocs/op
BenchmarkWHParallel/8-4 	  122364	     11716 ns/op	    2999 B/op	     164 allocs/op
BenchmarkWHParallel/8-8 	   99258	     22838 ns/op	    3000 B/op	     164 allocs/op
```
### Write
After random
```
$go test -bench=BenchmarkWrite -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWrite-4   	       1	18379591667 ns/op	4096934688 B/op	61332546 allocs/op

Write ~18s
```
After load
```
go test -bench=BenchmarkWrite -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkWrite-4   	       1	12592833076 ns/op	4096895096 B/op	61332484 allocs/op

Write ~13s
Showing nodes accounting for 6.28GB, 99.01% of 6.35GB total
Dropped 37 nodes (cum <= 0.03GB)
```
### Read
```
$go test -bench=BenchmarkRead -benchmem -cpuprofile=cpu.out -memprofile=mem.out
goos: linux
goarch: amd64
BenchmarkRead-4   	       1	11578677350 ns/op	2707435384 B/op	33914314 allocs/op

Read ~11s
```
