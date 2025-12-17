.PHONY: clean

all: hello ackermann factorial fib fib_bench hello_simple hello_test minimal primes sdl3example test_bool test_bool2 test_cli_args test_cstring_access test_fstring test_match

hello: add.c67 hello.c67
	c67 -o $@ $<

ackermann: ackermann.c67
	c67 -o $@ $<

factorial: ackermann.c67
	c67 -o $@ $<

fib: factorial.c67
	c67 -o $@ $<

fib_bench: fib.c67
	c67 -o $@ $<

hello_simple: fib_bench.c67
	c67 -o $@ $<

hello_test: hello_simple.c67
	c67 -o $@ $<

minimal: hello_test.c67
	c67 -o $@ $<

primes: minimal.c67
	c67 -o $@ $<

sdl3example: primes.c67
	c67 -o $@ $<

test_bool: sdl3example.c67
	c67 -o $@ $<

test_bool2: test_bool.c67
	c67 -o $@ $<

test_cli_args: test_bool2.c67
	c67 -o $@ $<

test_cstring_access: test_cli_args.c67
	c67 -o $@ $<

test_fstring: test_cstring_access.c67
	c67 -o $@ $<

test_match: test_fstring.c67
	c67 -o $@ $<

clean:
	rm -f hello ackermann factorial fib fib_bench hello_simple hello_test minimal primes sdl3example test_bool test_bool2 test_cli_args test_cstring_access test_fstring test_match
