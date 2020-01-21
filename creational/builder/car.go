package builder

//CarBuilder represents object used to build the car
type CarBuilder struct {
	v VehicleProduct
}

//SetWheels add respective wheels to the car
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

//SetSeats add seats to the car
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

//SetStructure comes up with the body form of a car
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car Structure"
	return c
}

//Build combines all the processes to come up with the finished car
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}
