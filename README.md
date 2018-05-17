# float256 - cryptocurrency ledger token denomination math library

Golang library built on top of math/big Float for high precision financial calculations, intended for use with a cryptocurrency ledger with a maximum integer part of the maximum supply of 4,398,046,511,103, and an unsigned fixed point 42.214 bit 256 bit form for storage and transmission.

This is mainly just convenience wrappers to make working with big.Float easier (functions simply take variables and return results instead of the confusing return-and-modify in `math/big`, with an added power function (Exp) and n<sup>th</sup> root function (Root) that automatically configures for 256 bits of precision.

It also has convenience converters to and from big.Int, string, int64 and uint64, and numerous other standard basic math library functions. There is no trigonometric functions as these are (pretty much) never used in financial calculations, and also largely absent from the base math/big library.

To make your code a little less wordy you may want to alias the import with a shorter name with or two letters like f, F, or fl or something similar. The return values are plain vanilla `*big.Float` variables so it is possible to use them in the normal `math/big` manner should there be any need to inspect internals such as precision, mantissa, exponent, accuracy, and so forth.