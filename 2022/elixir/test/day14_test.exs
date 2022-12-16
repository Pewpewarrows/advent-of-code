defmodule Day14Test do
  use ExUnit.Case
  doctest AdventOfCode.Day14

  test "simple" do
    input_data = """
    498,4 -> 498,6 -> 496,6
    503,4 -> 502,4 -> 502,9 -> 494,9
    """
    rock_grid = %{
      4 => %{498 => :rock, 502 => :rock, 503 => :rock},
      5 => %{498 => :rock, 502 => :rock},
      6 => %{496 => :rock, 497 => :rock, 498 => :rock, 502 => :rock},
      7 => %{502 => :rock},
      8 => %{502 => :rock},
      9 => %{494 => :rock, 495 => :rock, 496 => :rock, 497 => :rock, 498 => :rock, 499 => :rock, 500 => :rock, 501 => :rock, 502 => :rock},
    }

    assert AdventOfCode.Day14.rock_grid_from_input_data(input_data) == rock_grid
    assert AdventOfCode.Day14.resting_sand_count_before_abyss(rock_grid) == 24
    assert AdventOfCode.Day14.resting_sand_count_with_floor(rock_grid) == 93
  end
end
