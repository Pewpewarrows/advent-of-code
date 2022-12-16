defmodule AdventOfCode.Day12 do
  @moduledoc "Advent of Code 2022, Day 12"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    heightgrid = input_data |> heightgrid_from_input_data

    heightgrid
    |> fewest_steps_to_best_signal
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    heightgrid
    |> shortest_length_from_low_point_to_end
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def heightgrid_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.with_index
    |> Enum.reduce({nil, nil, %{}}, fn {line, y}, {start, finish, heightgrid} ->
      row =
        line
        |> String.to_charlist
        |> Enum.with_index
        |> Enum.into(%{}, fn {v, k} -> {k, v} end)

      start = start ||
        case Enum.find(row, fn {_, v} -> v == ?S end) do
          nil -> nil
          {x, _} -> {x, y}
        end

      row =
        case start do
          {start_x, start_y} when start_y == y -> Map.put(row, start_x, ?a)
          _ -> row
        end

      finish = finish ||
        case Enum.find(row, fn {_, v} -> v == ?E end) do
          nil -> nil
          {x, _} -> {x, y}
        end

      row =
        case finish do
          {finish_x, finish_y} when finish_y == y -> Map.put(row, finish_x, ?z)
          _ -> row
        end

      {start, finish, Map.put(heightgrid, y, row)}
    end)
  end

  def arrows_from_heightgrid(heightgrid) do
    Enum.flat_map(heightgrid, fn {y, row} ->
      Enum.flat_map(row, fn {x, elevation} ->
        neighbors = []

        # TODO: def prepend_if
        neighbors = if x != 0, do: [{heightgrid[y][x - 1], {x - 1, y}} | neighbors], else: neighbors
        neighbors = if y != 0, do: [{heightgrid[y - 1][x], {x, y - 1}} | neighbors], else: neighbors
        neighbors = if x != (Enum.count(row) - 1), do: [{heightgrid[y][x + 1], {x + 1, y}} | neighbors], else: neighbors
        neighbors = if y != (Enum.count(heightgrid) - 1), do: [{heightgrid[y + 1][x], {x, y + 1}} | neighbors], else: neighbors

        Enum.flat_map(neighbors, fn {e, coord} ->
          if e <= (elevation + 1) do
            [%Graph.Arrow{tail: {x, y}, head: coord, weight: 1}]
          else
            []
          end
        end)
      end)
    end)
  end

  # TODO: docs, spec
  def fewest_steps_to_best_signal(heightgrid) do
    {start, finish, heightgrid} = heightgrid

    heightgrid
    |> arrows_from_heightgrid
    # TODO: this was too slow, need to implement PriorityQueue w/ :gb_trees
    # |> Graph.dijkstra(start, finish)
    # |> elem(1)
    |> bfs(start, finish)
    |> then(&(Enum.count(&1) - 1))
  end

  def bfs(arrows, start, finish) do
    graph = :digraph.new()

    arrows
    |> Enum.reduce(MapSet.new(), fn arrow, vertices ->
      MapSet.put(vertices, arrow.tail)
      MapSet.put(vertices, arrow.head)
    end)
    |> Enum.each(fn vertex ->
      :digraph.add_vertex(graph, vertex)
    end)

    Enum.each(arrows, fn arrow ->
      :digraph.add_edge(graph, arrow.tail, arrow.head)
    end)

    :digraph.get_short_path(graph, start, finish)
  end

  # TODO: docs, spec
  def shortest_length_from_low_point_to_end(heightgrid) do
    {{start_x, start_y}, finish, heightgrid} = heightgrid

    arrows = arrows_from_heightgrid(heightgrid)
    start_elevation = heightgrid[start_y][start_x]

    # TODO: very slow, should further filter potential starts to remove ones
    #       with no viable neighbors other than 'a's, then look into other
    #       speedups
    heightgrid
    |> Enum.flat_map(fn {y, row} ->
      Enum.flat_map(row, fn {x, elevation} ->
        if elevation == start_elevation do
          [{x, y}]
        else
          []
        end
      end)
    end)
    |> Enum.flat_map(fn potential_start ->
      arrows
      |> bfs(potential_start, finish)
      |> case do
        :false -> []
        path -> [Enum.count(path) - 1]
      end
    end)
    |> Enum.min
  end
end
