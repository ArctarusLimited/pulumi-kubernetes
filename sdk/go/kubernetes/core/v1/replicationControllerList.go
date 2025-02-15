// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ReplicationControllerList is a collection of replication controllers.
type ReplicationControllerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicationControllerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewReplicationControllerList registers a new resource with the given unique name, arguments, and options.
func NewReplicationControllerList(ctx *pulumi.Context,
	name string, args *ReplicationControllerListArgs, opts ...pulumi.ResourceOption) (*ReplicationControllerList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ReplicationControllerList")
	var resource ReplicationControllerList
	err := ctx.RegisterResource("kubernetes:core/v1:ReplicationControllerList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationControllerList gets an existing ReplicationControllerList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationControllerList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationControllerListState, opts ...pulumi.ResourceOption) (*ReplicationControllerList, error) {
	var resource ReplicationControllerList
	err := ctx.ReadResource("kubernetes:core/v1:ReplicationControllerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationControllerList resources.
type replicationControllerListState struct {
}

type ReplicationControllerListState struct {
}

func (ReplicationControllerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationControllerListState)(nil)).Elem()
}

type replicationControllerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []ReplicationControllerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ReplicationControllerList resource.
type ReplicationControllerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicationControllerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReplicationControllerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationControllerListArgs)(nil)).Elem()
}

type ReplicationControllerListInput interface {
	pulumi.Input

	ToReplicationControllerListOutput() ReplicationControllerListOutput
	ToReplicationControllerListOutputWithContext(ctx context.Context) ReplicationControllerListOutput
}

func (*ReplicationControllerList) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationControllerList)(nil)).Elem()
}

func (i *ReplicationControllerList) ToReplicationControllerListOutput() ReplicationControllerListOutput {
	return i.ToReplicationControllerListOutputWithContext(context.Background())
}

func (i *ReplicationControllerList) ToReplicationControllerListOutputWithContext(ctx context.Context) ReplicationControllerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerListOutput)
}

// ReplicationControllerListArrayInput is an input type that accepts ReplicationControllerListArray and ReplicationControllerListArrayOutput values.
// You can construct a concrete instance of `ReplicationControllerListArrayInput` via:
//
//          ReplicationControllerListArray{ ReplicationControllerListArgs{...} }
type ReplicationControllerListArrayInput interface {
	pulumi.Input

	ToReplicationControllerListArrayOutput() ReplicationControllerListArrayOutput
	ToReplicationControllerListArrayOutputWithContext(context.Context) ReplicationControllerListArrayOutput
}

type ReplicationControllerListArray []ReplicationControllerListInput

func (ReplicationControllerListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationControllerList)(nil)).Elem()
}

func (i ReplicationControllerListArray) ToReplicationControllerListArrayOutput() ReplicationControllerListArrayOutput {
	return i.ToReplicationControllerListArrayOutputWithContext(context.Background())
}

func (i ReplicationControllerListArray) ToReplicationControllerListArrayOutputWithContext(ctx context.Context) ReplicationControllerListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerListArrayOutput)
}

// ReplicationControllerListMapInput is an input type that accepts ReplicationControllerListMap and ReplicationControllerListMapOutput values.
// You can construct a concrete instance of `ReplicationControllerListMapInput` via:
//
//          ReplicationControllerListMap{ "key": ReplicationControllerListArgs{...} }
type ReplicationControllerListMapInput interface {
	pulumi.Input

	ToReplicationControllerListMapOutput() ReplicationControllerListMapOutput
	ToReplicationControllerListMapOutputWithContext(context.Context) ReplicationControllerListMapOutput
}

type ReplicationControllerListMap map[string]ReplicationControllerListInput

func (ReplicationControllerListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationControllerList)(nil)).Elem()
}

func (i ReplicationControllerListMap) ToReplicationControllerListMapOutput() ReplicationControllerListMapOutput {
	return i.ToReplicationControllerListMapOutputWithContext(context.Background())
}

func (i ReplicationControllerListMap) ToReplicationControllerListMapOutputWithContext(ctx context.Context) ReplicationControllerListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerListMapOutput)
}

type ReplicationControllerListOutput struct{ *pulumi.OutputState }

func (ReplicationControllerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationControllerList)(nil)).Elem()
}

func (o ReplicationControllerListOutput) ToReplicationControllerListOutput() ReplicationControllerListOutput {
	return o
}

func (o ReplicationControllerListOutput) ToReplicationControllerListOutputWithContext(ctx context.Context) ReplicationControllerListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ReplicationControllerListOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerList) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
func (o ReplicationControllerListOutput) Items() ReplicationControllerTypeArrayOutput {
	return o.ApplyT(func(v *ReplicationControllerList) ReplicationControllerTypeArrayOutput { return v.Items }).(ReplicationControllerTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReplicationControllerListOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerList) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReplicationControllerListOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerList) metav1.ListMetaPtrOutput { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

type ReplicationControllerListArrayOutput struct{ *pulumi.OutputState }

func (ReplicationControllerListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationControllerList)(nil)).Elem()
}

func (o ReplicationControllerListArrayOutput) ToReplicationControllerListArrayOutput() ReplicationControllerListArrayOutput {
	return o
}

func (o ReplicationControllerListArrayOutput) ToReplicationControllerListArrayOutputWithContext(ctx context.Context) ReplicationControllerListArrayOutput {
	return o
}

func (o ReplicationControllerListArrayOutput) Index(i pulumi.IntInput) ReplicationControllerListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationControllerList {
		return vs[0].([]*ReplicationControllerList)[vs[1].(int)]
	}).(ReplicationControllerListOutput)
}

type ReplicationControllerListMapOutput struct{ *pulumi.OutputState }

func (ReplicationControllerListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationControllerList)(nil)).Elem()
}

func (o ReplicationControllerListMapOutput) ToReplicationControllerListMapOutput() ReplicationControllerListMapOutput {
	return o
}

func (o ReplicationControllerListMapOutput) ToReplicationControllerListMapOutputWithContext(ctx context.Context) ReplicationControllerListMapOutput {
	return o
}

func (o ReplicationControllerListMapOutput) MapIndex(k pulumi.StringInput) ReplicationControllerListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationControllerList {
		return vs[0].(map[string]*ReplicationControllerList)[vs[1].(string)]
	}).(ReplicationControllerListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationControllerListInput)(nil)).Elem(), &ReplicationControllerList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationControllerListArrayInput)(nil)).Elem(), ReplicationControllerListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationControllerListMapInput)(nil)).Elem(), ReplicationControllerListMap{})
	pulumi.RegisterOutputType(ReplicationControllerListOutput{})
	pulumi.RegisterOutputType(ReplicationControllerListArrayOutput{})
	pulumi.RegisterOutputType(ReplicationControllerListMapOutput{})
}
