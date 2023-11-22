// 此声明将会查找名为 `my.rs` 或 `my/mod.rs` 的文件，并将该文件的内容放到
// 此作用域中一个名为 `my` 的模块里面。
// rustc split.rs 会编译相关的代码，算很棒了。
//但是我单独编译了my这个lib，再用rustc split.rs 还是跳过my下面的lib重新编译一下，看来引用lib 要新的用法。
mod my;//应该要替换这个用法。

fn function() {
    println!("called `function()`");
}

fn main() {
    my::function();

    function();

    my::indirect_access();

    my::nested::function();
}
