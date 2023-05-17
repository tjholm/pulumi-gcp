// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Network Services Service Binding Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/networkservices"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicedirectory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultNamespace, err := servicedirectory.NewNamespace(ctx, "defaultNamespace", &servicedirectory.NamespaceArgs{
//				NamespaceId: pulumi.String("my-namespace"),
//				Location:    pulumi.String("us-central1"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			defaultService, err := servicedirectory.NewService(ctx, "defaultService", &servicedirectory.ServiceArgs{
//				ServiceId: pulumi.String("my-service"),
//				Namespace: defaultNamespace.ID(),
//				Metadata: pulumi.StringMap{
//					"stage":  pulumi.String("prod"),
//					"region": pulumi.String("us-central1"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = networkservices.NewServiceBinding(ctx, "defaultServiceBinding", &networkservices.ServiceBindingArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Description: pulumi.String("my description"),
//				Service:     defaultService.ID(),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # ServiceBinding can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:networkservices/serviceBinding:ServiceBinding default projects/{{project}}/locations/global/serviceBindings/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkservices/serviceBinding:ServiceBinding default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkservices/serviceBinding:ServiceBinding default {{name}}
//
// ```
type ServiceBinding struct {
	pulumi.CustomResourceState

	// Time the ServiceBinding was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Set of label tags associated with the ServiceBinding resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the ServiceBinding resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The full Service Directory Service name of the format
	// projects/*/locations/*/namespaces/*/services/*
	Service pulumi.StringOutput `pulumi:"service"`
	// Time the ServiceBinding was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewServiceBinding registers a new resource with the given unique name, arguments, and options.
func NewServiceBinding(ctx *pulumi.Context,
	name string, args *ServiceBindingArgs, opts ...pulumi.ResourceOption) (*ServiceBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource ServiceBinding
	err := ctx.RegisterResource("gcp:networkservices/serviceBinding:ServiceBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceBinding gets an existing ServiceBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceBindingState, opts ...pulumi.ResourceOption) (*ServiceBinding, error) {
	var resource ServiceBinding
	err := ctx.ReadResource("gcp:networkservices/serviceBinding:ServiceBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceBinding resources.
type serviceBindingState struct {
	// Time the ServiceBinding was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the ServiceBinding resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the ServiceBinding resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The full Service Directory Service name of the format
	// projects/*/locations/*/namespaces/*/services/*
	Service *string `pulumi:"service"`
	// Time the ServiceBinding was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type ServiceBindingState struct {
	// Time the ServiceBinding was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the ServiceBinding resource.
	Labels pulumi.StringMapInput
	// Name of the ServiceBinding resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The full Service Directory Service name of the format
	// projects/*/locations/*/namespaces/*/services/*
	Service pulumi.StringPtrInput
	// Time the ServiceBinding was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (ServiceBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceBindingState)(nil)).Elem()
}

type serviceBindingArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the ServiceBinding resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the ServiceBinding resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The full Service Directory Service name of the format
	// projects/*/locations/*/namespaces/*/services/*
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a ServiceBinding resource.
type ServiceBindingArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the ServiceBinding resource.
	Labels pulumi.StringMapInput
	// Name of the ServiceBinding resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The full Service Directory Service name of the format
	// projects/*/locations/*/namespaces/*/services/*
	Service pulumi.StringInput
}

func (ServiceBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceBindingArgs)(nil)).Elem()
}

type ServiceBindingInput interface {
	pulumi.Input

	ToServiceBindingOutput() ServiceBindingOutput
	ToServiceBindingOutputWithContext(ctx context.Context) ServiceBindingOutput
}

func (*ServiceBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBinding)(nil)).Elem()
}

func (i *ServiceBinding) ToServiceBindingOutput() ServiceBindingOutput {
	return i.ToServiceBindingOutputWithContext(context.Background())
}

func (i *ServiceBinding) ToServiceBindingOutputWithContext(ctx context.Context) ServiceBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBindingOutput)
}

// ServiceBindingArrayInput is an input type that accepts ServiceBindingArray and ServiceBindingArrayOutput values.
// You can construct a concrete instance of `ServiceBindingArrayInput` via:
//
//	ServiceBindingArray{ ServiceBindingArgs{...} }
type ServiceBindingArrayInput interface {
	pulumi.Input

	ToServiceBindingArrayOutput() ServiceBindingArrayOutput
	ToServiceBindingArrayOutputWithContext(context.Context) ServiceBindingArrayOutput
}

