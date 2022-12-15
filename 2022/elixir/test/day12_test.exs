defmodule Day12Test do
  use ExUnit.Case
  doctest AdventOfCode.Day12

  test "simple" do
    input_data = """
    Sabqponm
    abcryxxl
    accszExk
    acctuvwj
    abdefghi
    """

    assert AdventOfCode.Day12.fewest_steps_to_best_signal(AdventOfCode.Day12.heightgrid_from_input_data(input_data)) == 31
    assert AdventOfCode.Day12.shortest_length_from_low_point_to_end(AdventOfCode.Day12.heightgrid_from_input_data(input_data)) == 29
  end
end
