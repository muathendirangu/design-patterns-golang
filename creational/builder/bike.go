package builder

//BikeBuilder represents object used to build the bike
type BikeBuilder struct {
	v VehicleProduct
}

//SetWheels add respective wheels to the bike
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

//SetSeats add seats to the bike
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

//SetStructure comes up with the body form of a bike
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

//Build combines all the processes to come up with the finished bike
func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}
