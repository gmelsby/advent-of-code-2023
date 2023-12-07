def main():
    print(calculate_gear_ratios(*schematic_to_grid_dict("input.txt")))

def calculate_gear_ratios(gear_locations, grid_dict):
    total = 0
    print(grid_dict)
    for gear in gear_locations:
        direction_list = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0 ,1), (1, -1), (1, 0), (1, 1)]
        square_list = [(gear[0] + direction[0], gear[1] + direction[1]) for direction in direction_list]
        unique_parts = list(set([grid_dict[square] for square in square_list if square in grid_dict]))
        if len(unique_parts) == 2:
            total += unique_parts[0] * unique_parts[1]

    return total

def schematic_to_grid_dict(filename):
    grid = []
    grid_dict = {}
    gear_locations = []
    with open(filename) as schematic:
        next_line = schematic.readline()
        while next_line:
          grid.append(next_line)
          next_line = schematic.readline()

    for row_num, line in enumerate(grid):
        col_num, temp_num, symbol_flag = 0, 0, False
        for col_num in range(len(line)):
            if line[col_num] == '*':
                gear_locations.append((row_num, col_num))
            # case where we are in the middle of a number
            if temp_num != 0:
                # case where number continues
                if line[col_num].isnumeric():
                    temp_num *= 10
                    temp_num += int(line[col_num])
                    if not symbol_flag:
                        if (
                                row_num > 0 and grid[row_num-1][col_num] == '*' 
                                or row_num < len(grid) - 1 and grid[row_num+1][col_num] == '*'
                            ):
                            symbol_flag = True

                # case where number has ended
                if not line[col_num].isnumeric() or col_num == len(line) - 1:
                    # if we have already detected symbol no need to check again
                    if symbol_flag:
                        for i in range(len(str(temp_num))):
                            grid_dict[(row_num, col_num-1-i)] = temp_num
                    # if no symbol yet found check last 3 spots
                    elif (
                            row_num > 0 and grid[row_num-1][col_num] == '*' 
                            or row_num < len(grid) - 1 and grid[row_num+1][col_num] == '*'
                            or line[col_num] == '*'
                        ):
                        for i in range(len(str(temp_num))):
                            grid_dict[(row_num, col_num-1-i)] = temp_num

                    # regardless what happens reset things for next number
                    temp_num = 0
                    symbol_flag = False
            else:
                if line[col_num].isnumeric():
                    temp_num = int(line[col_num])
                    if (
                            row_num > 0 and grid[row_num-1][col_num] == '*' 
                            or row_num > 0 and col_num > 0 and grid[row_num-1][col_num-1] == '*' 
                            or row_num < len(grid) - 1 and grid[row_num+1][col_num] == '*'
                            or row_num < len(grid) - 1 and col_num > 0 and grid[row_num+1][col_num-1] == '*'
                            or col_num > 0 and grid[row_num][col_num-1] == '*'
                        ):
                        symbol_flag = True

                    # case where number ends (single digit at end of line)
                    if (col_num == len(line) - 1):
                        if symbol_flag:
                            grid_dict[(row_num, col_num)] = temp_num

    return gear_locations, grid_dict

if __name__ == "__main__":
    main()