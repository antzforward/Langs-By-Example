using System;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");

        string color = "green";

        // 闭包通过 Lambda 表达式打印 `color`。
        // 在 C# 中，闭包自然就捕获了对 color 的引用。
        Action print = () => Console.WriteLine("`color`: {0}", color);

        // 调用闭包
        print();

        // 在 C# 中，没有必要显式地进行“重新借用”操作。
        // 下面的代码只是为了演示在同一作用域中，`color` 可以被多次访问。
        string _reborrow = color;
        print();

        // 因为闭包已经捕获了 `color` 的引用，如果这个时候尝试更改 color 的值或者将其传出作用域，
        // 将会在编译时报错，因此下面的语句在 C# 中是非法的：
        // string _color_moved = color; // 错误：不能在闭包引用后移动 `color`

        int count = 0;
        // 这个闭包增加 `count` 的值。在 C# 中，如果你想修改闭包外部的局部变量，
        // 那么这个局部变量必须被声明为 `ref`。
        Action inc = () =>
        {
            count += 1;
            Console.WriteLine("`count`: {0}", count);
        };

        // 调用闭包
        inc();

        // 在闭包之外修改 `count` 也是允许的
        count += 1;

        // 再次调用闭包
        inc();

        // Box<T> 在 C# 中没有直接对应的类型，但可以考虑使用类似于“new int”的表达式，
        // 或者更通用的一个对象。
        var movable = new Box<int>(3);

        // 消耗 `movable` 的闭包
        // 在 C# 中，委托（或闭包）通常不“消耗”它们捕获的变量，
        // 因此你不需要像在 Rust 中那样担心这种情况。
        Action consume = () =>
        {
            Console.WriteLine("`movable`: {0}", movable);
            
            // 在 C# 中'Box'不是内置类型，所以我们假设有一个类似的
            // Box 类型和一个 Dispose 方法来模拟 Rust 中的 mem::drop。
            movable.Dispose();
        };

        // 调用闭包来“消耗”movable
        consume();

        // 在 C# 中，与 Rust 不同，由于闭包并不真正“消耗”它们捕获的变量，
        // 因此即使再调用一次 `consume` 也不会报错，尽管可能会得到运行时错误，
        // 因为 `movable` 已经被 Dispose 了。
        //consume(); // 运行时错误（如果没有实现重置或重新分配 movable） 结果错了。此时为default(T)了。并没有Rust那样在编译期就能给出提示。
    }

    // 这个类模拟 Rust 中的 Box<T>，包含一个 Dispose 方法来模拟 mem::drop。
    public class Box<T> : IDisposable
    {
        public T Value { get; private set; }

        public Box(T value)
        {
            Value = value;
        }

        public void Dispose()
        {
            // 模拟释放资源
            Value = default(T);
        }

        public override string ToString() => $"{Value}";
    }
}
/*Output
`color`: green
`color`: green
`count`: 1
`count`: 2
`movable`: 3
*/