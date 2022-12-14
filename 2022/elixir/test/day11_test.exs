defmodule Day11Test do
  use ExUnit.Case
  doctest AdventOfCode.Day11

  test "simple" do
    input_data = """
    Monkey 0:
      Starting items: 79, 98
      Operation: new = old * 19
      Test: divisible by 23
        If true: throw to monkey 2
        If false: throw to monkey 3

    Monkey 1:
      Starting items: 54, 65, 75, 74
      Operation: new = old + 6
      Test: divisible by 19
        If true: throw to monkey 2
        If false: throw to monkey 0

    Monkey 2:
      Starting items: 79, 60, 97
      Operation: new = old * old
      Test: divisible by 13
        If true: throw to monkey 1
        If false: throw to monkey 3

    Monkey 3:
      Starting items: 74
      Operation: new = old + 3
      Test: divisible by 17
        If true: throw to monkey 0
        If false: throw to monkey 1
    """
    monkeys = %{
      0 => %AdventOfCode.Day11.Monkey{id: 0, items: [79, 98], operation: {:multiply, 19}, test: {23, 2, 3}},
      1 => %AdventOfCode.Day11.Monkey{id: 1, items: [54, 65, 75, 74], operation: {:add, 6}, test: {19, 2, 0}},
      2 => %AdventOfCode.Day11.Monkey{id: 2, items: [79, 60, 97], operation: {:multiply, :old}, test: {13, 1, 3}},
      3 => %AdventOfCode.Day11.Monkey{id: 3, items: [74], operation: {:add, 3}, test: {17, 0, 1}}
    }

    assert AdventOfCode.Day11.monkeys_from_input_data(input_data) == monkeys
    assert AdventOfCode.Day11.monkey_business(monkeys, _round_count = 20) == 10_605
    assert AdventOfCode.Day11.monkey_business(monkeys, _round_count = 10_000, _worry_divisor = 1) == 2_713_310_158
  end
end
