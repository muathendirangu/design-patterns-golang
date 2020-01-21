package builder

import "testing"

func TestBuilder(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingDirector.SetBuilder(carBuilder)
	manufacturingDirector.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("A car is supposed to have 4 wheels but it has %d ", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("A car has 5 seats normally but now it has %d ", car.Seats)
	}

	if car.Structure != "Car Structure" {
		t.Errorf("A car is supposed to have a Car structure but the current structure is %s", car.Structure)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingDirector.SetBuilder(bikeBuilder)
	manufacturingDirector.Construct()

	motorbike := bikeBuilder.Build()

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'motorbike' and was %s\n", motorbike.Structure)
	}

	busBuilder := &BusBuilder{}

	manufacturingDirector.SetBuilder(busBuilder)
	manufacturingDirector.Construct()

	bus := busBuilder.Build()

	if bus.Wheels != 10 {
		t.Errorf("Wheels on a bus must be 10 and they were %d\n", bus.Wheels)
	}

	if bus.Structure != "bus" {
		t.Errorf("Structure on a bus must be 'bus' and was %s\n", bus.Structure)
	}

	if bus.Seats != 62 {
		t.Errorf("Seats on a bus must be 62 and they were %d\n", bus.Seats)
	}

}
