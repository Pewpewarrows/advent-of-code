defmodule Stack do
  @moduledoc "Wrapper over lists with stack-like operations"
  defstruct elements: []

  @typep t :: %__MODULE__{
    elements: [any()]
  }

  @spec new :: __MODULE__.t
  def new, do: %__MODULE__{}

  @spec new(Enum.t) :: __MODULE__.t
  def new(elements) do
    %__MODULE__{
      elements: Enum.to_list(elements),
    }
  end

  @spec push(__MODULE__.t, any) :: __MODULE__.t
  def push(%__MODULE__{} = stack, elem) do
    %__MODULE__{
      stack
      | elements: [elem | stack.elements]
    }
  end

  @spec pop(__MODULE__.t) :: {:error, String.t} | {:ok, any, __MODULE__.t}
  def pop(%__MODULE__{elements: []}), do: {:error, "Stack is empty"}
  def pop(%__MODULE__{elements: [top | rest]}) do
    {
      :ok,
      top,
      %__MODULE__{
        elements: rest
      },
    }
  end

  @spec depth(__MODULE__.t) :: non_neg_integer
  def depth(%__MODULE__{elements: elems}) do
    length(elems)
  end

  @spec empty?(__MODULE__.t) :: boolean
  def empty?(%__MODULE__{elements: [_head | _tail]}), do: false
  def empty?(%__MODULE__{elements: []}), do: true
end
