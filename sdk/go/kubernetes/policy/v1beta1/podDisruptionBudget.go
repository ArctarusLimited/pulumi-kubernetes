// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudget struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPtrOutput `pulumi:"spec"`
	// Most recently observed status of the PodDisruptionBudget.
	Status PodDisruptionBudgetStatusPtrOutput `pulumi:"status"`
}

// NewPodDisruptionBudget registers a new resource with the given unique name, arguments, and options.
func NewPodDisruptionBudget(ctx *pulumi.Context,
	name string, args *PodDisruptionBudgetArgs, opts ...pulumi.ResourceOption) (*PodDisruptionBudget, error) {
	if args == nil {
		args = &PodDisruptionBudgetArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("policy/v1beta1")
	args.Kind = pulumi.StringPtr("PodDisruptionBudget")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:policy/v1:PodDisruptionBudget"),
		},
	})
	opts = append(opts, aliases)
	var resource PodDisruptionBudget
	err := ctx.RegisterResource("kubernetes:policy/v1beta1:PodDisruptionBudget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodDisruptionBudget gets an existing PodDisruptionBudget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodDisruptionBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodDisruptionBudgetState, opts ...pulumi.ResourceOption) (*PodDisruptionBudget, error) {
	var resource PodDisruptionBudget
	err := ctx.ReadResource("kubernetes:policy/v1beta1:PodDisruptionBudget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodDisruptionBudget resources.
type podDisruptionBudgetState struct {
}

type PodDisruptionBudgetState struct {
}

func (PodDisruptionBudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetState)(nil)).Elem()
}

type podDisruptionBudgetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpec `pulumi:"spec"`
}

// The set of arguments for constructing a PodDisruptionBudget resource.
type PodDisruptionBudgetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPtrInput
}

func (PodDisruptionBudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetArgs)(nil)).Elem()
}

type PodDisruptionBudgetInput interface {
	pulumi.Input

	ToPodDisruptionBudgetOutput() PodDisruptionBudgetOutput
	ToPodDisruptionBudgetOutputWithContext(ctx context.Context) PodDisruptionBudgetOutput
}

func (*PodDisruptionBudget) ElementType() reflect.Type {
	return reflect.TypeOf((**PodDisruptionBudget)(nil)).Elem()
}

func (i *PodDisruptionBudget) ToPodDisruptionBudgetOutput() PodDisruptionBudgetOutput {
	return i.ToPodDisruptionBudgetOutputWithContext(context.Background())
}

func (i *PodDisruptionBudget) ToPodDisruptionBudgetOutputWithContext(ctx context.Context) PodDisruptionBudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetOutput)
}

// PodDisruptionBudgetArrayInput is an input type that accepts PodDisruptionBudgetArray and PodDisruptionBudgetArrayOutput values.
// You can construct a concrete instance of `PodDisruptionBudgetArrayInput` via:
//
//          PodDisruptionBudgetArray{ PodDisruptionBudgetArgs{...} }
type PodDisruptionBudgetArrayInput interface {
	pulumi.Input

	ToPodDisruptionBudgetArrayOutput() PodDisruptionBudgetArrayOutput
	ToPodDisruptionBudgetArrayOutputWithContext(context.Context) PodDisruptionBudgetArrayOutput
}

type PodDisruptionBudgetArray []PodDisruptionBudgetInput

func (PodDisruptionBudgetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodDisruptionBudget)(nil)).Elem()
}

func (i PodDisruptionBudgetArray) ToPodDisruptionBudgetArrayOutput() PodDisruptionBudgetArrayOutput {
	return i.ToPodDisruptionBudgetArrayOutputWithContext(context.Background())
}

func (i PodDisruptionBudgetArray) ToPodDisruptionBudgetArrayOutputWithContext(ctx context.Context) PodDisruptionBudgetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetArrayOutput)
}

