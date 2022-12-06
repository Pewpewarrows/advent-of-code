defmodule Input do
  # :io.get_password is not available from a mix task, cribbed from:
  # https://stackoverflow.com/questions/37720961/elixir-or-erlang-prompt-for-password-with-hidden-input
  def get_password(prompt) do
    # TODO: use Task/Agent instead?
    child = spawn_link(fn -> password_loop(prompt) end)
    ref = make_ref()
    value = IO.gets(prompt)

    send child, {:done, self(), ref}
    receive do: ({:done, ^child, ^ref} -> :ok)

    value |> String.trim_trailing("\n")
  end

  defp password_loop(prompt) do
    receive do
      {:done, parent, ref} ->
        send parent, {:done, self(), ref}
        IO.write(:standard_error, "\e[2K\r")
    after
      1 ->
        IO.write(:standard_error, "\e[2K\r#{prompt}")
        password_loop(prompt)
    end
  end
end
