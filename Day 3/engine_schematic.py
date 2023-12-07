def main():
    print(solve_schematic("input.txt"))

symbols = '!@#$%^&*()_+-=/'

def solve_schematic(filename):
    grid = []
    total = 0
    with open(filename) as schematic:
        next_line = schematic.readline()
        while next_line:
          grid.append(next_line)
          next_line = schematic.readline()

    for row_num, line in enumerate(grid):
        col_num, temp_num, symbol_flag = 0, 0, False
        for col_num in range(len(line)):
            # case where we are in the middle of a number
            if temp_num != 0:
                # case where number continues
                if line[col_num].isnumeric():
                    temp_num *= 10
                    temp_num += int(line[col_num])
                    if not symbol_flag:
                        if (
                                row_num > 0 and grid[row_num-1][col_num] in symbols 
                                or row_num < len(grid) - 1 and grid[row_num+1][col_num] in symbols
                            ):
                            symbol_flag = True

                # case where number has ended
                if not line[col_num].isnumeric() or col_num == len(line) - 1:
                    # if we have already detected symbol no need to check again
                    if symbol_flag:
                        total += temp_num 
                    # if no symbol yet found check last 3 spots
                    elif (
                            row_num > 0 and grid[row_num-1][col_num] in symbols 
                            or row_num < len(grid) - 1 and grid[row_num+1][col_num] in symbols
                            or line[col_num] in symbols
                        ):
                        total += temp_num

                    # regardless what happens reset things for next number
                    temp_num = 0
                    symbol_flag = False
            else:
                if line[col_num].isnumeric():
                    temp_num = int(line[col_num])
                    if (
                            row_num > 0 and grid[row_num-1][col_num] in symbols 
                            or row_num > 0 and col_num > 0 and grid[row_num-1][col_num-1] in symbols 
                            or row_num < len(grid) - 1 and grid[row_num+1][col_num] in symbols
                            or row_num < len(grid) - 1 and col_num > 0 and grid[row_num+1][col_num-1] in symbols
                            or col_num > 0 and grid[row_num][col_num-1] in symbols
                        ):
                        symbol_flag = True

                    # case where number ends (single digit at end of line)
                    if (col_num == len(line) - 1):
                        if symbol_flag:
                            total += temp_num

    return total


if __name__ == "__main__":
    main()