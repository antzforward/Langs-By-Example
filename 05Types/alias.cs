using System;

//注意以下的功能要在C# 12 版本的的新功能：类型的别名。
// `NanoSecond` 是 `u64` 的新名字。
using NanoSecond = ulong;
using Inch = ulong;
using u64_t = ulong;
// 试一试 ^ 移除上面那个属性
public class Program{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        // `NanoSecond` = `Inch` = `u64_t` = `u64`.
        NanoSecond  nanoseconds = 5 ;
        Inch        inches = 2 ;

        // 注意类型别名*并不能*提供额外的类型安全，因为别名*并不是*新的类型。
        Console.WriteLine("{0} nanoseconds + {1} inches = {2} unit?",
                nanoseconds,
                inches,
                nanoseconds + inches);
    }
}