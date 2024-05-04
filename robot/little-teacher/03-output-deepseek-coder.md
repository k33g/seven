Sure, I can provide an updated version of the previous code with a `Greet` method to say hello. Here's how you can do it:
```go
package main
import "fmt"
type Human struct {
    firstName string
    lastName  string
    age       int
}
func (h *Human) Greet() string {
    return "Hello, " + h.firstName + "!"
} 
func main() {
    var person Human
    fmt.Println("Enter first name:")
    fmt.Scan(&person.firstName)
    fmt.Println("Enter last name:")
    fmt.Scan(&person.lastName)
    fmt.Println("Enter age:")
    fmt.Scan(&person.age)
    fmt.Println(person.Greet())
} 
``` 
In this code:
- `func (h *Human) Greet() string` is a declaration that creates a new method named `Greet` on the `Human` struct.
- The `*Human` in the method signature means that the method is a method on the `Human` struct.
- `h.firstName` is a pointer to the `firstName` field of the `Human` struct.
- `return "Hello, " + h.firstName + "!"` is a string that uses the first name of the human to say hello.
- `person.Greet()` calls the `Greet` method on the `person` struct, which prints the greeting to the standard output.

