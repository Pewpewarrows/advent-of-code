defmodule Day06Test do
  use ExUnit.Case
  doctest AdventOfCode.Day06

  test "parsing" do
    input_data = """
    mjqjpqmgbljsphdztnvjfqwrcgsmlb
    """
    signal = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

    assert AdventOfCode.Day06.signal_from_input_data(input_data) == signal
  end

  test "packet_start_pos" do
    assert AdventOfCode.Day06.packet_start_pos("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 7
    assert AdventOfCode.Day06.packet_start_pos("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5
    assert AdventOfCode.Day06.packet_start_pos("nppdvjthqldpwncqszvftbrmjlhg") == 6
    assert AdventOfCode.Day06.packet_start_pos("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10
    assert AdventOfCode.Day06.packet_start_pos("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11
  end

  test "message_start_pos" do
    assert AdventOfCode.Day06.message_start_pos("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 19
    assert AdventOfCode.Day06.message_start_pos("bvwbjplbgvbhsrlpgdmjqwftvncz") == 23
    assert AdventOfCode.Day06.message_start_pos("nppdvjthqldpwncqszvftbrmjlhg") == 23
    assert AdventOfCode.Day06.message_start_pos("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29
    assert AdventOfCode.Day06.message_start_pos("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 26
  end
end
