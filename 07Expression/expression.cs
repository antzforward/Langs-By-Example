using System;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("这是c#语言");
        Console.WriteLine("世界你好!");
        var x = 5u;
        // 使用匿名方法同时声明并调用 比rust里面要麻烦多了
        var y = new Func<uint>(() =>
        {
            var x_squared = x * x;
            var x_cube = x_squared * x;

            // 将此表达式赋给 `y`
            return x_cube + x_squared + x;
        })();
        var z = 2*x;
        Console.WriteLine("x is {0}", x);
		Console.WriteLine("y is {0}", y);
		Console.WriteLine("z is {0}", z);
    }
}