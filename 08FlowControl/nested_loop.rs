#![allow(unreachable_code)]

fn main() {
	//loop是无限循环，因此，break跳出点如果不能从内到最外部，只是从一层进入另一层了
	//这里类似的goto语法很重要。当然goto 如果没有编译器支持，资源的管理很容易出现问题。
    'outer: loop {
        println!("Entered the outer loop");

        'inner: loop {
            println!("Entered the inner loop");

            // 这只是中断内部的循环
            //break;

            // 这会中断外层循环
            break 'outer;
        }

        println!("This point will never be reached");
    }

    println!("Exited the outer loop");
}
