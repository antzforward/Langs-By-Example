//既然能从A到B，就一定能从B到A，From和Into就代表这个相互的关系。
//From 通常的用法是 let B = T::From(A);
//Into 通常的用法是 let B:T = A.into();
//这两个都值需要实现 std::convert::From

use std::convert::From;

#[allow(dead_code)]
#[derive(Debug)]
struct Number {
    value: i32,
}

impl From<i32> for Number {
    fn from(item: i32) -> Self {
        Number { value: item }
    }
}

fn main() {
    let int = 5;
    // 试试删除类型说明
    let num: Number = int.into();
    println!("My number is {:?}", num);
    
    let my_str = "hello";
	let my_string = String::from(my_str);
	println!("My string is {} : {}", my_str,my_string);
	
	let num = Number::from(30);
    println!("My number is {:?}", num);
}
