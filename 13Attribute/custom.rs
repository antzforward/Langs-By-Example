#[cfg(some_condition)]
fn conditional_function() {
    println!("condition met!")
}

fn main() {
    conditional_function();
}

//编译指令加上 --cfg some_condition 就会显示出来。
//但是 直接使用 rustc custom.rs 放到报出编译错误了。cannot find function `conditional_function` in this scope 
//函数因为宏未定义所以不编译，但是调用的地方就出现上面的错误了。