defmodule AdventOfCode.Day07 do
  @moduledoc "Advent of Code 2022, Day 7"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    filesystem = input_data |> filesystem_from_input_data

    filesystem
    |> sum_dir_sizes(max_size = 100_000)
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    filesystem
    |> size_of_dir_to_delete(total_disk_size = 70_000_000, required_free_size = 30_000_000)
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def wip_filesystem_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce([], fn line, acc ->
      regexes = [
        {~r/^\$ ls/u, :ls},
        {~r/^\$ cd (?P<dirname>.+)$/u, :cd},
        {~r/^dir (?P<dirname>.+)$/u, :dir},
        {~r/^(?P<filesize>\d+) (?P<filename>.+)$/u, :file},
      ]

      command = Enum.find_value(regexes, {:warn}, fn {reg, type} ->
        case Regex.named_captures(reg, line) do
          nil -> nil
          values -> {type, values}
        end
      end)
      |> case do
        {:ls, values} -> {:ls}
        {:cd, values} -> {:cd, values["dirname"]}
        {:dir, values} -> {:dir, values["dirname"]}
        {:file, values} -> {:file, values["filesize"], values["filename"]}
        {:warn} -> {:warn, "unexpected line: #{line}"}
      end

      # TODO: try getting this way to work again?
      # TODO: alt, try case with string pattern matching instead of regexes
      # cond do
      #   Regex.match?(~r/^\$ ls/u, line) -> {:ls}
      #   %{} = values = Regex.named_captures(~r/^\$ cd (?P<dirname>.+)$/u, line) -> {:cd, values["dirname"]}
      #   %{} = values = Regex.named_captures(~r/^dir (?P<dirname>.+)$/u, line) -> {:dir, values["dirname"]}
      #   %{} = values = Regex.named_captures(~r/^(?P<filesize>\d+) (?P<filename>.+)$/u, line) -> {:file, values["filesize"], values["filename"]}
      #   true -> {:warn, "unexpected line: #{line}"}
      # end
      # |> IO.inspect

      [command | acc]
    end)
    |> List.foldr({Stack.new(["/"]), %Tree{root: %Tree.Node{label: "/", data: %{type: :dir}}}}, fn command, {cwd, acc} ->
      case elem(command, 0) do
        # TODO: maybe have an "in-ls" state?
        :ls -> {cwd, acc}
        :cd ->
          case elem(command, 1) do
            "/" -> {Stack.new(["/"]), acc}
            ".." -> {cwd |> Stack.pop |> elem(2), acc}  # TODO: Stack.pop error
            dirname -> {Stack.push(cwd, dirname), acc}
          end
        :dir ->
          {cwd, update_in(acc, [Access.key!(:root) | access_node_for_dir_stack(cwd)], fn(node) ->
            Tree.Node.add_child(node, %Tree.Node{label: elem(command, 1), data: %{type: :dir}})
          end)}
        :file ->
          {cwd, update_in(acc, [Access.key!(:root) | access_node_for_dir_stack(cwd)], fn(node) ->
            Tree.Node.add_child(node, %Tree.Node{label: elem(command, 2), data: %{type: :file, size: command |> elem(1) |> String.to_integer}})
          end)}
        :warn ->
          Logger.warning(elem(command, 1))
          {cwd, acc}
      end
      |> IO.inspect
    end)
    |> elem(1)
  end

  def filesystem_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce({[], Stack.new(["/"])}, fn line, {entries, cwd} ->
      case line do
        "$ ls" -> {entries, cwd}
        "$ cd .." -> {entries, cwd |> Stack.pop |> elem(2)}
        "$ cd /" <> dirname -> {entries, Stack.new(["/"])}
        "$ cd " <> dirname -> {entries, Stack.push(cwd, dirname)}
        "dir " <> dirname -> {[{:dir, dirname, cwd} | entries], cwd}
        s ->
          [filesize, filename] = String.split(s)
          {[{:file, filename, String.to_integer(filesize), cwd} | entries], cwd}
      end
    end)
    |> elem(0)
  end

  def access_node_for_dir_stack(stack) do
    stack
    |> Map.get(:elements)
    |> Enum.slice(0..-2)
    |> List.foldr([], fn dirname, acc ->
      [Access.filter(&(&1.label == dirname)) | [Access.key!(:children) | acc]]
    end)
    |> Enum.reverse
  end

  # TODO: docs, spec
  def sum_dir_sizes(filesystem, max_size \\ 1000) do
    filesystem
    |> dir_sizes
    |> Map.filter(fn {_, size} -> size <= max_size end)
    |> Map.values
    |> Enum.sum
  end

  def dir_sizes(filesystem) do
    filesystem
    |> Enum.filter(fn entry -> elem(entry, 0) == :file end)
    |> Enum.reduce(%{}, fn {:file, _, filesize, cwd}, dirsizes ->
      cwd
      |> Map.get(:elements)
      |> Enum.reverse
      |> Enum.scan(fn dirname, acc ->
        acc <> dirname <> "/"
      end)
      |> Enum.reduce(%{}, fn dirpath, sizes ->
        Map.merge(sizes, %{dirpath => filesize})
      end)
      |> Map.merge(dirsizes, fn _, v1, v2 -> v1 + v2 end)
    end)
  end

  # TODO: docs, spec
  def size_of_dir_to_delete(filesystem, total_disk_size \\ 1_000_000, required_free_size \\ 100_000) do
    fs_sizes = dir_sizes(filesystem)

    fs_sizes
    |> Map.filter(fn {_, size} -> size >= (required_free_size - (total_disk_size - fs_sizes["/"])) end)
    |> Map.values
    |> Enum.sort
    |> hd
  end
end
