import sys

def _solve1(input: list[int]) -> int:
    result = 0
    previous_value = sys.maxsize
    for current_value in input:
        if current_value > previous_value:
            result += 1
        previous_value = current_value
    return result

def _solve2(input: list[int]) -> int:
    result = 0
    first_value = None
    second_value = None
    third_value = None
    previous_sum = None
    for current_value in input:
        if first_value is not None and second_value is not None:
            previous_sum = first_value + second_value + third_value
        first_value = second_value
        second_value = third_value
        third_value = current_value
        if previous_sum is not None:
            current_sum = first_value + second_value + third_value
            if previous_sum < current_sum:
                result += 1
    return result

def solve(input: list[int]) -> tuple[int,int]:
    solution1 = _solve1(input)
    solution2 = _solve2(input)
    return (solution1,solution2)
