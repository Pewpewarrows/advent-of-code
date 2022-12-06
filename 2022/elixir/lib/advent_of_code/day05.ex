defmodule AdventOfCode.Day05 do
  @moduledoc "Advent of Code 2022, Day 5"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    {crates_by_position, instructions} = input_data |> rearrangement_from_input_data

    top_crates(crates_by_position, instructions)
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    top_crates_using_batch_moves(crates_by_position, instructions)
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  defmodule Move do
    defstruct count: 0, src: 0, dest: 0
  end

  # TODO: docs, spec
  def rearrangement_from_input_data(input_data) do
    {crate_text, instruction_text} = input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.split_while(fn line -> String.trim(line) != "" end)

    # NOTE: more performant version would probably prepend to strings instead
    #       of a full stack implementation
    crates_by_position = crate_text
    |> Enum.reverse
    |> List.delete_at(0)  # TODO: validate crate count using this line
    |> Enum.reduce(%{}, fn line, acc ->
      # TODO: this is disgusting, maybe something like:
      # https://stackoverflow.com/questions/66750240/elixir-update-a-specific-value-in-a-list-of-maps
      # or use reduce instead of final Enum.map/Map.new ?
      crates = line
      |> String.to_charlist
      |> Enum.chunk_every(4)
      |> Enum.map(&to_string(&1))
      |> Enum.with_index(fn elem, i -> {i + 1, elem} end)
      |> Enum.flat_map(fn {i, elem} ->
        case Regex.named_captures(~r/\s*\[(?P<label>[[:upper:]]+)\]/u, elem) do
          nil -> []
          caps -> [{i, caps["label"]}]
        end
      end)
      |> Enum.map(fn {i, label} ->
        crates = Map.update(acc, i, Stack.new([label]), fn stack -> Stack.push(stack, label) end)

        {i, Map.get(crates, i)}
      end)
      |> Map.new

      Map.merge(acc, crates)
    end)

    instructions = instruction_text
    |> List.delete_at(0)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce([], fn line, acc ->
      case Regex.named_captures(~r/^move (?P<count>\d+) from (?P<src>\d+) to (?P<dest>\d+)$/u, line) do
        nil -> acc  # TODO: warn
        # TODO: quick way to splat map keys into struct fields
        # https://stackoverflow.com/questions/41980358/convert-maps-to-struct
        caps -> acc ++ [%Move{count: String.to_integer(caps["count"]), src: String.to_integer(caps["src"]), dest: String.to_integer(caps["dest"])}]
      end
    end)

    {crates_by_position, instructions}
  end

  # TODO: docs, spec
  def top_crates(crates_by_position, instructions) do
    instructions
    |> Enum.reduce(crates_by_position, fn move, acc ->
      Enum.reduce(1..move.count, acc, fn _, inner_acc ->
        # TODO: error
        {:ok, label, stack} = inner_acc
        |> Map.get(move.src)
        |> Stack.pop

        inner_acc
        |> Map.put(move.src, stack)
        |> Map.update(move.dest, Stack.new([label]), fn x -> Stack.push(x, label) end)
      end)
    end)
    |> Enum.map(fn {_, stack} ->
      # TODO: error
      {:ok, label, _} = Stack.pop(stack)

      label
    end)
    |> to_string
  end

  # TODO: docs, spec
  def top_crates_using_batch_moves(crates_by_position, instructions) do
    instructions
    |> Enum.reduce(crates_by_position, fn move, acc ->
      {batch_stack, new_acc} = Enum.reduce(1..move.count, {Stack.new, acc}, fn _, {new_stack, inner_acc} ->
        # TODO: error
        {:ok, label, stack} = inner_acc
        |> Map.get(move.src)
        |> Stack.pop

        {Stack.push(new_stack, label), Map.put(inner_acc, move.src, stack)}
      end)

      {_, new_acc} = Enum.reduce(1..Stack.depth(batch_stack), {batch_stack, new_acc}, fn _, {stack, inner_acc} ->
        # TODO: error
        {:ok, label, new_stack} = Stack.pop(stack)
        inner_acc = Map.update(inner_acc, move.dest, Stack.new([label]), fn x -> Stack.push(x, label) end)

        {new_stack, inner_acc}
      end)

      new_acc
    end)
    |> Enum.map(fn {_, stack} ->
      # TODO: error
      {:ok, label, _} = Stack.pop(stack)

      label
    end)
    |> to_string
  end
end
