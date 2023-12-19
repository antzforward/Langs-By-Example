using System;
public class Program
{
    public struct Circle{
        public int radius;
        public override string ToString()
        {
            return $"Circle of radius {radius}";
        }
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        //to string
        var circle = new Circle { radius =  6 };
        Console.WriteLine("{0}", circle + " very good!");

        //from string
        var parsed  = Convert.ToInt32("5");
        var turbo_parsed = Convert.ToInt32("10");

        var sum = parsed + turbo_parsed;
        Console.WriteLine("Sum: {0}", sum);
    }
}
/*
Circle of radius 6 very good!
Sum: 15
*/