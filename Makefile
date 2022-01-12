# all:
# go test ./...

# specific:
# go test github.com/Pewpewarrows/advent-of-code/2021/golang/day01

# ad-hoc run for a day:
# go run 2021/golang/day01/main.go input.txt

# new day:
# cp -R 2021/golang/template 2021/golang/dayXX/

# debugging a day:
# dlv test github.com/Pewpewarrows/advent-of-code/2021/golang/day01
# dlv test github.com/Pewpewarrows/advent-of-code/2021/golang/day01 -- test.v -test.run <name>
# dlv debug 2021/golang/day01/main.go -- input.txt
