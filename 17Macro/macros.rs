// 这是一个简单的宏，名为 `say_hello`。
macro_rules! say_hello {
    // `()` 表示此宏不接受任何参数。
    () => (
        // 此宏将会展开成这个代码块里面的内容。
        println!("Hello!");
    )
}

fn main() {
    // 这个调用将会展开成 `println("Hello");`!
    say_hello!()
}
//问题1：C系列的宏被#[allow()]或#[disable()]之类的处理了
//问题2：Rust的macro实际上是元编程（metaprogramming）。宏并不产生函数调用，而是展开源码，并和程序的其余部分一起被编译。
//问题3：Rust的宏会展开为抽象语法树（AST），而不是像字符串预处理那样直接替换成代码，这样就不会产生无法预料的优先权错误。

//问题4：为什么宏是有效的？