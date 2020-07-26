using System;

namespace taskA
{
    class Program
    {
        static void Main(string[] args)
        {
	        var x = 0;
	        var y = 0;
	        var n = 0;
	        var count = 0;
	        count = int.Parse(Console.ReadLine());
	        for (int i = 0; i < count; i++)
	        {
		        string[] strs = Console.ReadLine().Split(' ');
		        x = int.Parse(strs[0]);
		        y = int.Parse(strs[1]);
		        n = int.Parse(strs[2]);

		        var k = n - (n % x) + y;
		        if (k <= n)
			        Console.WriteLine(k);
		        else
			        Console.WriteLine(k - x);
	        }
        }
    }
}
