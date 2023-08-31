go tool compile -o compile.o compile.go
# chmod +x compile.o
# ./compile.o
# go build compile.o


# go build -o x.out compile.go
# ./x.out

# Linking and building an executable: If you have multiple object files or packages that need to be linked together, you can use the go tool link command to link the object files and create an executable binary. Here's an example:
# go tool link -o executable main.o x.o

# Creating a package archive: If the object file represents a package that you want to distribute as a library, you can create a package archive (.a file) using the go tool pack command.
rm compile.a
go tool pack r compile.a compile.o

