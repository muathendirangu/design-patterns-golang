package builder

//BuildProcess defines the contract of building
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

//VehicleProduct structure
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//ManufacturingDirector is incharge of operations
type ManufacturingDirector struct {
	builder BuildProcess
}

//Construct the vehicle
func (m *ManufacturingDirector) Construct() {
	m.builder.SetStructure().SetWheels().SetSeats().Build()
}

//SetBuilder sets the methodology of building the product
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}