// PodDisruptionBudgetMapInput is an input type that accepts PodDisruptionBudgetMap and PodDisruptionBudgetMapOutput values.
// You can construct a concrete instance of `PodDisruptionBudgetMapInput` via:
//
//          PodDisruptionBudgetMap{ "key": PodDisruptionBudgetArgs{...} }
type PodDisruptionBudgetMapInput interface {
	pulumi.Input

	ToPodDisruptionBudgetMapOutput() PodDisruptionBudgetMapOutput
	ToPodDisruptionBudgetMapOutputWithContext(context.Context) PodDisruptionBudgetMapOutput
}

type PodDisruptionBudgetMap map[string]PodDisruptionBudgetInput

func (PodDisruptionBudgetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodDisruptionBudget)(nil)).Elem()
}

func (i PodDisruptionBudgetMap) ToPodDisruptionBudgetMapOutput() PodDisruptionBudgetMapOutput {
	return i.ToPodDisruptionBudgetMapOutputWithContext(context.Background())
}

func (i PodDisruptionBudgetMap) ToPodDisruptionBudgetMapOutputWithContext(ctx context.Context) PodDisruptionBudgetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetMapOutput)
}

type PodDisruptionBudgetOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodDisruptionBudget)(nil)).Elem()
}

func (o PodDisruptionBudgetOutput) ToPodDisruptionBudgetOutput() PodDisruptionBudgetOutput {
	return o
}

func (o PodDisruptionBudgetOutput) ToPodDisruptionBudgetOutputWithContext(ctx context.Context) PodDisruptionBudgetOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodDisruptionBudgetOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudget) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodDisruptionBudgetOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudget) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PodDisruptionBudgetOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudget) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Specification of the desired behavior of the PodDisruptionBudget.
func (o PodDisruptionBudgetOutput) Spec() PodDisruptionBudgetSpecPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudget) PodDisruptionBudgetSpecPtrOutput { return v.Spec }).(PodDisruptionBudgetSpecPtrOutput)
}

// Most recently observed status of the PodDisruptionBudget.
func (o PodDisruptionBudgetOutput) Status() PodDisruptionBudgetStatusPtrOutput {
	return o.ApplyT(func(v *PodDisruptionBudget) PodDisruptionBudgetStatusPtrOutput { return v.Status }).(PodDisruptionBudgetStatusPtrOutput)
}

type PodDisruptionBudgetArrayOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodDisruptionBudget)(nil)).Elem()
}

func (o PodDisruptionBudgetArrayOutput) ToPodDisruptionBudgetArrayOutput() PodDisruptionBudgetArrayOutput {
	return o
}

func (o PodDisruptionBudgetArrayOutput) ToPodDisruptionBudgetArrayOutputWithContext(ctx context.Context) PodDisruptionBudgetArrayOutput {
	return o
}

func (o PodDisruptionBudgetArrayOutput) Index(i pulumi.IntInput) PodDisruptionBudgetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodDisruptionBudget {
		return vs[0].([]*PodDisruptionBudget)[vs[1].(int)]
	}).(PodDisruptionBudgetOutput)
}

type PodDisruptionBudgetMapOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodDisruptionBudget)(nil)).Elem()
}

func (o PodDisruptionBudgetMapOutput) ToPodDisruptionBudgetMapOutput() PodDisruptionBudgetMapOutput {
	return o
}

func (o PodDisruptionBudgetMapOutput) ToPodDisruptionBudgetMapOutputWithContext(ctx context.Context) PodDisruptionBudgetMapOutput {
	return o
}

func (o PodDisruptionBudgetMapOutput) MapIndex(k pulumi.StringInput) PodDisruptionBudgetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodDisruptionBudget {
		return vs[0].(map[string]*PodDisruptionBudget)[vs[1].(string)]
	}).(PodDisruptionBudgetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodDisruptionBudgetInput)(nil)).Elem(), &PodDisruptionBudget{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodDisruptionBudgetArrayInput)(nil)).Elem(), PodDisruptionBudgetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodDisruptionBudgetMapInput)(nil)).Elem(), PodDisruptionBudgetMap{})
	pulumi.RegisterOutputType(PodDisruptionBudgetOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetArrayOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetMapOutput{})
}
