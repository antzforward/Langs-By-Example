using System;
using System.Runtime.InteropServices;
//public class Program{
//    public static void Main(string[] args)
{
    Console.WriteLine("This is C# Language");
    // 带后缀的字面量，其类型在初始化时已经知道了。
    byte x = 1; //byte 没有字面标识，只能这样了。
    var y = 2u;
    var z = 3f;

    // 无后缀的字面量，其类型取决于如何使用它们。
    var i = 1;
    var f = 1.0;//默认f64
    var g = 1.0f;

    // `size_of_val` 返回一个变量所占的字节数
    Console.WriteLine("size of `x` in bytes: {0}", Marshal.SizeOf(x));
    Console.WriteLine("size of `y` in bytes: {0}", Marshal.SizeOf(y));
    Console.WriteLine("size of `z` in bytes: {0}", Marshal.SizeOf(z));
    Console.WriteLine("size of `i` in bytes: {0}", Marshal.SizeOf(i));
    Console.WriteLine("size of `f` in bytes: {0}", Marshal.SizeOf(f));
    Console.WriteLine("size of `g` in bytes: {0}", Marshal.SizeOf(g));
}
//}
/*
size of `x` in bytes: 1
size of `y` in bytes: 4
size of `z` in bytes: 4
size of `i` in bytes: 4
size of `f` in bytes: 8
size of `g` in bytes: 4
*/