#[allow(unused)]
// 这个结构体不能使用 `fmt::Display` 或 `fmt::Debug` 来进行打印。
struct UnPrintable(i32);

// `derive` 属性会自动创建所需的实现，使这个 `struct` 能使用 `fmt::Debug` 打印。
#[allow(unused)]
#[derive(Debug)]
struct DebugPrintable(i32);

// 推导 `Structure` 的 `fmt::Debug` 实现。
// `Structure` 是一个包含单个 `i32` 的结构体。
#[derive(Debug)]
struct Structure(i32);

// 将 `Structure` 放到结构体 `Deep` 中。然后使 `Deep` 也能够打印。
#[derive(Debug)]
struct Deep(Structure);

#[allow(unused)]
#[derive(Debug)]
struct Person<'a> {
    name: &'a str,
    age: u8
}

fn main() {
    // 使用 `{:?}` 打印和使用 `{}` 类似。
    println!("{:?} months in a year.", 12);
    println!("{1:?} {0:?} is the {actor:?} name.",
             "Slater",
             "Christian",
             actor="actor's");
	//这样写才是最简单的，{:?}是用的Debug模式，会凸显出参数的类型 对于基本类型，不加？就是正常的用法
	println!("{1} {0} is the {actor} name.",
             "Slater",
             "Christian",
             actor="actor's");
    // `Structure` 也可以打印！
    println!("Now {:?} will print!", Structure(3));
    
    // 使用 `derive` 的一个问题是不能控制输出的形式。
    // 假如我只想展示一个 `7` 怎么办？
    println!("Now {:?} will print!", Deep(Structure(7)));
    
    let name = "Peter";
    let age = 27;
    let peter = Person { name, age };

    // 美化打印，类似plain json和 formatted Json 打印一样。
    // 能用:? 就可以用:#?
    println!("{:#?}", peter);
}
/*输出的内容如下
12 months in a year.
"Christian" "Slater" is the "actor's" name.
Christian Slater is the actor's name.
Now Structure(3) will print!
Now Deep(Structure(7)) will print!
Person {
    name: "Peter",
    age: 27,
}
*/
