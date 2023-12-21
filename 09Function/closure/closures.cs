using System;
public class Program
{
    /*
    在C#中，闭包、Func<>、Action<>和匿名函数是与委托相关的重要概念。这些都与定义和使用可以在运行时传递的代码块有关。

    委托（Delegate）
    委托是一种类型，它安全地封装了一个方法的引用。在C#中，委托是对函数的抽象，相当于是指向函数的指针或者是函数的容器。

    匿名函数
    匿名函数是没有名称的方法。在C#中，匿名函数通常以两种形式出现：匿名方法和Lambda表达式。

    匿名方法（Anonymous Methods）：在C# 2.0中引入，允许开发者直接在委托创建处指定一个代码块作为事件处理器。
    delegate(int x) { return x * x; }
    
    Lambda表达式（Lambda Expressions）：在C# 3.0中引入，提供了更简洁的方式编写内联代码。Lambda表达式可以有参数并返回值。
    x => x * x
    (x, y) => x + y
    () => SomeMethod()
    
    闭包（Closure）
    一个闭包是一种能够捕获外部变量值的函数。在C#中，当你在匿名函数中使用了其所在作用域之外的变量时，就创建了一个闭包。该匿名函数“关闭”了这些变量，并保持对它们的引用。

    int factor = 2;
    Func<int, int> multiplier = n => n * factor;
    在上面的例子中，multiplier是一个闭包，因为它捕获了外部变量factor的状态。

    Func<> 和 Action<>
    Func<>和Action<>是内置的委托类型，使得声明委托类型变得更简单。

    Action<>：这是不返回值的委托类型（void返回类型）。它可以有0到16个输入参数。
    Action<string> greet = name => Console.WriteLine("Hello, " + name);
    greet("World");
    Func<>：这是有返回值的委托类型。它可以有0到16个输入参数，最后一个类型参数是返回值类型。
    Func<int, int, int> sum = (a, b) => a + b;
    int result = sum(1, 2);
    使用Func<>和Action<>可以避免定义新的委托类型，并且它们非常适合用来编写具有表达式树功能的LINQ查询和其他需要用到表达式的情况。

    以上就是C#中闭包、Func<>、Action<>和匿名函数的基本知识。这些特性在实现如回调方法、事件处理器、LINQ表达式等场景下非常有用。
    */
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");

        // 使用方法实现自增
        Func<int, int> function = i => i + 1;//delegate (int i) { return i + 1; };
        // 将闭包赋值给变量。Lambda表达式会根据上下文推断类型。
        Func<int, int> closureAnnotated = (int i) => { return i + 1; };
        Func<int, int> closureInferred = i => i + 1;

        int i = 1;

        // 调用方法和闭包
        Console.WriteLine("function: " + function(i));
        Console.WriteLine("closureAnnotated: " + closureAnnotated(i));
        Console.WriteLine("closureInferred: " + closureInferred(i));

        // 无参数闭包，返回一个整数
        Func<int> one = () => 1;//delegate{return 1;};
        Console.WriteLine("closure returning one: " + one());
    }
}
/** Output
function: 2
closure_annotated: 2
closure_inferred: 2
closure returning one: 1
*/