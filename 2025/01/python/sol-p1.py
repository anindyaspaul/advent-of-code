import os
import sys

def debug(*kargs):
    if os.getenv("DEBUG") == "1":
        print(*kargs)


def main():
    assert len(sys.argv) == 2, "Usage: python3 sol-p1.py filename"
    
    cur_position = 50
    count_zero = 0

    debug("Input file:", sys.argv[1])

    with open(sys.argv[1], "r") as input_file:
        for instruction in input_file:
            instruction = instruction.strip()
            debug(instruction)

            direction = instruction[0]
            distance = int(instruction[1:])
            debug(direction, distance)

            distance = distance % 100
            if direction == 'L':
                distance = -distance
            debug(cur_position, distance)
            
            new_position = (cur_position + distance + 100) % 100

            if new_position == 0:
                count_zero += 1

            cur_position = new_position

    print("Password: ", count_zero)


if __name__ == "__main__":
    main()
