defmodule Day05Test do
  use ExUnit.Case
  doctest AdventOfCode.Day05

  test "simple" do
    input_data = """
        [D]
    [N] [C]
    [Z] [M] [P]
     1   2   3

    move 1 from 2 to 1
    move 3 from 1 to 3
    move 2 from 2 to 1
    move 1 from 1 to 2
    """
    crates_by_position = %{
      1 => Stack.new(["N", "Z"]),
      2 => Stack.new(["D", "C", "M"]),
      3 => Stack.new(["P"]),
    }
    instructions = [
      %AdventOfCode.Day05.Move{count: 1, src: 2, dest: 1},
      %AdventOfCode.Day05.Move{count: 3, src: 1, dest: 3},
      %AdventOfCode.Day05.Move{count: 2, src: 2, dest: 1},
      %AdventOfCode.Day05.Move{count: 1, src: 1, dest: 2},
    ]

    assert AdventOfCode.Day05.rearrangement_from_input_data(input_data) == {crates_by_position, instructions}
    assert AdventOfCode.Day05.top_crates(crates_by_position, instructions) == "CMZ"
    assert AdventOfCode.Day05.top_crates_using_batch_moves(crates_by_position, instructions) == "MCD"
  end
end
