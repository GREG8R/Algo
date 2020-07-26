using System;

namespace taskB
{
    class Program
    {
        static void Main(string[] args)
        {
            var count = 0;
            var currentNumber = 0;
            count = int.Parse(Console.ReadLine());
            var k = 0;
            for (int i = 0; i < count; i++)
            {
                currentNumber = int.Parse(Console.ReadLine());
                if (currentNumber == 1)
                {
                    Console.WriteLine(0);
                    continue;
                }
                if (currentNumber % 3 != 0)
                {
                    Console.WriteLine(-1);
                    continue;
                }

                while (currentNumber > 0)
                {
                    if (currentNumber % 6 == 0)
                    {
                        currentNumber /= 6;
                        k++;
                    }
                    else if (currentNumber % 3 == 0)
                    {
                        currentNumber *= 2;
                        k++;
                    }
                    else if (currentNumber == 1)
                    {
                        Console.WriteLine(k);
                        break;
                    }
                    else
                    {
                        Console.WriteLine(-1);
                        break;
                    }
                }
                k = 0;
            }
        }
    }
}