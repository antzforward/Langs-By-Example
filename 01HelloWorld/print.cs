using System;

public class Program
{
	//不支持在函数内定义类型。
	struct Structure{
		public int value;
		public override string ToString(){
			return $"Structure {{ value={value} }} ";
		}
	};
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
		//C#里面不能省略，没有这个语法糖
		// 变量内容会转化成字符串。
		Console.WriteLine("{0} days", 31);

		// 不加后缀的话，31 就自动成为 i32 类型。
		// 你可以添加后缀来改变 31 的类型（比如L或l 表示long，UL或ul表示ulong，U或u 表示u）。
		// 如果是浮点数，可以（F或f为float类型，注意这里默认是double类型，D或d，指定为double，M或m 指定为decimal类型）

		// 用变量替换字符串有多种写法。
		// 比如可以使用位置参数。
		Console.WriteLine("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

		// 可以使用命名参数。
		//这里使用interploted string  内插字符串，这里注意如果变量命是包留字符串可以用@来标记，下面就是用法
		var @object="the lazy dog";
		var subject="the quick brown fox";
		var verb="jumps over";
		//变量要提前声明并实现，这里效果为rust效果更好。 interpolated string要用$标注。
		Console.WriteLine($"{subject} {verb} {@object}");

		// 可以在 `:` 后面指定特殊的格式。默认支持的不多，{:X}是有的。其他的2，8进制要调用Convert.ToString来实现
		Console.WriteLine("{0} of {1} people know binary, the other half don't", 1, Convert.ToString(2,2));

		// 你可以按指定宽度来右对齐文本。
		// 下面语句输出 "     1"，5 个空格后面连着 1。
		//{index[,alignment][:formatString]} 其中index表示后续参数的序号。注意index数量要小于10
		//formatString表示为X 19进制，E 比如1.222e10，N表示千进制表示比如1，999.00
		//alignment，正数表示右对齐，负数表示左对齐
		var number = 1;
		var width = 6;
		var format = $"{{0,{width}}}";//此处展示特殊的{}形式，用\{是不行的
		Console.WriteLine(format, number);
		//$"{number:D{{width}}}" 是不行的，{}内部不进行内插解析了，要用其他的方式来写 比如 下面这句。

		// 你可以在数字左边补 0。下面语句输出 "000001"。
		Console.WriteLine(number.ToString($"D{width}"));

		// Console.WriteLine 会检查使用到的参数数量是否正确。
		Console.WriteLine("My name is {0}, {1} {0}", "Bond","James");
		// 改正 ^ 补上漏掉的参数："James"

		// 创建一个包含单个 `i32` 的结构体（structure）。命名为 `Structure`。
		var mystructure = new Structure{ value = 3 };
		Console.WriteLine(mystructure.ToString());//要打印，实现ToString的方法，或者找到parser json的工具等。
		var precision=3;
		var doubNumber = 3.141592;
		Console.WriteLine("{0}",doubNumber.ToString($"F{precision}")) ;
    }
}