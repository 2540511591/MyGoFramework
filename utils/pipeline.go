package utils

import "errors"

/*
管道创建对象
注入原料Material，返回产品Product
*/
type PipeLine[Material interface{}, Product interface{}] struct {
	pipes []func(Material, func(Material) Product) Product
	final func(Material) Product
}

// 设置管道
func (p *PipeLine[Material, Product]) SetPipes(pipes []func(Material, func(Material) Product) Product) {
	p.pipes = pipes
}

// 设置最后一个管道
func (p *PipeLine[Material, Product]) SetFinal(final func(Material) Product) {
	p.final = final
}

// 创建
func (p *PipeLine[Material, Product]) Create() (func(Material) Product, error) {
	if p.pipes == nil || p.final == nil {
		return nil, errors.New("参数不足")
	}

	next := p.final
	l := len(p.pipes)
	for ; l > 0; l-- {
		next = p.assemble(p.pipes[l-1], next)
	}

	return next, nil
}

// 管道组装
func (p *PipeLine[Material, Product]) assemble(current func(Material, func(Material) Product) Product, next func(Material) Product) func(Material) Product {
	return func(material Material) Product {
		return current(material, next)
	}
}
