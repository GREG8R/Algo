using System;
using System.Collections.Generic;
using System.Linq;

namespace taskE
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] strs = Console.ReadLine().Split(' ');
            var n = int.Parse(strs[0]);
            var k = int.Parse(strs[1]);
            
            var times = new List<int>[4]
            {
                new List<int>(), 
                new List<int>(), 
                new List<int>(), 
                new List<int>(), 
            };
            var sums  = new List<int>[4]
            {
                new List<int>(), 
                new List<int>(), 
                new List<int>(), 
                new List<int>(), 
            };
            
            for (int i = 0; i < n; i++)
            {
                string[] currentStrs = Console.ReadLine().Split(' ');
                var t = int.Parse(currentStrs[0]);
                var a = int.Parse(currentStrs[1]);
                var b = int.Parse(currentStrs[2]);
                times[a * 2 + b].Add(t);
            }

            for (int i = 0; i < 4; i++)
            {
                times[i].Sort();
                sums[i].Add(0);
                foreach (var t in times[i])
                {
                    sums[i].Add(sums[i].Last() + t);
                }
            }
            
            var result = Int64.MaxValue;

            for (int i = 0; i < Math.Min(k + 1, sums[3].Count); i++)
            {
                if (k - i < sums[1].Count && k - i < sums[2].Count)
                {
                    result = Math.Min(result, sums[3][i] + sums[1][k - i] + sums[2][k - i]);
                }
            }
            
            if (result == Int64.MaxValue)
                result = -1;
            Console.WriteLine(result);
        }
    }
}