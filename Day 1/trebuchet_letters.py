numbers = {
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9
}

def main():
    print(calibration("input.txt"))

def calibration(filename):
    with open(filename) as calibration_specs:
        next_line = calibration_specs.readline()
        total = 0
        while next_line:
            total += calculate_line(next_line)
            next_line = calibration_specs.readline()
    return total

def calculate_line(line: str) -> int:
    return 10 * first_digit(line) + last_digit(line)
    
# returns the first digit in the string
def first_digit(line: str) -> int:
    return next(get_numeric_value(line, idx) for idx in range(len(line)) if get_numeric_value(line, idx) is not None)

# returns last digit in the string
def last_digit(line: str) -> int:
    return next(get_numeric_value(line, idx) for idx in reversed(range(len(line))) if get_numeric_value(line, idx) is not None)

# returns integer if substring starts with a numeral or number string
# otherwise returns None
def get_numeric_value(line: str, index: int) -> int | None:
    if line[index].isnumeric():
        return int(line[index])

    number_generator = (numbers[key] for key in numbers.keys() if line.startswith(key, index)) 
    try:
        number = next(number_generator)
    except StopIteration:
        return None
    return number

if __name__ == "__main__":
    main()