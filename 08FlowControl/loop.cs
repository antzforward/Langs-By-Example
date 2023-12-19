using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        var count = 0u;

        Console.WriteLine("Let's count until infinity!");
        
        //既然if 支持let 赋值，那么loop也是支持的，当然break返回就要带值了，参考下一个例子。
        // 无限循环
        while(true) {
            count += 1;

            if ( count == 3 ) {
                Console.WriteLine("three");

                // 跳过这次迭代的剩下内容
                continue;
            }

            Console.WriteLine("{0}", count);

            if( count == 5 ){
                Console.WriteLine("OK, that's enough");

                // 退出循环
                break;
            }
        }
    }
}
/* Output
Let's count until infinity!
1
2
three
4
5
OK, that's enough
*/