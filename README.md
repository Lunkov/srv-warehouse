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
