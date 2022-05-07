package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(bulderType string) iBuilder {
	if bulderType == "normal" {
		return &normalBuilder{}
	}

	if bulderType == "igloo" {
		return &iglooBuilder{}
	}

	return nil
}
