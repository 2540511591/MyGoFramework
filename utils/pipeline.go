package utils

import "errors"

/*
管道创建对象
注入原料Material，返回产品Product
*/
type PipeLine[Material interface{}, Product interface{}] struct {
	pipes []func(Material, func(Material) Product) Product
	final func(Material) Product
	count uint32
}

// 设置管道
func (p *PipeLine[Material, Product]) SetPipes(pipes []func(Material, func(Material) Product) Product) {
	p.pipes = pipes
}

// 向管道末尾添加一个管道
func (p *PipeLine[Material, Product]) AddPipe(pipe func(Material, func(Material) Product) Product) {
	p.pipes = append(p.pipes, pipe)
}

// 设置最后一个管道
func (p *PipeLine[Material, Product]) SetFinal(final func(Material) Product) {
	p.final = final
}

// 全部组装
// 逆序组装
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

// 懒汉式组装(调用时才进行组装)
// 顺序组装
// 因需要count成员属性进行计数问题，同一实例不支持高并发
// 高并发请使用上面的create方法
func (p *PipeLine[Material, Product]) CreateIdler() func(material Material) Product {
	p.count = 0
	return func(material Material) Product {
		return p.assembleIdler(material)
	}
}

// 懒汉式组装
func (p *PipeLine[Material, Product]) assembleIdler(material Material) Product {
	if p.count >= uint32(len(p.pipes)) {
		return p.final(material)
	}

	current := p.pipes[p.count]
	p.count++
	return current(material, p.assembleIdler)
}
