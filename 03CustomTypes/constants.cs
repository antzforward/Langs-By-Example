using System;
public class Program
{
    static string LANGUAGE = "Rust";
    const  int  THRESHOLD = 10;

    public static bool is_big(int n ) {
        // 在一般函数中访问常量
        return n > THRESHOLD;
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        var n = 16;

        // 在 main 函数（主函数）中访问常量
        Console.WriteLine("This is {0}", LANGUAGE);
        Console.WriteLine("The threshold is {0}", THRESHOLD);
        Console.WriteLine("{0} is {1}", n, is_big(n)? "big" : "small" );
    }
}
/*Output
This is Rust
The threshold is 10
16 is big
*/