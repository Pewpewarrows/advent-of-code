defmodule AdventOfCode.Day09 do
  @moduledoc "Advent of Code 2022, Day 9"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    motions = input_data |> motions_from_input_data

    motions
    |> unique_tail_position_count
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    motions
    |> unique_tail_position_count(_total_knot_count = 10)
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def motions_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.reduce([], fn line, acc ->
      [dir_letter, step_count_text] = String.split(line)

      direction = case dir_letter do
        "U" -> :up
        "D" -> :down
        "L" -> :left
        "R" -> :right
      end

      [{direction, String.to_integer(step_count_text)} | acc]
    end)
    |> Enum.reverse
  end

  # TODO: docs, spec
  def unique_tail_position_count(motions, total_knot_count \\ 2) do
    knots = Enum.reduce(1..total_knot_count, [], fn _, acc ->
      [{0, 0} | acc]
    end)

    motions
    |> Enum.reduce({knots, MapSet.new([{0, 0}])}, fn {direction, step_count}, acc ->
      Enum.reduce(1..step_count, acc, fn _, {knots, tail_visits} ->
        {head_x, head_y} = hd(knots)

        {next_head_x, next_head_y} = case direction do
          :up -> {head_x, head_y - 1}
          :down -> {head_x, head_y + 1}
          :left -> {head_x - 1, head_y}
          :right -> {head_x + 1, head_y}
        end

        next_knots =
          knots
          |> tl
          |> Enum.reduce([{next_head_x, next_head_y}], fn knot, next_knots ->
            {knot_x, knot_y} = knot
            {follow_x, follow_y} = hd(next_knots)

            new_knot =
              if abs(follow_x - knot_x) >= 2 or abs(follow_y - knot_y) >= 2 do
                {
                  knot_x + round((follow_x - knot_x) / 2),
                  knot_y + round((follow_y - knot_y) / 2),
                }
              else
                knot
              end

            [new_knot | next_knots]
          end)
          |> Enum.reverse
          |> tl

        {[{next_head_x, next_head_y} | next_knots], MapSet.put(tail_visits, List.last(next_knots))}
      end)
    end)
    |> elem(1)
    |> Enum.count
  end
end
