// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Status is a return value for calls that don't return other objects.
type Status struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code pulumi.IntPtrOutput `pulumi:"code"`
	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	Details StatusDetailsPtrOutput `pulumi:"details"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// A human-readable description of the status of this operation.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata ListMetaPtrOutput `pulumi:"metadata"`
	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewStatus registers a new resource with the given unique name, arguments, and options.
func NewStatus(ctx *pulumi.Context,
	name string, args *StatusArgs, opts ...pulumi.ResourceOption) (*Status, error) {
	if args == nil {
		args = &StatusArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Status")
	var resource Status
	err := ctx.RegisterResource("kubernetes:meta/v1:Status", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatus gets an existing Status resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatusState, opts ...pulumi.ResourceOption) (*Status, error) {
	var resource Status
	err := ctx.ReadResource("kubernetes:meta/v1:Status", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Status resources.
type statusState struct {
}

type StatusState struct {
}

func (StatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*statusState)(nil)).Elem()
}

type statusArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code *int `pulumi:"code"`
	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	Details *StatusDetails `pulumi:"details"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// A human-readable description of the status of this operation.
	Message *string `pulumi:"message"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *ListMeta `pulumi:"metadata"`
	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason *string `pulumi:"reason"`
}

// The set of arguments for constructing a Status resource.
type StatusArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Suggested HTTP return code for this status, 0 if not set.
	Code pulumi.IntPtrInput
	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	Details StatusDetailsPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// A human-readable description of the status of this operation.
	Message pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata ListMetaPtrInput
	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason pulumi.StringPtrInput
}

func (StatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statusArgs)(nil)).Elem()
}

type StatusInput interface {
	pulumi.Input

	ToStatusOutput() StatusOutput
	ToStatusOutputWithContext(ctx context.Context) StatusOutput
}

func (*Status) ElementType() reflect.Type {
	return reflect.TypeOf((**Status)(nil)).Elem()
}

func (i *Status) ToStatusOutput() StatusOutput {
	return i.ToStatusOutputWithContext(context.Background())
}

func (i *Status) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusOutput)
}

// StatusArrayInput is an input type that accepts StatusArray and StatusArrayOutput values.
// You can construct a concrete instance of `StatusArrayInput` via:
//
//          StatusArray{ StatusArgs{...} }
type StatusArrayInput interface {
	pulumi.Input

	ToStatusArrayOutput() StatusArrayOutput
	ToStatusArrayOutputWithContext(context.Context) StatusArrayOutput
}

type StatusArray []StatusInput

func (StatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Status)(nil)).Elem()
}

func (i StatusArray) ToStatusArrayOutput() StatusArrayOutput {
	return i.ToStatusArrayOutputWithContext(context.Background())
}

func (i StatusArray) ToStatusArrayOutputWithContext(ctx context.Context) StatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusArrayOutput)
}

// StatusMapInput is an input type that accepts StatusMap and StatusMapOutput values.
// You can construct a concrete instance of `StatusMapInput` via:
//
//          StatusMap{ "key": StatusArgs{...} }
type StatusMapInput interface {
	pulumi.Input

	ToStatusMapOutput() StatusMapOutput
	ToStatusMapOutputWithContext(context.Context) StatusMapOutput
}

type StatusMap map[string]StatusInput

func (StatusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Status)(nil)).Elem()
}

func (i StatusMap) ToStatusMapOutput() StatusMapOutput {
	return i.ToStatusMapOutputWithContext(context.Background())
}

func (i StatusMap) ToStatusMapOutputWithContext(ctx context.Context) StatusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusMapOutput)
}

type StatusOutput struct{ *pulumi.OutputState }

func (StatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Status)(nil)).Elem()
}

func (o StatusOutput) ToStatusOutput() StatusOutput {
	return o
}

func (o StatusOutput) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o StatusOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Status) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Suggested HTTP return code for this status, 0 if not set.
func (o StatusOutput) Code() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Status) pulumi.IntPtrOutput { return v.Code }).(pulumi.IntPtrOutput)
}

// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
func (o StatusOutput) Details() StatusDetailsPtrOutput {
	return o.ApplyT(func(v *Status) StatusDetailsPtrOutput { return v.Details }).(StatusDetailsPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StatusOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Status) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// A human-readable description of the status of this operation.
func (o StatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Status) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StatusOutput) Metadata() ListMetaPtrOutput {
	return o.ApplyT(func(v *Status) ListMetaPtrOutput { return v.Metadata }).(ListMetaPtrOutput)
}

// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
func (o StatusOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Status) pulumi.StringPtrOutput { return v.Reason }).(pulumi.StringPtrOutput)
}

// Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o StatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Status) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

type StatusArrayOutput struct{ *pulumi.OutputState }

func (StatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Status)(nil)).Elem()
}

func (o StatusArrayOutput) ToStatusArrayOutput() StatusArrayOutput {
	return o
}

func (o StatusArrayOutput) ToStatusArrayOutputWithContext(ctx context.Context) StatusArrayOutput {
	return o
}

func (o StatusArrayOutput) Index(i pulumi.IntInput) StatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Status {
		return vs[0].([]*Status)[vs[1].(int)]
	}).(StatusOutput)
}

type StatusMapOutput struct{ *pulumi.OutputState }

func (StatusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Status)(nil)).Elem()
}

func (o StatusMapOutput) ToStatusMapOutput() StatusMapOutput {
	return o
}

func (o StatusMapOutput) ToStatusMapOutputWithContext(ctx context.Context) StatusMapOutput {
	return o
}

func (o StatusMapOutput) MapIndex(k pulumi.StringInput) StatusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Status {
		return vs[0].(map[string]*Status)[vs[1].(string)]
	}).(StatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatusInput)(nil)).Elem(), &Status{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatusArrayInput)(nil)).Elem(), StatusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatusMapInput)(nil)).Elem(), StatusMap{})
	pulumi.RegisterOutputType(StatusOutput{})
	pulumi.RegisterOutputType(StatusArrayOutput{})
	pulumi.RegisterOutputType(StatusMapOutput{})
}
