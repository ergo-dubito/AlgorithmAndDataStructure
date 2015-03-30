package main_test

import (
	Algo "AlgorithmAndDataStructure"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AlgorithmAndDataStructure", func() {

	Describe("non modifying sequence", func() {
		var a []int
		var b []int
		var equal func(int, int) bool

		BeforeEach(func() {
			a = []int{1, 2, 3}
			b = []int{1, 1, 3}
			equal = func(i int, j int) bool {
				return i == j
			}

		})

		Context("AllOf", func() {
			It("should be true", func() {
				Expect(Algo.AllOf(a, func(i int) bool {
					return i < 4
				})).To(BeTrue())
			})

			It("should be false", func() {
				Expect(Algo.AllOf(a, func(i int) bool {
					return i < 3
				})).To(BeFalse())
			})
		})

		It("AnyOf", func() {
			Expect(Algo.AnyOf(a, func(i int) bool {
				return i > 1
			})).To(BeTrue())
		})

		It("NoneOf", func() {
			Expect(Algo.NoneOf(a, func(i int) bool {
				return i > 4
			})).To(BeTrue())
		})

		It("ForEach", func() {
			tmp := make([]int, 0)
			Algo.ForEach(a, func(i int) {
				tmp = append(tmp, i)
			})
			Expect(tmp).To(Equal(a))
		})

		It("Count", func() {
			Expect(Algo.Count(a)).To(Equal(len(a)))
		})

		It("CountIf", func() {
			Expect(Algo.CountIf(a, func(i int) bool {
				return i%2 == 0
			})).To(Equal(1))
		})

		Context("Mismatch", func() {
			It("has misMatch", func() {
				first, second, hasMisMatch := Algo.Mismatch(a, b,
					func(i int, j int) bool {
						return i == j
					})
				Expect(hasMisMatch).To(BeTrue())
				Expect(first).To(Equal(2))
				Expect(second).To(Equal(1))
			})

			It("don't has mismatch", func() {
				c := make([]int, len(a))
				copy(c, a)
				first, second, hasMisMatch := Algo.Mismatch(a, c,
					func(i int, j int) bool {
						return i == j
					})
				Expect(hasMisMatch).To(BeFalse())
				Expect(first).To(Equal(0))
				Expect(second).To(Equal(0))
			})
		})

		It("Equal", func() {
			same := []int{1, 2, 3}
			Expect(Algo.Equal(a, same, func(first int, second int) bool {
				return first == second
			})).To(BeTrue())
		})

		It("Find", func() {
			Expect(Algo.Find(a, 2)).To(Equal(1))
		})

		It("FindIf", func() {
			Expect(Algo.FindIf(a, func(i int) bool {
				return i == 2
			})).To(Equal(1))
		})

		It("FindIfNot", func() {
			Expect(Algo.FindIfNot(a, func(i int) bool {
				return i == 2
			})).To(Equal(0))
		})

		Context("FindFirstOf", func() {
			It("Not existed", func() {
				tmp := []int{4, 5, 6}
				_, found := Algo.FindFirstOf(tmp, a, equal)
				Expect(found).To(BeFalse())
			})

			It("Only one existed", func() {
				tmp := []int{0, 1}
				idx, found := Algo.FindFirstOf(tmp, a, equal)
				Expect(found).To(BeTrue())
				Expect(idx).To(Equal(1))
			})

			It("existed many times", func() {
				tmp := []int{0, 1, 2, 3, 1, 2, 3}
				idx, _ := Algo.FindFirstOf(tmp, a, equal)
				Expect(idx).To(Equal(1))
			})
		})

		Context("FindEnd", func() {
			It("Not Existed", func() {
				tmp := []int{1, 2, 0}
				_, found := Algo.FindEnd(tmp, a, equal)
				Expect(found).To(BeFalse())
			})

			It("Existed", func() {
				tmp := []int{1, 2, 1, 2, 1}
				c := []int{1, 2, 1}
				idx, found := Algo.FindEnd(tmp, c, equal)
				Expect(found).To(BeTrue())
				Expect(idx).To(Equal(2))
			})
		})

		Context("AdjacentFind", func() {
			It("Not existed", func() {
				tmp := []int{1, 2, 3}
				_, found := Algo.AdjacentFind(tmp, equal)
				Expect(found).To(BeFalse())
			})

			It("Existed", func() {
				tmp := []int{1, 1, 3}
				idx, found := Algo.AdjacentFind(tmp, equal)
				Expect(found).To(BeTrue())
				Expect(idx).To(Equal(0))
			})
		})

		Context("Search", func() {
			It("Not existed", func() {
				tmp := []int{1, 2, 2, 4}
				_, found := Algo.Search(tmp, a, equal)
				Expect(found).To(BeFalse())
			})

			It("Existed", func() {
				tmp := []int{0, 1, 2, 3}
				idx, found := Algo.Search(tmp, a, equal)
				Expect(found).To(BeTrue())
				Expect(idx).To(Equal(1))
			})
		})

		Context("SearchN", func() {
			tmp := []int{0, 1, 1, 1}

			It("Existed", func() {
				idx, found := Algo.SearchN(tmp, 3, 1, equal)
				Expect(found).To(BeTrue())
				Expect(idx).To(Equal(1))
			})

			It("Not existed", func() {
				_, found := Algo.SearchN(tmp, 4, 1, equal)
				Expect(found).To(BeFalse())
			})
		})
	})

	Describe("modifying sequence", func() {
		var src []int
		var dest []int
		var isOdd func(int) bool

		BeforeEach(func() {
			src = []int{1, 2, 3}
			dest = make([]int, 3)
			isOdd = func(i int) bool {
				return i%2 != 0
			}
		})

		It("CopyIf", func() {
			Expect(Algo.CopyIf(dest, src, isOdd)).To(Equal(2))
			Expect(dest[0]).To(Equal(src[0]))
			Expect(dest[1]).To(Equal(src[2]))
		})

		It("CopyN", func() {
			Algo.CopyN(dest, src, 1)
			Expect(dest[0]).To(Equal(src[0]))
			Expect(dest[1]).NotTo(Equal(src[1]))
		})

		It("CopyBackward", func() {
			dest = make([]int, 2)
			Expect(Algo.CopyBackward(dest, src)).To(Equal(2))
			Expect(dest[0]).To(Equal(src[1]))
			Expect(dest[1]).To(Equal(src[2]))
		})

		It("Fill", func() {
			tmp := []int{1, 1, 1}
			Algo.Fill(dest, 1)
			Expect(dest).To(Equal(tmp))
		})

		It("FillN", func() {
			tmp := []int{1, 1, 1}
			Algo.FillN(dest, 3, 1)
			Expect(dest).To(Equal(tmp))
			Algo.FillN(dest, len(dest)+1, 1)
		})

		It("Transform", func() {
			dest = make([]int, 2)
			plusOne := func(prev int) int {
				return prev + 1
			}
			Expect(Algo.Transform(dest, src, plusOne)).
				To(Equal([]int{2, 3}))
		})

		Context("Generate", func() {
			allAreOne := func() int {
				return 1
			}

			It("Generate", func() {
				Algo.Generate(dest, allAreOne)
				Expect(dest).To(Equal([]int{1, 1, 1}))
			})

			It("GenerateN", func() {
				dest[2] = 2
				Algo.GenerateN(dest, 2, allAreOne)
				Expect(dest).To(Equal([]int{1, 1, 2}))
			})
		})

		It("Remove", func() {
			Expect(Algo.Remove(src, 1)).To(Equal([]int{2, 3}))
		})

		It("RemoveIf", func() {
			Expect(Algo.RemoveIf(src, isOdd)).To(Equal([]int{2}))
		})
	})
})
