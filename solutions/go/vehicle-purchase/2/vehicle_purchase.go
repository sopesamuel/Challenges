package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
}
// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    car := option2
	if option1 < option2 {
       car = option1 
    }
    return car + " is clearly the better choice."
    
}
// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    price := originalPrice
    switch {
        case age < 3:
        price = 0.8 * originalPrice
        case age >= 10:
        price = 0.5 * originalPrice
        case age >= 3 && age < 10:
        price = 0.7 * originalPrice
    }
    return price
}