package assert

func ThatString(t TestingT, actual string) *StringAssert {
	return &StringAssert{t, actual}
}

type StringAssert struct {
	t      TestingT
	actual string
}

func (assert *StringAssert) IsEqualTo(expected string) *StringAssert {
	if assert.actual != expected {
		assert.t.Errorf("Expected <%s>, but was <%s>.\n", expected, assert.actual)
	}
	return assert
}