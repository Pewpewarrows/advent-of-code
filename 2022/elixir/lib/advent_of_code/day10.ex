defmodule AdventOfCode.Day10 do
  @moduledoc "Advent of Code 2022, Day 10"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    instructions = input_data |> struct_from_input_data

    instructions
    |> periodic_signal_strength_sum
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    instructions
    |> crt_text
    |> String.replace_prefix("", "Part Two:\n")
    |> IO.puts
  end

  # TODO: docs, spec
  def struct_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce([], fn line, acc ->
      instruction =
        case line do
          "addx " <> delta_x -> {:addx, String.to_integer(delta_x)}
          "noop" -> {:noop}
        end

      [instruction | acc]
    end)
    |> Enum.reverse
  end

  # TODO: docs, spec
  def periodic_signal_strength_sum(instructions) do
    instructions
    |> cycle_registers
    |> Enum.reduce([], fn {x, cycle_num}, signal_strengths ->
      if (cycle_num == 20) or (Integer.mod(cycle_num, 40) == 20) do
        [(cycle_num * x) | signal_strengths]
      else
        signal_strengths
      end
    end)
    |> Enum.sum
  end

  def cycle_registers(instructions) do
    instructions
    |> Enum.reduce({1, []}, fn instruction, {x, cycle_registers} ->
      {next_x, next_cycles} =
        case instruction do
          {:addx, delta_x} -> {x + delta_x, [x, x]}
          {:noop} -> {x, [x]}
        end

      {next_x, next_cycles ++ cycle_registers}
    end)
    |> elem(1)
    |> Enum.reverse
    |> Enum.with_index(1)
  end

  # TODO: docs, spec
  def crt_text(instructions) do
    instructions
    |> cycle_registers
    |> Enum.reduce([], fn {x, cycle_num}, crt_chars ->
      pixel =
        if Integer.mod(cycle_num - 1, 40) in (x - 1)..(x + 1) do
          '#'
        else
          '.'
        end

      [pixel | crt_chars]
    end)
    |> Enum.reverse
    |> Enum.chunk_every(40)
    |> Enum.map(&Kernel.to_string/1)
    |> Enum.join("\n")
    |> Kernel.<>("\n")
  end
end
