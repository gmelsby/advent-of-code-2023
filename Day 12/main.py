def main():
    print(solve_springs("input.txt"))

def solve_springs(filename: str) -> int:
    with open(filename) as specs:
        next_line = specs.readline()
        result = 0
        while next_line:
            result += solve_line(next_line)
    return result

def solve_line(line: str) -> (int, int):
    springs, counts = line.split(' ')
    counts = [int(c) for c in counts.split(',')]
    return solve_step(springs, counts)

def solve_step(springs: str, counts: list[int]):
    print(springs)
    print(counts)
    # case where all counts have been satisfied and no more springs are left
    if not len(counts) and '#' not in springs: 
        return 1
    # case where counts are unfulfilled
    if len(counts) and len(springs) < counts[0]:
        return 0

    # check if start of springs can all be springs
    if '.' not in springs[:counts[0]]:
        if len(springs) == counts[0]:
            return 1
        return solve_step(springs[counts[0]+1:], counts[1:])

    # else continue
    return solve_step(springs[1:], counts)

if __name__ == "__main__":
    main()