package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}
	_, bill_exists := bill[item]
	if bill_exists {
		bill[item] += value
	} else {
		bill[item] = value
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	bill_value, bill_exists := bill[item]
	if !bill_exists {
		return false
	}
	value, unit_exists := units[unit]
	if !unit_exists {
		return false
	}

	if bill_value-value < 0 {
		return false
	} else if bill_value-value == 0 {
		delete(bill, item)
		return true
	}

	bill[item] -= value
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	bill_value, bill_exists := bill[item]
	if !bill_exists {
		return 0, false
	}
	return bill_value, true

}
