defmodule AdventOfCode.Day01 do
  # TODO: docs, spec
  def execute(input_data) do
    elf_snacks = input_data |> elf_snacks_from_input_data

    elf_snacks
    |> most_calories_held
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    elf_snacks
    |> most_calories_held(3)
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def elf_snacks_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> Enum.reduce([[]], fn line, acc ->
      case line do
        "" -> acc ++ [[]]
        int_text -> List.update_at(acc, -1, &(&1 ++ [String.to_integer(int_text)]))
      end
    end)
    |> List.delete_at(-1)
  end

  # TODO: docs, spec
  def most_calories_held(elf_snacks) do
    elf_snacks
    |> calories_per_elf
    |> Enum.max
  end

  # TODO: docs, spec
  def most_calories_held(elf_snacks, top_count) when top_count > 1 do
    elf_snacks
    |> calories_per_elf
    |> Enum.sort()
    |> Enum.slice(-top_count..-1)
    |> Enum.sum
  end

  # TODO: docs, spec
  def calories_per_elf(elf_snacks) do
    for elf <- elf_snacks do
      for snack <- elf, reduce: 0 do
        acc -> snack + acc
      end
    end
  end
end
