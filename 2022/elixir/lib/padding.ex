# see: https://groups.google.com/g/elixir-lang-talk/c/pil9caRXQnM
defmodule Padding do
  def pad_zip(enum_a, enum_b, padding \\ Stream.cycle([nil])) do
    count_a = Enum.count(enum_a)
    count_b = Enum.count(enum_b)

    {enum_a, enum_b} =
      cond do
        count_a < count_b -> {pad(enum_a, count_b, padding, count_a), enum_b}
        count_b < count_a -> {enum_a, pad(enum_b, count_a, padding, count_b)}
        true -> {enum_a, enum_b}
      end

    Enum.zip(enum_a, enum_b)
  end

  def pad(enum, target_count, padding, existing_count) do
    padding
    |> Enum.take(target_count - existing_count)
    |> then(&(enum ++ &1))
  end
end
