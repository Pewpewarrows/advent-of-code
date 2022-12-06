# copy this file to dayXX.ex
defmodule AdventOfCode.DayXX do
  @moduledoc "Advent of Code 2022, Day X"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    struct = input_data |> struct_from_input_data

    struct
    |> part_one_solution
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    struct
    |> part_two_solution
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def struct_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.reduce([], fn line, acc ->
      acc
    end)
  end

  # TODO: docs, spec
  def part_one_solution(struct) do
    42
  end

  # TODO: docs, spec
  def part_two_solution(struct) do
    42
  end
end
