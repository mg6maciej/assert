package assert

func ThatInt(t TestingT, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      TestingT
	actual int
}

func (assert *IntAssert) isTrue(condition bool, format string, args ...interface{}) *IntAssert {
	if !condition {
		assert.t.Errorf(format, args...)
	}
	return assert
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonzero() *IntAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected integer equal to <%d>, but was <%d>.", expected, assert.actual)
}

func (assert *IntAssert) IsPositive() *IntAssert {
	return assert.isTrue(assert.actual > 0,
		"Expected positive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNegative() *IntAssert {
	return assert.isTrue(assert.actual < 0,
		"Expected negative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsGreaterThan(smallerOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual > smallerOrEqual,
		"Expected integer greater than <%d>, but was <%d>.", smallerOrEqual, assert.actual)
}

func (assert *IntAssert) IsGreaterOrEqualTo(smaller int) *IntAssert {
	return assert.isTrue(assert.actual >= smaller,
		"Expected integer greater or equal to <%d>, but was <%d>.", smaller, assert.actual)
}
