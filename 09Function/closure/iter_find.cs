using System;
using System.Collections.Generic;
using System.Linq;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        List<int> vec1 = new() { 1, 2, 3 };
        List<int> vec2 = new() { 4, 5, 6 };

        Console.WriteLine("Find 2 in vec1: {0}", vec1.FirstOrDefault( x=>x == 2));
        Console.WriteLine("Find 2 in vec2: {0}", vec2.FirstOrDefault( x=>x == 2));

        var array1 = new[]{ 1, 2, 3 };
        var array2 = new[]{ 4, 5, 6 };

        Console.WriteLine("Find 2 in array1: {0}", array1.FirstOrDefault(x=>x == 2));
        Console.WriteLine("Find 2 in array2: {0}", array2.FirstOrDefault(x=>x == 2));
    }
}
/*
Find 2 in vec1: Some(2)
Find 2 in vec2: None
Find 2 in array1: Some(2)
Find 2 in array2: None
=========================
This is C# Language
Find 2 in vec1: 2
Find 2 in vec2: 0
Find 2 in array1: 2
Find 2 in array2: 0
*/