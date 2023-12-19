using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        while (true) { // 外层循环
            Console.WriteLine("Entered the outer loop");

            while (true) { // 内层循环
                Console.WriteLine("Entered the inner loop");

                // 这只是中断内部的循环
                //break;

                // 这会中断外层循环
                goto OuterLoopBreak;
            }

            // 实际上，由于上面的goto语句，这里的代码永远不会被执行。
            Console.WriteLine("This point will never be reached");
        }
        
    OuterLoopBreak:
        Console.WriteLine("Exited the outer loop");
    }
}
/*
Entered the outer loop
Entered the inner loop
Exited the outer loop
*/