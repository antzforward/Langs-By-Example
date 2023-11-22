fn main() {
    // 声明一个变量绑定
    let a_binding;

    {
        let x = 2;

        // 初始化一个绑定
        a_binding = x * x;
    }

    println!("a binding: {}", a_binding);

    let another_binding;

    // 报错！使用了未初始化的绑定
    // println!("another binding: {}", another_binding);//编译器禁止使用未经初始化的变量，因为会产生未定义行为
    // 改正 ^ 注释掉此行

    another_binding = 1;

    println!("another binding: {}", another_binding);
    
    let a_binding =
    {
        let x = 2;
        x * x//注意这里如果加了分号,a_binding 的值为unit，（）本身没有实现Display会显示出编译错误,当然如果你fmt设置的是{:?} 带？的打出Debug就可以。
    };
    
    println!("a binding: {}", a_binding);
}
