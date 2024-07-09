# usrgen

Package `usrgen` generates a username based on a person's first
and last name using the first letter of the first name plus the
entire last name.

- Consecutive calls to the `Generate()` function will return a
new username with an additional extra letter from the first name.

- Special characters will be translated to safe ASCII.

- An error will be generated if the `Generate()` function is called
more times than the count of the letters in the first name.

## Example

```go
package main

import (
	"fmt"

	"github.com/Alvaroalonsobabbel/usrgen"
)

func main() {
		var (
			firstName = "Joy"
			lastName  = "Perez"
			lang      = "de"
		)
		usrgen := usrgen.New(firstName, lastName, lang)

		username, err := usrgen.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(username) // Will print: "jperez"

		username, err = usrgen.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(username) // Will print: "joperez"

		username, err = usrgen.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(username) // Will print: "joyperez"

		username, err = usrgen.Generate()
		if err != nil {
			fmt.Println(err) // Will print: "User Generator error: Exceeded length of first name."
		}
		fmt.Println(username) // Will print: ""
}
```
