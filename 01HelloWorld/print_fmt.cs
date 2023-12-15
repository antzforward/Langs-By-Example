using System;

public class Program
{
    public struct City {
        public string name;
        public float  lat;
        public float  lon;

        public override string ToString(){
            var lat_c = this.lat >= 0.0f ? 'N': 'S';
            var lon_c = this.lon >= 0.0f ? 'E': 'W';
            //使用{index:F2}，注意不是.3F这种 这种C系的用法
            return string.Format("{0}: {1:F3}°{2} {3:F3}°{4}",name, Math.Abs(lat),lat_c, Math.Abs(lon),lon_c);
        }
    }
    public struct Color {
        public byte red;
        public byte green;
        public byte blue;

        public override string ToString(){
            //X2，不是2X
            return string.Format("RGB ({0},{1},{2}) 0x{0:X2}{1:X2}{2:X2}",
            red,green,blue);
        }
    }
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        //创建结构体的数组的形式之一，RAII形式最简模式了吧。但是类型推导失效 不能用var了
        City[] Cities = {
            new City { name = "Dublin",     lat = 53.347778f,    lon = -6.259722f },
            new City { name = "Oslo",       lat = 59.95f,        lon = 10.75f },
            new City { name = "Vancouver",  lat =  49.25f,       lon = -123.1f },
        };
        foreach( var city in Cities ){
            Console.WriteLine(city);
        }

        Color[] colors =  {
            new Color { red= 128,   green= 255, blue= 90 },
            new Color { red= 0,     green= 3,   blue= 254 },
            new Color { red= 0,     green= 0,   blue= 0 },
        };
        foreach (var color in colors)
        {
            Console.WriteLine(color);
        }
    }
}
/*output 目标
Dublin: 53.348°N 6.260°W
Oslo: 59.950°N 10.750°E
Vancouver: 49.250°N 123.100°W
RGB (128,255,90) 0x80FF5A
RGB (0,3,254) 0x0003FE
RGB (0,0,0) 0x000000
*/