defmodule Day15Test do
  use ExUnit.Case
  doctest AdventOfCode.Day15

  test "simple" do
    input_data = """
    Sensor at x=2, y=18: closest beacon is at x=-2, y=15
    Sensor at x=9, y=16: closest beacon is at x=10, y=16
    Sensor at x=13, y=2: closest beacon is at x=15, y=3
    Sensor at x=12, y=14: closest beacon is at x=10, y=16
    Sensor at x=10, y=20: closest beacon is at x=10, y=16
    Sensor at x=14, y=17: closest beacon is at x=10, y=16
    Sensor at x=8, y=7: closest beacon is at x=2, y=10
    Sensor at x=2, y=0: closest beacon is at x=2, y=10
    Sensor at x=0, y=11: closest beacon is at x=2, y=10
    Sensor at x=20, y=14: closest beacon is at x=25, y=17
    Sensor at x=17, y=20: closest beacon is at x=21, y=22
    Sensor at x=16, y=7: closest beacon is at x=15, y=3
    Sensor at x=14, y=3: closest beacon is at x=15, y=3
    Sensor at x=20, y=1: closest beacon is at x=15, y=3
    """
    sensor_beacons = [
      {{2, 18}, {-2, 15}},
      {{9, 16}, {10, 16}},
      {{13, 2}, {15, 3}},
      {{12, 14}, {10, 16}},
      {{10, 20}, {10, 16}},
      {{14, 17}, {10, 16}},
      {{8, 7}, {2, 10}},
      {{2, 0}, {2, 10}},
      {{0, 11}, {2, 10}},
      {{20, 14}, {25, 17}},
      {{17, 20}, {21, 22}},
      {{16, 7}, {15, 3}},
      {{14, 3}, {15, 3}},
      {{20, 1}, {15, 3}}
    ]

    assert AdventOfCode.Day15.struct_from_input_data(input_data) == sensor_beacons
    assert AdventOfCode.Day15.position_count_without_beacon(sensor_beacons, _row_index = 10) == 26
    # assert AdventOfCode.Day15.distress_beacon_tuning_frequency(sensor_beacons, _min_x = 14, _min_y = 11, _max_x = 20, _max_y = 20) == 56000011
    assert AdventOfCode.Day15.distress_beacon_tuning_frequency(sensor_beacons, _min_x = 0, _min_y = 0, _max_x = 20, _max_y = 20) == 56000011
  end
end
