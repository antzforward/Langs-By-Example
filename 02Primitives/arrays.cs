using System;
using System.Linq;
public class Program
{
    //在array上，都差不多的样子，slice上比较多不同的。
    //slice在C#上有两个不同的类型，一种是Span是结构类型，是一个ref struct，不能在堆上分配，因此不能作为类的属性
    //一种是Memory的类型，可以安全地存储在堆上，并且可用作异步操作的部分。
    //当前的例子都是Span的形式，只是引用的概念。

    // 此函数借用一个 slice
    public void analyze_slice(Span<int> slice ) {
        Console.WriteLine("first element of the slice: {0}", slice[0]);
        Console.WriteLine("the slice has {0} elements", slice.Length);
    }
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        // 定长数组（类型标记是多余的）
        var xs = new int[]{1, 2, 3, 4, 5};

        // 所有元素可以初始化成相同的值
        var ys = Enumerable.Repeat(0,500).ToArray();

        // 下标从 0 开始
        Console.WriteLine("first element of the array: {0}", xs[0]);
        Console.WriteLine("second element of the array: {0}", xs[1]);

        // `len` 返回数组的大小
        Console.WriteLine("array size: {0}", xs.Length);

        // 数组是在栈中分配的,C#里面是算不准的，除非是非托管的对象才能用近似C的方式来计算。
        Console.WriteLine("array occupies {0} bytes", sizeof(int)*xs.Length);

        // 数组可以自动被借用成为 slice
        Console.WriteLine("borrow the whole array as a slice");
        analyze_slice(xs.AsSpan().Slice(0));

        // slice 可以指向数组的一部分
        Console.WriteLine("borrow a section of the array as a slice");
        analyze_slice(ys.AsSpan().Slice(1,3));

        // 越界的下标会引发致命错误（panic）//运行时会出现 System.IndexOutOfRangeException: 索引超出了数组界限
        //Console.WriteLine("{0}", xs[5]);
    }
}
/*output
first element of the array: 1
second element of the array: 2
array size: 5
array occupies 20 bytes
borrow the whole array as a slice
first element of the slice: 1
the slice has 5 elements
borrow a section of the array as a slice
first element of the slice: 0
the slice has 3 elements
*/