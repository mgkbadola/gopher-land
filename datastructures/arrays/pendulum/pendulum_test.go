package pendulum

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPendulum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pendulum Suite")
}

var _ = Describe("Test Example", func() {
	It("Testing for []int{4, 10, 9}", func() { Expect(Pendulum([]int{4, 10, 9})).To(Equal([]int{10, 4, 9})) })
	It("Testing for []int{8, 7, 10, 3}", func() { Expect(Pendulum([]int{8, 7, 10, 3})).To(Equal([]int{8, 3, 7, 10})) })
	It("Testing for []int{6, 6, 8, 5, 10}", func() { Expect(Pendulum([]int{6, 6, 8, 5, 10})).To(Equal([]int{10, 6, 5, 6, 8})) })
	It("Testing for []int{9, 4, 6, 4, 10, 5}", func() { Expect(Pendulum([]int{9, 4, 6, 4, 10, 5})).To(Equal([]int{9, 5, 4, 4, 6, 10})) })
	It("Testing for []int{4, 6, 8, 7, 5}", func() { Expect(Pendulum([]int{4, 6, 8, 7, 5})).To(Equal([]int{8, 6, 4, 5, 7})) })
	It("Testing for []int{10, 5, 6, 10}", func() { Expect(Pendulum([]int{10, 5, 6, 10})).To(Equal([]int{10, 5, 6, 10})) })
	It("Testing for []int{11, 12, 12}", func() { Expect(Pendulum([]int{11, 12, 12})).To(Equal([]int{12, 11, 12})) })
	It("Testing for []int{27, 27, 19, 21, 22, 28, 24}", func() {
		Expect(Pendulum([]int{27, 27, 19, 21, 22, 28, 24})).To(Equal([]int{28, 27, 22, 19, 21, 24, 27}))
	})
	It("Testing for []int{20, 19, 25, 16, 19, 30, 18, 24}", func() {
		Expect(Pendulum([]int{20, 19, 25, 16, 19, 30, 18, 24})).To(Equal([]int{25, 20, 19, 16, 18, 19, 24, 30}))
	})
	It("Testing for []int{22, 26, 21, 27, 24, 21, 15, 26, 25}", func() {
		Expect(Pendulum([]int{22, 26, 21, 27, 24, 21, 15, 26, 25})).To(Equal([]int{27, 26, 24, 21, 15, 21, 22, 25, 26}))
	})
	It("Testing for []int{19, 30, 16, 19, 28, 26, 28, 17, 21, 17}", func() {
		Expect(Pendulum([]int{19, 30, 16, 19, 28, 26, 28, 17, 21, 17})).To(Equal([]int{28, 26, 19, 17, 16, 17, 19, 21, 28, 30}))
	})
	It("Testing for []int{27, 28, 26, 28, 24, 22, 18, 16, 15, 24}", func() {
		Expect(Pendulum([]int{27, 28, 26, 28, 24, 22, 18, 16, 15, 24})).To(Equal([]int{28, 26, 24, 18, 15, 16, 22, 24, 27, 28}))
	})
	It("Testing for []int{17, 26, 15, 26, 26, 20, 16, 18, 15}", func() {
		Expect(Pendulum([]int{17, 26, 15, 26, 26, 20, 16, 18, 15})).To(Equal([]int{26, 26, 18, 16, 15, 15, 17, 20, 26}))
	})
	It("Testing for []int{22, 21, 19, 27, 18, 15, 24, 24}", func() {
		Expect(Pendulum([]int{22, 21, 19, 27, 18, 15, 24, 24})).To(Equal([]int{24, 22, 19, 15, 18, 21, 24, 27}))
	})
	It("Testing for []int{33, 38, 38, 36, 43, 48, 32, 40, 47, 50, 33}", func() {
		Expect(Pendulum([]int{33, 38, 38, 36, 43, 48, 32, 40, 47, 50, 33})).To(Equal([]int{50, 47, 40, 38, 33, 32, 33, 36, 38, 43, 48}))
	})
	It("Testing for []int{49, 40, 41, 41, 39, 43, 40, 46, 30, 47, 46, 40}", func() {
		Expect(Pendulum([]int{49, 40, 41, 41, 39, 43, 40, 46, 30, 47, 46, 40})).To(Equal([]int{47, 46, 41, 40, 40, 30, 39, 40, 41, 43, 46, 49}))
	})
	It("Testing for []int{48, 41, 38, 35, 50, 46, 38, 42, 37, 49, 44, 32, 37}", func() {
		Expect(Pendulum([]int{48, 41, 38, 35, 50, 46, 38, 42, 37, 49, 44, 32, 37})).To(Equal([]int{50, 48, 44, 41, 38, 37, 32, 35, 37, 38, 42, 46, 49}))
	})
	It("Testing for []int{49, 30, 39, 30, 40, 44, 43, 48, 47, 50, 42, 48, 33}", func() {
		Expect(Pendulum([]int{49, 30, 39, 30, 40, 44, 43, 48, 47, 50, 42, 48, 33})).To(Equal([]int{50, 48, 47, 43, 40, 33, 30, 30, 39, 42, 44, 48, 49}))
	})
	It("Testing for []int{48, 43, 35, 47, 39, 38, 38, 46, 49, 32, 42}", func() {
		Expect(Pendulum([]int{48, 43, 35, 47, 39, 38, 38, 46, 49, 32, 42})).To(Equal([]int{49, 47, 43, 39, 38, 32, 35, 38, 42, 46, 48}))
	})
	It("Testing for []int{-9, -2, -10, -6}", func() { Expect(Pendulum([]int{-9, -2, -10, -6})).To(Equal([]int{-6, -10, -9, -2})) })
	It("Testing for []int{-3, -6, -7}", func() { Expect(Pendulum([]int{-3, -6, -7})).To(Equal([]int{-3, -7, -6})) })
	It("Testing for []int{-7, -8, -2, -3, -4}", func() { Expect(Pendulum([]int{-7, -8, -2, -3, -4})).To(Equal([]int{-2, -4, -8, -7, -3})) })
	It("Testing for []int{-8, -8, -9, -10, -10, -3}", func() { Expect(Pendulum([]int{-8, -8, -9, -10, -10, -3})).To(Equal([]int{-8, -9, -10, -10, -8, -3})) })
	It("Testing for []int{-7, -10, -1, -10, -8}", func() { Expect(Pendulum([]int{-7, -10, -1, -10, -8})).To(Equal([]int{-1, -8, -10, -10, -7})) })
	It("Testing for []int{-6, -2, -5}", func() { Expect(Pendulum([]int{-6, -2, -5})).To(Equal([]int{-2, -6, -5})) })
	It("Testing for []int{-2, -11, -6, -11, -4, -3, -5}", func() {
		Expect(Pendulum([]int{-2, -11, -6, -11, -4, -3, -5})).To(Equal([]int{-2, -4, -6, -11, -11, -5, -3}))
	})
	It("Testing for []int{-19, -9, -5, -6, -15, -16, -5, -12}", func() {
		Expect(Pendulum([]int{-19, -9, -5, -6, -15, -16, -5, -12})).To(Equal([]int{-5, -9, -15, -19, -16, -12, -6, -5}))
	})
	It("Testing for []int{-18, -2, -11, -10, -6, -7, -7, -12, -16}", func() {
		Expect(Pendulum([]int{-18, -2, -11, -10, -6, -7, -7, -12, -16})).To(Equal([]int{-2, -7, -10, -12, -18, -16, -11, -7, -6}))
	})
	It("Testing for []int{-10, -10, -12, -13, -5, -10, -4, -17, -5, -12}", func() {
		Expect(Pendulum([]int{-10, -10, -12, -13, -5, -10, -4, -17, -5, -12})).To(Equal([]int{-5, -10, -10, -12, -17, -13, -12, -10, -5, -4}))
	})
	It("Testing for []int{-33, -21, -6, -29, -24, -5, -50, -42, -43, -17, -17}", func() {
		Expect(Pendulum([]int{-33, -21, -6, -29, -24, -5, -50, -42, -43, -17, -17})).To(Equal([]int{-5, -17, -21, -29, -42, -50, -43, -33, -24, -17, -6}))
	})
	It("Testing for []int{-4, -50, -32, -23, -47, -44, -43, -24, -29, -44, -20, -35}", func() {
		Expect(Pendulum([]int{-4, -50, -32, -23, -47, -44, -43, -24, -29, -44, -20, -35})).To(Equal([]int{-20, -24, -32, -43, -44, -50, -47, -44, -35, -29, -23, -4}))
	})
	It("Testing for []int{-36, -38, -44, -47, -41, -27, -10, -30, -22, -11, -23, -50, -23}", func() {
		Expect(Pendulum([]int{-36, -38, -44, -47, -41, -27, -10, -30, -22, -11, -23, -50, -23})).To(Equal([]int{-10, -22, -23, -30, -38, -44, -50, -47, -41, -36, -27, -23, -11}))
	})
	It("Testing for []int{-15, 8, 11}", func() { Expect(Pendulum([]int{-15, 8, 11})).To(Equal([]int{11, -15, 8})) })
	It("Testing for []int{8, -1, -1, -10}", func() { Expect(Pendulum([]int{8, -1, -1, -10})).To(Equal([]int{-1, -10, -1, 8})) })
	It("Testing for []int{-8, 15, 8, -3, -11}", func() { Expect(Pendulum([]int{-8, 15, 8, -3, -11})).To(Equal([]int{15, -3, -11, -8, 8})) })
	It("Testing for []int{15, 17, 3, -20, -1, 3}", func() { Expect(Pendulum([]int{15, 17, 3, -20, -1, 3})).To(Equal([]int{15, 3, -20, -1, 3, 17})) })
	It("Testing for []int{-8, -13, -19, -8, 7, 15, -10}", func() {
		Expect(Pendulum([]int{-8, -13, -19, -8, 7, 15, -10})).To(Equal([]int{15, -8, -10, -19, -13, -8, 7}))
	})
	It("Testing for []int{11, -16, -18, 13, -11, -12, 3, 18}", func() {
		Expect(Pendulum([]int{11, -16, -18, 13, -11, -12, 3, 18})).To(Equal([]int{13, 3, -12, -18, -16, -11, 11, 18}))
	})
	It("Testing for []int{7, -5, -20, 15, 2, 10, 18, 4, -10}", func() {
		Expect(Pendulum([]int{7, -5, -20, 15, 2, 10, 18, 4, -10})).To(Equal([]int{18, 10, 4, -5, -20, -10, 2, 7, 15}))
	})
	It("Testing for []int{-6, 1, 2, 12, 19, 12, 19, -10, 13, 1}", func() {
		Expect(Pendulum([]int{-6, 1, 2, 12, 19, 12, 19, -10, 13, 1})).To(Equal([]int{19, 12, 2, 1, -10, -6, 1, 12, 13, 19}))
	})

})
