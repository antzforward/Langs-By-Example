using System;
using System.Diagnostics;
using System.Text;

class Program
{
    static void Main()
    {
        long times = 1_000_000;
		
		//记录初始内存使用量
		long initialMemory = Process.GetCurrentProcess().WorkingSet64;
        // 测试字符串插值
        var sw = Stopwatch.StartNew();
        string s = "" ;
        for (int i = 0; i < times; i++)
        {
            s = $"this is {i} testing;\nthis is {i} testing;\nthis is {i} testing;\nthis is {i} testing;\nthis is {i} testing;\n"; // 字符串插值
        }
        sw.Stop();
        Console.WriteLine($"String interpolation took: {sw.ElapsedMilliseconds} ms", s );
		long memoryAfterStringInterpolation = Process.GetCurrentProcess().WorkingSet64;
		GC.Collect();
		GC.WaitForPendingFinalizers();
		GC.Collect();
		
        // 测试StringBuilder
        sw.Restart();
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < times; i++)
        {
            sb.Clear(); // 清除StringBuilder的内容
            sb.AppendFormat("this is {0} testing;\nthis is {0} testing;\nthis is {0} testing;\nthis is {0} testing;\nthis is {0} testing;\n", i);
            s = sb.ToString();
        }
        sw.Stop();
        Console.WriteLine($"StringBuilder took: {sw.ElapsedMilliseconds} ms",s);
        
        // 计算StringBuilder后的内存使用量
        long memoryAfterStringBuilder = Process.GetCurrentProcess().WorkingSet64;

        Console.WriteLine($"Memory used by String Interpolation: {memoryAfterStringInterpolation - initialMemory} bytes");
        Console.WriteLine($"Memory used by StringBuilder: {memoryAfterStringBuilder - memoryAfterStringInterpolation} bytes");
        
    }
}