use std::fmt::{Display, Formatter, Result};

struct Structure{
  x: i32
}

impl Display for Structure {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "{}", self.x)
    }
}

fn main() {

  println!("My name is {0}, {1} {0}", "Bond", "James");

  // Create a structure named `Structure` which contains an `i32`.

  // However, custom types such as this structure require more complicated
  // handling. This will not work.
  println!("This struct `{}` won't print...", Structure{x:3});
  // FIXME ^ Comment out this line.

  let pi = 3.141592;
  println!("Pi is roughly {:.3}", pi);

}
