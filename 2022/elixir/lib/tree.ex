# TODO: tests
defmodule Tree do
  @moduledoc "TODO"

  defstruct root: nil
  @type t :: %__MODULE__{root: Node.t()}

  @spec new :: __MODULE__.t()
  def new, do: %__MODULE__{}

  defmodule Node do
    @moduledoc "TODO, also note that children are unordered"

    defstruct label: "", data: %{}, children: []

    @type t :: %__MODULE__{
            label: String.t(),
            data: %{optional(any) => any},
            children: [__MODULE__.t()],
          }

    @spec new :: __MODULE__.t()
    def new, do: %__MODULE__{}

    @spec add_child(__MODULE__.t(), __MODULE__.t()) :: __MODULE__.t()
    def add_child(%__MODULE__{} = parent, child) do
      %__MODULE__{
        parent
        | children: [child | parent.children]
      }
    end

    # TODO: override == comparison to make test work?
  end
end
