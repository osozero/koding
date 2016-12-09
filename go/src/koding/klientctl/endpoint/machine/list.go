package machine

import (
	"fmt"
	"os"
	"sort"
	"time"

	"koding/kites/kloud/machine"
	"koding/kites/kloud/machinestate"
	"koding/kites/kloud/stack"
	kmachine "koding/klient/machine"
	"koding/klient/machine/machinegroup"
	"koding/klientctl/endpoint/kloud"
	"koding/klientctl/klient"

	"github.com/koding/logging"
)

// ListOptions stores options for `machine list` call.
type ListOptions struct {
	Log logging.Logger
}

// List retrieves user's machines from kloud.
func List(options *ListOptions) ([]*Info, error) {
	var (
		req = stack.MachineListRequest{}
		res = stack.MachineListResponse{}
	)

	// Get info from kloud.
	if err := kloud.Call("machine.list", &req, &res); err != nil {
		fmt.Fprintln(os.Stderr, "Error communicating with Koding:", err)
		return nil, err
	}

	// Register machines to klient and get aliases.
	//
	// TODO(ppknap): this is copied from klientctl old list and will be reworked.
	k, err := klient.CreateKlientWithDefaultOpts()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating klient:", err)
		return nil, err
	}

	if err := k.Dial(); err != nil {
		fmt.Fprintln(os.Stderr, "Error dialing klient:", err)
		return nil, err
	}

	mgreq := machinegroup.CreateRequest{
		Addresses: make(map[kmachine.ID][]kmachine.Addr),
	}
	for _, m := range res.Machines {
		mgreq.Addresses[kmachine.ID(m.ID)] = []kmachine.Addr{
			{
				Network:   "ip",
				Value:     m.IP,
				UpdatedAt: time.Now(),
			},
		}
	}
	kresraw, err := k.Tell("machine.create", mgreq)
	if err != nil {
		return nil, err
	}

	mgres := machinegroup.CreateResponse{
		Statuses: make(map[kmachine.ID]kmachine.Status),
		Aliases:  make(map[kmachine.ID]string),
	}
	if err := kresraw.Unmarshal(&mgres); err != nil {
		return nil, err
	}

	infos := make([]*Info, len(res.Machines))
	for i, m := range res.Machines {
		infos[i] = &Info{
			ID:        m.ID,
			Alias:     mgres.Aliases[kmachine.ID(m.ID)],
			Team:      m.Team,
			Stack:     m.Stack,
			Provider:  m.Provider,
			Label:     m.Label,
			IP:        m.IP,
			CreatedAt: m.CreatedAt,
			Status: kmachine.MergeStatus(kmachine.Status{
				State:  fromMachineStateString(m.Status.State),
				Reason: m.Status.Reason,
				Since:  m.Status.ModifiedAt,
			}, mgres.Statuses[kmachine.ID(m.ID)]),
			Username: machineUserFromUsers(m.Users),
			Owner:    ownerFromUsers(m.Users),
		}
	}

	// Sort items before we return.
	sort.Sort(InfoSlice(infos))

	return infos, nil
}

// machineUserFromUsers gets machine user. Which may be different from machine
// owner when machine. is shared.
func machineUserFromUsers(users []machine.User) string {
	switch len(users) {
	case 1:
		return users[0].Username
	case 2:
		if !users[0].Owner {
			return users[0].Username
		}
		return users[1].Username
	default:
		return "<unknown>"
	}
}

// ownerFromUsers returns machine owner when machine is shared.
func ownerFromUsers(users []machine.User) string {
	if len(users) == 2 {
		if users[0].Owner {
			return users[0].Username
		}
		return users[1].Username
	}

	return ""
}

// ms2State maps machinestate states to State objects.
var ms2State = map[machinestate.State]kmachine.State{
	machinestate.NotInitialized: kmachine.StateOffline,
	machinestate.Building:       kmachine.StateOffline,
	machinestate.Starting:       kmachine.StateOffline,
	machinestate.Running:        kmachine.StateOnline,
	machinestate.Stopping:       kmachine.StateOffline,
	machinestate.Stopped:        kmachine.StateOffline,
	machinestate.Rebooting:      kmachine.StateOffline,
	machinestate.Terminating:    kmachine.StateOffline,
	machinestate.Terminated:     kmachine.StateOffline,
	machinestate.Snapshotting:   kmachine.StateOffline,
	machinestate.Pending:        kmachine.StateOffline,
	machinestate.Unknown:        kmachine.StateUnknown,
}

// FromMachineStateString converts machinestate string to State object.
func fromMachineStateString(raw string) kmachine.State {
	return ms2State[machinestate.States[raw]]
}
