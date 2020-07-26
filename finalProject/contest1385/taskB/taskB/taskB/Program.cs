using System;
using System.Collections.Generic;

namespace taskB
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = int.Parse(Console.ReadLine());

            for (int i = 0; i < count; i++)
            {
                var n = int.Parse(Console.ReadLine());
                string[] strs = Console.ReadLine().Split(' ');

                var map = new Dictionary<int, bool>();
                foreach (var str in strs)
                {
                    var c = int.Parse(str);
                    if (!map.ContainsKey(c))
                    {
                        Console.Write($"{c} ");
                        map.Add(c, true);
                    }
                }
                Console.WriteLine();
            }
            
        }
    }
}