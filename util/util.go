package util

func Merge(a []interface{}, b []interface{}) []interface{} {
	c := make([]interface{}, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
