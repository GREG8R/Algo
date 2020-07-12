using System;

namespace ConsoleApp2
{
    class Program
    {
        static void Main(string[] args)
        {
            int n;
            n = int.Parse(Console.ReadLine());
            int[,] m = new int[n, n];

            for (int i = 0; i < n; i++)
            {
                string[] strs = Console.ReadLine().Split(' ');
                int k = 0;
                foreach (var str in strs)
                    if (str == "")
                        k++;
                    else
                        break;

                for (int j = 0; j <= i; j++)
                    m[i, j] = int.Parse(strs[j + k]);
            }

            for (int i = n - 2; i >= 0; i--)
            {
                for (int j = 0; j <= i; j++)
                {
                    m[i, j] += Math.Max(m[i + 1, j], m[i + 1, j + 1]);
                }
            }

            Console.WriteLine(m[0, 0]);
        }
    }
}