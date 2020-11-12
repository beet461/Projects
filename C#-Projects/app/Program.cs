using System;

namespace app
{
    class Program
    {
        static void Main(string[] args)
        {
            char[] charsToRemove = {'/', '*', '&', '^', '<', '-', '>', ';', '@', '(', ')'};
            string one = $"//*&^&^&^<- Characters to remove : 18 ->;@()(";
            Console.WriteLine("Untrimmed: " + one);
            one = $"//*&^&^&^<- Characters to remove : 0 ->;@()(";
            string two = one.Trim(charsToRemove);
            Console.WriteLine($"Trimmed: {two}");
        }
    }
}
