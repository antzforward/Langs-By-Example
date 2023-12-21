using System;
using System.Linq;
public class Program
{
    
    static bool is_odd( int n )  {
        return n % 2 == 1;
    }

    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");

        Console.WriteLine("Find the sum of all the squared odd numbers under 1000");
        int upper = 1000;

        // 命令式（imperative）的写法
        // 声明累加器变量
        var acc = 0;
        // 迭代：0，1, 2, ... 到无穷大
        for(int n=0;;n++) {
            // 数字的平方
            var n_squared = n * n;

            if (n_squared >= upper) {
                // 若大于上限则退出循环
                break;
            } else if (is_odd(n_squared) ){
                // 如果是奇数就计数
                acc += n_squared;
            }
        }
        Console.WriteLine("imperative style: {0}", acc);

        // 函数式的写法 Linq的写法。
        var sumOfSquaredOddNumbers = Enumerable.Range(0, int.MaxValue)
                .Select( n=>(long)n*n)             // 所有自然数取平方
                .TakeWhile( n=> n < upper) // 取小于上限的
                .Where(n => n % 2== 1)     // 取奇数
                .Sum(); // 最后加起来
        Console.WriteLine("functional style: {0}", sumOfSquaredOddNumbers);
    }
}