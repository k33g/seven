Sure, I can update the previous code to include a `Human` struct with the fields `firstName`, `lastName`, and `age`. Here's how you can do it:

```go
package main
import "fmt"
type Human struct {
    firstName string
    lastName  string
    age       int
}
func main() {
    var person Human
    fmt.Println("Enter first name:")
    fmt.Scan(&person.firstName)
    fmt.Println("Enter last name:")
    fmt.Scan(&person.lastName)
    fmt.Println("Enter age:")
    fmt.Scan(&person.age)
    fmt.Printf("Your full name is %s %s and you are %d years old.\n", person.firstName, person.lastName, person.age)
} 
``` 
In this code:
- `type Human struct {...}` is a declaration that creates a new type named `Human`.
- `Human` is a struct type. Structs are used to group together fields of different types.
- `var person Human` declares a variable of type `Human`.
- `fmt.Println("Enter first name:")`, `fmt.Scan(&person.firstName)`, and similar lines read the first name from the user and assign it to the `firstName` field of the `person` struct.
- `fmt.Println("Enter last name:")`, `fmt.Scan(&person.lastName)`, and similar lines read the last name from the user and assign it to the `lastName` field of the `person` struct.
- `fmt.Println("Enter age:")`, `fmt.Scan(&person.age)`, and similar lines read the age from the user and assign it to the `age` field of the `person` struct.
- `fmt.Printf("Your full name is %s %s and you are %d years old.\n", person.firstName, person.lastName, person.age)` prints the full name and age of the `person` struct.

