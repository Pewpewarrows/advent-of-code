package main

import (
    "testing"
)

func TestDomainProblem(t *testing.T) {
    data := []displayPattern{{
        [10]string{
            "be",
            "cfbegad",
            "cbdgef",
            "fgaecd",
            "cgeb",
            "fdcge",
            "agebfd",
            "fecdb",
            "fabcd",
            "edb",
        },
        [4]string{
            "fdgacbe",
            "cefdb",
            "cefbgd",
            "gcbe",
        },
    },{
        [10]string{
            "edbfga",
            "begcd",
            "cbg",
            "gc",
            "gcadebf",
            "fbgde",
            "acbgfd",
            "abcde",
            "gfcbed",
            "gfec",
        },
        [4]string{
            "fcgedb",
            "cgb",
            "dgebacf",
            "gc",
        },
    },{
        [10]string{
            "fgaebd",
            "cg",
            "bdaec",
            "gdafb",
            "agbcfd",
            "gdcbef",
            "bgcad",
            "gfac",
            "gcb",
            "cdgabef",
        },
        [4]string{
            "cg",
            "cg",
            "fdcagb",
            "cbg",
        },
    },{
        [10]string{
            "fbegcd",
            "cbd",
            "adcefb",
            "dageb",
            "afcb",
            "bc",
            "aefdc",
            "ecdab",
            "fgdeca",
            "fcdbega",
        },
        [4]string{
            "efabcd",
            "cedba",
            "gadfec",
            "cb",
        },
    },{
        [10]string{
            "aecbfdg",
            "fbg",
            "gf",
            "bafeg",
            "dbefa",
            "fcge",
            "gcbea",
            "fcaegb",
            "dgceab",
            "fcbdga",
        },
        [4]string{
            "gecf",
            "egdcabf",
            "bgf",
            "bfgea",
        },
    },{
        [10]string{
            "fgeab",
            "ca",
            "afcebg",
            "bdacfeg",
            "cfaedg",
            "gcfdb",
            "baec",
            "bfadeg",
            "bafgc",
            "acf",
        },
        [4]string{
            "gebdcfa",
            "ecba",
            "ca",
            "fadegcb",
        },
    },{
        [10]string{
            "dbcfg",
            "fgd",
            "bdegcaf",
            "fgec",
            "aegbdf",
            "ecdfab",
            "fbedc",
            "dacgb",
            "gdcebf",
            "gf",
        },
        [4]string{
            "cefg",
            "dcbef",
            "fcge",
            "gbcadfe",
        },
    },{
        [10]string{
            "bdfegc",
            "cbegaf",
            "gecbf",
            "dfcage",
            "bdacg",
            "ed",
            "bedf",
            "ced",
            "adcbefg",
            "gebcd",
        },
        [4]string{
            "ed",
            "bcgafe",
            "cdgba",
            "cbgef",
        },
    },{
        [10]string{
            "egadfb",
            "cdbfeg",
            "cegd",
            "fecab",
            "cgb",
            "gbdefca",
            "cg",
            "fgcdab",
            "egfdb",
            "bfceg",
        },
        [4]string{
            "gbdfcae",
            "bgc",
            "cg",
            "cgb",
        },
    },{
        [10]string{
            "gcafb",
            "gcf",
            "dcaebfg",
            "ecagb",
            "gf",
            "abcdeg",
            "gaef",
            "cafbge",
            "fdbac",
            "fegbdc",
        },
        [4]string{
            "fgae",
            "cfgab",
            "fg",
            "bagce",
        },
    }}
    actual := uniqueSegmentedDigitCount(data)

    if actual != 26 {
        t.Errorf("example: expected 26 actual %d", actual)
    }

    actual = summingDisplayOutput(data)

    if actual != 61229 {
        t.Errorf("example summing: expected 61229 actual %d", actual)
    }
}