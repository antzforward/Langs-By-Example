using System;
public class Program
{
    static void apply( Action f ) {
        f();
    }
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");

        var x = 7;

        // 捕获 `x` 到匿名类型中，并为它实现 `Fn`。
        // 将闭包存储到 `print` 中。
        Action print = ()=>Console.WriteLine("{0}", x);
        
        apply(print);
    }
}
