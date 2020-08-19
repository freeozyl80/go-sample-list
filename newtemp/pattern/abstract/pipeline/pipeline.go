package pipeline

import (
	"pattern/abstract/plugin"
	"pattern/abstract/config"
)

type Pipeline struct {
    input  plugin.Input
    filter plugin.Filter
    output plugin.Output
}
// 一个消息的处理流程为 input -> filter -> output
func (p *Pipeline) Exec() {
    msg := p.input.Receive()
    msg = p.filter.Process(msg)
    p.output.Send(msg)
}

func DefaultConfig () *Config {
	return &config.Config {
		Input: plugin.Config {
			PluginType: plugin.InputType
		}
	}
}