# Factors and Primes
Factors and Primes is a Go CLI application that takes input from the user or generates values for the user.
Those values are used to determine the factors for each number and when the flag is set determines whether 
the numbers are prime numbers.

### Usage Factors and Primes
```
Usage:
$ ./factorsandprimes <-factor> <-prime> <-generate ''> <-generate '{nbrOfItemsToGenerate} {maxValue}'>

Explanation:
  * <-factor> - when this flag is set, the factors each number are returned
  * <-prime> - when this flag is set, prime numbers are returned
  * <-inputdata '{list of numbers separated by space}'> - when this flag is used, the next string is a list of numbers used as input
  * <-generate {nbrOfItemsToGenerate} {maxValue}> - when flag is used, the next string is 

Examples:
 The following command uses input data as input and returns the factors of each number:
 $ ./factorsandprimes -factor -inputdata '13 26 39' 
 
 The following command generates 10 number from 1-100, then returns the factors for each number and which numbers are prime numbers
 $ ./factorsandprimes -factor -prime -generate '10 100'
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
The following command will run the factors and primes application without building the application.

$ go run main.go
```

Run the build:
```
This command will generate 10 numbers from 1-100 and return the factors for each number

$ ./factorsandprime -factor -generate '10 100'
```

```
This command will generate 10 numbers from 1-100 and return the factors for each number and which numbers that are prime numbers

$ ./factorsandprime -factor -prime -generate '10 100'
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
