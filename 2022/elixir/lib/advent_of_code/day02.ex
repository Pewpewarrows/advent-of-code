defmodule AdventOfCode.Day02 do
  @moduledoc "Advent of Code 2022, Day 2"

  @shape_score %{
    rock: 1,
    paper: 2,
    scissors: 3,
  }
  @outcome_score %{
    loss: 0,
    draw: 3,
    win: 6,
  }
  @opponent_shape %{
    "A" => :rock,
    "B" => :paper,
    "C" => :scissors,
  }
  @player_shape %{
    "X" => :rock,
    "Y" => :paper,
    "Z" => :scissors,
  }
  @fixed_outcome %{
    "X" => :loss,
    "Y" => :draw,
    "Z" => :win,
  }

  # TODO: docs, spec
  def execute(input_data) do
    # TODO: before solving day 2's problems
    # - benchmark?
    # - Makefile
    # - vim tab error on insert mode?
    # - coc snippets?
    # - coc elixir ls?
    # - coc tab completion papercut? or tab+tab enough?
    # - mix format everywhere
    strategy_guide = input_data |> strategy_guide_from_input_data

    strategy_guide
    |> total_score_assuming_player_shape
    |> to_string
    |> String.replace_prefix("", "Part One: ")
    |> IO.puts

    strategy_guide
    |> total_score_given_outcome
    |> to_string
    |> String.replace_prefix("", "Part Two: ")
    |> IO.puts
  end

  # TODO: docs, spec
  def strategy_guide_from_input_data(input_data) do
    input_data
    |> String.split("\n")
    |> Enum.reduce([], fn line, acc ->
      round = line
      |> String.split
      |> List.to_tuple

      acc ++ [round]
    end)
    |> List.delete_at(-1)
  end

  def total_score_assuming_player_shape(strategy_guide) do
    strategy_guide
    |> Enum.map(fn round ->
      {@opponent_shape[elem(round, 0)], @player_shape[elem(round, 1)]}
    end)
    |> Enum.reduce(0, fn round, acc ->
      score = case round do
        {shape, shape} -> @shape_score[shape] + @outcome_score[:draw]
        {o_shape, p_shape} -> @shape_score[p_shape] + @outcome_score[game_result(o_shape, p_shape)]
      end

      acc + score
    end)
  end

  def total_score_given_outcome(strategy_guide) do
    strategy_guide
    |> Enum.map(fn round ->
      {@opponent_shape[elem(round, 0)], @fixed_outcome[elem(round, 1)]}
    end)
    |> Enum.reduce(0, fn round, acc ->
      score = case round do
        {shape, :loss} -> @shape_score[losing_shape(shape)] + @outcome_score[:loss]
        {shape, :draw} -> @shape_score[shape] + @outcome_score[:draw]
        {shape, :win} -> @shape_score[winning_shape(shape)] + @outcome_score[:win]
      end

      acc + score
    end)
  end

  def game_result(shape_a, shape_b) do
    # TODO: generalize into S > P > R > S
    case {shape_a, shape_b} do
      {:rock, :paper} -> :win
      {:rock, :scissors} -> :loss
      {:paper, :rock} -> :loss
      {:paper, :scissors} -> :win
      {:scissors, :rock} -> :win
      {:scissors, :paper} -> :loss
    end
  end

  def losing_shape(shape) do
    case shape do
      :rock -> :scissors
      :paper -> :rock
      :scissors -> :paper
    end
  end

  def winning_shape(shape) do
    case shape do
      :rock -> :paper
      :paper -> :scissors
      :scissors -> :rock
    end
  end
end
