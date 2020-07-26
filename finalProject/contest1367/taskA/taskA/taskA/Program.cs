using System;

namespace taskA
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = int.Parse(Console.ReadLine());

            for (int i = 0; i < count; i++)
            {
                var str = Console.ReadLine();
                Console.Write(str[0]);
                for (int j = 1; j < str.Length - 1; j += 2)
                {
                    Console.Write(str[j]);
                }
                Console.WriteLine(str[str.Length - 1]);
            }
        }
    }
}