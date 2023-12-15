using System;

//添加这行代码消除掉“CS0518: 预定义类型“System.Runtime.CompilerServices.IsExternalInit”未定义或导入”
namespace System.Runtime.CompilerServices
{
    internal static class IsExternalInit { }
}

public class Program
{
    public struct Structure{
        int value;
        public Structure(int value){
            this.value = value;
        }
        public override string ToString(){
            return value.ToString();
        }
    }

    //扩展方法必须在非泛型静态类中定义
    public struct MinMax{
        int a;
        int b;

        public MinMax(int a,int b){
            this.a = a;
            this.b = b;
        }
        public override string ToString(){
            return $"({a}, {b})";
        }
        public string ToDebugString(){
            //@开头表示逐字字符产，其中\不会被视为转移字符串起始符。
            // 在行尾加\ 是不行的。 当然这个也可以取消转义字符\的作用
            return $"MinMax({a}, {b})";
        }
    }
    
    public struct Point2D{
        double x;
        double y;
        public Point2D(double x,double y){
            this.x = x;
            this.y = y;
        }
        public override string ToString(){
            return $"x: {x}, y: {y}";
        }
        public string ToBinary(){
            var xbin =  Convert.ToString(BitConverter.DoubleToInt64Bits(x), toBase: 2).PadLeft(10, '0');
            var ybin =  Convert.ToString(BitConverter.DoubleToInt64Bits(y), toBase: 2).PadLeft(10, '0');
            return $"x: 0b{xbin}, y: 0b{ybin}";
        }
        public string ToDebugString(){
            //@开头表示逐字字符产，其中\不会被视为转移字符串起始符。
            // 在行尾加\ 是不行的。 当然这个也可以取消转义字符\的作用
            return $"Point2D{{ x: {x}, y: {y} }}";
        }
    }

    public struct Complex {
        float real;
        float imag;

        public Complex(float real, float imag ){
            this.real = real;
            this.imag = imag;
        }

        public override string ToString(){
            return $"{real}+{imag}i";
        }
        public string ToDebugString(){
            //@开头表示逐字字符产，其中\不会被视为转移字符串起始符。
            // 在行尾加\ 是不行的。 当然这个也可以取消转义字符\的作用
            return @$"Complex{{ real: {real}, imag: {imag} }}";
        }
    }
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        var minmax = new MinMax(0, 14);

        Console.WriteLine("Compare structures:");
        Console.WriteLine($"Display: {minmax}");
        Console.WriteLine("Debug: {0}", minmax.ToDebugString());

        var big_range =   new  MinMax(-300, 300);
        var small_range = new  MinMax(-3, 3);

        Console.WriteLine($"The big range is {big_range} and the small is {small_range}");

        var point = new  Point2D( 3.3, 7.2 );

        Console.WriteLine("Compare points:");
        Console.WriteLine("Display: {0}", point);
        Console.WriteLine("Debug: {0}", point.ToDebugString());

        // 报错。`Debug` 和 `Display` 都被实现了，但 `{:b}` 需要 `fmt::Binary`
        // 得到实现。这语句不能运行。
        Console.WriteLine("What does Point2D look like in binary: {0}?", point.ToBinary());
        
        var complex = new  Complex(  3.3f, 7.2f );
        Console.WriteLine("Compare Complex:");
        Console.WriteLine("Display: {0}", complex);
        Console.WriteLine("Debug: {0}", complex.ToDebugString());
    }
}
/*output:
Compare structures:
Display: (0, 14)
Debug: MinMax(0, 14)
The big range is (-300, 300) and the small is (-3, 3)
Compare points:
Display: x: 3.3, y: 7.2
Debug: Point2D { x: 3.3, y: 7.2 }
What does Point2D look like in binary: x: 0b100000000001010011001100110011001100110011001100110011001100110, y: 0b100000000011100110011001100110011001100110011001100110011001101?
Compare Complex:
Display: 3.3+7.2i
Debug: Complex { real: 3.3, imag: 7.2 }
*/