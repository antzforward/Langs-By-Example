using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        int counter = 0;
        int result = 0;

        while (true) {
            counter += 1;

            if (counter == 10) {
                result = counter * 2;
                break;
            }
        }

        Console.WriteLine($"Result is {result}");
        // 这行代码是检查结果是否为期望值
        if (result != 20) {
            Console.WriteLine("Assertion failed");
        } else {
            Console.WriteLine("Assertion succeeded");
        }
    }
}