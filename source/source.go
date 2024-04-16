package source

func Map[T](s []T, f func(T) T) []T {
    r := make([]T, len(s))

    for i, e := range s {
        r[i] = f(e)
    }
    return r
}


