go tool compile -o compile.o compile.go
./compile.o

go build -o x.out compile.go
./x.out