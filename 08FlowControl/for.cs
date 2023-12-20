using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        // `n` 将在每次迭代中分别取 1, 2, ..., 100
        for(int n=1;n<=100;n++) { //还是简单的for好些，用Enumerable.Range(1, 100) 要引入Linq
            //用?:的嵌套来实现一条语句的写法
            Console.WriteLine(n % 15 == 0 ? "fizzbuzz" : n % 3 == 0 ? "fizz" : n % 5 == 0 ? "buzz" : n.ToString());
        }
        
        string[] names = {"Bob", "Frank", "Ferris"};

        foreach(var name in names ) {
            var message  = name switch{
                "Ferris"=>"There is a rustacean among us!",
                _ => $"Hello {name}",
            };
            Console.WriteLine(message);
        }
        
        //c# 并没有所有权借出以及Move的概念，这里只是增加一些写法罢了。
        for(int i=0; i<names.Length;i++) {
            if( names[i] == "Ferris" )
                Console.WriteLine("There is a rustacean among us!");
            else
                Console.WriteLine($"Hello {names[i]}");
        }
        
        
        //重点的用法，自身引用 使用assign match .. => ...

        foreach(var name in names ) {
            if( name == "Ferris" )
                Console.WriteLine("There is a rustacean among us!");
            else
                Console.WriteLine($"Hello {name}");
        }
        Console.WriteLine("names: {0}", "["+string.Join(",",names)+"]");
    }
}