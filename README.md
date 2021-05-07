# Factors and Primes
Factors and Primes is a Go CLI application that displays factors and prime numbers from the given value(s).

### Usage Factors and Primes
**Note:** Assuming the following command was used to build the executable: `go build -o <output filename>`, .

```
Usage:
$ ./factorsandprimes <-factor> <-prime> <-generate {nbrOfItemsToGenerate} {maxValue}>

Explanation:
  * <-factor> - this flag is used to signify that factors are outputted 
  * <-prime> - this flag is used to signify that prime number are outputted
  * <-generate {nbrOfItemsToGenerate} {maxValue}> - this flag signifies that the number will be generated. **Note:** When this flag is used both valus are required

Examples:
 $ ./factorsandprimes -factor - the application will ask the user to enter numbers separated by spaces, then the factors for those numbers are returned
 $ ./factorsandprimes -prime - the application will ask the user to enter numbers separated by spaces, then the prime factors for those numbers are returned
 $ ./factorsandprimes -factor -prime - the application will ask the user to enter numbers separated by spaces, then the factors and prime factors for those numbers are returned
```

### Requirements

  * Git
  * Golang version 1.16

### Downloading Factors and Primes

Clone Factors and Primes (https://github.com/sflewis2970/factorsandprimes) project.

### Building Factors and Primes

Build Factors and Prime
```
The following command will create an output file called 'main.exe':

$ go build main.go
```

```
The following command will create an output file called '<filename>.exe'

$ go build -o <filename>.exe
```

### Running Factors and Primes

Run without building:
```
The following command will run the factors and primes application without building the application:

$ go run main.go
```

Run the build:
```
This command will ask the user to enter numbers separated by spaces, then return thet factors for that that number

$ ./factorsandprime -factor
```

```
This command will ask the user to enter numbers separated by spaces, then return thet factors AND primes for that that number

$ ./factorsandprime -factor -prime
```

```
The following command will generate 10 numbers from 1- 50 and display the factors and primes for each number

$ ./factorsandprime -factor -prime -generate 10 50
```

### Testing Factors and Primes

Test modules have been provided.

Testing the entire tree:

```
The following command runs all the tests in the project tree and provides verbose information

$ go test -v ./...
```

Testing the modules in the current directory:

```
The following command runs all the tests in the project tree and provides verbose information

$ go test -v
```

Testing a single module:

```
The following command runs all the tests in the algebra sub-directory

$ go test -v ./algebra
```
