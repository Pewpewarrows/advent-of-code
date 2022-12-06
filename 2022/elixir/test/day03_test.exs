defmodule Day03Test do
  use ExUnit.Case
  doctest AdventOfCode.Day03

  test "simple" do
    input_data = """
    vJrwpWtwJgWrhcsFMMfFFhFp
    jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
    PmmdzqPrVvPwwTWBwg
    wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
    ttgJtRGJQctTZtZT
    CrZsJsPPZsGzwwsLwLmpwMDw
    """
    rucksacks = [
      "vJrwpWtwJgWrhcsFMMfFFhFp",
      "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
      "PmmdzqPrVvPwwTWBwg",
      "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
      "ttgJtRGJQctTZtZT",
      "CrZsJsPPZsGzwwsLwLmpwMDw",
    ]

    assert AdventOfCode.Day03.rucksacks_from_input_data(input_data) == rucksacks
    assert AdventOfCode.Day03.sum_common_item_type_priorities(rucksacks) == 157
    assert AdventOfCode.Day03.sum_group_badge_priorities(rucksacks) == 70
  end
end
