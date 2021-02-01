using System;

namespace app
{
    class Program
    {
        static void Main(string[] args)
        { 
            double radius = 2.50;
            double area = Math.PI * (radius * radius);
            Console.WriteLine(area);
            if (area < 20) 
                Console.WriteLine($"Circle with radius of 2.50, has an area({area}) smaller than 20");
        }
    }
}
