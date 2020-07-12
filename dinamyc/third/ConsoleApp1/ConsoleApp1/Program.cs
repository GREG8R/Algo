using System;

namespace ConsoleApp1
{
    class Program
    {
        static void Main(string[] args)
        {
        int n;
        n = int.Parse(Console.ReadLine());

        long a = 1;
        long c = 0;
        long a1, c1;

        for (int i = 2; i <= n; i++)
        {
            a1 = a + c;
            c1 = a;
            a = a1;
            c = c1;
        }

        Console.WriteLine(a + a + c + c);
        }
    }
}