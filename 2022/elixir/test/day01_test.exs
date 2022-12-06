defmodule Day01Test do
  use ExUnit.Case
  doctest AdventOfCode.Day01

  test "simple" do
    input_data = """
    1000
    2000
    3000

    4000

    5000
    6000

    7000
    8000
    9000

    10000
    """
    elf_snacks = [
      [1000, 2000, 3000],
      [4000],
      [5000, 6000],
      [7000, 8000, 9000],
      [10000],
    ]

    assert AdventOfCode.Day01.elf_snacks_from_input_data(input_data) == elf_snacks
    assert AdventOfCode.Day01.most_calories_held(elf_snacks) == 24000
    assert AdventOfCode.Day01.most_calories_held(elf_snacks, 3) == 45000
  end
end
