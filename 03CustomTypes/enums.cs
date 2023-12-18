#define ENUM_EXTENSIENUM
using System;
//这里会有很多写法，但是无法写得像Rust一样简洁。

//type 类型，用来标记Event的类型

public static class Program
{
    public enum WebEventType{
        // 一个 `enum` 可以是单元结构体（称为 `unit-like` 或 `unit`），
        PageLoad,
        PageUnload,
        // 或者一个元组结构体，
        KeyPress,
        Paste,
        // 或者一个普通的结构体。
        Click
    }
    public interface IWebEvent<T> where T:notnull 
    {
        public T GetEventType();
    }

    public struct WE_PageLoad : IWebEvent<WebEventType> {
        public WebEventType GetEventType() => WebEventType.PageLoad;
    }

    public struct WE_PageUnload : IWebEvent<WebEventType> {
        public WebEventType GetEventType() => WebEventType.PageUnload;
    }

    public struct WE_KeyPress : IWebEvent<WebEventType> {
        public char  key;
        public WebEventType GetEventType() => WebEventType.KeyPress;
    }

    public struct WE_Paste  : IWebEvent<WebEventType> {
        public string content;
        public WebEventType GetEventType() => WebEventType.Paste;
    }

    public struct WE_Click  : IWebEvent<WebEventType> {
        public double x;
        public double y;
        public WebEventType GetEventType() => WebEventType.Click;
    }

    public static void inspect( IWebEvent<WebEventType> webEvent ){
        //改成新的写法 更简洁一点
        var eventType = webEvent.GetEventType(); // 获取事件类型
        // 假设eventType是枚举类型，根据其值输出不同的消息
        var message = eventType switch
        {
            WebEventType.PageLoad => "page loaded",
            WebEventType.PageUnload => "page unloaded",
            // 如果有更多类型，可以添加更多的case
            // 假设键盘按压、粘贴和点击事件对应于指定的WebEventType值
            WebEventType.KeyPress => $"pressed '{((WE_KeyPress)webEvent).key}'.", // 假设 WE_KeyPress 类型有 Key 属性
            WebEventType.Paste => $"pasted '{((WE_Paste)webEvent).content}'.", // 假设 WE_Paste 类型有 Content 属性
            WebEventType.Click => $"clicked at x={((WE_Click)webEvent).x}, y={((WE_Click)webEvent).y}.", // 假设 WE_Click 类型有 X 和 Y 属性
            _ => "Unknown event type."
        };

        Console.WriteLine( message );
    }

    public enum Operations
    {
        Add,
        Subtract
    }
    //注意这种重命名的方式只能用在文件头，这里就只能简化命名了。
    //using Operations = VeryVerboseEnumOfThingsToDoWithNumbers;
#if !ENUM_EXTENSIENUM
    //暂时没想到怎么写成泛型，调用方法不符合，还是run(op, x, y) 而不是op.run(x ,y );
    public static float run(  Operations operation, float x , float y ){
        float t = operation switch
        {
            Operations.Add=>x + y,
            Operations.Subtract => x - y,
            _ => throw new ArgumentOutOfRangeException(nameof(operation), "Unsupported operation") // 添加默认情况
        };
        return t;
    }

 #else
    //写个新的 更方便的,注意这个扩展方法只能放在顶层的static class里面。

    public static float run(this Operations operation,  float x,  float y )
    {
        float t = operation switch
        {
            Operations.Add=>x + y,
            Operations.Subtract => x - y,
            _ => throw new ArgumentOutOfRangeException(nameof(operation), "Unsupported operation") // 添加默认情况
        };
        return t;   
    }
 #endif   
    
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        var pressed = new WE_KeyPress{ key = 'x'};
        // `to_owned()` 从一个字符串切片中创建一个具有所有权的 `String`。
        var pasted  = new WE_Paste{ content = "my text"};
        var click   = new WE_Click{ x = 20, y = 80 };
        var load    = new WE_PageLoad{};
        var unload  = new WE_PageUnload{};

        inspect(pressed);
        inspect(pasted);
        inspect(click);
        inspect(load);
        inspect(unload);

        // 我们可以通过别名引用每个枚举变量，避免使用又长又不方便的枚举名字
        var x = Operations.Add;
        #if ENUM_EXTENSIENUM
        Console.WriteLine("add 3, 5 = {0}",x.run(3,5));
        #else
        Console.WriteLine("add 3, 5 = {0}",run( x,3,5));
        #endif
    }
}
/*output:
pressed 'x'.
pasted "my text".
clicked at x=20, y=80.
page loaded
page unloaded
add 3, 5 = 8
*/