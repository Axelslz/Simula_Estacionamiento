package scenes

import (
	"parking/models"
	"sync"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/render"
)

var (
	mutex sync.Mutex
)

type MainScene struct {
}

func NewParkingScene() *MainScene {
	return &MainScene{}
}

func runParking(ctx *scene.Context, handler *models.ManejoCarro, parking *models.Parking) {
	event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
		for {
			car := models.NewCar(ctx)
			go car.RunCar(&mutex, handler, parking)
			models.RandomSleep(300)
		}
	})
}

func (ps *MainScene) Start() {
    oak.AddScene("mainScene", scene.Scene{
        Start: func(ctx *scene.Context) {
            handler := models.NewManejoCarro()
            parking := models.NewParking(ctx)

            // Carga la imagen del letrero "Stop"
            stopSign, err := render.LoadSprite("assets/stop.png")
            if err != nil {
                panic(err)
            }

            // Ajusta la posición del letrero
            stopSign.SetPos(450, 100) // Ajusta las coordenadas según sea necesario

            // Agrega el letrero "Stop" al contexto
            ctx.Draw(stopSign)

            // Inicia la rutina para manejar los autos
            go runParking(ctx, handler, parking)
        },
    })
    oak.Init("mainScene")
}

