defmodule AdventOfCode.Day13 do
  @moduledoc "Advent of Code 2022, Day 13"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    packets = input_data |> packets_from_input_data

    packets
    |> sum_index_pairs_right_order
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    packets
    |> distress_signal_decoder_key
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def packets_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.flat_map(fn line ->
      if line == "" do
        []
      else
        [packet_from_line(line)]
      end
    end)
  end

  # NOTE: in future problems use json parser if applicable
  def packet_from_line([], _), do: []
  def packet_from_line(line, false = _has_inner_list?) do
    line
    |> to_string
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
  end

  def packet_from_line(line, true = _has_inner_list?) do
    {trail, _, _, packet} =
      line
      |> Enum.reduce({'', false, 0, []}, fn char, {text, in_list?, depth, acc} ->
        case in_list? do
          true ->
            case char do
              ?[ -> {[char | text], true, depth + 1, acc}
              ?] ->
                if depth > 1 do
                  {[char | text], true, depth - 1, acc}
                else
                  chars = Enum.reverse(text)
                  # TODO: use packet_from_line/1
                  {packet_from_line(chars, (?[ in chars) and (?] in chars)), false, 0, acc}
                end
              c -> {[c | text], true, depth, acc}
            end
          false ->
            case char do
              ?[ -> {'', true, 1, acc}
              ?] -> acc  # TODO: warning
              ?, ->
                next =
                  try do
                    text
                    |> Enum.reverse
                    |> int_from_charlist
                  rescue
                    ArgumentError -> text
                  end

                {'', false, 0, [next | acc]}
              n -> {[n | text], false, 0, acc}
            end
        end
      end)

    next =
      try do
        trail
        |> Enum.reverse
        |> int_from_charlist
      rescue
        ArgumentError -> trail
      end

    Enum.reverse([next | packet])
  end

  def packet_from_line(line) do
    chars =
      line
      |> String.to_charlist
      |> Enum.slice(1..-2)

    packet_from_line(chars, (?[ in chars) and (?] in chars))
  end

  def int_from_charlist(chars) do
    chars |> to_string |> String.to_integer
  end

  # TODO: docs, spec
  def sum_index_pairs_right_order(packets) do
    packets
    |> Enum.chunk_every(2)
    |> Enum.with_index
    |> Enum.flat_map(fn {[left, right], index} ->
      left
      |> Padding.pad_zip(right)
      |> Enum.reduce_while(0, fn {a, b}, _ ->
        case comparing_packets(a, b) do
          -1 -> {:halt, -1}
          0 -> {:cont, 0}
          1 -> {:halt, 1}
        end
      end)
      |> case do
        x when x in [-1, 0] -> [index + 1]
        1 -> []
      end
    end)
    |> Enum.sum
  end

  def comparing_packets(a, _) when a == nil, do: -1
  def comparing_packets(_, b) when b == nil, do: 1

  def comparing_packets(a, b) when is_integer(a) and is_integer(b) do
    cond do
      a < b -> -1
      a > b -> 1
      true -> 0  # a == b
    end
  end

  def comparing_packets(a, b) when is_list(a) and is_integer(b) do
    comparing_packets(a, [b])
  end

  def comparing_packets(a, b) when is_integer(a) and is_list(b) do
    comparing_packets([a], b)
  end

  def comparing_packets(a, b) when is_list(a) and is_list(b) do
    a
    |> Padding.pad_zip(b)
    |> Enum.reduce_while(0, fn {a, b}, _ ->
      case comparing_packets(a, b) do
        -1 -> {:halt, -1}
        0 -> {:cont, 0}
        1 -> {:halt, 1}
      end
    end)
  end

  # TODO: docs, spec
  def distress_signal_decoder_key(packets) do
    # TODO: no need to sort, can sum all packets <= each of the divider packets
    #       to find their indices
    sorted_packets = Enum.sort(packets ++ [[[2]], [[6]]], fn a, b ->
      case comparing_packets(a, b) do
        x when x in [-1, 0] -> true
        1 -> false
      end
    end)

    # TODO: handle nil for not found
    (Enum.find_index(sorted_packets, &(&1 == [[2]])) + 1) * (Enum.find_index(sorted_packets, &(&1 == [[6]])) + 1)
  end
end
