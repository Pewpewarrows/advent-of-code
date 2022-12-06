defmodule AdventOfCodeTest do
  use ExUnit.Case
  doctest AdventOfCode

  test "greets the world" do
    assert AdventOfCode.hello() == :world
  end

  @tag skip: "TODO: mock via mox: Input.get_password, httpc.request"
  test "given no env key when get_input then prompt user" do
    AdventOfCode.get_input(1)
  end
end
