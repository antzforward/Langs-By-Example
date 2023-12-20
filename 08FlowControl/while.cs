using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        // 计数器变量
        var n = 1;

        // 当 `n` 小于 101 时循环
        while( n < 101 ) {
            if( n % 15 == 0 ){
                Console.WriteLine("fizzbuzz");
            } else if( n % 3 == 0 ){
                Console.WriteLine("fizz");
            } else if( n % 5 == 0 ) {
                Console.WriteLine("buzz");
            } else {
                Console.WriteLine("{0}", n);
            }

            // 计数器值加 1
            n += 1;
        }
    }
}