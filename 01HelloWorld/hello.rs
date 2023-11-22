// 这是注释内容，将会被编译器忽略掉
// 可以单击那边的按钮 "Run" 来测试这段代码 ->
// 若想用键盘操作，可以使用快捷键 "Ctrl + Enter" 来运行

// 这段代码支持编辑，你可以自由地修改代码！
// 通过单击 "Reset" 按钮可以使代码恢复到初始状态 ->

// 这是主函数
fn main() {
    // 这是行注释的例子
    // 注意有两个斜线在本行的开头
    // 在这里面的所有内容都不会被编译器读取

    // println!("Hello, world!");

    // 请运行一下，你看到结果了吗？现在请将上述语句的两条斜线删掉，并重新运行。

    /*
     * 这是另外一种注释——块注释。一般而言，行注释是推荐的注释格式，
     * 不过块注释在临时注释大块代码特别有用。/* 块注释可以 /* 嵌套, */ */
     * 所以只需很少按键就可注释掉这些 main() 函数中的行。/*/*/* 自己试试！*/*/*/
     */

    /*
    注意，上面的例子中纵向都有 `*`，这只是一种风格，实际上这并不是必须的。
    */

    // 观察块注释是如何简单地对表达式进行修改的，行注释则不能这样。
    // 删除注释分隔符将会改变结果。
    let x = 5 + /* 90 + */ 5;
    println!("Is `x` 10 or 100? x = {}", x);
}
