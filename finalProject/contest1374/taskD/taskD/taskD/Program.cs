using System;
using System.Collections.Generic;

namespace taskD
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = 0;
            count = int.Parse(Console.ReadLine());
            for (int i = 0; i < count; i++)
            {
                string[] strs = Console.ReadLine().Split(' ');
                var n = int.Parse(strs[0]);
                var k = int.Parse(strs[1]);

                strs = Console.ReadLine().Split(' ');
                
                var m = new Dictionary<int, ulong>();
                ulong max = 0;
                for (int j = 0; j < n; j++)
                {
                    var current = int.Parse(strs[j]); 
                    if (current % k == 0)
                        continue;
                    if (!m.ContainsKey(k - current % k))
                        m.Add(k - current % k, 0);
                    
                    m[k - current % k]++;
                    if (m[k - current % k] > max)
                        max = m[k - current % k];
                }

                ulong result = 0;
                var maxI = 0;
                foreach (var c in m)
                {
                    if (c.Value == max && c.Key > maxI)
                    {
                        result = Convert.ToUInt64(k) * (c.Value - 1) + Convert.ToUInt64(c.Key) + 1;
                        maxI = c.Key;
                    }
                }
                Console.WriteLine(result);
            }
        }
    }
}
