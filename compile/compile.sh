go tool compile -o compile.o compile.go
chmod +x compile.o
./compile.o

go build -o x.out compile.go
./x.out