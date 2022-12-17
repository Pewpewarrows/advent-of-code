# (personal note: SSO auth via GitHub, not Google)
# Helpful links:
# https://rosettacode.org/
# https://programming-idioms.org/

# Degrees of solutions:
# 1. any (shortest development time)
# 2. shortest runtime
# 3. smallest memory usage
# 4. idiomatic/clean
# 5. clever alternatives

# all:
# go test ./...
# cd 2022/elixir/ && mix test --stale
# cd 2022/elixir/ && mix test --cover

# specific:
# go test github.com/Pewpewarrows/advent-of-code/2021/golang/day01
# cd 2022/elixir/ && mix test --stale test/day01_test.exs

# ad-hoc run for a day:
# go run 2021/golang/day01/main.go input.txt
# cd 2022/elixir/ && mix aoc --day 1

# new day:
# cp -R 2021/golang/template 2021/golang/dayXX/
# cp 2022/elixir/lib/advent_of_code/template.ex 2022/elixir/lib/advent_of_code/dayXX.ex
# cp 2022/elixir/test/template.exs 2022/elixir/test/dayXX_test.exs

# debugging a day:
# go test github.com/Pewpewarrows/advent-of-code/2021/golang/day01 -v
# dlv test github.com/Pewpewarrows/advent-of-code/2021/golang/day01
# dlv test github.com/Pewpewarrows/advent-of-code/2021/golang/day01 -- test.v -test.run <name>
# dlv debug 2021/golang/day01/main.go -- input.txt
# TODO: elixir
