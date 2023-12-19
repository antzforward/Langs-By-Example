using System;
using System.Collections.Generic;
public class Program{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");

        var an_integer = 1u;
        var a_boolean = true;
        var unit = new HashSet<int>();

        // 将 `an_integer` 复制到 `copied_integer`
        var copied_integer = an_integer;

        Console.WriteLine("An integer: {0}", copied_integer);
        Console.WriteLine("A boolean: {0}", a_boolean);
        Console.WriteLine("Meet the unit value: {0}", unit); //没有这种类型，得出（）的形式，（）只能是函数的调用。

        // 编译器会对未使用的变量绑定产生警告；可以给变量名加上下划线前缀来消除警告。
        var _unused_variable = 3u;

        var noisy_unused_variable = 2u;

    }
}
/* Output
An integer: 1
A boolean: true
Meet the unit value: ()
*/