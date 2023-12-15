using System;
using System.Text;//stringbuilder
using System.Collections.Generic; //for list
using System.Linq;//for select

public class Program
{
    public struct Vec{
        public List<int> vec;

        public override string ToString(){
            StringBuilder sb = new StringBuilder();
            if( vec != null )
            {
                sb.Append('[');
                foreach(var item in vec.Select((value,index)=>new {Index=index,Value=value})){
                    if( item.Index != 0 ) sb.Append(", ");
                    sb.Append(item.Value);
                }
                sb.Append(']');    
            }
            return sb.ToString();
        }
        public string ToDebugString(){
            StringBuilder sb = new StringBuilder();
            if( vec != null )
            {
                sb.Append('[');
                foreach(var item in vec.Select((value,index)=>new {Index=index,Value=value})){
                    if( item.Index != 0 ) sb.Append(", ");
                    sb.Append($"{item.Index}:{item.Value}");
                }
                sb.Append(']');    
            }
            return sb.ToString();
        }
    }
    public static void Main(string[] args)
    {
		Console.WriteLine("This is C# Language");
        var v = new Vec{vec = new List<int>{1, 2, 3}};
        Console.WriteLine(v);
        Console.WriteLine(v.ToDebugString());
    }
}