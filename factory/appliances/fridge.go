package appliances

type Fridge struct {
	typeName string
}

func (fr *Fridge) Start() {
	fr.typeName = " Fridge "
}

func (fr *Fridge) GetPurpose() string {
	return "i am a " + fr.typeName + " i cool stuff down!!"
}
