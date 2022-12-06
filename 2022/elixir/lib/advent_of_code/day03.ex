defmodule AdventOfCode.Day03 do
  @moduledoc "Advent of Code 2022, Day 3"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    rucksacks = input_data |> rucksacks_from_input_data

    rucksacks
    |> sum_common_item_type_priorities
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    rucksacks
    |> sum_group_badge_priorities
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def rucksacks_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
  end

  # TODO: docs, spec
  def sum_common_item_type_priorities(rucksacks) do
    rucksacks
    |> Enum.reduce([], fn line, acc ->
      acc ++ [String.split_at(line, div(String.length(line), 2))]
    end)
    |> Enum.reduce(0, fn rucksack, acc ->
      {first_compartment, second_compartment} = rucksack

      total_priority = common_chars(first_compartment, second_compartment)
      |> Enum.map(&item_priority/1)
      |> Enum.sum

      acc + total_priority
    end)
  end

  @spec common_chars(String.t, String.t) :: charlist
  def common_chars(text_a, text_b) do
    a_chars = text_a |> String.to_charlist |> Enum.uniq

    text_b
    |> String.to_charlist
    |> Enum.uniq
    |> Enum.reduce([], fn char, common_chars ->
      case char in a_chars do
        true -> common_chars ++ [char]
        false -> common_chars
      end
    end)
  end

  @spec item_priority(char) :: pos_integer
  def item_priority(item) do
    case item do
      item when item in ?A..?Z -> item - 65 + 27
      item when item in ?a..?z -> item - 97 + 1
    end
  end

  # TODO: docs, spec
  def sum_group_badge_priorities(rucksacks) do
    rucksacks
    |> Enum.chunk_every(3)
    |> Enum.reduce(0, fn group_rucksacks, acc ->
      total_priority = group_rucksacks
      |> hd
      |> common_chars(Enum.at(group_rucksacks, 1))
      |> to_string
      |> common_chars(Enum.at(group_rucksacks, 2))
      |> Enum.map(&item_priority/1)
      |> Enum.sum

      acc + total_priority
    end)
  end
end
