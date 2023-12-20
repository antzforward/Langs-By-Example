using System;
namespace System.Runtime.CompilerServices
{
    internal static class IsExternalInit {}
}
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        var triple = (0, -2, 3);
        // 试一试 ^ 将不同的值赋给 `triple`

        Console.WriteLine("Tell me about {0}", triple);
        // match 可以解构一个元组，按照匹配来，这个跟C#新的switch语句和匹配，也用了match
        var message = triple switch {
            // 解构出第二个和第三个元素
            (0,var y,var z) => $"First is `0`, `y` is {y}, and `z` is {z}",
            (1, _, _)  => "First is `1` and the rest doesn't matter",
            // `..` 可用来忽略元组的其余部分
            _      => "It doesn't matter what they are",
            // `_` 表示不将值绑定到变量，确实存在x 在0，1以外的情况，默认匹配必须的。
        };
        Console.WriteLine( message );
    }
}