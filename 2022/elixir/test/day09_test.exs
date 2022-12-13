defmodule Day09Test do
  use ExUnit.Case
  doctest AdventOfCode.Day09

  test "simple" do
    input_data = """
    R 4
    U 4
    L 3
    D 1
    R 4
    D 1
    L 5
    R 2
    """
    motions = [
      {:right, 4},
      {:up, 4},
      {:left, 3},
      {:down, 1},
      {:right, 4},
      {:down, 1},
      {:left, 5},
      {:right, 2},
    ]

    assert AdventOfCode.Day09.motions_from_input_data(input_data) == motions
    assert AdventOfCode.Day09.unique_tail_position_count(motions) == 13
    assert AdventOfCode.Day09.unique_tail_position_count(motions, _total_knot_count = 10) == 1
  end

  test "larger" do
    input_data = """
    R 5
    U 8
    L 8
    D 3
    R 17
    D 10
    L 25
    U 20
    """

    assert AdventOfCode.Day09.unique_tail_position_count(AdventOfCode.Day09.motions_from_input_data(input_data), _total_knot_count = 10) == 36
  end
end
