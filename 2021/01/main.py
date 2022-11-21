from file_reader import read_int
import solver

def main():
    asdaf = read_int("input")
    (solution1, solution2) = solver.solve(asdaf)
    print(f'Day 01 - Part 1 - {solution1}')
    print(f'Day 01 - Part 2 - {solution2}')

if __name__ == "__main__":
    main()
