package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *iglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *iglooBuilder) setNumFloor() {
	i.floor = 3
}

func (i *iglooBuilder) getHouse() house {
	return house{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
