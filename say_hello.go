package go_module_first

type FilterString func(string, string) bool

func FilterName(firstName, lastName string, filterString FilterString) string {
	var status string
	if filterString(firstName, lastName) {
		status = "You are blocked"
	} else {
		status = "Hello my name is " + firstName + " " + lastName
	}
	return status
}
