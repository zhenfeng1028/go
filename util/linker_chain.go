package main

import "fmt"

type Handler func(ctx *[]string) error

type Linker func(ctx *[]string, handler Handler) error

func LinkerChain(linkers ...Linker) Linker {
	n := len(linkers)

	if n > 1 {
		lastI := n - 1
		return func(ctx *[]string, handler Handler) error {
			var (
				chainHandler Handler
				curI         int
			)

			chainHandler = func(currentCtx *[]string) error {
				if curI == lastI {
					if handler != nil {
						return handler(currentCtx)
					}
					return nil
				}
				curI++
				err := linkers[curI](currentCtx, chainHandler)
				curI--
				return err
			}

			return linkers[0](ctx, chainHandler)
		}
	}

	if n == 1 {
		return linkers[0]
	}

	return func(ctx *[]string, handler Handler) error {
		return handler(ctx)
	}
}

func main() {
	ctx := make([]string, 0)
	handler := func(ctx *[]string) error {
		*ctx = append(*ctx, "hello world")
		return nil
	}

	linker_0 := func(ctx *[]string, handler Handler) error {
		if handler != nil {
			return handler(ctx)
		}
		return nil
	}
	linker_1 := func(ctx *[]string, handler Handler) error {
		if handler != nil {
			return handler(ctx)
		}
		return nil
	}
	linker_2 := func(ctx *[]string, handler Handler) error {
		if handler != nil {
			return handler(ctx)
		}
		return nil
	}

	chain := LinkerChain(linker_0, linker_1, linker_2)
	err := chain(&ctx, handler)
	if err != nil {
		panic(err)
	}
	fmt.Println(ctx)
}
