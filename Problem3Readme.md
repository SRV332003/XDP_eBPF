# Explaination of Go Concurency Program (Problem 3)

This is the assessment for recruitment process at Accuknox. Current Readme.md describes the solution of 3rd problem statement. For problem 1 solution, visit [`Readme.md`](https://github.com/SRV332003/XDP_eBPF/blob/main/Readme.md "Readme.md")

## Code to Explain
The code demonstrates the concept of concurrency and concurrency control in golang using `Go routines` and `Go channels`.

```go
package main

import "fmt"

func main() {

    // makes channel that can store 10 functions
    // with no return value and no argument
    cnp := make(chan func(), 10) 

    //this for loop creates 4 go routines
    for i := 0; i < 4; i++ {

        // each go routine wait for channel to 
        // have some function inserted in it.
        go func() {
            for f := range cnp {
                f()
            }
        }()
    }

    // insert the anonymous function in the 
    // channel which prints HERE1
    cnp <- func() {
        fmt.Println("HERE1")
    }

    // prints Hello
    fmt.Println("Hello")

    // program exits closing all go routines
}
```

The program performs the following operations:
1. The execution starts from the main function.
2. The line `cnp := make(chan func(), 10)` creates a channel that can store 10 functions of prototype `func()`.
3. Next, a for loop stars and iterate 4 times. 
   - At each time, it creates a go routine calling anonymous function that reads from `cnp` channel and executes the function in it, if its empty then it waits for some function to be inserted.    
4. This loop creates 4 such go routines which are now waiting for cnp (currently empty) to contain some value. 

5. Now the function `func() {fmt.Println("HERE1")}` is inserted into channel cnp. Now, cnp has 1 function in it.
This function prints the string `HERE1` into stdout.
6. Now that function will be extracted out by any of the go routines listening for cnp.
7. The Next line `fmt.Println("Hello")` in `main` function, prints `Hello` to stdout.

### Output
Despite the simple explaination as above, the code seems to be printing both `HERE1` and `Hello`, but that's **not the actual output**. 
#### Actual Output
```bash
Hello
```
#### Explaination
Because the goroutines are running concurrently, it's possible that `Hello` could be printed before `HERE1`, even though the function to print `HERE1` is sent to the channel first. This is because the goroutines are scheduled to run at the discretion of the Go runtime, and their execution order is not guaranteed. 

Moreover, after execution of `fmt.Println("Hello")` in main program, the whole program exits closing all go routines before printing "HERE1".

This is a race condition actually, the code can end with 3 different output.

When main ends too early than go routine func execution:
```bash
Hello
```
Above content will be the output most of the time.

When go routine executes the function after print of "Hello" but before exit of program. The output will be:
```bash
Hello
HERE1
```
You will see the above output once or twice if you run the program 5-10 times.

When go routine executes the print statement before print execution of main thread. The output will be
```bash
HERE1
Hello
```
Theoretically, the above output seems possible, but practically, this cannot occure or has chances close to `0%` because, it take some time to read from channel and then calling the function and then executing the statements in the function. The main thread, on the other hand, has to just execute the print statement. 

Hence pracically, the output of `HERE1` before `Hello` is nearly immpossible.


## Use Cases of this construct
- `Implementation of Network Buffers` : The above construct can be used for implementation of Network Buffers that concurrently store the incomming requests in a `channel` and allow multiple subscibers to fetch this concurrently for processing.
- `Implementation of Publisher-Subsciber Topics` : The above construct effectively demonstrates the `publisher-subscriber model` that can be used in variety of applications and services.

