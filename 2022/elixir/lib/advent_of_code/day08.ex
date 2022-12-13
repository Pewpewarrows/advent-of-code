defmodule AdventOfCode.Day08 do
  @moduledoc "Advent of Code 2022, Day 8"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    forest_grid = input_data |> forest_grid_from_input_data

    forest_grid
    |> visible_tree_count
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    forest_grid
    |> highest_scenic_score
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def forest_grid_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.with_index
    |> Enum.reduce(%{}, fn {line, index}, grid ->
      row = line
            |> String.to_charlist
            |> Enum.with_index
            |> Enum.reduce(%{}, fn {elem, index}, row ->
              Map.merge(row, %{index => ([elem] |> to_string |> String.to_integer)})
            end)

      Map.merge(grid, %{index => row})
    end)
  end

  # TODO: docs, spec
  def visible_tree_count(forest_grid) do
    forest_grid
    |> Enum.flat_map(fn {row_index, row} ->
      row
      |> Enum.map(fn {column_index, _} ->
        coord_visible?(forest_grid, column_index, row_index)
      end)
    end)
    |> Enum.count(&Function.identity/1)
  end

  def coord_visible?(forest_grid, x, y) do
    grid_width = forest_grid |> Map.get(0) |> map_size
    grid_height = map_size(forest_grid)
    coord_tree_height = forest_grid[y][x]

    cond do
      x == 0 -> true
      x == (grid_width - 1) -> true
      y == 0 -> true
      y == (grid_height - 1) -> true
      true ->
        # interior coordinate

        directional_views(forest_grid, x, y)
        |> Tuple.to_list
        |> Enum.map(&Tuple.to_list/1)
        |> Enum.concat
        |> Enum.map(fn direction_trees ->
          direction_trees |> Enum.all?(fn tree_height ->
            tree_height < coord_tree_height
          end)
        end)
        |> Enum.any?
    end
  end

  def directional_views(forest_grid, x, y) do
    left_right =
      forest_grid[y]
      |> Map.to_list
      |> List.keysort(0)
      |> Keyword.values
      |> Enum.split(x)

    up_down =
      forest_grid
      |> Enum.map(fn {y, row} ->
        {y, row[x]}
      end)
      |> List.keysort(0)
      |> Keyword.values
      |> Enum.split(y)

    # remove the coord's elem from the above Enum.split's
    left_right = update_in(left_right, [Access.elem(1)], &(tl(&1)))
    up_down = update_in(up_down, [Access.elem(1)], &(tl(&1)))

    {left_right, up_down}
  end

  # TODO: docs, spec
  def highest_scenic_score(forest_grid) do
    forest_grid
    |> Enum.flat_map(fn {y, row} ->
      row
      |> Enum.map(fn {x, _} ->
        scenic_score(forest_grid, x, y)
      end)
    end)
    |> Enum.max
  end

  def scenic_score(forest_grid, x, y) do
    grid_width = forest_grid |> Map.get(0) |> map_size
    grid_height = map_size(forest_grid)
    coord_tree_height = forest_grid[y][x]

    cond do
      x == 0 -> 0
      x == (grid_width - 1) -> 0
      y == 0 -> 0
      y == (grid_height - 1) -> 0
      true ->
        # interior coordinate

        {{left, right}, {up, down}} = directional_views(forest_grid, x, y)

        [Enum.reverse(left), right, Enum.reverse(up), down]
        |> Enum.map(fn view ->
          view
          # TODO: is Enum.reduce_while sufficient?
          |> List.foldl({:cont, 0}, fn tree, {status, score} ->
            case status do
              :cont ->
                cond do
                  tree < coord_tree_height -> {:cont, score + 1}
                  true -> {:halt, score + 1}
                end
              :halt -> {status, score}
            end
          end)
          |> elem(1)
        end)
        |> Enum.product
    end
  end
end
