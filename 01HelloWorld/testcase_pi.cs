using System;
using System.Diagnostics;
using System.Text;

class Program
{
    static void Main()
    {
		//记录初始内存使用量
		long initialMemory = Process.GetCurrentProcess().WorkingSet64;
		const long test_count = 1000;
        // 测试PI的近似值的数值积分代码
        const long num_rects = 100_000;
        double mid,height,width=1.0/(double)num_rects,sum=0.0;
        var sw = Stopwatch.StartNew();
        sum = 0.0;
        for(int c=0; c<test_count ;c++){
			for (int i = 0; i < num_rects; i++)//积分方式误差更大
			{
				mid = (i + 0.5) * width;//这个取中值，是积分的常规误差
				height = 4.0/(1.0 + mid * mid );
				sum += height ;
			}
        }
        sum *= width;
        sum /= test_count ;
        sw.Stop();
        Console.WriteLine($"积分方法时间占用 took: {sw.ElapsedMilliseconds} ms ,{Math.PI} : {{0}}", sum.ToString("F12") );
		long memoryAfterStringInterpolation = Process.GetCurrentProcess().WorkingSet64;
		GC.Collect();
		GC.WaitForPendingFinalizers();
		GC.Collect();
		
        // 测试StringBuilder
        sw.Restart();
        sum = 0.0;
        for(int c=0; c<test_count ;c++){
			sum += Math.Atan(1.0)*4.0;
		}
		sum /= test_count;
        sw.Stop();
        Console.WriteLine($"函数调用 Atan(1.0)*4.0 took: {sw.ElapsedMilliseconds} ms ,{Math.PI} :  {{0}}",sum.ToString("F12"));
        
        // 计算StringBuilder后的内存使用量
        long memoryAfterStringBuilder = Process.GetCurrentProcess().WorkingSet64;

        Console.WriteLine($"Memory used by 积分方法: {memoryAfterStringInterpolation - initialMemory} bytes");
        Console.WriteLine($"Memory used by 函数调用: {memoryAfterStringBuilder - memoryAfterStringInterpolation} bytes");
        
    }
}