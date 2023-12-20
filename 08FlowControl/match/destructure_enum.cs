using System;
namespace System.Runtime.CompilerServices
{
    internal static class IsExternalInit {}
}
//C# 里面的enum 没有支持带数据的enum 用class继承吧。
//试试struct继承看看（不行的，包括abstract的修饰符 包括构造函数必须为public等限制，struct不支持继承。支持实现接口）
public abstract class Color{
    private Color(){}//防止外部继承
    public class Red : Color{};
    public class Blue:Color{};
    public class Green:Color{};
    public class RGB : Color { public int R, G, B; public RGB(int r, int g, int b) => (R, G, B) = (r, g, b); }
    public class HSV : Color { public int H, S, V; public HSV(int h, int s, int v) => (H, S, V) = (h, s, v); }
    public class HSL : Color { public int H, S, L; public HSL(int h, int s, int l) => (H, S, L) = (h, s, l); }
    public class CMY : Color { public int C, M, Y; public CMY(int c, int m, int y) => (C, M, Y) = (c, m, y); }
    public class CMYK : Color { public int C, M, Y, K; public CMYK(int c, int m, int y, int k) => (C, M, Y, K) = (c, m, y, k); }
}
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");

        Color color = new Color.RGB(122, 17, 40);
        Console.WriteLine("What color is it?");
        //当然写成正常的switch也可以
        var message = color switch{
            Color.Red=>"The color is Red!",
            Color.Blue=>"The color is Blue!",
            Color.Green=>"The color is Green!",
            Color.RGB rgb=>$"Red: {rgb.R}, green: {rgb.G}, and blue: {rgb.B}!",
            Color.HSV hsv=>$"Hue: {hsv.H}, saturation: {hsv.S}, value: {hsv.V}!",
            Color.HSL { H: var hue, S: var saturation, L: var lightness }=>$"Hue: {hue}, saturation: {saturation}, lightness: {lightness}!",
            Color.CMY cmy=>$"Cyan: {cmy.C}, magenta: {cmy.M}, yellow: {cmy.Y}!",
            Color.CMYK cmyk=>$"Cyan: {cmyk.C}, magenta: {cmyk.M}, yellow: {cmyk.Y}, key (black): {cmyk.K}!",
            _=>throw new ArgumentOutOfRangeException(nameof(color)),
        };
        Console.WriteLine( message );

    }
}
/*Output
What color is it?
Red: 122, green: 17, and blue: 40!
*/