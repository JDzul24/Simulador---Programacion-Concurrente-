package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Vehiculo struct {
	id              int
	tiempoLim       time.Duration
	espacioAsignado int
	imagenEntrada   *canvas.Image
	imagenSalida    *canvas.Image
}

func NuevoVehiculo(id int) *Vehiculo {
	imagenEntrada := canvas.NewImageFromURI(storage.NewFileURI("./assets/carEntra.png"))
	imagenSalida := canvas.NewImageFromURI(storage.NewFileURI("./assets/carSalida.png"))

	return &Vehiculo{
		id:              id,
		tiempoLim:       time.Duration(rand.Intn(50)+50) * time.Second,
		espacioAsignado: 0,
		imagenEntrada:   imagenEntrada,
		imagenSalida:    imagenSalida,
	}
}

func (v *Vehiculo) Ingresar(p *Estacionamiento, contenedor *fyne.Container) {
	//ID del vehículo al canal de espacios en el estacionamiento.
	p.ObtenerEspacio() <- v.ObtenerID()
	// Adquiere el mutex
	p.ObtenerPuerta().Lock()
	//Estado actual de ocupación de espacios en el estacionamiento.
	espacios := p.ObtenerArrayEspacios()
	const (
		columnasPorGrupo  = 10
		espacioHorizontal = 57
		espacioVertical   = 320
	)
	// Iterar sobre los espacios para encontrar uno disponible.
	for i := 0; i < len(espacios); i++ {
		if !espacios[i] {
			// Marca el espacio como ocupado y asignar el número de espacio al vehículo.
			espacios[i] = true
			v.espacioAsignado = i
			// Calcula la posición en la interfaz según la fila y columna del espacio.
			fila := i / (columnasPorGrupo * 1)
			columna := i % (columnasPorGrupo * 1)
			if columna >= columnasPorGrupo {
				columna += 1
			}
			x := float32(133 + columna*espacioHorizontal)
			y := float32(10 + fila*espacioVertical)
			// Mover la imagen del vehículo a la posición
			v.imagenEntrada.Move(fyne.NewPos(x, y))
			break
		}
	}
	// Actualizar el estado de ocupación de espacios en el estacionamiento.
	p.EstablecerArrayEspacios(espacios)
	// Liberar el mutex
	p.ObtenerPuerta().Unlock()
	contenedor.Refresh()
	fmt.Printf("Auto %d ocupó el lugar %d.\n", v.ObtenerID(), v.espacioAsignado)
	//Simular el tiempo que el vehículo permanece estacionado.
	time.Sleep(5 * time.Second)
}

func (v *Vehiculo) Salir(p *Estacionamiento, contenedor *fyne.Container) {
	//Recibe el espacio asignado al vehículo
	<-p.ObtenerEspacio()
	//Adquiere el mutex de la puerta
	p.ObtenerPuerta().Lock()
	espacios := p.ObtenerArrayEspacios()
	//Marca el espacio asignado como disponible.
	espacios[v.espacioAsignado] = false
	fmt.Printf("Auto %d salió. Espacio %d marcado como disponible.\n", v.ObtenerID(), v.espacioAsignado)
	p.EstablecerArrayEspacios(espacios)
	p.ObtenerPuerta().Unlock() //libera
	contenedor.Remove(v.imagenEntrada)
	contenedor.Refresh()
	v.imagenSalida.Resize(fyne.NewSize(30, 50))
	v.imagenSalida.Move(fyne.NewPos(-30, 290))
	contenedor.Add(v.imagenSalida)
	contenedor.Refresh()
	for i := 0; i < 10; i++ {
		v.imagenSalida.Move(fyne.NewPos(v.imagenSalida.Position().X, v.imagenSalida.Position().Y-20))
		time.Sleep(time.Millisecond * 100)
	}
	contenedor.Remove(v.imagenSalida)
	contenedor.Refresh()
}

func (v *Vehiculo) Iniciar(p *Estacionamiento, contenedor *fyne.Container, wg *sync.WaitGroup) {
	v.Avanzar(10)
	v.Ingresar(p, contenedor)
	//Simulando la estancia del vehículo en el estacionamiento.
	time.Sleep(5 * time.Second)
	timer := time.NewTimer(5 * time.Second)
	//Cuando el temporizador expire.
	select {
	case <-timer.C:
		contenedor.Remove(v.imagenEntrada)
		contenedor.Refresh()
		contenedor.Add(v.imagenSalida)
		contenedor.Refresh()
		p.ColaSalida(contenedor, v.imagenEntrada)
		v.Salir(p, contenedor)
		//Secrementar el contador de espera del WaitGroup(rutina termina).
		wg.Done()
	}
}

// Simula el avance del vehículo.
func (v *Vehiculo) Avanzar(pasos int) {
	for i := 0; i < pasos; i++ {
		v.imagenEntrada.Move(fyne.NewPos(v.imagenEntrada.Position().X, v.imagenEntrada.Position().Y+20))
		time.Sleep(time.Millisecond * 100)
	}
}

func (v *Vehiculo) ObtenerID() int {
	return v.id
}

func (v *Vehiculo) ObtenerTiempoLim() time.Duration {
	return v.tiempoLim
}

func (v *Vehiculo) ObtenerImagenEntrada() *canvas.Image {
	return v.imagenEntrada
}
