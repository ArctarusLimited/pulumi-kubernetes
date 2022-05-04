// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a "flow distinguisher".
type FlowSchema struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec FlowSchemaSpecPtrOutput `pulumi:"spec"`
	// `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status FlowSchemaStatusPtrOutput `pulumi:"status"`
}

// NewFlowSchema registers a new resource with the given unique name, arguments, and options.
func NewFlowSchema(ctx *pulumi.Context,
	name string, args *FlowSchemaArgs, opts ...pulumi.ResourceOption) (*FlowSchema, error) {
	if args == nil {
		args = &FlowSchemaArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("flowcontrol.apiserver.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("FlowSchema")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1beta1:FlowSchema"),
		},
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:FlowSchema"),
		},
	})
	opts = append(opts, aliases)
	var resource FlowSchema
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowSchema gets an existing FlowSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowSchemaState, opts ...pulumi.ResourceOption) (*FlowSchema, error) {
	var resource FlowSchema
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowSchema resources.
type flowSchemaState struct {
}

type FlowSchemaState struct {
}

func (FlowSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaState)(nil)).Elem()
}

type flowSchemaArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *FlowSchemaSpec `pulumi:"spec"`
}

// The set of arguments for constructing a FlowSchema resource.
type FlowSchemaArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec FlowSchemaSpecPtrInput
}

func (FlowSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaArgs)(nil)).Elem()
}

type FlowSchemaInput interface {
	pulumi.Input

	ToFlowSchemaOutput() FlowSchemaOutput
	ToFlowSchemaOutputWithContext(ctx context.Context) FlowSchemaOutput
}

func (*FlowSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSchema)(nil)).Elem()
}

func (i *FlowSchema) ToFlowSchemaOutput() FlowSchemaOutput {
	return i.ToFlowSchemaOutputWithContext(context.Background())
}

func (i *FlowSchema) ToFlowSchemaOutputWithContext(ctx context.Context) FlowSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaOutput)
}

// FlowSchemaArrayInput is an input type that accepts FlowSchemaArray and FlowSchemaArrayOutput values.
// You can construct a concrete instance of `FlowSchemaArrayInput` via:
//
//          FlowSchemaArray{ FlowSchemaArgs{...} }
type FlowSchemaArrayInput interface {
	pulumi.Input

	ToFlowSchemaArrayOutput() FlowSchemaArrayOutput
	ToFlowSchemaArrayOutputWithContext(context.Context) FlowSchemaArrayOutput
}

type FlowSchemaArray []FlowSchemaInput

func (FlowSchemaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowSchema)(nil)).Elem()
}

func (i FlowSchemaArray) ToFlowSchemaArrayOutput() FlowSchemaArrayOutput {
	return i.ToFlowSchemaArrayOutputWithContext(context.Background())
}

func (i FlowSchemaArray) ToFlowSchemaArrayOutputWithContext(ctx context.Context) FlowSchemaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaArrayOutput)
}

// FlowSchemaMapInput is an input type that accepts FlowSchemaMap and FlowSchemaMapOutput values.
// You can construct a concrete instance of `FlowSchemaMapInput` via:
//
//          FlowSchemaMap{ "key": FlowSchemaArgs{...} }
type FlowSchemaMapInput interface {
	pulumi.Input

	ToFlowSchemaMapOutput() FlowSchemaMapOutput
	ToFlowSchemaMapOutputWithContext(context.Context) FlowSchemaMapOutput
}

type FlowSchemaMap map[string]FlowSchemaInput

func (FlowSchemaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowSchema)(nil)).Elem()
}

func (i FlowSchemaMap) ToFlowSchemaMapOutput() FlowSchemaMapOutput {
	return i.ToFlowSchemaMapOutputWithContext(context.Background())
}

func (i FlowSchemaMap) ToFlowSchemaMapOutputWithContext(ctx context.Context) FlowSchemaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaMapOutput)
}

type FlowSchemaOutput struct{ *pulumi.OutputState }

func (FlowSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSchema)(nil)).Elem()
}

func (o FlowSchemaOutput) ToFlowSchemaOutput() FlowSchemaOutput {
	return o
}

func (o FlowSchemaOutput) ToFlowSchemaOutputWithContext(ctx context.Context) FlowSchemaOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o FlowSchemaOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSchema) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o FlowSchemaOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSchema) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o FlowSchemaOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *FlowSchema) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o FlowSchemaOutput) Spec() FlowSchemaSpecPtrOutput {
	return o.ApplyT(func(v *FlowSchema) FlowSchemaSpecPtrOutput { return v.Spec }).(FlowSchemaSpecPtrOutput)
}

// `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o FlowSchemaOutput) Status() FlowSchemaStatusPtrOutput {
	return o.ApplyT(func(v *FlowSchema) FlowSchemaStatusPtrOutput { return v.Status }).(FlowSchemaStatusPtrOutput)
}

type FlowSchemaArrayOutput struct{ *pulumi.OutputState }

func (FlowSchemaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowSchema)(nil)).Elem()
}

func (o FlowSchemaArrayOutput) ToFlowSchemaArrayOutput() FlowSchemaArrayOutput {
	return o
}

func (o FlowSchemaArrayOutput) ToFlowSchemaArrayOutputWithContext(ctx context.Context) FlowSchemaArrayOutput {
	return o
}

func (o FlowSchemaArrayOutput) Index(i pulumi.IntInput) FlowSchemaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlowSchema {
		return vs[0].([]*FlowSchema)[vs[1].(int)]
	}).(FlowSchemaOutput)
}

type FlowSchemaMapOutput struct{ *pulumi.OutputState }

func (FlowSchemaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowSchema)(nil)).Elem()
}

func (o FlowSchemaMapOutput) ToFlowSchemaMapOutput() FlowSchemaMapOutput {
	return o
}

func (o FlowSchemaMapOutput) ToFlowSchemaMapOutputWithContext(ctx context.Context) FlowSchemaMapOutput {
	return o
}

func (o FlowSchemaMapOutput) MapIndex(k pulumi.StringInput) FlowSchemaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlowSchema {
		return vs[0].(map[string]*FlowSchema)[vs[1].(string)]
	}).(FlowSchemaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaInput)(nil)).Elem(), &FlowSchema{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaArrayInput)(nil)).Elem(), FlowSchemaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaMapInput)(nil)).Elem(), FlowSchemaMap{})
	pulumi.RegisterOutputType(FlowSchemaOutput{})
	pulumi.RegisterOutputType(FlowSchemaArrayOutput{})
	pulumi.RegisterOutputType(FlowSchemaMapOutput{})
}
