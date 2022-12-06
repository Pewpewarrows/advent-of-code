defmodule AdventOfCode.Day06 do
  @moduledoc "Advent of Code 2022, Day 6"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    signal = input_data |> signal_from_input_data

    signal
    |> packet_start_pos
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    signal
    |> message_start_pos
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def signal_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> Enum.at(0)
  end

  # TODO: docs, spec
  def packet_start_pos(signal) do
    distinct_char_seq_end_pos(signal)
  end

  # TODO: docs, spec
  def message_start_pos(signal) do
    distinct_char_seq_end_pos(signal, 14)
  end

  def distinct_char_seq_end_pos(signal, length \\ 4) do
    signal
    |> String.to_charlist
    |> Enum.with_index
    |> Enum.reduce_while({'', 0}, fn {char, i}, {acc, _} ->
      acc = case Enum.count(acc) do
        x when x in 0..(length - 1) -> acc ++ [char]
        x when x == length -> Enum.slice(acc, 1..(length - 1)) ++ [char]
      end

      case acc |> Enum.uniq |> Enum.count do
        x when x == length -> {:halt, {acc, i + 1}}
        _ -> {:cont, {acc, 0}}
      end
    end)
    |> elem(1)
  end
end
