package main

func main() {
	mainBad2()
}

func mainBad() {
	go func() {
		print("in go routine 1\n")
		for {

		}
	}()
	//uncomment me to see the print above in the output
	//runtime.Gosched()
	go func() {
		print("in go routine 2\n")
		for {

		}
	}()
	for {
	}
}

func mainBad2() {
	testme(func() {
		print("in go routine 1\n")
		for {

		}
	})
	testme(func() {
		print("in go routine 2\n")
		for {

		}
	})
	for {
	}

}
func testme(fn func()) {
	go fn()
	//runtime.Gosched()
}
