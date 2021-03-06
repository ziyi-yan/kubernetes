/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodeSelectorWrapper wraps a NodeSelector inside.
type NodeSelectorWrapper struct{ v1.NodeSelector }

// MakeNodeSelector creates a NodeSelector wrapper.
func MakeNodeSelector() *NodeSelectorWrapper {
	return &NodeSelectorWrapper{v1.NodeSelector{}}
}

// In injects a matchExpression (with an operator IN) as a selectorTerm
// to the inner nodeSelector.
// NOTE: appended selecterTerms are ORed.
func (s *NodeSelectorWrapper) In(key string, vals []string) *NodeSelectorWrapper {
	expression := v1.NodeSelectorRequirement{
		Key:      key,
		Operator: v1.NodeSelectorOpIn,
		Values:   vals,
	}
	selectorTerm := v1.NodeSelectorTerm{}
	selectorTerm.MatchExpressions = append(selectorTerm.MatchExpressions, expression)
	s.NodeSelectorTerms = append(s.NodeSelectorTerms, selectorTerm)
	return s
}

// NotIn injects a matchExpression (with an operator NotIn) as a selectorTerm
// to the inner nodeSelector.
func (s *NodeSelectorWrapper) NotIn(key string, vals []string) *NodeSelectorWrapper {
	expression := v1.NodeSelectorRequirement{
		Key:      key,
		Operator: v1.NodeSelectorOpNotIn,
		Values:   vals,
	}
	selectorTerm := v1.NodeSelectorTerm{}
	selectorTerm.MatchExpressions = append(selectorTerm.MatchExpressions, expression)
	s.NodeSelectorTerms = append(s.NodeSelectorTerms, selectorTerm)
	return s
}

// Obj returns the inner NodeSelector.
func (s *NodeSelectorWrapper) Obj() *v1.NodeSelector {
	return &s.NodeSelector
}

// LabelSelectorWrapper wraps a LabelSelector inside.
type LabelSelectorWrapper struct{ metav1.LabelSelector }

// MakeLabelSelector creates a LabelSelector wrapper.
func MakeLabelSelector() *LabelSelectorWrapper {
	return &LabelSelectorWrapper{metav1.LabelSelector{}}
}

// Label applies a {k,v} pair to the inner LabelSelector.
func (s *LabelSelectorWrapper) Label(k, v string) *LabelSelectorWrapper {
	if s.MatchLabels == nil {
		s.MatchLabels = make(map[string]string)
	}
	s.MatchLabels[k] = v
	return s
}

// In injects a matchExpression (with an operator In) to the inner labelSelector.
func (s *LabelSelectorWrapper) In(key string, vals []string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      key,
		Operator: metav1.LabelSelectorOpIn,
		Values:   vals,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// NotIn injects a matchExpression (with an operator NotIn) to the inner labelSelector.
func (s *LabelSelectorWrapper) NotIn(key string, vals []string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      key,
		Operator: metav1.LabelSelectorOpNotIn,
		Values:   vals,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// Exists injects a matchExpression (with an operator Exists) to the inner labelSelector.
func (s *LabelSelectorWrapper) Exists(k string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      k,
		Operator: metav1.LabelSelectorOpExists,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// NotExist injects a matchExpression (with an operator NotExist) to the inner labelSelector.
func (s *LabelSelectorWrapper) NotExist(k string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      k,
		Operator: metav1.LabelSelectorOpDoesNotExist,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// Obj returns the inner LabelSelector.
func (s *LabelSelectorWrapper) Obj() *metav1.LabelSelector {
	return &s.LabelSelector
}

// PodWrapper wraps a Pod inside.
type PodWrapper struct{ v1.Pod }

// MakePod creates a Pod wrapper.
func MakePod() *PodWrapper {
	return &PodWrapper{v1.Pod{}}
}

// Obj returns the inner Pod.
func (p *PodWrapper) Obj() *v1.Pod {
	return &p.Pod
}

// Name sets `s` as the name of the inner pod.
func (p *PodWrapper) Name(s string) *PodWrapper {
	p.SetName(s)
	return p
}

// Namespace sets `s` as the namespace of the inner pod.
func (p *PodWrapper) Namespace(s string) *PodWrapper {
	p.SetNamespace(s)
	return p
}

// Node sets `s` as the nodeName of the inner pod.
func (p *PodWrapper) Node(s string) *PodWrapper {
	p.Spec.NodeName = s
	return p
}

// NodeSelector sets `m` as the nodeSelector of the inner pod.
func (p *PodWrapper) NodeSelector(m map[string]string) *PodWrapper {
	p.Spec.NodeSelector = m
	return p
}

// NodeAffinityIn creates a HARD node affinity (with the operator In)
// and injects into the innner pod.
func (p *PodWrapper) NodeAffinityIn(key string, vals []string) *PodWrapper {
	if p.Spec.Affinity == nil {
		p.Spec.Affinity = &v1.Affinity{}
	}
	if p.Spec.Affinity.NodeAffinity == nil {
		p.Spec.Affinity.NodeAffinity = &v1.NodeAffinity{}
	}
	nodeSelector := MakeNodeSelector().In(key, vals).Obj()
	p.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = nodeSelector
	return p
}

// NodeAffinityNotIn creates a HARD node affinity (with the operator NotIn)
// and injects into the innner pod.
func (p *PodWrapper) NodeAffinityNotIn(key string, vals []string) *PodWrapper {
	if p.Spec.Affinity == nil {
		p.Spec.Affinity = &v1.Affinity{}
	}
	if p.Spec.Affinity.NodeAffinity == nil {
		p.Spec.Affinity.NodeAffinity = &v1.NodeAffinity{}
	}
	nodeSelector := MakeNodeSelector().NotIn(key, vals).Obj()
	p.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = nodeSelector
	return p
}

// SpreadConstraint constructs a TopologySpreadConstraint object and injects
// into the inner pod.
func (p *PodWrapper) SpreadConstraint(maxSkew int, tpKey string, mode v1.UnsatisfiableConstraintAction, selector *metav1.LabelSelector) *PodWrapper {
	c := v1.TopologySpreadConstraint{
		MaxSkew:           int32(maxSkew),
		TopologyKey:       tpKey,
		WhenUnsatisfiable: mode,
		LabelSelector:     selector,
	}
	p.Spec.TopologySpreadConstraints = append(p.Spec.TopologySpreadConstraints, c)
	return p
}

// Label sets a {k,v} pair to the inner pod.
func (p *PodWrapper) Label(k, v string) *PodWrapper {
	if p.Labels == nil {
		p.Labels = make(map[string]string)
	}
	p.Labels[k] = v
	return p
}

// NodeWrapper wraps a Node inside.
type NodeWrapper struct{ v1.Node }

// MakeNode creates a Node wrapper.
func MakeNode() *NodeWrapper {
	return &NodeWrapper{v1.Node{}}
}

// Obj returns the inner Node.
func (n *NodeWrapper) Obj() *v1.Node {
	return &n.Node
}

// Name sets `s` as the name of the inner pod.
func (n *NodeWrapper) Name(s string) *NodeWrapper {
	n.SetName(s)
	return n
}

// Label applies a {k,v} label pair to the inner node.
func (n *NodeWrapper) Label(k, v string) *NodeWrapper {
	if n.Labels == nil {
		n.Labels = make(map[string]string)
	}
	n.Labels[k] = v
	return n
}
