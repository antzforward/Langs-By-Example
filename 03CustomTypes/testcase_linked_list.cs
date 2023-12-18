using System;
using System.Text;
public class Program
{
    // 可以为 enum 定义方法
    public class SingleLinkedList<T> where T:notnull 
    {
        public class Node {
            public T  Value{get;set;}
            public Node Next {get;set;}
            public Node(T value){
                Value = value;
                Next  = null;
            }
        }
        private Node Head{get; set;}

        private SingleLinkedList() {
            Head = null;
        }
        // 创建一个空的 List 实例
        public static SingleLinkedList<T> New()  {
            return new SingleLinkedList<T>();
        }

        // 处理一个 List，在其头部插入新元素，并返回该 List
        public void  prepend( T elem)  {
            var newNode = new Node(elem) { Next = Head };
            Head = newNode;
        }

        // 返回 List 的长度
        public uint len() {
            var tail = Head;
            uint count = 0;
            while( tail != null )
            {
                count++;
                tail = tail.Next;    
            }
            return count;
        }

        // 返回列表的字符串表示（该字符串是堆分配的）
        public string stringify() {
            var tail = Head;
            StringBuilder sb = new StringBuilder();
            while( tail != null )
            {
                sb.Append($"{tail.Value}, ");
                tail = tail.Next;    
            }
            sb.Append("Nil");
            return sb.ToString();
        }
    }

    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        
        // 创建一个空链表
        var list = SingleLinkedList<int>.New();

        // 追加一些元素
        list.prepend(1);
        list.prepend(2);
        list.prepend(3);

        // 显示链表的最后状态
        Console.WriteLine("linked list has length: {0}", list.len());
        Console.WriteLine("{0}", list.stringify());
    }
}
/*output
linked list has length: 3
3, 2, 1, Nil
*/