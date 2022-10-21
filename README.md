# lalalint
Lalalint is a static code analyzer (a.k.a. linter) for LaTeX documents - specifically for .tex, .bib and .tikz-files.
Lalalint can output any errors found to the terminal, write a linted version of the input file(s) to a new file or
overwrite the input file(s). It works for both individual files and for directories. The linter rules can be enabled or
disabled at will via JSON.

## (Pseudo)-installation
Lalalint does not currently get installed into your path. In order to run the program, you need to be in the 
lalalint directory. However, if you have a 64bit processor and your OS is Linux, Windows or macOS, you're in luck. 
Getting it up and running is super easy.

### macOS
1. Open up your terminal and navigate to the lalalint directory
2. Run ./install_macos.bash
3. Done! Run ./lalalint --help to see that everything is working correctly

### Windows
1. Open up your terminal and navigate to the lalalint directory
2. Run ./install_windows.bash
3. Done! Run ./lalalint --help to see that everything is working correctly

### Linux
1. Open up your terminal and navigate to the lalalint directory
2. Run ./install_linux.bash
3. Done! Run ./lalalint --help to see that everything is working correctly

## Build it yourself (any OS)
If you'd like to build lalalint yourself, follow these steps:

1. Install [Go](https://go.dev/) and [Git](https://git-scm.com/)
2. Run git clone https://github.com/jf-Lindberg/lalalint.git
3. Run cd lalalint
4. Run go build -o lalalint main.go
5. Done! Run ./lalalint --help to see that everything is working correctly

## Usage
```
Usage:

  lalalint [inputfile | directory] [outputfile] [flags]


Flags:

  -d, --directory   Runs the linter on all .tex, .bib and .tikz files in the input directory (changed via config).
                    An argument can be passed to override the default directory.
                    Example: lalalint -d <directory>
  -h, --help        help for lalalint
  -o, --overwrite   Overwrites the input file(s). If run together with the --directory flag, all files in the directory will be overwritten.
  -v, --verbose     Prints all linter problems found to the terminal.

```

#### Linting an individual file:
In order to lint a file, at least one argument must be passed - the input file. If lalalint is called without any
other arguments or flags, it will output how many problems it found in the file. If you wish to save the output to a
new file, you can pass in the output name as a second argument.

#### Linting a directory:
Lalalint can be run in directory mode without any arguments. By default, lalalint will run on the directory specified
under "inputdirectory" in the configuration file. If a path is specified as argument, it will run on that directory
instead. Please note that subdirectories will **not** be linted.

#### Simple demo
If you've installed lalalint as per the instructions above, you can try running lalalint on data/demo.tex. Try running:

```
./lalalint demo.tex --verbose
```

You should get a terminal output similar to this one:

```
Using config file: PATH/lalalint/.lalalint.json
Checking demo.tex for errors
Problems detected:

        no space after comment on row 1, character 18: '%Inli'
        no space after comment on row 2, character 1: '%This'
        no blank line(s) before section on row 4
        no newline after sentence on row 6
        environment not indented correctly on row 10
        no newline after sentence on row 11
        environment not indented correctly on row 11
        no space after comment on row 11, character 29: '%This'
        environment not indented correctly on row 12
        environment not indented correctly on row 13
        environment not indented correctly on row 14
        environment not indented correctly on row 15

✔ Check finished
Check took 899.208µs
To fix the problems, please specify an output file as a second argument or use the overwrite flag.
```

As mentioned at the end of the terminal output, you can specify an output file to write the linter results to a new 
file, like so:

```
./lalalint demo.tex output.tex
```

Now check /data/output/. A file called output.tex should be written with a linted version of the input. Neat, huh?

Lalalint can do other things like overwriting as well, but now you get the basics. Happy linting!

## Configuration
Lalalint is configured via a .json file, .lalalint.json. The default configuration file looks like this:
```json
{
  "global": {
    "verbose": false,
    "inputdirectory": "./data/",
    "outputdirectory": "./data/output/",
    "trimwhitespace": false
  },
  "rules": {
    "indentenvironments": {
      "enabled": true,
      "tabs": 1,
      "indentlevel": 0,
      "excluded": [
        "document"
      ]
    },
    "spaceaftercomments": {
      "enabled": true
    },
    "blanklinesbeforesection": {
      "enabled": true,
      "lines": 1
    },
    "newlineaftersentence": {
      "enabled": true
    }
  }
}
```

## Global parameters
#### Verbose
If *verbose* is enabled, the linter will explicitly output any linter problems found to the terminal. This can also 
be enabled temporarily via the -v | --verbose flag. Disabled by default.

#### Inputdirectory / outputdirectory
*Input-* and *output-directory* specifies the directories where the input files are stored and where output files 
should be stored, respectively. When linting files, there is unfortunately no way to change the 
input/output-directories via arguments (coming at a later date if I continue working on this as a side project). If 
lalalint is called with the --directory flag and no arguments, lalalint will run on the directory specified under 
inputdirectory. If it's called with an argument, lalalint will run on the directory specified in the argument.

#### Trimwhitespace
If *trimwhitespace* is enabled, any whitespace surrounding a line in the file (like spaces or tabs) will be removed 
prior to linting. This can be useful if your file, for example, has messy indentation. **However, enabling this will 
result in false linter problems output in regard to the indentation rule.** Disabled by default.

## Rules parameters
### All rules
*Enabled* will turn the rule on or off. If all rules and trimwhitespace are disabled (set to false), the output of the 
program will be exactly the same as the input file.

### Indentenvironments
#### Tabs
*Tabs* sets the amount of tabs that should be prepended to lines inside environments. 1 by default.

#### Indentlevel
*Indentlevel* defines the baseline indentation level in a file. If you want each line to be indented for some reason, 
set this to a positive value. 0 by default, meaning no indentation except for environments.

#### Excluded
*Excluded* is an array of environments that should be excluded from the indentation rule, meaning that lalalint will 
not indent any content inside these environments. The only environment excluded by default is "document".

### Blanklinesbeforesection
#### Lines
*Lines* defines the amount of lines that lalalint will set as a minimum before a section. 1 by default.

## Bill of materials
Other than Go's built in library, lalalint also uses:

### Cobra
[Cobra](https://cobra.dev/) is a framework for creating CLI apps. This project uses it for parsing arguments and 
flags.

### Viper
[Viper](https://github.com/spf13/viper) is a configuration library which this project uses for everything 
configuration; reading and writing to the .json-file, making sure that the Cobra flags overrides default 
configuration value, etc.

### ColoredCobra
[ColoredCobra](https://github.com/ivanpirog/coloredcobra) formats the Cobra output, so that it's not just plain text.
Colors are more fun. 

### Color by Fatih
[Color by Fatih](https://github.com/fatih/color) also helps with adding color to the output of lalalint. Any color 
that is added to non-Cobra output is added through Color by Fatih.
