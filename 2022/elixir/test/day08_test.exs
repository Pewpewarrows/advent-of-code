defmodule Day08Test do
  use ExUnit.Case
  doctest AdventOfCode.Day08

  test "simple" do
    input_data = """
    30373
    25512
    65332
    33549
    35390
    """
    forest_grid = %{
      0 => %{0 => 3, 1 => 0, 2 => 3, 3 => 7, 4 => 3},
      1 => %{0 => 2, 1 => 5, 2 => 5, 3 => 1, 4 => 2},
      2 => %{0 => 6, 1 => 5, 2 => 3, 3 => 3, 4 => 2},
      3 => %{0 => 3, 1 => 3, 2 => 5, 3 => 4, 4 => 9},
      4 => %{0 => 3, 1 => 5, 2 => 3, 3 => 9, 4 => 0},
    }

    assert AdventOfCode.Day08.forest_grid_from_input_data(input_data) == forest_grid
    assert AdventOfCode.Day08.visible_tree_count(forest_grid) == 21
    assert AdventOfCode.Day08.highest_scenic_score(forest_grid) == 8
  end

  test "edge cases" do
    input_data = """
    000
    050
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    950
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    050
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    059
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    050
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    950
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    950
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    959
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    050
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    059
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    059
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    959
    000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    000
    959
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    950
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    059
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    090
    959
    090
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 8

    input_data = """
    060
    656
    060
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 8

    input_data = """
    050
    555
    050
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 8

    input_data = """
    040
    454
    040
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 9

    input_data = """
    00000
    01230
    04560
    07890
    00000
    """
    assert AdventOfCode.Day08.visible_tree_count(AdventOfCode.Day08.forest_grid_from_input_data(input_data)) == 25
  end
end
