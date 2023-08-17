quantum=0x200000
echo quantum is $quantum
go build -ldflags "-R $quantum" -o i62064 i62064.go
du -sh i62064
./i62064