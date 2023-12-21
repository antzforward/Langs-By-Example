using System;
public class Program
{
    // 这个方法接受一个没有输入和输出的 Action 委托
    static void Apply(Action f)
    {
        f();
    }

    // 这个方法接受一个 Func<T, TResult> 委托，即接受一个参数并返回一个值的函数
    static int ApplyTo3(Func<int, int> f)
    {
        return f(3);
    }

    static int apply_to_fact(Func<int, int,int> f,int n ){
        return f(n,1);
    }

//这里主要是闭包不支持递归调用，只用具名函数来递归。
//这里学习的不是闭包的内容，而是Func的用法。
    static int fact(int n, int a )=>
       n switch{
            <=1=>a,
            _=>fact(n-1, n*a),
       };

    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");

        string greeting = "hello";
        string farewell = "goodbye";

        // 定义一个匿名方法（闭包），它捕获了 greeting 和 farewell 变量
        Action diary = () =>
        {
            Console.WriteLine("I said {0}.", greeting);

            farewell += "!!!";
            Console.WriteLine("Then I screamed {0}.", farewell);
            Console.WriteLine("Now I can sleep. zzzzz");

            // 在 C# 中，对象通常通过垃圾回收来释放，不过可以调用 Dispose 来手动释放资源，
            // 如果farewell是一个IDisposable对象的话。这里没有这个必要，因为string不占用非托管资源。
        };

        // 应用闭包
        Apply(diary);

        // 定义另一个闭包 double，它符合 Func<int,int> 的签名
        Func<int, int> dfunc = x => 2 * x;

        // 输出结果
        Console.WriteLine("3 doubled: {0}", ApplyTo3(dfunc));

        Console.WriteLine("LambdaFactorial: {0}", apply_to_fact(fact,3));

    }
}
/*Output
I said hello.
Then I screamed goodbye!!!.
Now I can sleep. zzzzz
3 doubled: 6
LambdaFactorial: 6
*/