defmodule AdventOfCode.Day15 do
  @moduledoc "Advent of Code 2022, Day 15"

  @doc """
  Displays the solution(s) to this day's problem(s).
  """
  @spec execute(String.t) :: nil
  def execute(input_data) do
    sensor_beacons = input_data |> struct_from_input_data

    sensor_beacons
    |> position_count_without_beacon(_row_index = 2_000_000)
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    sensor_beacons
    |> distress_beacon_tuning_frequency(_min_x = 0, _min_y = 0, _max_x = 4_000_000, _max_y = 4_000_000)
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def struct_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> List.delete_at(-1)
    |> Enum.map(&String.trim(&1))
    |> Enum.flat_map(fn line ->
      case Regex.named_captures(~r/^Sensor at x=(?P<sensor_x>\-?\d+), y=(?P<sensor_y>\-?\d+): closest beacon is at x=(?P<beacon_x>\-?\d+), y=(?P<beacon_y>\-?\d+)$/u, line) do
        nil -> []  # TODO: warn
        caps -> [{{String.to_integer(caps["sensor_x"]), String.to_integer(caps["sensor_y"])}, {String.to_integer(caps["beacon_x"]), String.to_integer(caps["beacon_y"])}}]
      end
    end)
  end

  # TODO: docs, spec
  def position_count_without_beacon(sensor_beacons, row_index) do
    beacons_in_row = Enum.flat_map(sensor_beacons, fn {_, {beacon_x, beacon_y}} ->
      if beacon_y == row_index do
        [beacon_x]
      else
        []
      end
    end)

    sensor_beacons
    |> Enum.flat_map(fn {{sensor_x, sensor_y} = sensor, beacon} ->
      # TODO: rename to radius?
      spread = manhattan_distance(sensor, beacon)

      if row_index in (sensor_y - spread)..(sensor_y + spread) do
        delta_tip = abs(abs(row_index - sensor_y) - spread)
        [(sensor_x - delta_tip)..(sensor_x + delta_tip)]
      else
        []
      end
    end)
    |> Enum.concat
    |> then(&(MapSet.new(&1)))
    |> MapSet.difference(MapSet.new(beacons_in_row))
    |> Enum.count
  end

  def manhattan_distance({a_x, a_y}, {b_x, b_y}) do
    abs(a_x - b_x) + abs(a_y - b_y)
  end

  def distress_beacon_tuning_frequency(sensor_beacons, min_x, min_y, max_x, max_y) do
    square_coverages =
      sensor_beacons
      |> Enum.flat_map(fn {{sensor_x, sensor_y} = sensor, beacon} ->
        spread = manhattan_distance(sensor, beacon)

        # TODO: ick, use rectangle overlap algo
        # https://stackoverflow.com/questions/40795709/checking-whether-two-rectangles-overlap-in-python-using-two-bottom-left-corners
        if (((sensor_y < min_y) and ((sensor_y + spread) > min_y))
          or ((sensor_y > max_y) and ((sensor_y - spread) < max_y))
          or ((sensor_y >= min_y) or (sensor_y <= max_y)))
          and (((sensor_x < min_x) and ((sensor_x + spread) > min_x))
          or ((sensor_x > max_x) and ((sensor_x - spread) < max_x))
          or ((sensor_x >= min_x) or (sensor_x <= max_x))) do

          [{sensor, spread}]
        else
          []
        end
      end)

    # TODO: account for nil
    {x, y} =
      square_coverages
      |> Stream.flat_map(fn {{sensor_x, sensor_y}, spread} ->
        rectangle_perimeter_points({sensor_x - spread - 1, sensor_y}, {sensor_x, sensor_y - spread - 1}, {sensor_x + spread + 1, sensor_y})
      end)
      |> Stream.filter(&(point_in_rectangle?(&1, {min_x, max_y}, {min_x, min_y}, {max_x, min_y})))
      # TODO: MapSet to remove dupes?
      |> Enum.reduce_while(nil, fn {x, y}, _ ->
        square_coverages
        |> Enum.map(&(coord_in_square_coverage({x, y}, &1)))
        |> Enum.any?
        |> case do
          true -> {:cont, nil}
          false -> {:halt, {x, y}}
        end
      end)

    (x * 4_000_000) + y
  end

  # assumes whole integer coordinates
  def rectangle_perimeter_points({a_x, a_y} = a, {b_x, b_y} = b, {c_x, c_y} = c) do
    # NOTE: this only works for parallel axes, should generalize later
    d = {a_x |> Bitwise.bxor(b_x) |> Bitwise.bxor(c_x), a_y |> Bitwise.bxor(b_y) |> Bitwise.bxor(c_y)}

    Stream.concat([points_between_points(a, b), points_between_points(b, c), points_between_points(c, d), points_between_points(d, a), [a, b, c, d]])
  end

  # assumes whole integer coordinates
  def points_between_points({a_x, a_y}, {b_x, b_y}) do
    Enum.zip(a_x..b_x, a_y..b_y) |> Enum.slice(1..-2)
  end

  def coord_in_square_coverage(coord, {{sensor_x, sensor_y}, spread}) do
    point_in_rectangle?(coord, {sensor_x - spread, sensor_y}, {sensor_x, sensor_y - spread}, {sensor_x + spread, sensor_y})
  end

  # for more generic polygon solution, see:
  # https://stackoverflow.com/questions/2752725/finding-whether-a-point-lies-inside-a-rectangle-or-not
  def point_in_rectangle?(m, a, b, d) do
    # IO.inspect({m, a, b, d})
    am = vector(a, m)
    ab = vector(a, b)
    # ac = vector(a, d)
    # IO.inspect({dot_product(am, ab), dot_product(ab, ab), dot_product(am, ad), dot_product(ad, ad)})

    # (0 < dot_product(am, ab)) and (dot_product(am, ab) < dot_product(ab, ab))
    # or
    # (0 < dot_product(am, ad)) and (dot_product(am, ad) < dot_product(ad, ad))

    bd = vector(b, d)
    bm = vector(b, m)
    # TODO: figure out why this wasn't working, something to do with a/b/c/d
    #       ordering and spatial relationship to one another?
    # IO.inspect({dot_product(ab, am), dot_product(ab, ab), dot_product(bd, bm), dot_product(bd, bd)})
    (0 <= dot_product(ab, am)) and (dot_product(ab, am) <= dot_product(ab, ab))
    and
    (0 <= dot_product(bd, bm)) and (dot_product(bd, bm) <= dot_product(bd, bd))
  end

  def vector({a_x, a_y}, {b_x, b_y}) do
    {b_x - a_x, b_y - a_y}
  end

  # TODO: generalize to any length
  def dot_product({a_x, a_y}, {b_x, b_y}) do
    (a_x * b_x) + (a_y * b_y)
  end
end
