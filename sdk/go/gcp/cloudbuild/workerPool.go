// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudbuild

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// WorkerPool can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default projects/{{project}}/locations/{{location}}/workerPools/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{location}}/{{name}}
// ```
type WorkerPool struct {
	pulumi.CustomResourceState

	// Output only. Time at which the request to create the `WorkerPool` was received.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. Time at which the request to delete the `WorkerPool` was received.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// User-defined name of the `WorkerPool`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configuration for the `WorkerPool`.
	NetworkConfig WorkerPoolNetworkConfigPtrOutput `pulumi:"networkConfig"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. WorkerPool state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
	State pulumi.StringOutput `pulumi:"state"`
	// Output only. Time at which the request to update the `WorkerPool` was received.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Configuration to be used for a creating workers in the `WorkerPool`.
	WorkerConfig WorkerPoolWorkerConfigOutput `pulumi:"workerConfig"`
}

// NewWorkerPool registers a new resource with the given unique name, arguments, and options.
func NewWorkerPool(ctx *pulumi.Context,
	name string, args *WorkerPoolArgs, opts ...pulumi.ResourceOption) (*WorkerPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource WorkerPool
	err := ctx.RegisterResource("gcp:cloudbuild/workerPool:WorkerPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkerPool gets an existing WorkerPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkerPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkerPoolState, opts ...pulumi.ResourceOption) (*WorkerPool, error) {
	var resource WorkerPool
	err := ctx.ReadResource("gcp:cloudbuild/workerPool:WorkerPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkerPool resources.
type workerPoolState struct {
	// Output only. Time at which the request to create the `WorkerPool` was received.
	CreateTime *string `pulumi:"createTime"`
	// Output only. Time at which the request to delete the `WorkerPool` was received.
	DeleteTime *string `pulumi:"deleteTime"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// User-defined name of the `WorkerPool`.
	Name *string `pulumi:"name"`
	// Network configuration for the `WorkerPool`.
	NetworkConfig *WorkerPoolNetworkConfig `pulumi:"networkConfig"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. WorkerPool state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
	State *string `pulumi:"state"`
	// Output only. Time at which the request to update the `WorkerPool` was received.
	UpdateTime *string `pulumi:"updateTime"`
	// Configuration to be used for a creating workers in the `WorkerPool`.
	WorkerConfig *WorkerPoolWorkerConfig `pulumi:"workerConfig"`
}

type WorkerPoolState struct {
	// Output only. Time at which the request to create the `WorkerPool` was received.
	CreateTime pulumi.StringPtrInput
	// Output only. Time at which the request to delete the `WorkerPool` was received.
	DeleteTime pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// User-defined name of the `WorkerPool`.
	Name pulumi.StringPtrInput
	// Network configuration for the `WorkerPool`.
	NetworkConfig WorkerPoolNetworkConfigPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. WorkerPool state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
	State pulumi.StringPtrInput
	// Output only. Time at which the request to update the `WorkerPool` was received.
	UpdateTime pulumi.StringPtrInput
	// Configuration to be used for a creating workers in the `WorkerPool`.
	WorkerConfig WorkerPoolWorkerConfigPtrInput
}

func (WorkerPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*workerPoolState)(nil)).Elem()
}

type workerPoolArgs struct {
	// The location for the resource
	Location string `pulumi:"location"`
	// User-defined name of the `WorkerPool`.
	Name *string `pulumi:"name"`
	// Network configuration for the `WorkerPool`.
	NetworkConfig *WorkerPoolNetworkConfig `pulumi:"networkConfig"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Configuration to be used for a creating workers in the `WorkerPool`.
	WorkerConfig *WorkerPoolWorkerConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a WorkerPool resource.
type WorkerPoolArgs struct {
	// The location for the resource
	Location pulumi.StringInput
	// User-defined name of the `WorkerPool`.
	Name pulumi.StringPtrInput
	// Network configuration for the `WorkerPool`.
	NetworkConfig WorkerPoolNetworkConfigPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Configuration to be used for a creating workers in the `WorkerPool`.
	WorkerConfig WorkerPoolWorkerConfigPtrInput
}

func (WorkerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workerPoolArgs)(nil)).Elem()
}

type WorkerPoolInput interface {
	pulumi.Input

	ToWorkerPoolOutput() WorkerPoolOutput
	ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput
}

func (*WorkerPool) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil))
}

func (i *WorkerPool) ToWorkerPoolOutput() WorkerPoolOutput {
	return i.ToWorkerPoolOutputWithContext(context.Background())
}

func (i *WorkerPool) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolOutput)
}

func (i *WorkerPool) ToWorkerPoolPtrOutput() WorkerPoolPtrOutput {
	return i.ToWorkerPoolPtrOutputWithContext(context.Background())
}

