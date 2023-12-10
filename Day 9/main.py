def main():
    print(solve_oasis("input.txt"))

def solve_oasis(filename):
    with open(filename) as specs:
        next_line = specs.readline()
        left_total, right_total = 0, 0
        while next_line:
            solution = solve_line(next_line)
            left_total += solution[0]
            right_total += solution[1]
            next_line = specs.readline()
    return left_total, right_total

def solve_line(line: str) -> (int, int):
    nums = [int(num) for num in line.split()]
    return solve_nums(nums)

def solve_nums(nums: list[int]) -> (int, int):
    if len(set(nums)) == 1:
        return nums[0], nums[0]
    left_delta, right_delta = solve_nums([num - nums[i-1] for i, num in enumerate(nums[1:], 1)])
    return nums[0] - left_delta, nums[-1] + right_delta


if __name__ == "__main__":
    main()