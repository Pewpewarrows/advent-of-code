# copy this file to dayXX_test.exs
defmodule DayXXTest do
  use ExUnit.Case
  doctest AdventOfCode.DayXX

  test "simple" do
    input_data = """
    """
    struct = [
    ]

    assert AdventOfCode.DayXX.struct_from_input_data(input_data) == struct
    assert AdventOfCode.DayXX.part_one_solution(struct) == 42
    assert AdventOfCode.DayXX.part_two_solution(struct) == 42
  end
end
