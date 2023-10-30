package notifier

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	"sync"
)

type PodStateMachine struct {
	lastState map[string]PodState
	modify    chan PodState
	stop      chan struct{}
	once      sync.Once
	sync.RWMutex
}

type PodState struct {
	PodName string
	App     string
	State   State
}

var stateMachine *PodStateMachine
var once sync.Once

func NewStateMachine() *PodStateMachine {
	once.Do(func() {
		stateMachine = &PodStateMachine{
			lastState: map[string]PodState{},
			modify:    make(chan PodState),
			stop:      make(chan struct{}),
			once:      sync.Once{},
			RWMutex:   sync.RWMutex{},
		}
	})
	return stateMachine
}

func (m *PodStateMachine) Modify(state PodState) {
	m.modify <- state
}

func (m *PodStateMachine) Start() {
	m.once.Do(func() {
		log.Info().Msgf("Start pod state machine")
		go func() {
			for {
				podState := PodState{}
				select {
				case newState := <-m.modify:
					m.RLock()
					if last, ok := m.lastState[newState.PodName]; ok {
						podState.App = newState.App
						podState.PodName = newState.PodName
						switch last.State {
						case PodCreating:
							switch newState.State {
							// New pod created.
							case PodCreated:
								podState.State = PodCreated
							}
						case PodCreated:
							switch newState.State {
							// old deleted.
							case PodDeleted:
								podState.State = PodDeleted
							}
						}
					} else {
						switch newState.State {
						// Pod deleted.
						case PodDeleted:
							podState.State = PodDeleted
						}
					}
					m.RUnlock()

				case <-m.stop:
					close(m.modify)
					close(m.stop)
					return
				}

				if podState.State != None {
					NotifyPodModified(context.Background(), podState)
					m.Lock()
					if podState.State == PodDeleted {
						delete(m.lastState, podState.PodName)
					} else {
						m.lastState[podState.PodName] = podState
					}
					m.Unlock()
				}
			}
		}()
	})

}
