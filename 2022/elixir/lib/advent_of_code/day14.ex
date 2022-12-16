defmodule AdventOfCode.Day14 do
  @moduledoc "Advent of Code 2022, Day 14"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    rock_grid = input_data |> rock_grid_from_input_data

    rock_grid
    |> resting_sand_count_before_abyss
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    rock_grid
    |> resting_sand_count_with_floor
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def rock_grid_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce(MapSet.new, fn line, acc ->
      line
      |> String.split(" -> ")
      |> Enum.map(fn x -> x |> String.split(",") |> Enum.map(&String.to_integer/1) |> List.to_tuple end)
      |> Enum.chunk_every(2, 1)
      |> Enum.slice(0..-2)
      |> Enum.reduce(MapSet.new, fn [{a_x, a_y}, {b_x, b_y}], range_coords ->
        if a_x == b_x do
          Enum.map(a_y..b_y, &({a_x, &1}))
        else
          Enum.map(a_x..b_x, &({&1, a_y}))
        end
        |> MapSet.new
        |> MapSet.union(range_coords)
      end)
      |> MapSet.union(acc)
    end)
    |> Enum.reduce(%{}, fn {x, y}, grid ->
      updating_2d_grid(grid, x, y, :rock)
    end)
  end

  def updating_2d_grid(grid, x, y, v) do
    Map.update(grid, y, %{x => v}, fn x_map ->
      Map.put(x_map, x, v)
    end)
  end

  # TODO: docs, spec
  def resting_sand_count_before_abyss(rock_grid, sand_source \\ {500, 0}) do
    # TODO: could keep track of resting_sand_count during sim instead of
    #       counting at the end
    rock_grid
    |> simulating_until_abyss(sand_source)
    |> counting_sand
  end

  def counting_sand(rock_grid) do
    rock_grid
    |> Enum.map(fn {_, row} ->
      row
      |> Map.values
      |> Enum.filter(&(&1 == :sand))
      |> Enum.count
    end)
    |> Enum.sum
  end

  def simulating_until_abyss(rock_grid, {source_x, source_y}) do
    {{sand_x, sand_y}, state} = next_sand_step(rock_grid, {source_x, source_y + 1})

    case state do
      :rest -> simulating_until_abyss(updating_2d_grid(rock_grid, sand_x, sand_y, :sand), {source_x, source_y})
      :abyss -> rock_grid
    end
  end

  def simulating_until_block(rock_grid, {source_x, source_y}, implied_floor) do
    {{sand_x, sand_y}, _} = next_sand_step(rock_grid, {source_x, source_y + 1}, :down, implied_floor)

    if {sand_x, sand_y} == {source_x, source_y} do
      # TODO: dedup
      updating_2d_grid(rock_grid, sand_x, sand_y, :sand)
    else
      simulating_until_block(updating_2d_grid(rock_grid, sand_x, sand_y, :sand), {source_x, source_y}, implied_floor)
    end
  end

  def grid_floor(rock_grid), do: rock_grid |> Map.keys |> Enum.max

  def next_sand_step(rock_grid, coord, state \\ :down, implied_floor \\ nil)
  def next_sand_step(_, {x, y}, state, _) when state in [:rest, :abyss] do
    {{x, y}, state}
  end

  def next_sand_step(rock_grid, {x, y}, state, implied_floor) do
    # returns coord, state

    default_item =
      if (implied_floor == nil) or (y < implied_floor) do
        :air
      else
        :rock
      end

    rock_grid
    |> Map.get(y, %{})
    |> Map.get(x, default_item)
    |> case do
      :air ->
        # TODO: nil check necessary?
        if (implied_floor == nil) and ((y + 1) > grid_floor(rock_grid)) do
          {{x, y + 1}, :abyss}
        else
          next_sand_step(rock_grid, {x, y + 1}, :down, implied_floor)
        end

      s when s in [:rock, :sand] ->
        case state do
          :down -> next_sand_step(rock_grid, {x - 1, y}, :left, implied_floor)
          :left -> next_sand_step(rock_grid, {x + 2, y}, :right, implied_floor)
          :right -> next_sand_step(rock_grid, {x - 1, y - 1}, :rest, implied_floor)
        end
    end
  end

  # TODO: docs, spec
  def resting_sand_count_with_floor(rock_grid, sand_source \\ {500, 0}) do
    rock_grid
    |> simulating_until_block(sand_source, grid_floor(rock_grid) + 2)
    |> counting_sand
  end
end
