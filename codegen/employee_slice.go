// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Employee

package main

// EmployeeSlice is a slice of type Employee. Use it where you would use []Employee.
type EmployeeSlice []Employee

// AggregateString iterates over EmployeeSlice, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (rcv EmployeeSlice) AggregateString(fn func(string, Employee) string) (result string) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}
