using System;
public class Program{
    public static void Main(string[] args)
    {
        // 声明一个变量绑定
        int a_binding;

        {
            var x = 2;

            // 初始化一个绑定
            a_binding = x * x;
        }

        Console.WriteLine("a binding: {0}", a_binding);

        //let another_binding;

        // 报错！使用了未初始化的绑定
        // println!("another binding: {}", another_binding);//编译器禁止使用未经初始化的变量，因为会产生未定义行为
        // 改正 ^ 注释掉此行

        var another_binding = 1;

        Console.WriteLine("another binding: {0}", another_binding);
        
        Func<int,int> twice = (int x) => x * x;
        a_binding = twice(2);
        
        Console.WriteLine("a binding: {0}", a_binding);
    }
}