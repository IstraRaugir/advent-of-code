namespace AdventOfCode2025.Solvers;

public class Day01Solver
{
    private readonly List<string> _input;

    public Day01Solver(bool isTest = false)
    {
        var fileName = isTest ? "Day01Test" : "Day01";
        _input = File.ReadLines($"Inputs/{fileName}.txt").ToList();
    }
    
    public string SolvePartOne()
    {
        var currentPosition = 50;
        var timesEndedAtZero = 0;
        _input.ForEach(row =>
        {
            var rotationAmount = int.Parse(row[1..]);
            if (row.Contains('R'))
            {
                currentPosition += rotationAmount;
                currentPosition %= 100;
            }
            else
            {
                currentPosition -= rotationAmount;
                while (currentPosition < 0) currentPosition += 100;
            }
            if (currentPosition == 0) timesEndedAtZero++; 
        });
        
        return timesEndedAtZero.ToString();
    }

    public string SolvePartTwo()
    {
        throw new NotImplementedException();
    }
}