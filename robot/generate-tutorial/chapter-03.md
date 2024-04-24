Sure, I can provide a way to create such `Human` struct and use it in Rust. Here is how you could do this:
```rust
struct Human {  // Declaration of a new type 'Human'. This struct will contain three fields, which are defined here as first_name: String (which stores the name), last_name: String(for storing surname) and age: u8 which is for years.
    first_name: String,  // Data type of 'first name' field - a string (the data this will store) is enclosed in double quotes.
    last_name: String, // Data type of 'last name' field - a string (the data this will store) is enclosed in double quotes.
    age: u8, // Data type of 'age' field - an unsigned integer (the data this will store) is enclosed in a single quote.
} 
// This line signals the end of defining 'Human' struct which marks our structure as complete now we can start using it by creating an instance.
// To create a new human, you would do something like this:  let someone = Human { first_name:"John", last_name : "Doe". age :25};
fn main() {   // The 'main' function.  It is the entry point of any program in Rust which will be called by default when you run your application
    let someone = Human { // Creating a new instance. We are using the 'structured' syntax to create an object (a human). 
        first_name: String::from("John"), // We are using the `String` struct's method, "from", to turn a string into an actual 'string'.
        last_name: String::from("Doe"), // We are doing the same for each of our fields.  These will be "John" and so on...
        age: 25, // ...and these are the ages (assumed to always exist).   This will be stored in 'age'.
    }; 
     println!("{} {} is {} years old.", someone.first_name,someone.last_name , someone.age); // Printing out a message using the fields of 'Human'
}  /* End main function */  
```    This code creates an instance `someone` and assigns it a new value of the fields. Then, print out "John Doe is 25 years old." to console using `println!` macro which uses string formatting features of Rust.

