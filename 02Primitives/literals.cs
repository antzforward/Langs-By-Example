using System;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
            // 整数相加
        Console.WriteLine("1 + 2 = {0}", 1u + 2);

        // 整数相减
        Console.WriteLine("1 - 2 = {0}", 1 - 2);
        // 试一试 ^ 尝试将 `1` 改为 `1u`，C#编译器会提醒在“error CS0220: 在 checked 模式下，运算在编译时溢出”

        // 短路求值的布尔逻辑
        Console.WriteLine("true AND false is {0}", true && false);
        Console.WriteLine("true OR false is {0}", true || false);
        Console.WriteLine("NOT true is {0}", !true);

        // 位运算
        Console.WriteLine("0011 AND 0101 is {0}", Convert.ToString(0b0011 & 0b0101,2).PadLeft(4,'0'));
        Console.WriteLine("0011 OR 0101 is {0}", Convert.ToString(0b0011 | 0b0101,2).PadLeft(4,'0'));
        Console.WriteLine("0011 XOR 0101 is {0}", Convert.ToString(0b0011 ^ 0b0101,2).PadLeft(4,'0'));
        Console.WriteLine("1 << 5 is {0}", Convert.ToString(1u << 5,2).PadLeft(32,'0'));
        Console.WriteLine("0x80 >> 2 is 0x{0}", Convert.ToString(0x80 >> 2,16));

        // 使用下划线改善数字的可读性！
        Console.WriteLine("One million is written as {0}", 1_000_000u);
    }
}
/*output 
1 + 2 = 3
1 - 2 = -1
true AND false is false
true OR false is true
NOT true is false
0011 AND 0101 is 0001
0011 OR 0101 is 0111
0011 XOR 0101 is 0110
1 << 5 is 00000000000000000000000000100000
0x80 >> 2 is 0x20
One million is written as 1000000
*/