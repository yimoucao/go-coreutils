package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
GNU uses Pollard-Brent rho algorithm

> The Pollard-Brent rho algorithm used by factor is particularly effective for numbers with relatively small factors.
> https://en.wikipedia.org/wiki/Pollard%27s_rho_algorithm
The current implementation is using trial algorithm for small numbers. wheel algorithm for big numbers
*/

var (
	showVersion bool
	showHelp    bool
	version     string = "dev"
)

func init() {
	flag.BoolVar(&showHelp, "help", false, "display this help and exit")
	flag.BoolVar(&showVersion, "version", false, "output version information and exit")
}

func printHelp() {
	fmt.Println(`Usage: factor [NUMBER]...
or:  factor OPTION
Print the prime factors of each specified integer NUMBER.  If none are specified on the command line, read them from standard input.`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		return
	}
	if showHelp {
		printHelp()
		return
	}

	args := flag.Args()
	if len(args) > 0 {
		doOneline(args)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		doOneline(fields)
	}
}

// TODO: large integer out of range
func doOneline(fields []string) {
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			// fmt.Println(err)
			if err, ok := err.(*strconv.NumError); !ok || err.Err != strconv.ErrRange {
				fmt.Printf("factor: '%s' is not a valid postive integer\n", field)
				continue
			}
			if err = doBigNumWheel(field); err != nil {
				fmt.Printf("factor: '%s' is not a valid postive integer\n", field)
			}
			continue
		}
		if num < 0 {
			// fmt.Println("< 0")
			fmt.Printf("factor: '%s' is not a valid postive integer\n", field)
			continue
		}
		facs := factor(num)
		fmt.Printf("%d: %s\n", num, strings.Join(iToASlice(facs), " "))
	}
}

func iToASlice(i []int) []string {
	var res []string
	for _, x := range i {
		res = append(res, strconv.Itoa(x))
	}
	return res
}

// trial division
func factor(n int) []int {
	var res []int
	p := 2
	for n > 1 {
		if n%p == 0 {
			res = append(res, p)
			n /= p
		} else {
			p++
		}
	}
	return res
}

// https://en.wikipedia.org/wiki/Wheel_factorization
// https://en.wikipedia.org/wiki/Integer_factorization
//12452345134513453451345346451345
func doBigNumWheel(field string) error {
	n := new(big.Int)
	if _, ok := n.SetString(field, 10); !ok {
		return errors.New("can't parse big num")
	}
	var facs []string
	var (
		// one   = big.NewInt(1)
		zero  = big.NewInt(0)
		two   = big.NewInt(2)
		three = big.NewInt(3)
		four  = big.NewInt(4)
		five  = big.NewInt(5)
		six   = big.NewInt(6)
		wheel = []*big.Int{four, two, four, two, four, six, two, six}
	)

	mod := new(big.Int)
	for mod.Mod(n, two).Cmp(zero) == 0 {
		facs = append(facs, two.String())
		n.Div(n, two)
	}
	for mod.Mod(n, three).Cmp(zero) == 0 {
		facs = append(facs, three.String())
		n.Div(n, three)
	}
	for mod.Mod(n, five).Cmp(zero) == 0 {
		facs = append(facs, five.String())
		n.Div(n, five)
	}
	k := big.NewInt(7)
	i := 0
	kSqaure := new(big.Int)
	for kSqaure.Mul(k, k).Cmp(n) <= 0 {
		// fmt.Println(n, k)
		if mod.Mod(n, k).Cmp(zero) == 0 {
			facs = append(facs, k.String())
			n.Div(n, k)
		} else {
			k.Add(k, wheel[i])
			i = (i + 1) % len(wheel)
		}
	}
	fmt.Printf("%s: %s\n", field, strings.Join(facs, " "))
	return nil
}
