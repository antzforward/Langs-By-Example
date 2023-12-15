using System;
//引入核心库，实现ToDebugString，用来实现#[derive(Debug)]
using System.Reflection;
using System.Text;
//一般显示形式,用csc 编译单一文件的模式下 找不到对应的dll文件
//using System.Text.Json;
//其实C#并没有Rust中类似#[derive(Debug)]的编译属性
//通常有实现class的ToString方法，或者将对象序列化为Json字符串等
//当然Rust的对象以struct为主，要使用上面的功能就只能是class了。

//要求测试用例输出是
/*输出的内容如下
12 months in a year.
"Christian" "Slater" is the "actor's" name.
Christian Slater is the actor's name.
Now Structure(3) will print!
Now Deep(Structure(7)) will print!
Person {
    name: "Peter",
    age: 27,
}
*/

//使用csc -langversion:?能查询到当前.net framework 支持的c#版本的列表。
//可以使用csc -langversion:8.0 将语言版本设置成8.0

//添加这行代码消除掉“CS0518: 预定义类型“System.Runtime.CompilerServices.IsExternalInit”未定义或导入”
namespace System.Runtime.CompilerServices
{
    internal static class IsExternalInit { }
}
public class Program
{
    //申明record
    public record UnPrintable(int value);
    public record DebugPrintable(int value);
    public record Structure(int value);

    public record Deep(Structure value);

    public struct Person{
        public string name;//Rust有主动生命周期的管理，C# 是基于gc的，这里写法有不同
        public byte   age;
        public override string ToString(){
            var sb = new StringBuilder();
            //这里的\t显示是8个空格，比rust下的效果更难看。
            sb.Append("Person ").AppendLine("{");
            sb.Append("    name:").AppendLine($"\"{name}\",");
            sb.Append("    age:").AppendLine($"{age},");
            sb.Append("}");
            return sb.ToString();
        }
    }
    //实现通用的ToDebugString的方法,写成模板函数的形式
    public static string ToDebugString<T>(T record) where T : notnull{
        var type = typeof(T);
        var properties = type.GetProperties(BindingFlags.Instance | BindingFlags.Public);
        var sb = new StringBuilder();
        
        sb.Append(type.Name).Append("(");
        
        foreach (var property in properties)
        {
            var value = property.GetValue(record);
            var formattedValue = FormatValue(value, property.PropertyType);
            sb.Append(formattedValue).Append(", ");
        }
        
        // Remove the trailing comma and space if there are properties.
        if (properties.Length > 0)
        {
            sb.Length -= 2;
        }
        sb.Append(")");
        return sb.ToString();
    }
    private static string FormatValue(object value, Type type)
    {
        // Check for null values first.
        if (value == null)
            return "null";

        // Add special formatting for strings to include double quotes.
        if (type == typeof(string))
            return $"\"{value}\"";

        // Additional type checks and formatting can be added here.
        if( type == typeof(Structure))
            return ToDebugString<Structure>(value as Structure);
        
        return value.ToString();
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        Console.WriteLine("{0} months in a year.", FormatValue(12,typeof(int)));
        // 使用 ToDebugString 打印
        Console.WriteLine("{1} {0} is the {2} name.",
             FormatValue("Slater",typeof(string)),
             FormatValue("Christian",typeof(string)),
             FormatValue("actor's",typeof(string)));
        Console.WriteLine("{1} {0} is the {2} name.",
             "Slater",
             "Christian",
             "actor's");
        // `Structure` 也可以打印！
        //不可调用的成员“Program.Structure”不能像方法一样使用 只能用
        Console.WriteLine("Now {0} will print!", ToDebugString(new Structure(3 )) );

        Console.WriteLine("Now {0} will print!",ToDebugString(new  Deep(new Structure(7))));

        var name = "Peter";
        byte age = 27;
        var peter = new Person{name=name, age=age };
        Console.WriteLine("{0}", peter);
    }
}
/*
This is C# Language
12 months in a year.
"Christian" "Slater" is the "actor's" name.
Christian Slater is the actor's name.
Now Structure(3) will print!
Now Deep(Structure(7)) will print!
Person {
    name:"Peter",
    age:27,
}
*/