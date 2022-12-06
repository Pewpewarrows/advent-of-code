defmodule AdventOfCode.Day04 do
  @moduledoc "Advent of Code 2022, Day 4"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    section_assignment_pairs = input_data |> assignments_from_input_data

    section_assignment_pairs
    |> redundant_assignment_count
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    section_assignment_pairs
    |> overlapping_assignment_count
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def assignments_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.reduce([], fn line, acc ->
      assignment = line
      |> String.split(",")
      |> Enum.map(&String.split(&1, "-"))
      |> Enum.map(&List.to_tuple/1)
      |> Enum.map(fn {x, y} -> {String.to_integer(x), String.to_integer(y)} end)
      |> List.to_tuple

      acc ++ [assignment]
    end)
  end

  # TODO: docs, spec
  def redundant_assignment_count(section_assignment_pairs) do
    section_assignment_pairs
    |> Enum.reduce([], fn pair, acc ->
      {{a_start, a_end}, {b_start, b_end}} = pair

      redundant? = cond do
        (a_start >= b_start) and (a_end <= b_end) -> true
        (b_start >= a_start) and (b_end <= a_end) -> true
        true -> false
      end

      acc ++ [redundant?]
    end)
    |> Extensions.Enum.count_bools
  end

  # TODO: docs, spec
  def overlapping_assignment_count(section_assignment_pairs) do
    section_assignment_pairs
    |> Enum.reduce([], fn pair, acc ->
      {{a_start, a_end}, {b_start, b_end}} = pair

      overlapping? = cond do
        (a_end >= b_start) and (a_start <= b_end) -> true
        (b_end >= a_start) and (b_start <= a_end) -> true
        true -> false
      end

      acc ++ [overlapping?]
    end)
    |> Extensions.Enum.count_bools
  end
end
