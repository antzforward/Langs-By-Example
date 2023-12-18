using System;
public class Program
{
    public struct Person{
        public string name;
        public byte   age;
    }

    // 单元结构体
    public struct Unit{}

    // 元组结构体
    public struct Pair{
        public int Item1;
        public float Item2;
        public Pair(int _0, float _1 ){
            Item1 = _0;
            Item2 = _1;
        }
    }

    public struct  Point{
        public float x;
        public float y;

        public override string ToString()
        {
            return $"x: {x},y: {y}";
        }
    }

    public struct Rectangle{
        public Point top_left;
        public Point bottom_right;

        public float rect_area{
            var width = top_left.x - bottom_right.x;
            var height = top_left.y - bottom_right.y;
            return Math.Abs( width * height );
        }
        public static Rectangle square( Point point, float width ){
            return new Rectangle{
                top_left = point,//use struct copy
                bottom_right = new Point{ x = point.x + width, y = point.y + width }
            };
        }
    }
    
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        // 使用简单的写法初始化字段，并创建结构体
        var name = "Peter";
        byte age = 27;
        var peter = new Person { name, age };

        // 以 Debug 方式打印结构体
        Console.WriteLine("{0}", peter);

        // 实例化结构体 `Point`
        Point point = Point { x = 10.3, y = 0.4 };

        // 访问 point 的字段
        Console.WriteLine("point coordinates: ({0}, {1})", point.x, point.y);

        // 使用结构体更新语法创建新的 point，
        // 这样可以用到之前的 point 的字段,这里就没办法了，要逐一指定
        var bottom_right = Point { x =  5.2, y = point.y };

        // `bottom_right.y` 与 `point.y` 一样，因为这个字段就是从 `point` 中来的
        Console.WriteLine("second point: ({0}, {1})", bottom_right.x, bottom_right.y);

        // 使用 `let` 绑定来解构 point,只能用tuple来解构了
        var ( left_edge,  top_edge ) = point;

        let _rectangle = Rectangle {
            // 结构体的实例化也是一个表达式
            top_left =  Point { x= left_edge, y= top_edge },
            bottom_right = bottom_right,
        };
        
        Console.WriteLine("rectangle area:{0}",_rectangle.rect_area());
        Console.WriteLine("rectangle area:{0}",Rectangle::square(Point{x=0.0,y=0f}, 35.0f).rect_area());
        // 实例化一个单元结构体
        var _unit = Unit;

        // 实例化一个元组结构体
        var pair = new Pair(1, 0.1);

        // 访问元组结构体的字段
        Console.WriteLine("pair contains {:?} and {:?}", pair.0, pair.1);

        // 解构一个元组结构体
        let Pair(integer, decimal) = pair;

        Console.WriteLine("pair contains {:?} and {:?}", integer, decimal);
    }
}
/*output
Person { name: "Peter", age: 27 }
point coordinates: (10.3, 0.4)
second point: (5.2, 0.4)
rectangle area:50.490005
rectangle area:1225
pair contains 1 and 0.1
pair contains 1 and 0.1
*/