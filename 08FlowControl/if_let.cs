using System;
public class Program
{
    public enum Foo
    {
        Bar,
        Baz
    }

    public class Qux
    {
        public int Value { get; }

        public Qux(int value) => Value = value;
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        // 将 `optional` 定为 `Option<i32>` 类型
        int? optional = 7;

        var result = optional switch {
            int value => $"This is a really long string and `{value}`", //value居然是全范围的，明显有问题 应该scope非常窄才对，不如rust！！！
            _ => "",
            // ^ 必须有，因为 `match` 需要覆盖全部情况。不觉得这行很多余吗？
        };
        if( result !="") Console.WriteLine(result);
        
        //使用if let来做match匹配，当然，用？表示更简洁。
        // 全部都是 `Option<i32>` 类型
        int? number = 7;
        int? letter = null;
        int? emoticon = null;

        // `if let` 结构读作：若 `let` 将 `number` 解构成 `Some(i)`，则执行
        // 语句块（`{}`）
        if( number is int  n){//这个写法 秒啊， if a is int va，唯一的缺点是va的scope范围太大了。
            Console.WriteLine("Matched {0}!", n);
        }

        // 如果要指明失败情形，就使用 else：
        if( letter is int le ) {
            Console.WriteLine("Matched {0}!", le);
        } else {
            // 解构失败。切换到失败情形。
            Console.WriteLine("Didn't match a number. Let's go with a letter!");
        }

        // 提供另一种失败情况下的条件。
        var i_like_letters = false;
        
        // 并没有进行条件判断，这里认为None自动处理。
        if( emoticon is int em ){
            Console.WriteLine("Matched {0}!", em);
        // 解构失败。使用 `else if` 来判断是否满足上面提供的条件。
        } else if ( i_like_letters ) {
            Console.WriteLine("Didn't match a number. Let's go with a letter!");
        } else {
            // 条件的值为 false。于是以下是默认的分支：
            Console.WriteLine("I don't like letters. Let's go with an emoticon :)!");
        }
    
        var a = Foo.Bar;
        var b = Foo.Baz;
        var c = new Qux(100);

        // 变量 a 匹配到了 Foo.Bar
        if (a is Foo.Bar)
        {
            Console.WriteLine("a is foobar");
        }

        // 变量 b 没有匹配到 Foo.Bar，因此什么也不会打印。
        if (b is Foo.Bar)
        {
            Console.WriteLine("b is foobar");
        }

        // 如果 c 是 Qux 类型，并且其中包含的值是 100
        if (c is Qux { Value: 100 })
        {
            Console.WriteLine("c is Qux with value 100");
        }
        
        var d = Foo.Bar;

        // 变量匹配 Foo::Bar
        if (d is Foo.Bar)
        {
            Console.WriteLine("a is foobar");
        }
        
    }
}