using System;
using System.Diagnostics;
public class Program
{
    // 定义一个函数，可以接受一个由 `Fn` 限定的泛型 `F` 参数并调用它。
    static void call_me(Action f) {
        f();
    }

    // 定义一个满足 `Fn` 约束的封装函数（wrapper function）。
    static void function() {
        Console.WriteLine("I'm a function!");
    }

    static void fib_root(Func<int,int,int,int>f,int n ) {
        Console.WriteLine("call a recursive function! {0} ", f(n,1,1));
    }

    static int fib(int n,int pre , int ppre) => 
        n switch {
            (1 or 2)=>pre,
            _ => fib(n-1, pre+ppre, pre),
        };
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        // 定义一个满足 `Fn` 约束的闭包。
        Action closure = ()=>Console.WriteLine("I'm a closure!");
        
        call_me(closure);
        call_me(function);
        Stopwatch stopwatch = Stopwatch.StartNew();
        fib_root(fib, 45);
        stopwatch.Stop();
        // 如果你需要更精细的时间单位，可以这样做：
        Console.WriteLine("Nanoseconds elapsed: {0}", (double)stopwatch.ElapsedTicks/Stopwatch.Frequency * 1_000_000_000);
    }
}
/*
I'm a closure!
I'm a function!
call a recursive function! 1134903170
Elapsed milliseconds: 335500 //C# 这个数值为308600，尾递归只防止overflow 但是执行时间并没有少太多啊。
*/