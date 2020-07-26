using System;
using System.Collections.Generic;

namespace taskA
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = int.Parse(Console.ReadLine());

            for (int i = 0; i < count; i++)
            {
                string[] strs = Console.ReadLine().Split(' ');
                var x= ulong.Parse(strs[0]);
                var y = ulong.Parse(strs[1]);
                var z = ulong.Parse(strs[2]);
                var arr = new List<ulong>{x, y, z};
                arr.Sort();
                if (arr[1] != arr[2])
                    Console.WriteLine("NO");
                else
                {
                    Console.WriteLine("YES");
                    Console.WriteLine($"{arr[0]} {arr[0]} {arr[2]}");
                }
            }
        }
    }
}