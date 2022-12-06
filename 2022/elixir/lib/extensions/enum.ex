defmodule Extensions.Enum do
  def count_bools(enumerable), do: Enum.count(enumerable, &(&1))
end
