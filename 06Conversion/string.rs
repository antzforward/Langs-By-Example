//这次解释的是针对String和其他类型的from和into 以及实现方式。
//一个类型如果要Into String的写法，实现fmt.Display即可。然后调用A.to_string即可
//一个字符串要Into到其他类型，需要实现FromStr这个特质（Trait）然后 let B:T = A.parse().unwrap() 或者let B = A.parse::<T>().unwrap()
use std::fmt;

struct Circle {
    radius: i32
}
// 实现了fmt::Display,会自动提供ToString。
impl fmt::Display for Circle {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Circle of radius {}", self.radius)
    }
}
/*
impl ToString for Circle {
    fn to_string(&self) -> String {
        format!("Circle of radius {:?}", self.radius)
    }
}
*/

fn main() {
	//to string
    let circle = Circle { radius: 6 };
    println!("{}", circle.to_string() + " very good!");
    
    //from string
    let parsed: i32 = "5".parse().unwrap();
    let turbo_parsed = "10".parse::<i32>().unwrap();

    let sum = parsed + turbo_parsed;
    println!{"Sum: {}", sum};
}
