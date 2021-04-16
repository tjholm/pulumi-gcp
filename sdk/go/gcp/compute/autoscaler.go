// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an Autoscaler resource.
//
// Autoscalers allow you to automatically scale virtual machine instances in
// managed instance groups according to an autoscaling policy that you
// define.
//
// To get more information about Autoscaler, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
// * How-to Guides
//     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
//
// ## Example Usage
// ### Autoscaler Single Instance
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		debian9, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultInstanceTemplate, err := compute.NewInstanceTemplate(ctx, "defaultInstanceTemplate", &compute.InstanceTemplateArgs{
// 			MachineType:  pulumi.String("e2-medium"),
// 			CanIpForward: pulumi.Bool(false),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("foo"),
// 				pulumi.String("bar"),
// 			},
// 			Disks: compute.InstanceTemplateDiskArray{
// 				&compute.InstanceTemplateDiskArgs{
// 					SourceImage: pulumi.String(debian9.Id),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceTemplateNetworkInterfaceArray{
// 				&compute.InstanceTemplateNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 			Metadata: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			ServiceAccount: &compute.InstanceTemplateServiceAccountArgs{
// 				Scopes: pulumi.StringArray{
// 					pulumi.String("userinfo-email"),
// 					pulumi.String("compute-ro"),
// 					pulumi.String("storage-ro"),
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		defaultTargetPool, err := compute.NewTargetPool(ctx, "defaultTargetPool", nil, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		defaultInstanceGroupManager, err := compute.NewInstanceGroupManager(ctx, "defaultInstanceGroupManager", &compute.InstanceGroupManagerArgs{
// 			Zone: pulumi.String("us-central1-f"),
// 			Versions: compute.InstanceGroupManagerVersionArray{
// 				&compute.InstanceGroupManagerVersionArgs{
// 					InstanceTemplate: defaultInstanceTemplate.ID(),
// 					Name:             pulumi.String("primary"),
// 				},
// 			},
// 			TargetPools: pulumi.StringArray{
// 				defaultTargetPool.ID(),
// 			},
// 			BaseInstanceName: pulumi.String("autoscaler-sample"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewAutoscaler(ctx, "defaultAutoscaler", &compute.AutoscalerArgs{
// 			Zone:   pulumi.String("us-central1-f"),
// 			Target: defaultInstanceGroupManager.ID(),
// 			AutoscalingPolicy: &compute.AutoscalerAutoscalingPolicyArgs{
// 				MaxReplicas:    pulumi.Int(5),
// 				MinReplicas:    pulumi.Int(1),
// 				CooldownPeriod: pulumi.Int(60),
// 				Metrics: compute.AutoscalerAutoscalingPolicyMetricArray{
// 					&compute.AutoscalerAutoscalingPolicyMetricArgs{
// 						Name:                     pulumi.String("pubsub.googleapis.com/subscription/num_undelivered_messages"),
// 						Filter:                   pulumi.String("resource.type = pubsub_subscription AND resource.label.subscription_id = our-subscription"),
// 						SingleInstanceAssignment: pulumi.Float64(65535),
// 					},
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Autoscaler Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		debian9, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		foobarInstanceTemplate, err := compute.NewInstanceTemplate(ctx, "foobarInstanceTemplate", &compute.InstanceTemplateArgs{
// 			MachineType:  pulumi.String("e2-medium"),
// 			CanIpForward: pulumi.Bool(false),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("foo"),
// 				pulumi.String("bar"),
// 			},
// 			Disks: compute.InstanceTemplateDiskArray{
// 				&compute.InstanceTemplateDiskArgs{
// 					SourceImage: pulumi.String(debian9.Id),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceTemplateNetworkInterfaceArray{
// 				&compute.InstanceTemplateNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 			Metadata: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			ServiceAccount: &compute.InstanceTemplateServiceAccountArgs{
// 				Scopes: pulumi.StringArray{
// 					pulumi.String("userinfo-email"),
// 					pulumi.String("compute-ro"),
// 					pulumi.String("storage-ro"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foobarTargetPool, err := compute.NewTargetPool(ctx, "foobarTargetPool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		foobarInstanceGroupManager, err := compute.NewInstanceGroupManager(ctx, "foobarInstanceGroupManager", &compute.InstanceGroupManagerArgs{
// 			Zone: pulumi.String("us-central1-f"),
// 			Versions: compute.InstanceGroupManagerVersionArray{
// 				&compute.InstanceGroupManagerVersionArgs{
// 					InstanceTemplate: foobarInstanceTemplate.ID(),
// 					Name:             pulumi.String("primary"),
// 				},
// 			},
// 			TargetPools: pulumi.StringArray{
// 				foobarTargetPool.ID(),
// 			},
// 			BaseInstanceName: pulumi.String("foobar"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewAutoscaler(ctx, "foobarAutoscaler", &compute.AutoscalerArgs{
// 			Zone:   pulumi.String("us-central1-f"),
// 			Target: foobarInstanceGroupManager.ID(),
// 			AutoscalingPolicy: &compute.AutoscalerAutoscalingPolicyArgs{
// 				MaxReplicas:    pulumi.Int(5),
// 				MinReplicas:    pulumi.Int(1),
// 				CooldownPeriod: pulumi.Int(60),
// 				CpuUtilization: &compute.AutoscalerAutoscalingPolicyCpuUtilizationArgs{
// 					Target: pulumi.Float64(0.5),
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
//
// ## Import
//
// Autoscaler can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/autoscaler:Autoscaler default projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/autoscaler:Autoscaler default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/autoscaler:Autoscaler default {{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/autoscaler:Autoscaler default {{name}}
// ```
type Autoscaler struct {
	pulumi.CustomResourceState

	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy AutoscalerAutoscalingPolicyOutput `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier for this object. Format specified above.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringOutput `pulumi:"target"`
	// URL of the zone where the instance group resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewAutoscaler(ctx *pulumi.Context,
	name string, args *AutoscalerArgs, opts ...pulumi.ResourceOption) (*Autoscaler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoscalingPolicy == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingPolicy'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("gcp:compute/autoscalar:Autoscalar"),
		},
	})
	opts = append(opts, aliases)
	var resource Autoscaler
	err := ctx.RegisterResource("gcp:compute/autoscaler:Autoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscaler gets an existing Autoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalerState, opts ...pulumi.ResourceOption) (*Autoscaler, error) {
	var resource Autoscaler
	err := ctx.ReadResource("gcp:compute/autoscaler:Autoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Autoscaler resources.
type autoscalerState struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy *AutoscalerAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The identifier for this object. Format specified above.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target *string `pulumi:"target"`
	// URL of the zone where the instance group resides.
	Zone *string `pulumi:"zone"`
}

type AutoscalerState struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy AutoscalerAutoscalingPolicyPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The identifier for this object. Format specified above.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringPtrInput
	// URL of the zone where the instance group resides.
	Zone pulumi.StringPtrInput
}

func (AutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalerState)(nil)).Elem()
}

type autoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy AutoscalerAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The identifier for this object. Format specified above.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target string `pulumi:"target"`
	// URL of the zone where the instance group resides.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Autoscaler resource.
type AutoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy AutoscalerAutoscalingPolicyInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The identifier for this object. Format specified above.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringInput
	// URL of the zone where the instance group resides.
	Zone pulumi.StringPtrInput
}

func (AutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalerArgs)(nil)).Elem()
}

type AutoscalerInput interface {
	pulumi.Input

	ToAutoscalerOutput() AutoscalerOutput
	ToAutoscalerOutputWithContext(ctx context.Context) AutoscalerOutput
}

func (*Autoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscaler)(nil))
}

func (i *Autoscaler) ToAutoscalerOutput() AutoscalerOutput {
	return i.ToAutoscalerOutputWithContext(context.Background())
}

func (i *Autoscaler) ToAutoscalerOutputWithContext(ctx context.Context) AutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalerOutput)
}

func (i *Autoscaler) ToAutoscalerPtrOutput() AutoscalerPtrOutput {
	return i.ToAutoscalerPtrOutputWithContext(context.Background())
}

func (i *Autoscaler) ToAutoscalerPtrOutputWithContext(ctx context.Context) AutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalerPtrOutput)
}

type AutoscalerPtrInput interface {
	pulumi.Input

	ToAutoscalerPtrOutput() AutoscalerPtrOutput
	ToAutoscalerPtrOutputWithContext(ctx context.Context) AutoscalerPtrOutput
}

type autoscalerPtrType AutoscalerArgs

func (*autoscalerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscaler)(nil))
}

func (i *autoscalerPtrType) ToAutoscalerPtrOutput() AutoscalerPtrOutput {
	return i.ToAutoscalerPtrOutputWithContext(context.Background())
}

func (i *autoscalerPtrType) ToAutoscalerPtrOutputWithContext(ctx context.Context) AutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalerPtrOutput)
}

// AutoscalerArrayInput is an input type that accepts AutoscalerArray and AutoscalerArrayOutput values.
// You can construct a concrete instance of `AutoscalerArrayInput` via:
//
//          AutoscalerArray{ AutoscalerArgs{...} }
type AutoscalerArrayInput interface {
	pulumi.Input

	ToAutoscalerArrayOutput() AutoscalerArrayOutput
	ToAutoscalerArrayOutputWithContext(context.Context) AutoscalerArrayOutput
}

type AutoscalerArray []AutoscalerInput

func (AutoscalerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Autoscaler)(nil))
}

func (i AutoscalerArray) ToAutoscalerArrayOutput() AutoscalerArrayOutput {
	return i.ToAutoscalerArrayOutputWithContext(context.Background())
}

func (i AutoscalerArray) ToAutoscalerArrayOutputWithContext(ctx context.Context) AutoscalerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalerArrayOutput)
}

// AutoscalerMapInput is an input type that accepts AutoscalerMap and AutoscalerMapOutput values.
// You can construct a concrete instance of `AutoscalerMapInput` via:
//
//          AutoscalerMap{ "key": AutoscalerArgs{...} }
type AutoscalerMapInput interface {
	pulumi.Input

	ToAutoscalerMapOutput() AutoscalerMapOutput
	ToAutoscalerMapOutputWithContext(context.Context) AutoscalerMapOutput
}

type AutoscalerMap map[string]AutoscalerInput

func (AutoscalerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Autoscaler)(nil))
}

func (i AutoscalerMap) ToAutoscalerMapOutput() AutoscalerMapOutput {
	return i.ToAutoscalerMapOutputWithContext(context.Background())
}

func (i AutoscalerMap) ToAutoscalerMapOutputWithContext(ctx context.Context) AutoscalerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalerMapOutput)
}

type AutoscalerOutput struct {
	*pulumi.OutputState
}

func (AutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscaler)(nil))
}

func (o AutoscalerOutput) ToAutoscalerOutput() AutoscalerOutput {
	return o
}

func (o AutoscalerOutput) ToAutoscalerOutputWithContext(ctx context.Context) AutoscalerOutput {
	return o
}

func (o AutoscalerOutput) ToAutoscalerPtrOutput() AutoscalerPtrOutput {
	return o.ToAutoscalerPtrOutputWithContext(context.Background())
}

func (o AutoscalerOutput) ToAutoscalerPtrOutputWithContext(ctx context.Context) AutoscalerPtrOutput {
	return o.ApplyT(func(v Autoscaler) *Autoscaler {
		return &v
	}).(AutoscalerPtrOutput)
}

type AutoscalerPtrOutput struct {
	*pulumi.OutputState
}

func (AutoscalerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscaler)(nil))
}

func (o AutoscalerPtrOutput) ToAutoscalerPtrOutput() AutoscalerPtrOutput {
	return o
}

func (o AutoscalerPtrOutput) ToAutoscalerPtrOutputWithContext(ctx context.Context) AutoscalerPtrOutput {
	return o
}

type AutoscalerArrayOutput struct{ *pulumi.OutputState }

func (AutoscalerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Autoscaler)(nil))
}

func (o AutoscalerArrayOutput) ToAutoscalerArrayOutput() AutoscalerArrayOutput {
	return o
}

func (o AutoscalerArrayOutput) ToAutoscalerArrayOutputWithContext(ctx context.Context) AutoscalerArrayOutput {
	return o
}

func (o AutoscalerArrayOutput) Index(i pulumi.IntInput) AutoscalerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Autoscaler {
		return vs[0].([]Autoscaler)[vs[1].(int)]
	}).(AutoscalerOutput)
}

type AutoscalerMapOutput struct{ *pulumi.OutputState }

func (AutoscalerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Autoscaler)(nil))
}

func (o AutoscalerMapOutput) ToAutoscalerMapOutput() AutoscalerMapOutput {
	return o
}

func (o AutoscalerMapOutput) ToAutoscalerMapOutputWithContext(ctx context.Context) AutoscalerMapOutput {
	return o
}

func (o AutoscalerMapOutput) MapIndex(k pulumi.StringInput) AutoscalerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Autoscaler {
		return vs[0].(map[string]Autoscaler)[vs[1].(string)]
	}).(AutoscalerOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoscalerOutput{})
	pulumi.RegisterOutputType(AutoscalerPtrOutput{})
	pulumi.RegisterOutputType(AutoscalerArrayOutput{})
	pulumi.RegisterOutputType(AutoscalerMapOutput{})
}
