using System;

namespace taskC
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = int.Parse(Console.ReadLine());
            for (int i = 0; i < count; i++)
            {
                var len = int.Parse(Console.ReadLine());
                var k = 0;
                string strs = Console.ReadLine();
                var balance = 0;
                for (int j = 0; j < len; j++)
                {
                    if (strs[j] == '(')
                        balance++;
                    else
                        balance--;
                    
                    if (balance < 0)
                    {
                        k++;
                        balance++;
                    }
                }
                Console.WriteLine(k);
            }
        }
    }
}