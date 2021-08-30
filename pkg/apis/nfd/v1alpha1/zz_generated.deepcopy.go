//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Expressions) DeepCopyInto(out *Expressions) {
	{
		in := &in
		*out = make(Expressions, len(*in))
		for key, val := range *in {
			var outVal *MatchExpression
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MatchExpression)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Expressions.
func (in Expressions) DeepCopy() Expressions {
	if in == nil {
		return nil
	}
	out := new(Expressions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FeatureMatcher) DeepCopyInto(out *FeatureMatcher) {
	{
		in := &in
		*out = make(FeatureMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMatcher.
func (in FeatureMatcher) DeepCopy() FeatureMatcher {
	if in == nil {
		return nil
	}
	out := new(FeatureMatcher)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureMatcherTerm) DeepCopyInto(out *FeatureMatcherTerm) {
	*out = *in
	in.MatchExpressions.DeepCopyInto(&out.MatchExpressions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMatcherTerm.
func (in *FeatureMatcherTerm) DeepCopy() *FeatureMatcherTerm {
	if in == nil {
		return nil
	}
	out := new(FeatureMatcherTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchAnyElem) DeepCopyInto(out *MatchAnyElem) {
	*out = *in
	if in.MatchFeatures != nil {
		in, out := &in.MatchFeatures, &out.MatchFeatures
		*out = make(FeatureMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchAnyElem.
func (in *MatchAnyElem) DeepCopy() *MatchAnyElem {
	if in == nil {
		return nil
	}
	out := new(MatchAnyElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchExpression) DeepCopyInto(out *MatchExpression) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(MatchValue, len(*in))
		copy(*out, *in)
	}
	in.valueRe.DeepCopyInto(&out.valueRe)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpression.
func (in *MatchExpression) DeepCopy() *MatchExpression {
	if in == nil {
		return nil
	}
	out := new(MatchExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchExpressionSet) DeepCopyInto(out *MatchExpressionSet) {
	*out = *in
	if in.Expressions != nil {
		in, out := &in.Expressions, &out.Expressions
		*out = make(Expressions, len(*in))
		for key, val := range *in {
			var outVal *MatchExpression
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MatchExpression)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpressionSet.
func (in *MatchExpressionSet) DeepCopy() *MatchExpressionSet {
	if in == nil {
		return nil
	}
	out := new(MatchExpressionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MatchValue) DeepCopyInto(out *MatchValue) {
	{
		in := &in
		*out = make(MatchValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchValue.
func (in MatchValue) DeepCopy() MatchValue {
	if in == nil {
		return nil
	}
	out := new(MatchValue)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureRule) DeepCopyInto(out *NodeFeatureRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureRule.
func (in *NodeFeatureRule) DeepCopy() *NodeFeatureRule {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureRuleList) DeepCopyInto(out *NodeFeatureRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFeatureRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureRuleList.
func (in *NodeFeatureRuleList) DeepCopy() *NodeFeatureRuleList {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureRuleSpec) DeepCopyInto(out *NodeFeatureRuleSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureRuleSpec.
func (in *NodeFeatureRuleSpec) DeepCopy() *NodeFeatureRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchFeatures != nil {
		in, out := &in.MatchFeatures, &out.MatchFeatures
		*out = make(FeatureMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchAny != nil {
		in, out := &in.MatchAny, &out.MatchAny
		*out = make([]MatchAnyElem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}
