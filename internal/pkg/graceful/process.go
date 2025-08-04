package graceful

// NewProcess создает новый процесс с переданным стартером.
func NewProcess(starter starter) Process {
	return Process{
		starter:  starter,
		disabled: false,
	}
}

// Disable возвращает копию процесса с обновленным флагом disabled.
func (p Process) Disable(d bool) Process {
	p.disabled = d

	return p
}
