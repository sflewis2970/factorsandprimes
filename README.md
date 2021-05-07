# Factors and Primes
Factors and Primes is a Go CLI application that displays factors and prime numbers from the given value(s).

### Usage Factors and Primes
**Note:** Assuming the following command was used to build the executable: `go build -o <output filename>`, .

Display the usage message:
$ ./factorsandprimes <-factor> <-prime> <-generate {nbrOfItemsToGenerate} {maxValue}>

Explanation:
  * <-factor> - this flag is used to signify that factors are outputted 
  * <-prime> - this flag is used to signify that prime number are outputted
  * <-generate {nbrOfItemsToGenerate} {maxValue}> - this flag signifies that the number will be generated. **Note:** When this flag is used both valus are required

When -factor or -prime is used the user will be asked to enter the values (separated by spaces).

### Requirements

  * Git
  * Golang version 1.16

### Downloading Factors and Primes

Clone Factors and Primes (https://github.com/sflewis2970/factorsandprimes) project.

### Building Factors and Primes

Build Factors and Prime
```
**Note:** Make sure you're in the root directory of the project
$ go build main.go

This command will create an outpuyt file called 'main.exe'
```
**Note:** Make sure you're in the root directory of the project
$ go build -o <filename>.exe
  
This command will create an output file called '<filename>.exe'
```

### Running Factors and Primes

Run without building:
```
**Note:** Make sure you're in the root directory of the project
$ go run main.go

This command will run the factors and primes without the need to build an executable
```

Run the build:
```
**Note:** Assuming the output file is 'factorsandprimes.exe'
$ ./factorsandprime -factor

This command will ask the user to enter numbers separated by spaces, then return thet factors for that that number
```
$ ./factorsandprime -factor -prime

This command will ask the user to enter numbers separated by spaces, then return thet factors AND primes for that that number
```
$ ./factorsandprime -factor -prime -generate 10 50

This command will generate 10 numbers from 1- 50 and display the factors and primes for ech number
```
### Testing Factors and Primes

Test modules have been provided.

Testing the entire tree:
```
**Note:** Make sure you're in the root directory of the project
$ go test -v ./...

This command runs all the tests in the project tree and provides verbose information
```
Testing a single module:
```
**Note:** Make sure you're in the root directory of the project
$ go test -v ./algebra

This command runs all the tests in the algebra sub-directory
```
