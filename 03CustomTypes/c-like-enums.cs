using System;
public class Program
{
    // 拥有隐式辨别值（implicit discriminator，从 0 开始）的 enum
    enum Number {
        Zero,
        One,
        Two,
    }

    // 拥有显式辨别值（explicit discriminator）的 enum
    enum Color {
        Red = 0xff0000,
        Green = 0x00ff00,
        Blue = 0x0000ff,
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        
        Console.WriteLine("zero is {0}", (int)Number.Zero);//enum 不能在这里使用as，如果是Enum的对象就可以，as 用于动态检查
        Console.WriteLine("one is {0}", (int)Number.One );

        Console.WriteLine("roses are #{0:x6}", (int)Color.Red);
        Console.WriteLine("violets are #{0:x6}", (int)Color.Blue);
    }
}
/*output
zero is 0
one is 1
roses are #ff0000
violets are #0000ff
*/