package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    bestOption := option1

    if option2 < option1 {
        bestOption = option2
    }

    return bestOption + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    cost := float64(originalPrice) * 0.8

    if age >= 3 && age < 10 {
        cost = float64(originalPrice) * 0.7
    }

    if age >= 10 {
        cost = float64(originalPrice) * 0.5
    }

    return cost
}
