package patterns

import "fmt"

type IState interface {
	render()
	publish()
}

type Draft struct {
}

func (d *Draft) render() {
	fmt.Println("Draft render")
}

func (d *Draft) publish() {
	fmt.Println("Draft publish")
}

type Moderated struct {
}

func (m *Moderated) render() {
	fmt.Println("Moderated render")
}

func (m *Moderated) publish() {
	fmt.Println("Moderated publish")
}

type Published struct {
}

func (p *Published) render() {
	fmt.Println("Published render")
}

func (p *Published) publish() {
	fmt.Println("Published publish")
}

type Document struct {
	state IState
}

func (d *Document) render() {
	d.state.render()
}

func (d *Document) publish() {
	d.state.publish()
}

func (d *Document) newState(state IState) {
	d.state = state
}

func state() {
	d := &Document{}
	d.newState(&Draft{})
	d.render()
	d.publish()
	d.newState(&Moderated{})
	d.render()
	d.publish()
	d.newState(&Published{})
	d.render()
	d.publish()
}
