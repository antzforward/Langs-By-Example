using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");

        var n = 5;

        if (n < 0) {
            Console.Write("{0} is negative", n);
        } else if (n > 0) {
            Console.Write("{0} is positive", n);
        } else {
            Console.Write("{0} is zero", n);
        }

        //这里还是不如rust，语句分段了。
        int big_n =  n < 10 && n > -10? 10 *n : n/2;
        // 根据条件输出相应的信息
        Console.WriteLine((n < 10 && n > -10) ? ", and is a small number, increase ten-fold"
                                               : ", and is a big number, half the number");
        //   ^ 不要忘记在这里加上一个分号！所有的 `let` 绑定都需要它。

        Console.WriteLine("{0} -> {1}", n, big_n);
    }
}
/* output
5 is positive, and is a small number, increase ten-fold
5 -> 50
*/