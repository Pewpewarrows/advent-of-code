defmodule Graph do
  defmodule Arrow do
    defstruct tail: nil, head: nil, weight: 0
  end

  # TODO: generalize to take a list of undirected edges, by duping them all
  #       into arrows and calling this def?
  def dijkstra(arrows, start, finish) do
    shortest_path(arrows, [{0, [start]}], finish, MapSet.new())
  end

  def shortest_path(_, [], _, _), do: {0, []}
  def shortest_path(_, [{weight, [finish | _] = path} | _], finish, _) do
    # base case
    {weight, Enum.reverse(path)}
  end
  def shortest_path(arrows, [{weight, [vertex | _] = path} | routes], finish, visited) do
    new_routes =
      arrows
      |> Enum.flat_map(fn arrow ->
        if (arrow.tail == vertex) and not MapSet.member?(visited, arrow.head) do
          [{weight + arrow.weight, [arrow.head | path]}]
        else
          []
        end
      end)

    shortest_path(arrows, Enum.sort(new_routes ++ routes), finish, MapSet.put(visited, vertex))
  end
end
