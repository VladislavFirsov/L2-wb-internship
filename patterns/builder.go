package main

import (
	"fmt"
)

type computerBuilderI interface {
	setCpu(value string) computerBuilderI
	setVideoCard(value string) computerBuilderI
	setRam(value int) computerBuilderI

	build() *computer
}

type computer struct {
	cpu       string
	videoCard string
	ram       int
}

func (c computer) showStats() {
	fmt.Printf("cpu: %s\nvideocard: %s\nram: %d gb", c.cpu, c.videoCard, c.ram)
}

type computerBuilder struct {
	cpu       string
	videoCard string
	ram       int
}

func (c *computerBuilder) setCpu(value string) computerBuilderI {
	c.cpu = value
	return c
}
func (c *computerBuilder) setVideoCard(value string) computerBuilderI {
	c.videoCard = value
	return c
}
func (c *computerBuilder) setRam(value int) computerBuilderI {
	c.ram = value
	return c
}
func (c *computerBuilder) build() *computer {
	return &computer{
		cpu:       c.cpu,
		videoCard: c.videoCard,
		ram:       c.ram,
	}
}

func newComputerBuilder() computerBuilderI {
	return &computerBuilder{}
}
func main() {
	compBuilder := newComputerBuilder()
	computer := compBuilder.setRam(16).setCpu("core i7").setVideoCard("nvidia GTX 3080").build()
	computer.showStats()
}
