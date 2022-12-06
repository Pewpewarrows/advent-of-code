require Logger

defmodule Mix.Tasks.Aoc do
  use Mix.Task

  @shortdoc "TODO"
  def run(args) do
    # (TODO: by default, run all, or run latest?)
    # mix aoc

    # mix aoc --day 1 --file input.txt
    {parsed, opts, invalid} = OptionParser.parse(args, aliases: [t: :day, f: :file], strict: [day: :integer, file: :string])

    if (opts != []) or (invalid != []) do
      Logger.error("invalid args, these options were unexpected: #{opts} #{inspect(invalid)}")
      exit({:shutdown, 2})
    end

    # TODO: figure out how to have a case statement where one branch assigns
    #       and the other errors
    if parsed[:day] not in 1..25 do
      Logger.error("invalid day, must be between 1 and 25 inclusive: #{parsed[:day]}")
      exit({:shutdown, 2})
    end

    input_data = case parsed[:file] do
      # TODO: error handling on get_input
      nil -> AdventOfCode.get_input(parsed[:day])
      file_path -> File.read!(file_path)
    end

    padded_day = parsed[:day] |> Integer.to_string |> String.pad_leading(2, "0")
    day_module_text = "Day#{padded_day}"
    day_module = String.to_existing_atom("Elixir.AdventOfCode.#{day_module_text}")
    day_module.execute(input_data)
  end
end
