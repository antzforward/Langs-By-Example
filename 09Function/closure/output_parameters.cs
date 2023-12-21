using System;

class Program
{
    static Action CreateFn()
    {
        string text = "Fn";
        return () => Console.WriteLine($"This is a: {text}");
    }

    static Action CreateFnMut()
    {
        string text = "FnMut";
        return () => Console.WriteLine($"This is a: {text}");
    }

    static Action CreateFnOnce()
    {
        string text = "FnOnce";
        bool called = false;
        return () =>
        {
            if (!called)
            {
                Console.WriteLine($"This is a: {text}");
                called = true;
            }
            else
            {
                throw new InvalidOperationException("This action can only be called once.");
            }
        };
    }

    static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        var fnPlain = CreateFn();
        var fnMut = CreateFnMut();
        var fnOnce = CreateFnOnce();

        fnPlain(); // Can be invoked multiple times.
        fnMut();   // Can also be invoked multiple times.
        fnOnce();  // Designed to be invoked only once.

        // Uncommenting the following line will throw an exception because fnOnce() can only be called once.
        //fnOnce();
    }
}