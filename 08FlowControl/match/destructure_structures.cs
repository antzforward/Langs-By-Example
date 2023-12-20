using  System;
public struct Foo {
    public (int, int) X { get; }
    public int Y { get; }

    public Foo((int, int) x, int y) {
        X = x;
        Y = y;
    }

    // 这是一个解构方法
    public void Deconstruct(out (int, int) x, out int y) {
        x = X;
        y = Y;
    }
}

class Program {
    static void Main() {
        Console.WriteLine("This is C# Language");

        var foo = new Foo((1, 2), 3);

        // 利用解构方法解构Foo
        var (x, y) = foo;
        var ((a, b), _) = foo; // 使用_忽略不需要的值

        Console.WriteLine($"a = {a}, b = {b},  y = {y}");

        // 如果你想要重命名变量，可以这样做：
        var (j, i) = foo;
        Console.WriteLine($"i = {i}, j = ({j.Item1}, {j.Item2})");

        // 忽略x只取y
        var (_, yOnly) = foo;
        Console.WriteLine($"y = {yOnly}");
    }
}
/*Output
a = 1, b = 2,  y = 3
i = 3, j = (1, 2)
y = 3
*/