# textwrap

I have a variety of projects that:

 * **need** a textwrapping library, even if it is overly simple
 * **could benefit from** a thorough textwrapping library
 * I want to work on **instead of** a thorough textwrapping library

This package serves as:

 * a handy utility
 * a simple library that works for now
 * a distribution mechanism for making incremental updates from a simple
   library to a thorough one that all of my projects can benefit from


## Development Status

Very alpha


## Usage (utility)

```sh
$ echo '12345' | textwrap -length 2
12
34
5
```


## Usage (library)

```go
import (
	"fmt"
	"regexp"
	"go.dominic-ricottone.com/textwrap"
)

func main() {
	a := []string{
		"> hello!",
		"> this is quoted",
		"hi there!",
		"this is not",
	}
	for _, line := range textwrap.wrap_array(a,10) {
		fmt.Println(line)
	}
}
```


## Licensing

GPLv3


