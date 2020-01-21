package builder

//BusBuilder represents object used to build the bus
type BusBuilder struct {
	v VehicleProduct
}

//SetWheels add respective wheels to the bus
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 10
	return b
}

//SetSeats add seats to the bus
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 62
	return b
}

//SetStructure comes up with the body form of a bus
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "bus"
	return b
}

//Build combines all the processes to come up with the finished bus
func (b *BusBuilder) Build() VehicleProduct {
	return b.v
}
