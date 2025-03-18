package mflag

import (
	"reflect"

	"github.com/spf13/cobra"
)

type Ts struct {
	cc    *cobra.Command
	flags any
}

func New(flags any) *Ts {
	ref := &Ts{
		cc: &cobra.Command{
			CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
		},
	}

	ref.flags = flags
	return ref
}

func (t *Ts) Cobra() *cobra.Command {
	return t.cc
}

func (t *Ts) SetName(name, short string) *cobra.Command {
	t.cc.Use = name
	t.cc.Short = short
	return t.cc
}

func (t *Ts) UsePackageName(short string) *Ts {
	t.SetName(GetPackageName(2), short)
	return t
}

func (t *Ts) AddCobra(cc *cobra.Command) {
	t.cc.AddCommand(cc)
}

func (t *Ts) AddCmd(
	runFunc func(cmd *cobra.Command, args []string),
	name, short string,
	group ...string,
) {
	cmd := &cobra.Command{
		Use:   name,
		Short: short,
		Run:   runFunc,
	}
	if len(group) > 0 {
		if t.flags != nil {
			bindFieldTag(cmd, reflect.ValueOf(t.flags).Elem(), "", group)
		}
	}

	t.cc.AddCommand(cmd)

}

func (c *Ts) Execute() error {
	return c.cc.Execute()
}
