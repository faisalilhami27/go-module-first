package go_module_first

type FilterName func(string) bool

func SayHello(name string, filterName FilterName) string {
	var status string
	if filterName(name) {
		status = "You are blocked"
	} else {
		status = "Hello " + name
	}
	return status
}
