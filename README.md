# lalalint

## How to download
Just run this in your terminal:
```
git clone https://github.com/jf-Lindberg/lalalint
```

## How to use
Currently, binaries are compiled for macOS which means you can run it without setting anything else up.
### macOS x86 (Intel processors)
Run lalalint on macOS with an Intel processor:
```
./lalalint_intel <command> [argument] [option]
```

### macOS ARM (Silicon/MX processors)
Run lalalint on macOS with a Silicon processor:
```
./lalalint_silicon <command> [argument] [option]
```

### Support for Linux and Windows
Binaries will be compiled for Linux and Windows as well. If you wish to install lalalint on your Linux/Windows 
machine, make sure you have installed Go, that you are in the root folder of the cloned directory and then run:
```
go build -o lalalint main.go
```

You should then be able to run lalalint via: 
```
./lalalint <command> [argument] [option]
```
