fn main() {
    let mut count = 0u32;

    println!("Let's count until infinity!");
	
	//既然if 支持let 赋值，那么loop也是支持的，当然break返回就要带值了，参考下一个例子。
    // 无限循环
    loop {
        count += 1;

        if count == 3 {
            println!("three");

            // 跳过这次迭代的剩下内容
            continue;
        }

        println!("{}", count);

        if count == 5 {
            println!("OK, that's enough");

            // 退出循环
            break;
        }
    }
}
