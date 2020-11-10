package main

import (
  "sync"
  "encoding/gob"
  "os"
  "github.com/golang/glog"
)

type Region struct {
  ID           int64     `db:"id"             json:"id"              yaml:"id"`
  Name         string    `db:"title"          json:"title"          yaml:"title"`
}

/*
type District struct {
  r   []Region `json:"regions"`
}

type Districts struct {
  d    []District `json:"federalDistricts"`
}

type Regions struct {
}
*/

var memRegion = make(map[int64]Region)

func RegionCount() int64 {
  return int64(len(memRegion))
}

func RegionAppend(info *Region) {
  memRegion[info.ID] = *info
}

func GetRegionByID(id int64) (*Region) {
  item, ok := memRegion[id]
  if ok {
    return &item
  }
  return nil
}

func loadRegions(filename string) {
  
}

func WriteFileRegions(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, _ := os.Create(filename)
  defer file.Close()
  encoder := gob.NewEncoder(file)
  encoder.Encode(memRegion)
}

func LoadFileRegions(wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  file, err := os.Open(filename)
  if err !=nil {
    glog.Errorf("ERR: Load(%s): %v", filename, err)
    return
  }
  defer file.Close()
  
  decoder := gob.NewDecoder(file)
  err = decoder.Decode(&memRegion)
  if err != nil {
    glog.Errorf("ERR: Decoder(%s): %v", filename, err)
    return
  }
}
