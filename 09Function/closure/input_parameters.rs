//这里说的是将函数作为参数的模式。类似的写法要定义Function 是第一性的对象才行。
//显然要对Function自身进行假定。在C#中应该能找到同等的语法。
//从这个Example来看，例子简单了。一般函数定义为 fn Func<T1,T2,R>( a:T1,b:T2)->R{}
//这个就是C# 里面的Func<T1,T2,R>实践了。当然 C#中的所有权不用处理，这个就简单一点，在Rust里面要标注一下。

/**
虽然 Rust 无需类型说明就能在大多数时候完成变量捕获，但在编写函数时，这种模糊写法是不允许的。当以闭包作为输入参数时，必须指出闭包的完整类型，它是通过使用以下 trait 中的一种来指定的。其受限制程度按以下顺序递减：

Fn：表示捕获方式为通过引用（&T）的闭包
FnMut：表示捕获方式为通过可变引用（&mut T）的闭包
FnOnce：表示捕获方式为通过值（T）的闭包
*/

// 该函数将闭包作为参数并调用它。
fn apply<F>(f: F) where
    // 闭包没有输入值和返回值。
    F: FnOnce() {
    // ^ 试一试：将 `FnOnce` 换成 `Fn` 或 `FnMut`。

    f();
}

// 输入闭包，返回一个 `i32` 整型的函数。
fn apply_to_3<F>(f: F) -> i32 where
    // 闭包处理一个 `i32` 整型并返回一个 `i32` 整型。
    F: Fn(i32) -> i32 {

    f(3)
}

fn apply_to_fact<F>(f:F,n:i32)->i32 where F: Fn(i32,i32) -> i32{
	f(n,1)
}

//这里主要是闭包不支持递归调用，只用具名函数来递归。
//这里学习的不是闭包的内容，而是Func的用法。
fn fact(n:i32,a:i32)->i32{
	if n<=1	{
		a
	}else{
		fact(n-1,n*a)
	}
}


fn main() {
    use std::mem;

    let greeting = "hello";
    // 不可复制的类型。
    // `to_owned` 从借用的数据创建有所有权的数据。
    let mut farewell = "goodbye".to_owned();

    // 捕获 2 个变量：通过引用捕获 `greeting`，通过值捕获 `farewell`。
    let diary = || {
        // `greeting` 通过引用捕获，故需要闭包是 `Fn`。
        println!("I said {}.", greeting);

        // 下文改变了 `farewell` ，因而要求闭包通过可变引用来捕获它。
        // 现在需要 `FnMut`。
        farewell.push_str("!!!");
        println!("Then I screamed {}.", farewell);
        println!("Now I can sleep. zzzzz");

        // 手动调用 drop 又要求闭包通过值获取 `farewell`。
        // 现在需要 `FnOnce`。
        mem::drop(farewell);
    };

    // 以闭包作为参数，调用函数 `apply`。
    apply(diary);

    // 闭包 `double` 满足 `apply_to_3` 的 trait 约束。
    let double = |x| 2 * x;

    println!("3 doubled: {}", apply_to_3(double));
    
    println!("LambdaFactorial: {}", apply_to_fact(fact,3));
}
