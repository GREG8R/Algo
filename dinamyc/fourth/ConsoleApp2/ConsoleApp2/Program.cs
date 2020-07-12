using System;

namespace ConsoleApp2
{
    class Program
    {
        private static int[,] map;
        private static int n;

        static void Main(string[] args)
        {
            n = int.Parse(Console.ReadLine());
            map = new int[n, n];
            for (int y = 0; y < n; y++)
            {
                string[] strs = Console.ReadLine().Split(' ');
                for (int x = 0; x < n; x++)
                {
                    map[x, y] = int.Parse(strs[x]);
                }
            }

            int islands = 0;
            for (int y = 0; y < n; y++)
            {
                for (int x = 0; x < n; x++)
                {
                    if (map[x, y] == 1)
                    {
                        islands++;
                        go(x, y);
                    }
                }
            }

            Console.WriteLine(islands);
        }

        static void go(int x, int y)
        {
            if (x < 0 || x >= n) return;
            if (y < 0 || y >= n) return;
            if (map[x, y] == 0) return;
            map[x, y] = 0;
            go(x - 1, y);
            go(x + 1, y);
            go(x, y - 1);
            go(x, y + 1);
        }
    }
}