type ServiceBindingArray []ServiceBindingInput

func (ServiceBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceBinding)(nil)).Elem()
}

func (i ServiceBindingArray) ToServiceBindingArrayOutput() ServiceBindingArrayOutput {
	return i.ToServiceBindingArrayOutputWithContext(context.Background())
}

func (i ServiceBindingArray) ToServiceBindingArrayOutputWithContext(ctx context.Context) ServiceBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBindingArrayOutput)
}

// ServiceBindingMapInput is an input type that accepts ServiceBindingMap and ServiceBindingMapOutput values.
// You can construct a concrete instance of `ServiceBindingMapInput` via:
//
//	ServiceBindingMap{ "key": ServiceBindingArgs{...} }
type ServiceBindingMapInput interface {
	pulumi.Input

	ToServiceBindingMapOutput() ServiceBindingMapOutput
	ToServiceBindingMapOutputWithContext(context.Context) ServiceBindingMapOutput
}

type ServiceBindingMap map[string]ServiceBindingInput

func (ServiceBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceBinding)(nil)).Elem()
}

func (i ServiceBindingMap) ToServiceBindingMapOutput() ServiceBindingMapOutput {
	return i.ToServiceBindingMapOutputWithContext(context.Background())
}

func (i ServiceBindingMap) ToServiceBindingMapOutputWithContext(ctx context.Context) ServiceBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBindingMapOutput)
}

type ServiceBindingOutput struct{ *pulumi.OutputState }

func (ServiceBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBinding)(nil)).Elem()
}

func (o ServiceBindingOutput) ToServiceBindingOutput() ServiceBindingOutput {
	return o
}

func (o ServiceBindingOutput) ToServiceBindingOutputWithContext(ctx context.Context) ServiceBindingOutput {
	return o
}

// Time the ServiceBinding was created in UTC.
func (o ServiceBindingOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o ServiceBindingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Set of label tags associated with the ServiceBinding resource.
func (o ServiceBindingOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the ServiceBinding resource.
func (o ServiceBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ServiceBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The full Service Directory Service name of the format
// projects/*/locations/*/namespaces/*/services/*
func (o ServiceBindingOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// Time the ServiceBinding was updated in UTC.
func (o ServiceBindingOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceBinding) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ServiceBindingArrayOutput struct{ *pulumi.OutputState }

func (ServiceBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceBinding)(nil)).Elem()
}

func (o ServiceBindingArrayOutput) ToServiceBindingArrayOutput() ServiceBindingArrayOutput {
	return o
}

func (o ServiceBindingArrayOutput) ToServiceBindingArrayOutputWithContext(ctx context.Context) ServiceBindingArrayOutput {
	return o
}

func (o ServiceBindingArrayOutput) Index(i pulumi.IntInput) ServiceBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceBinding {
		return vs[0].([]*ServiceBinding)[vs[1].(int)]
	}).(ServiceBindingOutput)
}

type ServiceBindingMapOutput struct{ *pulumi.OutputState }

func (ServiceBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceBinding)(nil)).Elem()
}

func (o ServiceBindingMapOutput) ToServiceBindingMapOutput() ServiceBindingMapOutput {
	return o
}

func (o ServiceBindingMapOutput) ToServiceBindingMapOutputWithContext(ctx context.Context) ServiceBindingMapOutput {
	return o
}

func (o ServiceBindingMapOutput) MapIndex(k pulumi.StringInput) ServiceBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceBinding {
		return vs[0].(map[string]*ServiceBinding)[vs[1].(string)]
	}).(ServiceBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceBindingInput)(nil)).Elem(), &ServiceBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceBindingArrayInput)(nil)).Elem(), ServiceBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceBindingMapInput)(nil)).Elem(), ServiceBindingMap{})
	pulumi.RegisterOutputType(ServiceBindingOutput{})
	pulumi.RegisterOutputType(ServiceBindingArrayOutput{})
	pulumi.RegisterOutputType(ServiceBindingMapOutput{})
}
