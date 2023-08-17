rm i62064
quantum=0x010A
echo quantum is $quantum
go build -ldflags "-R $quantum" -o i62064 i62064.go
du -sh i62064
./i62064
readelf -l i62064