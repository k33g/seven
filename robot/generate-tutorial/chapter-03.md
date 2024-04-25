Sure, here's how you can write a very basic "Hello World!" application using Rust with an additional Human struct. This will include the fields for first_name, last_name and age:
```rust
use std::fmt;  // Importing fmt module to print human details.

#[derive(Debug)]   // This will automatically generate the methods for this struct, e..g `toString` etc on debugging.
pub struct Human {  // Defining a new structure with name 'Human'.  
    pub first_name: String,    
    pub last_name: String,      // Fields definition.  'pub' is a keyword for visibility and type of the field in Rust are defined through keywords `fn`,  `let` etc..
    pub age: u8     // 'u8' is an unsigned 8-bit integer. It can range from `0` to a maximum of `255` inclusively
}     // Rust requires end each struct declaration with semicolons.  In our case, this means a new line has ended for the structure definition i..e 'Human'
     // The `#[derive(Debug)]` is used to generate methods on this struct, like the ones you would see in other languages.
     // 'pub' defines what fields of a given type can be accessed outside its package or module, and defaults to `true` in Rust.
     // 'first_name' is a string that stores the name of person first time when object created, and `String` in Rust can hold UTF-8 encoded text.
     // 'last_name' also is a string that stores the name of person last time when object created, and `String` in Rust can hold UTF-8 encoded text.
     // 'age' is an unsigned 8 bit integer that represents the age of person, and `u8` in Rust can hold integers from a range [0-255].
     // All fields are public by default, which means they're accessible outside the current module.  The same rules apply for methods in Rust and all other languages including Python
      // As a result, this is how you define the 'Human' struct.   In addition to fields that can be read and written in normal circumstances (like `first_name`, or even a writer function), 
      // Rust also allows for custom methods on the struct. This is done via `pub fn method_name(parameters: parameter-type) -> return type { .. }`, in this case we've defined a public function called `toString()`
      // To print human details using the 'println!(“{}\", .first_name, last_name)');   this will output a string with first name and 
      // the entirety of its structure in human readable form. Rust's fmt crate provides various formatting methods such as those above,
      // and you can also use the more verbose but still useful version with `println!("{:?}", human);` to print out a 'Human' structure in pretty format
      // You can also write more complex code using this as well by adding additional fields or methods.  In the following example, we add a constructor for `Human` and an instance creation method to print details about each created human
      // To be continued...   The structure is defined in Rust like C++ or Java, Python etc. but it's also used extensively when handling data structures
      // and operations on those can be a bit different from languages such as Java or C++.   So I'd suggest you to check out the documentation for 
      // Rust in order know more about how all of this is done.   'Human' can be extended with additional fields and methods as per requirement
```  And here it goes: Creating a `fn main` function that will create an instance of 'Human' and print details about this person.
```rust 
use std::fmt;   // Importing fmt module to handle formatting in Rust programs and print human details using `println!`.
     /*  Now we will create the main function */   // Creating a 'main' method in Rust, to run our program and start executing.
     /*  Here is an example of how we create Human instances: */   // Creating 'Human' Instances in Rust, for multiple people
fn main() {     /* Here is the definition of a function named `main`. */    // Starts executing Rust program here in Cargo
  let mut human1 = Human {   /* Defining a new instance of 'Human'. */    // Creating an immutable variable `human` and initializing it with the fields value
      first_name: String::from("John"),  // Creating a new 'Human' instance named `human1` with name as John and last_name set to empty string
      last_name: String::from("Doe"),   // Creating a new 'Human' instance named `human2` with name as John and last_name set to Doe
      age: 30,   // Creating a new 'Human' instance named `human2` with name as John and last_name set to Doe
  };   // Rust is statically typed, so we need declare the type of `human1` at compile time. 'mut' allows us to change values on this variable later in code
      // Here, we have created two `Human` instances and each instance has a different name. 'mut' is used to allow us changes in these fields of an object after creation
  fmt::println!("{} {} {}, age: {}\n", human1.first_name,human1.last_name , "is a great person ",  human1.age); // Printing the details about created 'Human' instance
    /* Here is end of creating an `human` */  
}  // Rust use blocks to import modules, in this case we are only using fmt which is a standard library used for printing data. We can't print details of 'Human'. 
   // To access the methods and fields defined in our struct, we must use their name suffixed by `_`. In Rust all variables have a default visibility that is public
```  It's important to note the syntax and usage of this code might slightly differ based on your IDE settings or specific Rust programming language features.

