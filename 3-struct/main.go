package main

import (
	"fmt"
	"time"
)

// Bin представляет один бин (хранилище) в облаке
type Bin struct {
    ID        string    `json:"id"`
    Private   bool      `json:"private"`
    CreatedAt time.Time `json:"createdAt"`
    Name      string    `json:"name"`
}

// BinList — список бинов
type BinList struct {
    Bins []Bin `json:"bins"`
}

// NewBin создаёт новый бин с указанными параметрами
func NewBin(id, name string, private bool) Bin {
    return Bin{
        ID:        id,
        Private:   private,
        CreatedAt: time.Now(), // текущая дата и время
        Name:      name,
    }
}

// NewEmptyBin создаёт пустой бин с нулевыми значениями
func NewEmptyBin() Bin {
    return Bin{}
}

// NewBinList создаёт новый пустой список бинов
func NewBinList() BinList {
    return BinList{
        Bins: make([]Bin, 0),
    }
}

// AddBin добавляет бин в список
func (bl *BinList) AddBin(bin Bin) {
    bl.Bins = append(bl.Bins, bin)
}