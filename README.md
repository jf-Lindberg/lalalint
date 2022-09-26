# lalalint

## How to download
Just run this in your terminal:
```
git clone https://github.com/jf-Lindberg/lalalint
```

## How to use
Currently, a binary is compiled for macOS which means you can run it without setting anything else up.
### macOS (works for both x86 and ARM processors)
```
./lalalint_macos <command> [argument] [option]
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
