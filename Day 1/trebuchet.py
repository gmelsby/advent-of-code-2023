def main():
    print(calibration("input.txt"))

def calibration(filename):
    with open(filename) as calibration_specs:
        next_line = calibration_specs.readline()
        total = 0
        while next_line:
            total += first_digit(next_line) * 10 + last_digit(next_line)
            next_line = calibration_specs.readline()
    return total

# returns the first digit in the string
def first_digit(line: str) -> int:
    return int(next(char for char in line if char.isnumeric()))

# returns last digit in the string
def last_digit(line: str) -> int:
    return int(next(char for char in reversed(line) if char.isnumeric()))


if __name__ == "__main__":
    main()