defmodule Day02Test do
  use ExUnit.Case
  doctest AdventOfCode.Day02

  test "simple" do
    input_data = """
    A Y
    B X
    C Z
    """
    strategy_guide = [
      {"A", "Y"},
      {"B", "X"},
      {"C", "Z"},
    ]

    assert AdventOfCode.Day02.strategy_guide_from_input_data(input_data) == strategy_guide
    assert AdventOfCode.Day02.total_score_assuming_player_shape(strategy_guide) == 15
    assert AdventOfCode.Day02.total_score_given_outcome(strategy_guide) == 12
  end
end
