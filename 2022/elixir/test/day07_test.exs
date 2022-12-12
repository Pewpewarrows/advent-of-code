defmodule Day07Test do
  use ExUnit.Case
  doctest AdventOfCode.Day07

  @tag :skip
  test "wip simple" do
    input_data = """
    $ cd /
    $ ls
    dir a
    14848514 b.txt
    8504156 c.dat
    dir d
    $ cd a
    $ ls
    dir e
    29116 f
    2557 g
    62596 h.lst
    $ cd e
    $ ls
    584 i
    $ cd ..
    $ cd ..
    $ cd d
    $ ls
    4060174 j
    8033020 d.log
    5626152 d.ext
    7214296 k
    """
    filesystem = %Tree{
      root: %Tree.Node{
        label: "/",
        data: %{type: :dir},
        children: [
          %Tree.Node{
            label: "a",
            data: %{type: :dir},
            children: [
              %Tree.Node{
                label: "e",
                data: %{type: :dir},
                children: [
                  %Tree.Node{
                    label: "i",
                    data: %{type: :file, size: 584},
                  },
                ],
              },
              %Tree.Node{
                label: "f",
                data: %{type: :file, size: 29_116},
              },
              %Tree.Node{
                label: "g",
                data: %{type: :file, size: 2557},
              },
              %Tree.Node{
                label: "h.lst",
                data: %{type: :file, size: 62_596},
              },
            ],
          },
          %Tree.Node{
            label: "b.txt",
            data: %{type: :file, size: 14_848_514},
          },
          %Tree.Node{
            label: "c.dat",
            data: %{type: :file, size: 8_504_156},
          },
          %Tree.Node{
            label: "d",
            data: %{type: :dir},
            children: [
              %Tree.Node{
                label: "j",
                data: %{type: :file, size: 4_060_174},
              },
              %Tree.Node{
                label: "d.log",
                data: %{type: :file, size: 8_033_020},
              },
              %Tree.Node{
                label: "d.ex",
                data: %{type: :file, size: 5_626_152},
              },
              %Tree.Node{
                label: "k",
                data: %{type: :file, size: 7_214_296},
              },
            ],
          },
        ],
      }
    }

    assert AdventOfCode.Day07.filesystem_from_input_data(input_data) == filesystem
    assert AdventOfCode.Day07.sum_dir_sizes(filesystem, _max_size = 100_000) == 95_437
    assert AdventOfCode.Day07.size_of_dir_to_delete(filesystem, _total_disk_size = 70_000_000, _required_free_size = 30_000_000) == 24_933_642
  end

  test "simple" do
    input_data = """
    $ cd /
    $ ls
    dir a
    14848514 b.txt
    8504156 c.dat
    dir d
    $ cd a
    $ ls
    dir e
    29116 f
    2557 g
    62596 h.lst
    $ cd e
    $ ls
    584 i
    $ cd ..
    $ cd ..
    $ cd d
    $ ls
    4060174 j
    8033020 d.log
    5626152 d.ext
    7214296 k
    """
    filesystem = [
      {:file, "k", 7214296, %Stack{elements: ["d", "/"]}},
      {:file, "d.ext", 5626152, %Stack{elements: ["d", "/"]}},
      {:file, "d.log", 8033020, %Stack{elements: ["d", "/"]}},
      {:file, "j", 4060174, %Stack{elements: ["d", "/"]}},
      {:file, "i", 584, %Stack{elements: ["e", "a", "/"]}},
      {:file, "h.lst", 62596, %Stack{elements: ["a", "/"]}},
      {:file, "g", 2557, %Stack{elements: ["a", "/"]}},
      {:file, "f", 29116, %Stack{elements: ["a", "/"]}},
      {:dir, "e", %Stack{elements: ["a", "/"]}},
      {:dir, "d", %Stack{elements: ["/"]}},
      {:file, "c.dat", 8504156, %Stack{elements: ["/"]}},
      {:file, "b.txt", 14848514, %Stack{elements: ["/"]}},
      {:dir, "a", %Stack{elements: ["/"]}}
    ]

    assert AdventOfCode.Day07.filesystem_from_input_data(input_data) == filesystem
    assert AdventOfCode.Day07.sum_dir_sizes(filesystem, _max_size = 100_000) == 95_437
    assert AdventOfCode.Day07.size_of_dir_to_delete(filesystem, _total_disk_size = 70_000_000, _required_free_size = 30_000_000) == 24_933_642
  end
end
