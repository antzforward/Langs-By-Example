using System;
using System.Linq;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        var vec1 = new[]{ 1, 2, 3};
        var vec2 = new[]{ 4, 5, 6};

        // 使用 LINQ 的 Any() 方法检查 vec1 是否包含数字 2
        Console.WriteLine("2 in vec1: {0}", vec1.Any(x=> x == 2));
        // vec2 是一个新数组，所以直接调用 Any() 方法即可
        Console.WriteLine("2 in vec2: {0}", vec2.Any(x=> x == 2));

        var array1 = new[]{ 1, 2, 3 };
        var array2 = new[]{ 4, 5, 6 };

        // 对于数组也是同样的方法
        Console.WriteLine("2 in array1: " +  array1.Any(x=> x == 2));
        // 对数组的 `into_iter()` 举出 `i32`。//使用array2.into_iter().any(...)会出现错误，提示换成下面写法。
        Console.WriteLine("2 in array2: " +  array2.Any(x=> x == 2));
    }
}
/*
This is C# Language
2 in vec1: True
2 in vec2: False
2 in array1: True
2 in array2: False
*/