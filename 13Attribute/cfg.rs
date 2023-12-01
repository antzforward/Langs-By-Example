// 这个函数仅当目标系统是 Linux 的时候才会编译
#[cfg(target_os = "linux")]
fn are_you_on_linux() {
    println!("You are running linux!")
}

//写法类似C#中的Complication 属性，当不属于当前的属性 就不会编译，并且对应的调用自动忽略。
//因此 这个函数，应该没有返回值才合适。
// 而这个函数仅当目标系统 **不是** Linux 时才会编译
#[cfg(not(target_os = "linux"))]
fn are_you_on_linux() {
    println!("You are *not* running linux!")
}

fn main() {
    are_you_on_linux();
    
    println!("Are you sure?");
    if cfg!(target_os = "linux") {
        println!("Yes. It's definitely linux!");
    } else {
        println!("Yes. It's definitely *not* linux!");
    }
}
//问题1：target_os 可以直接在语句中使用吗？比如println！中作为参数传递？
//问题2：系统级别的默认的宏都有哪些？如何查找？还有复杂的用法吗？