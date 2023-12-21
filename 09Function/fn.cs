using System;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        fizzbuzz_to(100);
    }
    // 一个返回布尔值的函数
    public static bool is_divisible_by(uint lhs, uint rhs) =>
        rhs switch{
            0=>false,
            _ =>lhs % rhs == 0,
        };
        
    

    // 一个 “不” 返回值的函数。实际上会返回一个单元类型 `()`。
    public static void fizzbuzz(uint n) =>Console.WriteLine(n switch {
        uint _ when is_divisible_by(n, 15) => "fizzbuzz",
        uint _ when is_divisible_by(n, 3) => "fizz",
        uint _ when is_divisible_by(n, 5) => "buzz",
        _ => $"{n}"
    });

    // 当函数返回 `()` 时，函数签名可以省略返回类型
    public static void fizzbuzz_to(uint n) {
        for(uint i=1u;i<=n; i++){
            fizzbuzz(i);
        }
    }
}