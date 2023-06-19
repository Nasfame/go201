package main


// +gen slice:"Aggregate[string]"
type Employee struct{
	Name 	   string
	Department string
}


func main(){
	employees := EmployeeSlice {
		{"Alice", "Accounting"},
		{"Bob", "Back Office"},
		{"Carly", "Containers"},
	}
	
	join := func(state string, e Employee) string {
		if state != "" {
			state += ", "
		}
		return state + e.Name
	}
	
	employees.AggregateString(join)
}