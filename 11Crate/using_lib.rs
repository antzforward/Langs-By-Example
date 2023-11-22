//extern crate rary; // 在 Rust 2015 版或更早版本需要这个导入语句
//现在使用^的语句，会报错can't find crate for `rary` 实际上library.rlib已经存在了。

fn main() {
    rary::public_function();

    // 报错！ `private_function` 是私有的
    //rary::private_function();

    rary::indirect_access();
}
