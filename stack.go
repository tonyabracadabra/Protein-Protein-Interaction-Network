package main

type Stack []Edge

func (s *Stack) Push(e Edge) {
    *s = append(*s, e)
}

func (s *Stack) Pop() Edge {
    e := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return e
}
