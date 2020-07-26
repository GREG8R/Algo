using System;
using System.Collections.Generic;
using System.Linq;

namespace taskE
{
    class Program
    {
        static void Main(string[] args)
        {
            var go = new Go();
            go.GoGo();
        }
    }

    class Go
    {
        public List<int> ord { get; set; }
        public List<bool> used { get; set; }
        public List<List<int>> g { get; set; }
        public void DFS(int k) {
            used[k] = true;
            foreach (var c in g[k])
            {
                if (!used[c]) DFS(c);
            }
            ord.Add(k);
        }
        
        public void GoGo()
        {
            var count = int.Parse(Console.ReadLine());
            for (int i = 0; i < count; i++)
            {
                var nm = Console.ReadLine().Split(' ').ToList().Select(xx => int.Parse(xx)).ToList();
                var n = nm[0];
                var m = nm[1];
                g = new List<List<int>>();
                ord = new List<int>();
                used = new List<bool>();
                var pos = new List<int>();
                for (int k = 0; k < n; k++)
                {
                    g.Add(new List<int>());
                    ord.Add(0);
                    used.Add(false);
                    pos.Add(0);
                }
                var edges = new List<Edge>();
    
                for (int j = 0; j < m; j++)
                {
                    var input = Console.ReadLine().Split(' ').ToList().Select(xx => int.Parse(xx)).ToList();
                    var t = input[0];
                    var x = input[1];
                    var y = input[2];
                    x--; y--;
                    if (t == 1)
                        g[x].Add(y);
                    edges.Add(new Edge(x, y));
                }
    
                for (int k = 0; k < n; k++)
                {
                    if (!used[k])
                    {
                        DFS(k);
                    }
                }
                
                ord.Reverse();
                for (int k = 0; k < n; k++)
                    pos[ord[k]] = k;
                
                bool bad = false;
                for (int k = 0; k < n; k++)
                    foreach (var c in g[k])
                        if (pos[k] > pos[c])
                        {
                            bad = true;
                        }

                if (bad) 
                {
                    Console.WriteLine("NO");
                } 
                else 
                {
                    Console.WriteLine("YES");
                    foreach (var edge in edges)
                    {
                        if (pos[edge.X] < pos[edge.Y])
                        {
                            Console.WriteLine($"{edge.X + 1} {edge.Y + 1}");
                        }
                        else
                        {
                            Console.WriteLine($"{edge.Y + 1} {edge.X + 1}");
                        }
                    }
                }
            }
        }
    }

    public class Edge
    {
        public Edge(int x, int y)
        {
            X = x;
            Y = y;
        }
        public int X { get; set; }
        public int Y { get; set; }
    }
}