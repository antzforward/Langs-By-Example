#[allow(unused)]
fn main() {
    let x = 5u32;
	//这个类似c语言的逗号，最后一个是结果。顺序是从左到右。
    let y = {
        let x_squared = x * x;
        let x_cube = x_squared * x;

        // 将此表达式赋给 `y`
        x_cube + x_squared + x
    };
	//默认是unit类型，表示为空的tuple。这是Rust特殊的。
    let z = {
        // 分号结束了这个表达式，于是将 `()` 赋给 `z`
        2 * x;
    };

    println!("x is {:?}", x);
    println!("y is {:?}", y);
    println!("z is {:?}", z);
}
