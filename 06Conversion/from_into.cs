using System;
public class Program
{
    public struct Number{
        public int value;
        public static Number from(int item ){
            return new Number{value = item};
        }
        public override string ToString()
        {
            return $"Number {{ value: {value} }}";
        }
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");

        var intv = 5;
        // 试试删除类型说明
        Number num = Number.from(intv);
        Console.WriteLine("My number is {0}", num);

        var my_str = "hello";
        var my_string = my_str.Clone();
        Console.WriteLine("My string is {0} : {1}", my_str,my_string);

        var num2 = Number.from(30);
        Console.WriteLine("My number is {0}", num2);

    }
}
/*
My number is Number { value: 5 }
My string is hello : hello
My number is Number { value: 30 }
*/