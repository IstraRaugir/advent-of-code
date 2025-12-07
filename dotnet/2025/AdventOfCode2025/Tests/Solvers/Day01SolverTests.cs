using Shouldly;

namespace AdventOfCode2025.Solvers;

public class Day01SolverTests
{
    private readonly Day01Solver _sut = new(isTest: true);
    
    [Fact]
    public void SolvePart1_TestInput_ShouldSolveCorrectly()
    {
        // Arrange
        // Act
        var solution = _sut.SolvePartOne();
        
        // Assert
        solution.ShouldBe("3");
    }
    
    [Fact]
    public void SolvePart2_TestInput_ShouldSolveCorrectly()
    {
        // Arrange
        // Act
        var solution = _sut.SolvePartTwo();
        
        // Assert
        solution.ShouldBe("6");
    }
}