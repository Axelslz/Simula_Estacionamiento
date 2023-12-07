package models

import "sync"

type ManejoCarro struct {
	Cars  []*Car
	Mutex sync.Mutex
}

func NewManejoCarro() *ManejoCarro {
	return &ManejoCarro{
		Cars: make([]*Car, 0),
	}
}

func (manejoCarro *ManejoCarro) Add(car *Car) {
	manejoCarro.Mutex.Lock()
	defer manejoCarro.Mutex.Unlock()
	manejoCarro.Cars = append(manejoCarro.Cars, car)
}

func (manejoCarro *ManejoCarro) Remove(car *Car) {
	manejoCarro.Mutex.Lock()
	defer manejoCarro.Mutex.Unlock()
	for i, c := range manejoCarro.Cars {
		if c == car {
			manejoCarro.Cars = append(manejoCarro.Cars[:i], manejoCarro.Cars[i+1:]...)
			break
		}
	}
}

func (manejoCarro *ManejoCarro) GetCars() []*Car {
	manejoCarro.Mutex.Lock()
	defer manejoCarro.Mutex.Unlock()
	return manejoCarro.Cars
}
