package tutorial

import "fmt"



func checkType (typeVehicle string) (int,int,int) {
	
	if (typeVehicle == "motorcycle") {
		return 3000,2000,20000
	}else {
		return 7000,5000,50000
	}

}

func ParkingLand (typeVehicle string, parkDuration int) int {
	if (parkDuration < 2) {
		a,b,c := checkType(typeVehicle)
		b = b*0
		c = c*0
		return a+b+c
	}else if (parkDuration <25) {
		a,b,c := checkType(typeVehicle)
		b = b*(parkDuration-1)
		c = c*0
		return a+b+c
	}else{
		a,b,c := checkType(typeVehicle)
		b = b*(parkDuration-1)
		return a+b+c
	}
}

func main() {
	testCases := []struct {
		vehicle string
		hours   int
	}{
		{"motorcycle", 1},
		{"motorcycle", 2},
		{"motorcycle", 24},
		{"motorcycle", 25},
		{"car", 1},
		{"car", 5},
		{"car", 24},
		{"car", 25},
	}

	for _, tc := range testCases {
		fee := ParkingLand(tc.vehicle, tc.hours)
		fmt.Printf("Vehicle=%s, Hours=%d => Fee=%d\n", tc.vehicle, tc.hours, fee)
	}
}