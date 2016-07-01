//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
