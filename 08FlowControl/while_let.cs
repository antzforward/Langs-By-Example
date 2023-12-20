using System;
public class Program
{
    //强行的凑样子，远不如Rust语法啊
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        int? optional = 0;

        while (optional.HasValue) {
            int i = optional.Value;
            if (i > 9) {
                Console.WriteLine("Greater than 9, quit!");
                optional = null;
            } else {
                Console.WriteLine("`i` is `{0}`. Try again.", i);
                optional = i + 1;
            }
        }

        int? optional2 = 0;
        while(true)
        {
            if( optional2.HasValue )
            {
                int i = optional2.Value;
                if (i > 9) {
                    Console.WriteLine("Greater than 9, quit!");
                    break;
                } else {
                    Console.WriteLine("`i` is `{0}`. Try again.", i);
                    optional2 = i + 1;
                }     
            }
            else
                break;
        }
    }
}
/*
*/