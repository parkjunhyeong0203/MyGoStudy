package ex192

import (
	"time"
)

type Courier struct {
	Name string
}
type Product struct {
	Name  string
	Price int
	ID    int
}
type Parcel struct {
	Pdt          *Product
	ShippedTime  time.Time
	DeliverdTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	pc := &Parcel{Pdt: p, ShippedTime: time.Now()}
	return pc
}
func (pc *Parcel) Deliverd() *Product {
	pc.DeliverdTime = time.Now()
	return pc.Pdt
}
