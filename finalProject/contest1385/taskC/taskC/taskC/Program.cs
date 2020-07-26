using System;
using System.Linq;

namespace taskC
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = int.Parse(Console.ReadLine());
            for (int i = 0; i < count; i++)
            {
                var n = int.Parse(Console.ReadLine());
                var inputArray = Console.ReadLine().Split(' ').Select(x => int.Parse(x)).ToArray();
                var p = n - 1;
                while (p > 0 && inputArray[p - 1] >= inputArray[p]) p--;
                while (p > 0 && inputArray[p - 1] <= inputArray[p]) p--;
                Console.WriteLine(p);
            }
        }
    }
}