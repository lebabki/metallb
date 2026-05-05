// SPDX-License-Identifier:Apache-2.0

package speakerlist

import "github.com/prometheus/client_golang/prometheus"

// Prometheus metrics for memberlist node events. All three metrics share
// namespace "metallb" / subsystem "speaker" with the existing speaker
// metric `metallb_speaker_announced` (see speaker/main.go).
var stats = metrics{
	nodeEvents: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "metallb",
		Subsystem: "speaker",
		Name:      "memberlist_node_events_total",
		Help:      "Number of memberlist node events received by this speaker, broken down by event type and the affected node name. Includes NodeJoin / NodeLeave / NodeUpdate.",
	}, []string{
		"event_type",
		"node",
	}),

	members: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "metallb",
		Subsystem: "speaker",
		Name:      "memberlist_members",
		Help:      "Current number of members in the memberlist cluster as observed by this speaker (includes self). Updated after every memberlist event.",
	}, []string{
		"self_node",
	}),

	transitions: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "metallb",
		Subsystem: "speaker",
		Name:      "memberlist_transitions_total",
		Help:      "Number of memberlist membership transitions observed by this speaker. transition=\"gain\" on NodeJoin, transition=\"loss\" on NodeLeave. NodeUpdate is not counted here; use metallb_speaker_memberlist_node_events_total{event_type=\"NodeUpdate\"} for that.",
	}, []string{
		"transition",
		"self_node",
	}),
}

type metrics struct {
	nodeEvents  *prometheus.CounterVec
	members     *prometheus.GaugeVec
	transitions *prometheus.CounterVec
}

func init() {
	prometheus.MustRegister(stats.nodeEvents)
	prometheus.MustRegister(stats.members)
	prometheus.MustRegister(stats.transitions)
}
