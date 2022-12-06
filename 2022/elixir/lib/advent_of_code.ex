defmodule AdventOfCode do
  @moduledoc """
  Documentation for `AdventOfCode`.
  """

  @doc """
  Hello world.

  ## Examples

      iex> AdventOfCode.hello()
      :world

  """
  def hello do
    :world
  end

  # TODO: docs, spec
  def get_input(day, year \\ 2022) when day in 1..25 and is_integer(year) do
    # TODO: :eof or {:error, reason}
    session_key = Input.get_password("adventofcode.com session key: ")
    # TODO: optionally get session key from env

    # for more serious HTTP usage, look into HTTPoison/Mint/Finch
    :ssl.start()
    :application.start(:inets)
    response = :httpc.request(:get, {'https://adventofcode.com/#{year}/day/#{day}/input', [{'Cookie', 'session=#{session_key}'}]}, [], [])

    # TODO: caching

    case response do
      # TODO: 500
      {:ok, {{_, 200, 'OK'}, _, body}} -> body |> String.Chars.to_string
      {:error, reason} -> IO.inspect(reason)
    end
  end
end
