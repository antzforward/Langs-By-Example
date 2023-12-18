using System;
public class Program
{
    // 元组可以充当函数的参数和返回值
    public static (bool,int ) reverse((int , bool) pair)  {
        // 可以使用 `let` 把一个元组的成员绑定到一些变量
        (int integer, bool boolean) = pair;

        return (boolean, integer);
    }

    // 在 “动手试一试” 的练习中要用到下面这个结构体。
    public struct Matrix{
        public float Item1;
        public float Item2;
        public float Item3;
        public float Item4;
        public Matrix(float _0,float _1, float _2, float _3){
            Item1 = _0;
            Item2 = _1;
            Item3 = _2;
            Item4 = _3;
        }
        public override string ToString()
        {
            return string.Format($"( {Item1} {Item2} )\n( {Item3} {Item4} )");
        }
       
    };

     public static Matrix transpose(Matrix mat){
        return new Matrix{
            Item1= mat.Item1,
            Item2 = mat.Item3,
            Item3 = mat.Item2,
            Item4 = mat.Item4
            };
    }
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");

        // 包含各种不同类型的元组,由于C#不能添加太多字面标识，所以要用变量来定义,rust有这么多细节 应该是打算支持FPGA之类的设备吧
        ushort      u16 = 2;
        byte        u8  = 1;
        sbyte       i8  = -1;
        short       i16 = -2;
        var long_tuple = (u8, u16, 3u, 4ul,
                        i8, i16, -3, -4L,
                        0.1f, 0.2,
                        'a', true);

        // 通过元组的下标来访问具体的值
        Console.WriteLine("long tuple first value: {0}", long_tuple.Item11);
        Console.WriteLine("long tuple second value: {0}", long_tuple.Item12);
        Console.WriteLine("long tuple :{0}",long_tuple);
        // 元组也可以充当元组的元素
        var tuple_of_tuples = ((u8, u16, 2u), (4ul, i8), i16);

        // 元组可以打印
        Console.WriteLine("tuple of tuples: {0}", tuple_of_tuples);

        // 但很长的元组无法打印:但是第一个的long_tuple 有 12个elements,下面这个最多到12，13个就超过模板了。
        var too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12);//13个就编译出错了。
        Console.WriteLine("too long tuple: {0}", too_long_tuple);
        // 试一试 ^ 取消上面两行的注释，阅读编译器给出的错误信息。

        var pair = (1, true);
        Console.WriteLine("pair is {0}", pair);

        Console.WriteLine("the reversed pair is {0}", reverse(pair));

        // 创建单元素元组需要一个额外的逗号，这是为了和被括号包含的字面量作区分。
        //Console.WriteLine("one element tuple: {0}", (5, ));//这是tuple,要把空格空的明显一点。(5, )不行啊 在Python里面应该可以的
        Console.WriteLine("one element tuple: {0}", new ValueTuple<int>(5));
        Console.WriteLine("just an integer: {0}", (5u));//这是普通字面量

        // 元组可以被解构（deconstruct），从而将值绑定给变量
        var tuple = (1, "hello", 4.5, true);

        var (a, b, c, d) = tuple;
        Console.WriteLine("{0}, {1}, {2}, {3}", a, b, c, d);
        
        var (a1, b1, c1, _) = tuple;
        Console.WriteLine("{0}, {1}, {2}, {3} ", a1, b1, c1 ,d);
        
        var matrix = new Matrix(1.1f, 1.2f, 2.1f, 2.2f);
        Console.WriteLine("{0}", matrix);
        Console.WriteLine("Matrix:\n{0}", matrix);
        Console.WriteLine("Transpose:\n{0}", transpose(matrix));
    }
}
/*
long tuple first value: 1
long tuple second value: 2
long tuple :(1, 2, 3, 4, -1, -2, -3, -4, 0.1, 0.2, 'a', true)
tuple of tuples: ((1, 2, 2), (4, -1), -2)
too long tuple: (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
pair is (1, true)
the reversed pair is (true, 1)
one element tuple: (5,)
just an integer: 5
1, "hello", 4.5, true
1, "hello", 4.5, true
Matrix(1.1, 1.2, 2.1, 2.2)
Matrix:
( 1.1 1.2 )
( 2.1 2.2 )
Transpose:
( 1.1 2.1 )
( 1.2 2.2 )
*/