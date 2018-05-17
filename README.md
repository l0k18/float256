# float32
Golang library built on top of math/big Float for basic high precision financial calculations

This is mainly just convenience wrappers to make working with big.Float easier (functions simply take variables and return results instead of the confusing return-and-modify in `math/big`, with an added power function (Exp) and n<sup>th</sup> root function (Root) that automatically configures for 256 bits of precision. The purpose is for implementing super preecise financial calculations.