package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Elem struct {
	Id     uint32
	Weight uint32
}

func (e *Elem) reset() {
	e.Id, e.Weight = 0, 0
}

type ElemPool struct {
	elems       []Elem
	tail        int
	totalWeight uint32
}

func (e *ElemPool) Count() uint32 {
	return uint32(len(e.elems[:e.tail]))
}

func (e *ElemPool) IsExist(id uint32) bool {
	if e.tail == 0 {
		return false
	}
	for i := 0; i < e.tail; i++ {
		if e.elems[i].Id == id {
			return true
		}
	}
	return false
}

func (e *ElemPool) AddElem(id, weight uint32) bool {
	if e.IsExist(id) {
		return false
	}

	if e.tail < len(e.elems) {
		e.elems[e.tail] = Elem{Id: id, Weight: weight}
		e.tail++
	} else {
		e.elems = append(e.elems, Elem{Id: id, Weight: weight})
		e.tail++
	}
	e.totalWeight += weight
	return true
}

func (e *ElemPool) DelElem(id uint32) bool {
	if e.tail == 0 {
		return false
	}

	for i := 0; i < e.tail; i++ {
		if e.elems[i].Id == id {
			e.totalWeight -= e.elems[i].Weight
			e.elems[i] = e.elems[e.tail-1]
			e.elems[e.tail-1].reset()
			e.tail--
			break
		}
	}
	return true
}

func (e *ElemPool) GetIds() []uint32 {

	ids := make([]uint32, 0, e.tail)
	for i := 0; i < e.tail; i++ {
		ids = append(ids, e.elems[i].Id)
	}
	return ids
}

func (e *ElemPool) RandomElems(count uint32) ([]uint32, error) {

	if e.Count() <= count {
		return e.GetIds(), nil
	}

	ids := make([]uint32, 0, count)
	totalWeight := e.totalWeight

	for i := 0; i < int(count); i++ {
		rdm := uint32(rand.Int31n(int32(totalWeight)))
		for k := len(ids); k < e.tail; k++ {
			if rdm < e.elems[k].Weight {
				// 交换元素
				e.elems[len(ids)], e.elems[k] = e.elems[k], e.elems[len(ids)]
				totalWeight -= e.elems[len(ids)].Weight
				ids = append(ids, e.elems[len(ids)].Id)
				break
			}
			rdm -= e.elems[k].Weight
		}
	}
	return ids, nil
}

func main() {

	var pool = ElemPool{}
	pool.AddElem(1, 10)
	pool.AddElem(2, 20)
	pool.AddElem(3, 30)
	pool.AddElem(4, 10)
	pool.AddElem(5, 20)
	pool.AddElem(6, 30)
	pool.AddElem(7, 10)
	pool.AddElem(8, 20)
	pool.AddElem(9, 30)
	pool.AddElem(10, 30)
	pool.AddElem(11, 30)
	pool.DelElem(11)
	pool.DelElem(10)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := uint32(0); i <= pool.Count(); i++ {
		perks, err := pool.RandomElems(i)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(pool.GetIds(), perks)
	}

}
