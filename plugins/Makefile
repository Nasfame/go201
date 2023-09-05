build-eng:
	go build -buildmode=plugin -o eng/eng.so eng/greeter.go

build-chi:
	go build -buildmode=plugin -o chi/chi.so chi/greeter.go

build-ru:
	go build -buildmode=plugin -o ru/ru.so ru/greeter.go

build-chi:
	go build -buildmode=plugin -o swe/swe.so swe/greeter.go

run: 
	go run greeter.go

# Currently supports only Unix, MacOs, Linux.. not windows