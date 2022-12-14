defmodule AdventOfCode.Day11 do
  @moduledoc "Advent of Code 2022, Day 11"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    monkeys = input_data |> monkeys_from_input_data

    monkeys
    |> monkey_business
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    monkeys
    |> monkey_business(_round_count = 10_000, _worry_divisor = 1)
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  defmodule Monkey do
    defstruct id: 0, items: [], operation: {:multiply, 1}, test: {0, 0, 0}, inspection_count: 0

    @typep t :: %__MODULE__{
      id: non_neg_integer(),
      # TODO: type hint for IO.inspect, tell it to not print as charlist?
      items: [integer()],
      operation: {:add | :subtract | :multiply, :old | integer()},
      # {:divisible_by, :true_monkey_id, :false_monkey_id}
      test: {integer(), non_neg_integer(), non_neg_integer()},
      inspection_count: non_neg_integer()
    }
  end

  # TODO: docs, spec
  def monkeys_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce({0, %{}}, fn line, {current_id, monkeys} ->
      case line do
        "Monkey " <> id -> {id |> String.slice(0..-2) |> String.to_integer, monkeys}
        _ ->
          Map.get_and_update(monkeys, current_id, fn m ->
            new_m =
              if m == nil do
                %Monkey{id: current_id}
              else
                m
              end

            new_m =
              case line do
                "Starting items: " <> items_text ->
                  %{new_m | items: items_text |> String.split(", ") |> Enum.map(&String.to_integer/1)}
                "Operation: new = old " <> operation ->
                  [type_text, operand_text] = String.split(operation)

                  type =
                    case type_text do
                      "+" -> :add
                      "-" -> :subtract
                      "*" -> :multiply
                    end

                  operand =
                    case operand_text do
                      "old" -> :old
                      x -> String.to_integer(x)
                    end

                  %{new_m | operation: {type, operand}}
                "Test: divisible by " <> divisor ->
                  %{new_m | test: put_elem(new_m.test, 0, String.to_integer(divisor))}
                "If true: throw to monkey " <> true_monkey_id ->
                  %{new_m | test: put_elem(new_m.test, 1, String.to_integer(true_monkey_id))}
                "If false: throw to monkey " <> false_monkey_id ->
                  %{new_m | test: put_elem(new_m.test, 2, String.to_integer(false_monkey_id))}
                "\n" -> new_m
                _ ->
                  # TODO: warning
                  new_m
              end

            {m, new_m}
          end)
          |> elem(1)
          |> then(&([current_id | [&1]]))
          |> List.to_tuple
      end
    end)
    |> elem(1)
  end

  # TODO: docs, spec
  def monkey_business(monkeys, round_count \\ 20, worry_divisor \\ 3) do
    lcm = lcm_monkey_tests(monkeys)

    1..round_count
    |> Enum.reduce(monkeys, fn _, round_monkeys ->
      Enum.reduce(round_monkeys, round_monkeys, fn {id, _}, turn_monkeys ->
        # cannot use `_` from reduce fn, is stale, do fresh lookup:
        m = turn_monkeys[id]

        {true_monkey_items, false_monkey_items} =
          m
          |> Map.get(:items)
          |> Enum.map(fn worry ->
            case m.operation do
              {:add, :old} -> worry + worry
              {:add, x} -> worry + x
              {:subtract, :old} -> worry - worry
              {:subtract, x} -> worry - x
              {:multiply, :old} -> worry * worry
              {:multiply, x} -> worry * x
            end
            |> div(worry_divisor)
            |> Integer.mod(lcm)
          end)
          |> Enum.split_with(fn worry ->
            rem(worry, elem(m.test, 0)) == 0
          end)

        turn_monkeys
        |> Map.update!(id, fn m ->
          %{m | items: [], inspection_count: m.inspection_count + Enum.count(m.items)}
        end)
        |> Map.update!(elem(m.test, 1), fn true_monkey ->
          %{true_monkey | items: true_monkey.items ++ true_monkey_items}
        end)
        |> Map.update!(elem(m.test, 2), fn false_monkey ->
          %{false_monkey | items: false_monkey.items ++ false_monkey_items}
        end)
      end)
    end)
    |> Map.values
    |> Enum.map(&(&1.inspection_count))
    |> Enum.sort
    |> Enum.slice(-2..-1)
    |> Enum.product
  end

  def lcm_monkey_tests(monkeys) do
    monkeys
    |> Enum.map(fn {_, m} ->
      m |> Map.get(:test) |> elem(0)
    end)
    |> Enum.reduce(1, fn divisor, acc ->
      Math.lcm(divisor, acc)
    end)
  end
end
