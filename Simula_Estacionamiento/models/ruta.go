package models

type Ruta struct {
	ruta string
	spot  float64
}

func newRuta(ruta string, spot float64) *Ruta {
	return &Ruta{
		ruta: ruta,
		spot:  spot,
	}
}
