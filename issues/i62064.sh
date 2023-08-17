rm i62064
quantum=0x2000
echo quantum is $quantum
go build -ldflags "-R $quantum" -o i62064 i62064.go
du -sh i62064
echo "\n"
./i62064
readelf -l i62064

# View the ELF header:
#    $ readelf -h <binary-file>
# The ELF header contains general information about the binary, such as the entry point address, section and program header table offsets, and the machine architecture.

# View the program header table:
#    $ readelf -l <binary-file>
# The program header table describes the various segments of the binary, including the loadable segments, their memory addresses, file offsets, and permissions.

# View the section header table:
#    $ readelf -S <binary-file>
# The section header table provides information about the binary's sections, such as the section names, types, sizes, virtual addresses, and file offsets.

# View symbols:
#    $ readelf -s <binary-file>
# This option displays the symbol table, which contains information about the binary's defined and referenced symbols. It includes function names, variable names, and their corresponding addresses.

# View relocations:
#    $ readelf -r <binary-file>
# The -r option displays the relocation entries, which are necessary for linking and resolving addresses at runtime.

