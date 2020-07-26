using System;
using System.Linq;

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
                var input = Console.ReadLine().Split(' ').Select(xx => int.Parse(xx)).ToList();
                
                var countOddOutOfPlace = 0;
                var countEvenOutOfPlace = 0;
                for (int k = 0; k < n; k++)
                {
                    if (k % 2 != input[k] % 2 && input[k] % 2 == 0)
                        countEvenOutOfPlace++;
                    if (k % 2 != input[k] % 2 && input[k] % 2 != 0)
                        countOddOutOfPlace++;
                }
                
                if (countEvenOutOfPlace != countOddOutOfPlace)
                    Console.WriteLine(-1);
                else
                    Console.WriteLine(countEvenOutOfPlace);
            }
        }
    }
}