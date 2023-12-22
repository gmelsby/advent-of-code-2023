from functools import cache

def main():
    print(solve_springs("input.txt"))

def solve_springs(filename: str) -> (int, int):
    with open(filename) as lines:
        next_line = lines.readline()
        result1 = 0
        result2 = 0
        while next_line:
            result1 += solve_line(next_line)
            result2 += solve_5_times_line(next_line)
            next_line = lines.readline()
    return result1, result2

def solve_5_times_line(line: str) -> int:
    lava, counts = line.split(' ')
    lava = '?'.join([lava] * 5)
    counts = tuple([int(c) for c in counts.split(',')]) * 5
    result = solve_step(lava, counts)
    return result


def solve_line(line: str) -> int:
    lava, counts = line.split(' ')
    counts = tuple([int(c) for c in counts.split(',')])
    return solve_step(lava, counts)

@cache
def solve_step(lava: str, counts: tuple[int]) -> int:
    # case where all counts have been satisfied
    if not counts: 
        return int('#' not in lava)

    # ignore dots
    lava = lava.lstrip('.')
    # case where we are out of lava
    if not lava:
        return int(not counts)

    # check if start of lava is a spring
    if '#' == lava[0]:
        if len(lava) < counts[0] or '.' in lava[:counts[0]]:
            return 0
        elif '.' not in lava[:counts[0]]:
            # spring ends the lava
            if len(lava) == counts[0]:
                # returns 1 if 1 spring left, else 0
                return int(len(counts) == 1)
            # case where spring does not fit because it is too long on the end
            elif lava[counts[0]] == '#':
                return 0
            else:
                return solve_step(lava[counts[0]+1:], counts[1:])

    # else its a question mark, so check both cases
    return solve_step(f'#{lava[1:]}', counts) + solve_step(lava[1:], counts)

if __name__ == "__main__":
    main()