import unittest
import solver
from file_reader import read_int

class TestAoc(unittest.TestCase):
    def test(self):
        input = read_int("test_input")
        print(input)
        (part1, part2) = solver.solve(input)
        self.assertEqual(part1, 7)
        self.assertEqual(part2, 5)
unittest.main()
