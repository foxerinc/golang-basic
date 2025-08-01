package tutorial

import "testing"

func TestParkingLand_Motorcycle_1Hour(t *testing.T) {
	got := ParkingLand("motorcycle", 1)
	want := 3000

	if got != want {
		t.Errorf("ParkingLand(motorcycle, 1) = %d, want %d", got, want)
	}
}

func TestParkingLand_Motorcycle_2Hours(t *testing.T) {
	got := ParkingLand("motorcycle", 2)
	// 3,000 + 2,000 = 5,000
	want := 5000

	if got != want {
		t.Errorf("ParkingLand(motorcycle, 2) = %d, want %d", got, want)
	}
}

func TestParkingLand_Car_1Hour(t *testing.T) {
	got := ParkingLand("car", 1)
	want := 7000

	if got != want {
		t.Errorf("ParkingLand(car, 1) = %d, want %d", got, want)
	}
}

func TestParkingLand_Car_5Hours(t *testing.T) {
	got := ParkingLand("car", 5)
	// 7,000 + 4 * 5,000 = 27,000
	want := 27000

	if got != want {
		t.Errorf("ParkingLand(car, 5) = %d, want %d", got, want)
	}
}

func TestParkingLand_Motorcycle_24Hours(t *testing.T) {
	got := ParkingLand("motorcycle", 24)
	// No extra charge because not > 24
	// 3,000 + 23 * 2,000 = 49,000
	want := 49000

	if got != want {
		t.Errorf("ParkingLand(motorcycle, 24) = %d, want %d", got, want)
	}
}

func TestParkingLand_Car_25Hours(t *testing.T) {
	got := ParkingLand("car", 25)
	// From problem statement:
	// 7,000 + 24 * 5,000 + 50,000 = 177,000
	want := 177000

	if got != want {
		t.Errorf("ParkingLand(car, 25) = %d, want %d", got, want)
	}
}

func TestParkingLand_Motorcycle_25Hours(t *testing.T) {
	got := ParkingLand("motorcycle", 25)
	// 3,000 + 24 * 2,000 + 20,000 = 71,000
	want := 71000

	if got != want {
		t.Errorf("ParkingLand(motorcycle, 25) = %d, want %d", got, want)
	}
}

func TestParkingLand_UnknownType_TreatedAsCar(t *testing.T) {
	got := ParkingLand("truck", 1)
	// In current checkType, non-"motorcycle" uses car pricing
	want := 7000

	if got != want {
		t.Errorf("ParkingLand(truck, 1) = %d, want %d (treated as car)", got, want)
	}
}
