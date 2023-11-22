//rustc --crate-type=lib filename
//能不能添加main fn？
//要想定制名字要用 crate_type 和 crate_name
//然后编译指令就可以退回到普通的。rustc filename
//lib的大小可不小啊。
#![crate_type="lib"]
#![crate_name="rary"]
pub fn public_function() {
    println!("called rary's `public_function()`");
}

fn private_function() {
    println!("called rary's `private_function()`");
}

pub fn indirect_access() {
    print!("called rary's `indirect_access()`, that\n> ");

    private_function();
}
