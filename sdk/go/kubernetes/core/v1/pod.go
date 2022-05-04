// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.
//
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
//
// 1. The Pod is scheduled ("PodScheduled"" '.status.condition' is true).
// 2. The Pod is initialized ("Initialized" '.status.condition' is true).
// 3. The Pod is ready ("Ready" '.status.condition' is true) and the '.status.phase' is
//    set to "Running".
//    Or (for Jobs): The Pod succeeded ('.status.phase' set to "Succeeded").
//
// If the Pod has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
//
// ## Example Usage
// ### Create a Pod with auto-naming
// ```go
// package main
//
// import (
// 	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
// 	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := corev1.NewPod(ctx, "nginxPod", &corev1.PodArgs{
// 			Spec: &corev1.PodSpecArgs{
// 				Containers: corev1.ContainerArray{
// 					&corev1.ContainerArgs{
// 						Name:  pulumi.String("nginx"),
// 						Image: pulumi.String("nginx:1.14.2"),
// 						Ports: corev1.ContainerPortArray{
// 							&corev1.ContainerPortArgs{
// 								ContainerPort: pulumi.Int(80),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a Pod with a user-specified name
// ```go
// package main
//
// import (
// 	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
// 	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := corev1.NewPod(ctx, "nginxPod", &corev1.PodArgs{
// 			Metadata: &metav1.ObjectMetaArgs{
// 				Name: pulumi.String("nginx"),
// 			},
// 			Spec: &corev1.PodSpecArgs{
// 				Containers: corev1.ContainerArray{
// 					&corev1.ContainerArgs{
// 						Name:  pulumi.String("nginx"),
// 						Image: pulumi.String("nginx:1.14.2"),
// 						Ports: corev1.ContainerPortArray{
// 							&corev1.ContainerPortArgs{
// 								ContainerPort: pulumi.Int(80),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// {% /examples %}}
type Pod struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PodSpecPtrOutput `pulumi:"spec"`
	// Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status PodStatusPtrOutput `pulumi:"status"`
}

// NewPod registers a new resource with the given unique name, arguments, and options.
func NewPod(ctx *pulumi.Context,
	name string, args *PodArgs, opts ...pulumi.ResourceOption) (*Pod, error) {
	if args == nil {
		args = &PodArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Pod")
	var resource Pod
	err := ctx.RegisterResource("kubernetes:core/v1:Pod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPod gets an existing Pod resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodState, opts ...pulumi.ResourceOption) (*Pod, error) {
	var resource Pod
	err := ctx.ReadResource("kubernetes:core/v1:Pod", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pod resources.
type podState struct {
}

type PodState struct {
}

func (PodState) ElementType() reflect.Type {
	return reflect.TypeOf((*podState)(nil)).Elem()
}

type podArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PodSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Pod resource.
type PodArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PodSpecPtrInput
}

func (PodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podArgs)(nil)).Elem()
}

type PodInput interface {
	pulumi.Input

	ToPodOutput() PodOutput
	ToPodOutputWithContext(ctx context.Context) PodOutput
}

func (*Pod) ElementType() reflect.Type {
	return reflect.TypeOf((**Pod)(nil)).Elem()
}

func (i *Pod) ToPodOutput() PodOutput {
	return i.ToPodOutputWithContext(context.Background())
}

func (i *Pod) ToPodOutputWithContext(ctx context.Context) PodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodOutput)
}

// PodArrayInput is an input type that accepts PodArray and PodArrayOutput values.
// You can construct a concrete instance of `PodArrayInput` via:
//
//          PodArray{ PodArgs{...} }
type PodArrayInput interface {
	pulumi.Input

	ToPodArrayOutput() PodArrayOutput
	ToPodArrayOutputWithContext(context.Context) PodArrayOutput
}

type PodArray []PodInput

func (PodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pod)(nil)).Elem()
}

func (i PodArray) ToPodArrayOutput() PodArrayOutput {
	return i.ToPodArrayOutputWithContext(context.Background())
}

func (i PodArray) ToPodArrayOutputWithContext(ctx context.Context) PodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodArrayOutput)
}

// PodMapInput is an input type that accepts PodMap and PodMapOutput values.
// You can construct a concrete instance of `PodMapInput` via:
//
//          PodMap{ "key": PodArgs{...} }
type PodMapInput interface {
	pulumi.Input

	ToPodMapOutput() PodMapOutput
	ToPodMapOutputWithContext(context.Context) PodMapOutput
}

type PodMap map[string]PodInput

func (PodMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pod)(nil)).Elem()
}

func (i PodMap) ToPodMapOutput() PodMapOutput {
	return i.ToPodMapOutputWithContext(context.Background())
}

func (i PodMap) ToPodMapOutputWithContext(ctx context.Context) PodMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodMapOutput)
}

type PodOutput struct{ *pulumi.OutputState }

func (PodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pod)(nil)).Elem()
}

func (o PodOutput) ToPodOutput() PodOutput {
	return o
}

func (o PodOutput) ToPodOutputWithContext(ctx context.Context) PodOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pod) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pod) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PodOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Pod) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o PodOutput) Spec() PodSpecPtrOutput {
	return o.ApplyT(func(v *Pod) PodSpecPtrOutput { return v.Spec }).(PodSpecPtrOutput)
}

// Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o PodOutput) Status() PodStatusPtrOutput {
	return o.ApplyT(func(v *Pod) PodStatusPtrOutput { return v.Status }).(PodStatusPtrOutput)
}

type PodArrayOutput struct{ *pulumi.OutputState }

func (PodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pod)(nil)).Elem()
}

func (o PodArrayOutput) ToPodArrayOutput() PodArrayOutput {
	return o
}

func (o PodArrayOutput) ToPodArrayOutputWithContext(ctx context.Context) PodArrayOutput {
	return o
}

func (o PodArrayOutput) Index(i pulumi.IntInput) PodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pod {
		return vs[0].([]*Pod)[vs[1].(int)]
	}).(PodOutput)
}

type PodMapOutput struct{ *pulumi.OutputState }

func (PodMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pod)(nil)).Elem()
}

func (o PodMapOutput) ToPodMapOutput() PodMapOutput {
	return o
}

func (o PodMapOutput) ToPodMapOutputWithContext(ctx context.Context) PodMapOutput {
	return o
}

func (o PodMapOutput) MapIndex(k pulumi.StringInput) PodOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pod {
		return vs[0].(map[string]*Pod)[vs[1].(string)]
	}).(PodOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodInput)(nil)).Elem(), &Pod{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodArrayInput)(nil)).Elem(), PodArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodMapInput)(nil)).Elem(), PodMap{})
	pulumi.RegisterOutputType(PodOutput{})
	pulumi.RegisterOutputType(PodArrayOutput{})
	pulumi.RegisterOutputType(PodMapOutput{})
}
