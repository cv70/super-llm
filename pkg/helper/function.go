package helper

import (
	"context"
)

func IFunction[I any](fn func(ctx context.Context, input I) error) func(context.Context, I) (struct{}, error) {
	return func(ctx context.Context, input I) (struct{}, error) {
		return struct{}{}, fn(ctx, input)
	}
}

func OFunction[O any](fn func(ctx context.Context) (O, error)) func(context.Context, struct{}) (O, error) {
	return func(ctx context.Context, input struct{}) (O, error) {
		return fn(ctx)
	}
}

func Function(fn func(ctx context.Context) error) func(context.Context, struct{}) (struct{}, error) {
	return func(ctx context.Context, input struct{}) (struct{}, error) {
		return struct{}{}, fn(ctx)
	}
}

func FunctionInject[D any](data D, fn func(context.Context, D) error) func(context.Context, struct{}) (struct{}, error) {
	return func(ctx context.Context, _ struct{}) (struct{}, error) {
		return struct{}{}, fn(ctx, data)
	}
}

func IFunctionInject[D any, I any](data D, fn func(context.Context, D, I) error) func(context.Context, I) (struct{}, error) {
	return func(ctx context.Context, input I) (struct{}, error) {
		return struct{}{}, fn(ctx, data, input)
	}
}

func OFunctionInject[D any, O any](data D, fn func(context.Context, D) (O, error)) func(context.Context, struct{}) (O, error) {
	return func(ctx context.Context, _ struct{}) (O, error) {
		return fn(ctx, data)
	}
}

func IOFunctionInject[D any, I any, O any](data D, fn func(context.Context, D, I) (O, error)) func(context.Context, I) (O, error) {
	return func(ctx context.Context, input I) (O, error) {
		return fn(ctx, data, input)
	}
}
