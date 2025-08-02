# Lot Billing — Parking Fee Calculator

## Objectives
- Use conditional statements (`if`, `else if`, `else`).
- Work with basic function parameters and return values.
- Implement simple business rules in Go.

## Problem Description

You are tasked with creating a function to calculate parking fees for a vehicle based on:

- Vehicle type
- Parking duration in hours

The system supports two vehicle types:

- `motorcycle`
- `car` (and any other string is treated as car in the current implementation)

## Pricing Rules

### 1. Base Price (First 1 Hour)

| Duration   | Motorcycle | Car  |
|-----------|-----------:|-----:|
| First 1 h |      3,000 | 7,000 |

### 2. After the First Hour

| Duration after 1 h | Motorcycle    | Car         |
|--------------------|--------------:|------------:|
| Per hour           | 2,000 / hour  | 5,000 / hour|

### 3. Extra Charge for Parking > 24 Hours

If a vehicle is parked for **more than 24 hours** (`parkDuration > 24`), an extra fee is added:

| Vehicle Type | Extra Charge |
|--------------|-------------:|
| Motorcycle   |       20,000 |
| Car          |       50,000 |

### Summary of Rules

Let:
- `typeVehicle` be `"motorcycle"` or `"car"`.
- `parkDuration` be an integer number of hours.

Then:

1. If `parkDuration < 2`  
   Charge only the **first-hour price**.
2. If `2 <= parkDuration <= 24`  
   Charge:  
   `first_hour_price + (parkDuration - 1) * per_hour_price`
3. If `parkDuration > 24`  
   Charge:  
   `first_hour_price + (parkDuration - 1) * per_hour_price + extra_charge`

> Note: In the provided implementation, any `typeVehicle` other than `"motorcycle"` is handled using car pricing.

## Function Signature

Main function to implement:

```go
func ParkingLand(typeVehicle string, parkDuration int) int
