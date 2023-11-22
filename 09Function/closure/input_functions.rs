//这个应该在C#中用的delegate的方式，对比一下，C#的delegate支持闭包吗？试试。

// 定义一个函数，可以接受一个由 `Fn` 限定的泛型 `F` 参数并调用它。
fn call_me<F: Fn()>(f: F) {
    f()
}

// 定义一个满足 `Fn` 约束的封装函数（wrapper function）。
fn function() {
    println!("I'm a function!");
}

fn fib_root<F:Fn(i32,i32,i32)->i32>(f:F,n:i32){
	println!("call a recursive function! {} ", f(n,1,1));
}

fn fib(n:i32,pre:i32,ppre:i32)->i32 {
	if n == 1 || n == 2 {
		pre
	}else {
		fib( n-1, pre+ppre, pre)
	}
}

fn main() {
    // 定义一个满足 `Fn` 约束的闭包。
    let closure = || println!("I'm a closure!");
    
    call_me(closure);
    call_me(function);
    fib_root(fib, 5);
}
