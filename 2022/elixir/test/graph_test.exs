defmodule GraphTest do
  use ExUnit.Case
  doctest Graph

  test "rosettacode.org" do
    arrows = [
      %Graph.Arrow{tail: "a", head: "b", weight: 7},
      %Graph.Arrow{tail: "a", head: "c", weight: 9},
      %Graph.Arrow{tail: "a", head: "f", weight: 14},
      %Graph.Arrow{tail: "b", head: "c", weight: 10},
      %Graph.Arrow{tail: "b", head: "d", weight: 15},
      %Graph.Arrow{tail: "c", head: "d", weight: 11},
      %Graph.Arrow{tail: "c", head: "f", weight: 2},
      %Graph.Arrow{tail: "d", head: "e", weight: 6},
      %Graph.Arrow{tail: "e", head: "f", weight: 9},
      # TODO: after generalizing algo for undirected edge list, remove this
      #       "dupe"
      %Graph.Arrow{tail: "f", head: "e", weight: 9}
    ]

    assert Graph.dijkstra(arrows, "a", "e") == {20, ["a", "c", "f", "e"]}
  end
end
