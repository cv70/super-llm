package gmath

import (
	"math"
	"strconv"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

func ToFixed(v float64, precision int) float64 {
	a, _ := decimal.NewFromFloat(v).Round(int32(precision)).Float64()
	return a
}

func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var ret T
	for _, n := range nums {
		ret += n
	}
	return ret
}

func Exp[T constraints.Float](x T) T {
	return T(math.Exp(float64(x)))
}

func Divide[T constraints.Integer | constraints.Float](dividend, divisor T) float64 {
	epsilon := 1e-10
	if math.Abs(float64(divisor)) <= epsilon {
		return 0.0
	}
	return float64(dividend) / float64(divisor)
}

func Sigmoid[T constraints.Float](x T) T {
	return T(1.0 / (1.0 + math.Exp(float64(-x))))
}

// GreatestCommonDivisor 最大公约数(确保参数为正数)
func GreatestCommonDivisor[T constraints.Integer](a, b T) T {
	var (
		quotient  T
		remainder T
	)
	for {
		quotient = a / b
		remainder = a - quotient*b
		if remainder == 0 {
			return b
		}
		a, b = b, remainder
	}
}

// LeastCommonMultiple 最小公倍数(请确保参数为正数, 且两数的公倍数不会导致溢出)
func LeastCommonMultiple[T constraints.Integer](a, b T) T {
	return a / GreatestCommonDivisor(a, b) * b
}

func GetYByXFunc(x1, y1, x2, y2 float64) func(float64) float64 {
	return func(x float64) float64 {
		return y1 + (y2-y1)/(x2-x1)*(x-x1)
	}
}

// ExponentialDecay Exp时间衰减
// baseValue = 1, decayRate = 0.001
// elapsedTime=30，  衰减后的权重: 0.97
// elapsedTime=100， 衰减后的权重: 0.90
// elapsedTime=200， 衰减后的权重: 0.82
// elapsedTime=400， 衰减后的权重: 0.67
// elapsedTime=1000，衰减后的权重: 0.37
// elapsedTime=1799，衰减后的权重: 0.17
func ExponentialDecay(baseValue float64, decayRate float64, elapsedTime int64, minValue float64) float64 {
	decayFactor := math.Exp(-decayRate * float64(elapsedTime))
	weightedValue := baseValue * decayFactor
	return max(weightedValue, minValue)
}

// HalfLifeDecay 半衰期时间衰减
// halfLife 半衰期 单位为天
// scala 不衰减的时间窗口
func HalfLifeDecay(baseValue float64, scala int64, halfLife int64, elapsedTime int64) float64 {
	decayFactor := math.Pow(0.5, math.Max(0, float64(elapsedTime-scala))/float64(halfLife))
	weightedValue := baseValue * decayFactor
	return weightedValue
}

// StepsDecay 阶梯时间衰减
// linearRate 线性衰减率
// expRate 指数衰减率
// scala 不衰减的时间窗口
// linearWindow 线性衰减时间窗口
func StepsDecay(baseValue float64, scala int64, linearRate float64, linearWindow int64, expRate float64, elapsedTime int64) float64 {
	if elapsedTime <= scala {
		return baseValue
	} else if elapsedTime <= linearWindow {
		return baseValue * (1 - linearRate*float64(elapsedTime-scala))
	} else if elapsedTime <= 2*linearWindow {
		// 分段点平滑处理
		valuePre := baseValue * (1 - linearRate*float64(linearWindow-scala))
		valueNext := baseValue * math.Exp(-expRate*float64(elapsedTime-scala))
		return valuePre + float64(elapsedTime-linearWindow)/float64(linearWindow)*(valueNext-valuePre)
	} else {
		return baseValue * math.Exp(-expRate*float64(elapsedTime-scala))
	}
}

// LogBase 自定义底数
func LogBase(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func SafeDivisionToFloat64(numerator int, denominator int) float64 {
	if denominator == 0 {
		return 0
	} else {
		return float64(numerator) / float64(denominator)
	}
}

// AtanNormalize 反正切归一化，将[0,∞)的值归一化到[0,1)
func AtanNormalize(num int, dis int) float64 {
	if num == 0 || dis == 0 {
		return 0
	}
	return math.Atan(float64(num/dis)) * 2 / math.Pi
}

func SafeDivisionToStr(numerator int32, denominator int32, prec int) string {
	if denominator == 0 {
		return "0"
	} else {
		return strconv.FormatFloat(float64(numerator)/float64(denominator), 'f', prec, 64)
	}
}

// L2Normalize 对切片进行L2归一化，直接修改原数组
func L2Normalize[T constraints.Float](arr []T) {
	if len(arr) == 0 {
		return
	}
	var sum float64
	for _, v := range arr {
		sum += float64(v * v)
	}
	l2Norm := math.Sqrt(sum)
	if l2Norm == 0 {
		return
	}
	for i := range arr {
		arr[i] = T(float64(arr[i]) / l2Norm)
	}
}
