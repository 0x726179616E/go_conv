# go_conv

A simple commandline tool for converting values between some common physical units.

Inspired by Exercise 2.2 from the [GOPL book](https://www.gopl.io/).

### Usage
- Build executable: `go build main.go`
- Run program: `./main <measurement> <units> <value>`
- `<measurement>` may be either `-t`, `-w`, `-d`, or `-v` which represent conversions for temperature, weight, distance, or volume respectively 
- `<units>` may be any of the valid units for the specified `<measurement>` flag
    - valid units for`-t` are "c", "f", and "k"
    - valid units for`-w` are "kg" and "lbs"
    - valid units for`-d` are "m" and "ft"
    - valid units for`-v` are "l" and "gal"
- `<value>` may be any floating point value
- if `<units>` and/or `<value>` are omitted from the commandline arguments then the program will try to read them from standard input (stdin)
- Example usage: `./main -t -from=f 80.5` 

    ```
    converting temperature:
    80.50 F = 26.94 C = 300.09 K 
    ```
