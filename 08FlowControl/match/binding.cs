using System;
public class Program
{
    public static uint Age() {
        return 15u;
    }

    public static int? SomeNumber(){
        return 42;
    }
    //guard,就是match 然后用if的方式。
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        var age = Age();
        var message = age switch {
            0 => "I haven't celebrated my first birthday yet",
            >= 1 and <= 12 => $"I'm a child of age {age}",
            >= 13 and <= 19 => $"I'm a teen of age {age}",
            _ => $"I'm an old person of age {age}"
        };
        Console.WriteLine(message);

        int? number = SomeNumber();
        message = number switch{
            42=>$"The Answer: {number}!",
            int n=>$"Not interesting... {n}",
            _=>"",
        };
        if( message != "") Console.WriteLine(message);
    }
}