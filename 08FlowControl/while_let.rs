
#![allow(unused)]
fn main() {
	// 将 `optional` 设为 `Option<i32>` 类型
	let mut optional = Some(0);

	// 重复运行这个测试。
	loop {
		match optional {
			// 如果 `optional` 解构成功，就执行下面语句块。
			Some(i) => {
				if i > 9 {
					println!("Greater than 9, quit!");
					optional = None;
				} else {
					println!("`i` is `{:?}`. Try again.", i);
					optional = Some(i + 1);
				}
				// ^ 需要三层缩进！
			},
			// 当解构失败时退出循环：
			_ => { break; }
			// ^ 为什么必须写这样的语句呢？肯定有更优雅的处理方式！
		}
	}
	
    // 将 `optional` 设为 `Option<i32>` 类型
    let mut optional = Some(0);

    // 这读作：当 `let` 将 `optional` 解构成 `Some(i)` 时，就
    // 执行语句块（`{}`）。否则就 `break`。
    while let Some(i) = optional {//符合Python的:= walrus operator 内部定义变量，但是python中 if while 都要求是个conditional expression 不能是上面的简单形式。
        if i > 9 {
            println!("Greater than 9, quit!");
            optional = None;
        } else {
            println!("`i` is `{:?}`. Try again.", i);
            optional = Some(i + 1);
        }
        // ^ 使用的缩进更少，并且不用显式地处理失败情况。
    }
    // ^ `if let` 有可选的 `else`/`else if` 分句，
    // 而 `while let` 没有。

}
