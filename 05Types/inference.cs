using System;
using System.Collections.Generic;
public class Program{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        //可以说 C# 并没有这个特性
        // 因为有类型说明，编译器知道 `elem` 的类型是 u8。
        byte elem = 5;

        // 创建一个空向量（vector，即不定长的，可以增长的数组）。
        var vec = new List<int>();
        // 现在编译器还不知道 `vec` 的具体类型，只知道它是某种东西构成的向量（`Vec<_>`）

        // 在向量中插入 `elem`。
        vec.Add(elem);
        // 啊哈！现在编译器知道 `vec` 是 u8 的向量了（`Vec<u8>`）。
        // 试一试 ^ 注释掉 `vec.push(elem)` 这一行。
        // answer：去掉了类型推断的第一句，vec类型不确定，编译时出错了。type annotations needed for `Vec<T>`
        Console.WriteLine("{0}", "["+string.Join(",",vec)+"]");
    }
}
/*
5
*/