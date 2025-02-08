# gobigger

go bigger or go home-r?

## What

List of wrapper functions for Go's math/big, plus additional ease-of-use functions with variadics and more!

I avoided writing every single math/big function as a wrapper for two reasons: namespace cleanliness and lack of need.

Certain functions make more sense to be used in the normal sense (e.g., `a.blah()`) vs. with a wrapper (`blah(a)`). Typically, these functions are those without a parameter (e.g., `a.IsInt64()`), although some were left out for simplicity.

## Why

I had originally created my gobig repository. The problems with this are as follows:

- I found the implementation a little awkward and it required the importer to basically just write the file themselves to actually reap the benefits of the wrappers.
- It also didn't come with additional features so it was not a worth while endeavor.
- It did not wrap every function in math/big.

So, to mitigate these weaknesses, this repository does the following:

- Wraps each and every one of Go's math/big functions!
- Creates additional functions with variadics (like `addall(...)`) to extend the math/big package.
- Doesn't act as a module. Instead, "import" by copying the file into the appropriate package folder and use as needed. This enables better "portability" and doesn't require rewriting of each function like gobig.

These functions are unexported (meant to be private). I made this design decision for two reasons: enable the importer full control of what is imported (and write their own functions in the file if needed) and to reduce keystrokes / improve readability.

There are use cases behind making all the functions unexported (less typing, e.g., utils.mul vs mul) vs exported (avoid polluting the namespace). Ultimately, this is personal preference so I made it entirely up to the user.

You might look at the code and say, "These wrapper functions' names aren't much shorter, and perform the same functions; how is this wrapper worth it?"

Let me demonstrate an example. I will write A000139 from the OEIS database using both math/big only and with these wrapper functions to demonstrate the difference in readability, keystrokes, etc.

First, using only math/big:

```go
// let's write the following function using math/big:
// A000139(n) = 2(3n)!/((2n+1)!*(n+1)!) (this is from OEIS)
func A000139(seqlen int64) []*big.Int {
    a := bigSlice(seqlen)
    for n := int64(0); n < seqlen; n++ {
        nplus1 := fact(big.NewInt(n+1))        // (n+1)!
        twonplus1 := fact(big.NewInt(2*n+1))   // (2n+1)!
        threen := fact(big.NewInt(3*n))        // (3n)!
        numer := big.NewInt(0).Mul(big.NewInt(2), threen)    // 2(3n)!
        denom := big.NewInt(0).Mul(twonplus1, nplus1)

        // convert to floats!
        numer_f := new(big.Float)
        numer_f.SetInt(numer)
        numer_f.SetPrec(53)         // default prec is 53

        denom_f := new(big.Float)
        denom_f.SetInt(denom)
        denom_f.SetPrec(53)         // default prec is 53

        // a(n) = 2(3n)!/((2n+1)!*(n+1)!)
        a[n] = big.NewFloat(0).Div(numer_f, denom_f)
    }
    return a
}

// must write your own factorial function
func fact(a *big.Int) *big.Int {
    prod := big.NewInt(1)
    for i := big.NewInt(2); i.Cmp(a) == -1; i = i.Add(1) {
        prod = prod.Mul(prod, i)
    }
    return prod
}
```

Now, I'll demonstrate by using gobigger.go:

```go
func A000139(seqlen int64) []*bint {
    // a(n) = 2(3n)!/((2n+1)!*(n+1)!)
    a := bigSlice(seqlen)
    for n := int64(0); n < seqlen; n++ {
        nplus1 := fact(inew(n+1))        // (n+1)!
        twonplus1 := fact(inew(2*n+1))   // (2n+1)!
        threen := fact(inew(3*n))        // (3n)!
        numer := mul(inew(2), threen)    // 2(3n)!
        denom := mul(twonplus1, nplus1)
        a[n] = floor(fdiv(tofloat(numer), tofloat(denom)))
    }
    return a
}
```

That goes from about 1000 characters to less than 500 characters; a reduction by over 50%! Plus, it is much easier to read and enables cleaner code for maintainability.

## How

This can be used by following 3 easy steps:

1. Copying the gobigger.go file into your package.
2. Renaming the package in gobigger.go to whatever package it now resides in.
3. These functions are now available to every file in that package! These are private (so they aren't exported/visible), but now they can be used nice and easy!

## Future

Updates would be pushed to this repo, but the base functionality is already here, so updates aren't necessary after copying the file (unless you want new stuff).

TODO: I'm going to build a testing package to ensure these functions work as expected. Once the tester is complete, I will generate a release of v1.0.0
