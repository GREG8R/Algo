using System;
using System.Linq;

namespace taskC
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = int.Parse(Console.ReadLine());

            for (int c = 0; c < count; c++)
            {
                var strs = Console.ReadLine().Split(' ').Select(xx => int.Parse(xx)).ToList();
                var n = strs[0];
                var k = strs[1];
                var str = Console.ReadLine();
                var result = 0;

                for (int i = 0; i < n;) {
                    int j = i + 1;

                    for (; j < n && str[j] != '1'; j++);

                    int left = str[i] == '1' ? k : 0;
                    int right = j < n && str[j] == '1' ? k : 0;
                    int len = j - i;

                    if (left == k) {
                        len--;
                    }

                    len -= left + right;

                    if (len > 0) {
                        result += (len + k) / (k + 1);
                    }

                    i = j;
                }
                Console.WriteLine(result);
            }
        }
    }
}