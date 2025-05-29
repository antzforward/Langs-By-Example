using System;
//public class Program{
//public static void Main(string[] args)
{
  Console.WriteLine("This is C# Language");

  var decimalv = 65.4321f;

  // 错误！不提供隐式转换
  //let integer: u8 = decimal;
  // 改正 ^ 注释掉这一行

  // 可以显式转换
  var integer = (byte)decimalv;
  var character = (char)integer;

  Console.WriteLine("Casting: {0} -> {1} -> {2}", decimalv, integer, character);

  // 当把任何类型转换为无符号类型 T 时，会不断加上或减去 (std::T::MAX + 1)
  // 直到值位于新类型 T 的范围内。

  //先把1000 转换成16进制的表示一下。
  Console.WriteLine("1000 as 0x is: 0x{0:x8}", 1000);
  // 1000 已经在 u16 的范围内
  Console.WriteLine("1000 as a u16 is: {0}", (ushort)1000);

  // 1000 - 256 - 256 - 256 = 232
  // 事实上的处理方式是：从最低有效位（LSB，least significant bits）开始保留
  // 8 位，然后剩余位置，直到最高有效位（MSB，most significant bit）都被抛弃。
  // 译注：MSB 就是二进制的最高位，LSB 就是二进制的最低位，按日常书写习惯就是
  // 最左边一位和最右边一位。
  Console.WriteLine("1000 as a u8 is : {0}", unchecked((byte)1000));
  // -1 + 256 = 255
  Console.WriteLine("  -1 as a u8 is : {0}", unchecked((byte)(-1)));

  // 对正数，这就和取模一样。
  Console.WriteLine("1000 mod 256 is : {0}", 1000 % 256);

  // 当转换到有符号类型时，（位操作的）结果就和 “先转换到对应的无符号类型，
  // 如果 MSB 是 1，则该值为负” 是一样的。

  // 当然如果数值已经在目标类型的范围内，就直接把它放进去。
  Console.WriteLine(" 128 as a i16 is: {0}", (short)128);
  // 128 转成 u8 还是 128，但转到 i8 相当于给 128 取八位的二进制补码，其值是：
  Console.WriteLine(" 128 as a i8 is : {0}", unchecked((sbyte)128));

  // 重复之前的例子
  // 1000 as u8 -> 232
  Console.WriteLine("1000 as a u8 is : {0}", unchecked((byte)1000));
  // 232 的二进制补码是 -24
  Console.WriteLine(" 232 as a i8 is : {0}", unchecked((sbyte)232));
  // u32-u32 当小减去大的话，会让结果溢出
  Console.WriteLine("1u - 2u is : {0}", unchecked((int)1u - (int)2u));

  //除了as 之外 还有 From 和Into 两个函数。 待到06 再处理。一个是编辑期决定，一个是运行时处理？？
}
//}
/*Output
Casting: 65.4321 -> 65 -> A
1000 as 0x is: 0x000003e8
1000 as a u16 is: 1000
1000 as a u8 is : 232
  -1 as a u8 is : 255
1000 mod 256 is : 232
 128 as a i16 is: 128
 128 as a i8 is : -128
1000 as a u8 is : 232
 232 as a i8 is : -24
1u - 2u is : -1
*/