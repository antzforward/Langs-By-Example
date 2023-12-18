using System;
public class Program
{
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        #pragma warning disable CS0219 //变量“xxx”已被赋值，但从未使用过它的值
        // 变量可以给出类型说明。这里不能用readonly，readonly限制于属性
        //并且 隐式类型化的变量不能是常量 必须写清类型。
        const bool  logical = true;
        var  logical2 = true;
        
        const double  a_float = 1.0;  // 常规说明
        var     an_integer   = 5; // 后缀说明 默认就是i32

        // 否则会按默认方式决定类型。
        var  default_float   = 3.0; // `f64` 隐式类型化的变量不能是常量
        var  default_integer = 7;   // `i32`varvar

        // 类型也可根据上下文自动推断。
        var inferred_type = 12L; // C# 没有这么牛逼 “根据下一行的赋值推断为 i64 类型”
        inferred_type = 4294967296L;

        // 可变的（mutable）变量，其值可以改变。
        var mutable = 12; // Mutable `i32`var
        mutable = 21;

        // 报错！变量的类型并不能改变。
        //mutable = true; //无法将类型“bool”隐式转换为“int”

        // 但可以用遮蔽（shadow）来覆盖前面的变量。不能重名定义，不存在遮蔽的功能。
        // var mutable = true;

        #pragma warning restore CS0219 
    }
}