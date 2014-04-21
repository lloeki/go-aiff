package aiff_test

import (
    "testing"
    check "gopkg.in/check.v1"
    aiff "."
)

func Test(t *testing.T) { check.TestingT(t) }

type IffSuite struct{}

var _ = check.Suite(&IffSuite{})
