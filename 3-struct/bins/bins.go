package bins

import "time"

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewEmptyBin() Bin {
	return Bin{}
}

func NewBinList() BinList {
	return BinList{Bins: make([]Bin, 0)}
}

func (bl *BinList) AddBin(bin Bin) {
	bl.Bins = append(bl.Bins, bin)
}