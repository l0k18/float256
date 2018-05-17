# float256
Golang library built on top of math/big Float for high precision financial calculations.

This is mainly just convenience wrappers to make working with big.Float easier (functions simply take variables and return results instead of the confusing return-and-modify in `math/big`, with an added power function (Exp) and n<sup>th</sup> root function (Root) that automatically configures for 256 bits of precision.

It also has convenience converters to and from big.Int, string, int64 and uint64, and numerous other standard basic math library functions. There is no trigonometric functions as these are (pretty much) never used in financial calculations.

To make your code a little less wordy you may want to alias the import with a shorter name with or two letters like f, F, or fl or something similar. The return values are plain vanilla `*big.Float` variables so it is possible to use them in the normal `math/big` manner to variables created by assignments from these functions.

As this library is intended to be used for a cryptocurrency token denomination, it will also have an encoder/decoder that encodes the values as a fixed point value stored in 32 bytes with no sign (cryptocurrency ledgers do not have negative values).