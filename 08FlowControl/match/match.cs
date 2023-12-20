using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        var number = 11;
        // 试一试 ^ 将不同的值赋给 `number`

        Console.WriteLine("Tell me about {0}", number);
        var message = number switch {
            // 匹配单个值
            1 => "One!",
            // 匹配多个值
            2 or 3 or 5 or 7 or 11 => "This is a prime",
            // 试一试 ^ 将 13 添加到质数列表中
            // 匹配一个闭区间范围  start..end->[start,end) start..=end->[start,end]
            (>=13) and (<19) => "A teen",
            // 处理其他情况
            _ => "Ain't special",
            // 试一试 ^ 注释掉这个总括性的分支,提示没有完全covered，不完整。
        };
        Console.WriteLine( message );

        var boolean = true;
        // match 也是一个表达式
        var binary = boolean switch {
            // match 分支必须覆盖所有可能的值
            false => 0,
            true => 1,
            // 试一试 ^ 将其中一条分支注释掉，不带_的情况要全covered，类似switch的default
        };

        Console.WriteLine("{0} -> {1}", boolean, binary);
    }
}