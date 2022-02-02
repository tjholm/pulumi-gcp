// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of a single API service for a Google Cloud Platform project.
//
// For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
// or run `gcloud services list --available`.
//
// This resource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
// to use.
//
// To get more information about `projects.Service`, see:
//
// * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
// * How-to Guides
//     * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewService(ctx, "project", &projects.ServiceArgs{
// 			DisableDependentServices: pulumi.Bool(true),
// 			Project:                  pulumi.String("your-project-id"),
// 			Service:                  pulumi.String("iam.googleapis.com"),
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
// Project services can be imported using the `project_id` and `service`, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/service:Service my_project your-project-id/iam.googleapis.com
// ```
//
//  Note that unlike other resources that fail if they already exist, `terraform apply` can be successfully used to verify already enabled services. This means that when importing existing resources into Terraform, you can either import the `google_project_service` resources or treat them as new infrastructure and run `terraform apply` to add them to state.
type Service struct {
	pulumi.CustomResourceState

	// If `true`, services that are enabled
	// and which depend on this service should also be disabled when this service is
	// destroyed. If `false` or unset, an error will be generated if any enabled
	// services depend on this service when destroying it.
	DisableDependentServices pulumi.BoolPtrOutput `pulumi:"disableDependentServices"`
	// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy pulumi.BoolPtrOutput `pulumi:"disableOnDestroy"`
	// The project ID. If not provided, the provider project
	// is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The service to enable.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource Service
	err := ctx.RegisterResource("gcp:projects/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:projects/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// If `true`, services that are enabled
	// and which depend on this service should also be disabled when this service is
	// destroyed. If `false` or unset, an error will be generated if any enabled
	// services depend on this service when destroying it.
	DisableDependentServices *bool `pulumi:"disableDependentServices"`
	// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy *bool `pulumi:"disableOnDestroy"`
	// The project ID. If not provided, the provider project
	// is used.
	Project *string `pulumi:"project"`
	// The service to enable.
	Service *string `pulumi:"service"`
}

type ServiceState struct {
	// If `true`, services that are enabled
	// and which depend on this service should also be disabled when this service is
	// destroyed. If `false` or unset, an error will be generated if any enabled
	// services depend on this service when destroying it.
	DisableDependentServices pulumi.BoolPtrInput
	// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy pulumi.BoolPtrInput
	// The project ID. If not provided, the provider project
	// is used.
	Project pulumi.StringPtrInput
	// The service to enable.
	Service pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// If `true`, services that are enabled
	// and which depend on this service should also be disabled when this service is
	// destroyed. If `false` or unset, an error will be generated if any enabled
	// services depend on this service when destroying it.
	DisableDependentServices *bool `pulumi:"disableDependentServices"`
	// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy *bool `pulumi:"disableOnDestroy"`
	// The project ID. If not provided, the provider project
	// is used.
	Project *string `pulumi:"project"`
	// The service to enable.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// If `true`, services that are enabled
	// and which depend on this service should also be disabled when this service is
	// destroyed. If `false` or unset, an error will be generated if any enabled
	// services depend on this service when destroying it.
	DisableDependentServices pulumi.BoolPtrInput
	// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy pulumi.BoolPtrInput
	// The project ID. If not provided, the provider project
	// is used.
	Project pulumi.StringPtrInput
	// The service to enable.
	Service pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
