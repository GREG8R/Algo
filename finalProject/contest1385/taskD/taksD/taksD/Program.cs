using System;
using System.Linq;

namespace taksD
{
    class Program
    {
        static void Main(string[] args)
        {
            var count =  int.Parse(Console.ReadLine());
            for (int i = 0; i < count; i++)
            {
                var n = int.Parse(Console.ReadLine());
                var inputStr = Console.ReadLine();
                Console.WriteLine(CalculateCount(inputStr, 'a'));
            }
        }

        static int CalculateCount(string str, char c)
        {
            if (str.Length == 1)
                return str[0] != c ? 1 : 0;

            var middle = str.Length / 2;
            var countLeftPart = CalculateCount(str.Substring(0, middle), Next(c));
            countLeftPart += str.Length / 2 - 
                             str.Substring(middle, str.Length - middle)
                                 .Count(x => x == c);
            var countRightPart = CalculateCount(str.Substring(middle, str.Length - middle), Next(c));
            countRightPart += str.Length / 2 -
                              str.Substring(0, middle)
                                  .Count(x => x == c);
            return Math.Min(countLeftPart, countRightPart);
        }

        static char Next(char c)
        {
            return (char)(c + 1);
        }
    }
}