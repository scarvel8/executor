package types

// Runner type prelimiary struct for running stuff
type Runner struct {
	command string
	shell []string
}

// Command ...
func (r *Runner) Command() string {
	return r.command
}

// Shell ...
func (r *Runner) Shell() []string {
	return r.shell
}

// Init will initialize the runner (bash if no shell provided)
func (r *Runner) Init(shell []string, command string) error {
	if shell[0] != "" {
		r.shell = shell
	} else {
		r.shell = []string {"sh","-c"}
	}
	r.command = command

	return nil
}
