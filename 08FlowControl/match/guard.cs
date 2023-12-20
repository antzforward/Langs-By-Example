using System;
public class Program
{
    //guard,就是match 然后用if的方式。
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");

        var pair = (3, -1);
        Console.WriteLine("Tell me about {0}", pair);

        var message = pair switch{
            (var x,var y) when x == y =>"These are twins",
            (var x,var y) when x + y == 0 =>"Antimatter, kaboom!",
            (var x, _) when x % 2 == 1=>"The first one is odd",
            _=>"No correlation...",
        };
        Console.WriteLine( message );
    }
}
/*Output
Tell me about (3, -1)
The first one is odd
*/