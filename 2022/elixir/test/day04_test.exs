defmodule Day04Test do
  use ExUnit.Case
  doctest AdventOfCode.Day04

  test "simple" do
    input_data = """
    2-4,6-8
    2-3,4-5
    5-7,7-9
    2-8,3-7
    6-6,4-6
    2-6,4-8
    """
    section_assignment_pairs = [
      {{2, 4}, {6, 8}},
      {{2, 3}, {4, 5}},
      {{5, 7}, {7, 9}},
      {{2, 8}, {3, 7}},
      {{6, 6}, {4, 6}},
      {{2, 6}, {4, 8}},
    ]

    assert AdventOfCode.Day04.assignments_from_input_data(input_data) == section_assignment_pairs
    assert AdventOfCode.Day04.redundant_assignment_count(section_assignment_pairs) == 2
    assert AdventOfCode.Day04.overlapping_assignment_count(section_assignment_pairs) == 4
  end
end
