﻿using System;

namespace bank
{
    class Program
    {
        static void Main(string[] args)
        {
            var account = new BankAccount("Rushvith", 1000);
            Console.WriteLine($"Account {account.Number} was created for {account.Owner} with {account.Balance} initial balance.");
        }
    }
}
