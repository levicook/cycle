# cycle

cycle is a tiny library for cycling through slices of string. I might
add other basic types in the future.

I created this library because I want to cycle through a set of host names.


## Usage

Create a StringCycle by calling Strings with the string values you want
to cycle through. 

Here's an application that creates a cycle containing two hex color
codes. It loops 10 times and printing the next value from the cycle
to standard out. The program illustrates a brute force pattern for
alternating row colors in html generators.

```go
package main

import (
	"fmt"

	"github.com/levicook/cycle"
)

func main() {
	colors := cycle.Strings("#fff", "#ccc")

	for i := 0; i < 10; i++ {
		fmt.Printf("Cycle %v: %q\n", i, colors.Next())
	}
}
```

The output for this application should look like this:

```bash
$ go run demo/main.go
Cycle 0: "#fff"
Cycle 1: "#ccc"
Cycle 2: "#fff"
Cycle 3: "#ccc"
Cycle 4: "#fff"
Cycle 5: "#ccc"
Cycle 6: "#fff"
Cycle 7: "#ccc"
Cycle 8: "#fff"
Cycle 9: "#ccc"
```

Boring, but useful. My favorite kind of software :)


## Benchmarks and Race Detection:

Nothing interesting here, just a sanity check because we are using the sync package. 

```bash
levi@needle:~/go/src/github.com/levicook/cycle
$ go test -bench=. -race
PASS
Benchmark_Strings_with_simpsons  5000000               643 ns/op
ok      github.com/levicook/cycle       4.882s
```
