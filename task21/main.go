// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// USB-C интерфейс
type USBC interface {
	InsertIntoUSBCPort()
}

// Micro USB интерфейс
type MicroUSB interface {
	InsertIntoMicroUSBPort()
}

// Device с USB-C разъемом
type PhoneWithUSBC struct{}

func (p *PhoneWithUSBC) InsertIntoUSBCPort() {
	fmt.Println("USB-C connector is plugged into the phone.")
}

// Device с Micro USB разъемом
type PhoneWithMicroUSB struct{}

func (p *PhoneWithMicroUSB) InsertIntoMicroUSBPort() {
	fmt.Println("Micro USB connector is plugged into the phone.")
}

// Adapter позволяет использовать USB-C устройство через Micro USB интерфейс
type USBCToMicroUSBAdapter struct {
	device *PhoneWithUSBC
}

func (a *USBCToMicroUSBAdapter) InsertIntoMicroUSBPort() {
	fmt.Println("Adapter converts Micro USB to USB-C.")
	a.device.InsertIntoUSBCPort()
}

func main() {
	phoneWithMicroUSB := PhoneWithMicroUSB{}
	phoneWithMicroUSB.InsertIntoMicroUSBPort()

	phoneWithUSBC := PhoneWithUSBC{}
	phoneWithUSBC.InsertIntoUSBCPort()

	adpter := USBCToMicroUSBAdapter{
		device: &phoneWithUSBC,
	}
	adpter.InsertIntoMicroUSBPort()
}

// адаптер позволяет телефону с разъемом USB-C работать через
// зарядное устройство с разъемом Micro USB.