func (i *WorkerPool) ToWorkerPoolPtrOutputWithContext(ctx context.Context) WorkerPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolPtrOutput)
}

type WorkerPoolPtrInput interface {
	pulumi.Input

	ToWorkerPoolPtrOutput() WorkerPoolPtrOutput
	ToWorkerPoolPtrOutputWithContext(ctx context.Context) WorkerPoolPtrOutput
}

type workerPoolPtrType WorkerPoolArgs

func (*workerPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkerPool)(nil))
}

func (i *workerPoolPtrType) ToWorkerPoolPtrOutput() WorkerPoolPtrOutput {
	return i.ToWorkerPoolPtrOutputWithContext(context.Background())
}

func (i *workerPoolPtrType) ToWorkerPoolPtrOutputWithContext(ctx context.Context) WorkerPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolPtrOutput)
}

// WorkerPoolArrayInput is an input type that accepts WorkerPoolArray and WorkerPoolArrayOutput values.
// You can construct a concrete instance of `WorkerPoolArrayInput` via:
//
//          WorkerPoolArray{ WorkerPoolArgs{...} }
type WorkerPoolArrayInput interface {
	pulumi.Input

	ToWorkerPoolArrayOutput() WorkerPoolArrayOutput
	ToWorkerPoolArrayOutputWithContext(context.Context) WorkerPoolArrayOutput
}

type WorkerPoolArray []WorkerPoolInput

func (WorkerPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WorkerPool)(nil))
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return i.ToWorkerPoolArrayOutputWithContext(context.Background())
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolArrayOutput)
}

// WorkerPoolMapInput is an input type that accepts WorkerPoolMap and WorkerPoolMapOutput values.
// You can construct a concrete instance of `WorkerPoolMapInput` via:
//
//          WorkerPoolMap{ "key": WorkerPoolArgs{...} }
type WorkerPoolMapInput interface {
	pulumi.Input

	ToWorkerPoolMapOutput() WorkerPoolMapOutput
	ToWorkerPoolMapOutputWithContext(context.Context) WorkerPoolMapOutput
}

type WorkerPoolMap map[string]WorkerPoolInput

func (WorkerPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WorkerPool)(nil))
}

func (i WorkerPoolMap) ToWorkerPoolMapOutput() WorkerPoolMapOutput {
	return i.ToWorkerPoolMapOutputWithContext(context.Background())
}

func (i WorkerPoolMap) ToWorkerPoolMapOutputWithContext(ctx context.Context) WorkerPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolMapOutput)
}

type WorkerPoolOutput struct {
	*pulumi.OutputState
}

func (WorkerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil))
}

func (o WorkerPoolOutput) ToWorkerPoolOutput() WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ToWorkerPoolPtrOutput() WorkerPoolPtrOutput {
	return o.ToWorkerPoolPtrOutputWithContext(context.Background())
}

func (o WorkerPoolOutput) ToWorkerPoolPtrOutputWithContext(ctx context.Context) WorkerPoolPtrOutput {
	return o.ApplyT(func(v WorkerPool) *WorkerPool {
		return &v
	}).(WorkerPoolPtrOutput)
}

type WorkerPoolPtrOutput struct {
	*pulumi.OutputState
}

func (WorkerPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkerPool)(nil))
}

func (o WorkerPoolPtrOutput) ToWorkerPoolPtrOutput() WorkerPoolPtrOutput {
	return o
}

func (o WorkerPoolPtrOutput) ToWorkerPoolPtrOutputWithContext(ctx context.Context) WorkerPoolPtrOutput {
	return o
}

type WorkerPoolArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil))
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) Index(i pulumi.IntInput) WorkerPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPool {
		return vs[0].([]WorkerPool)[vs[1].(int)]
	}).(WorkerPoolOutput)
}

type WorkerPoolMapOutput struct{ *pulumi.OutputState }

func (WorkerPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkerPool)(nil))
}

func (o WorkerPoolMapOutput) ToWorkerPoolMapOutput() WorkerPoolMapOutput {
	return o
}

func (o WorkerPoolMapOutput) ToWorkerPoolMapOutputWithContext(ctx context.Context) WorkerPoolMapOutput {
	return o
}

func (o WorkerPoolMapOutput) MapIndex(k pulumi.StringInput) WorkerPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkerPool {
		return vs[0].(map[string]WorkerPool)[vs[1].(string)]
	}).(WorkerPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkerPoolOutput{})
	pulumi.RegisterOutputType(WorkerPoolPtrOutput{})
	pulumi.RegisterOutputType(WorkerPoolArrayOutput{})
	pulumi.RegisterOutputType(WorkerPoolMapOutput{})
}
