// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows configuring a single GCP resource that should be inside of a service perimeter.
// This resource is intended to be used in cases where it is not possible to compile a full list
// of projects to include in a `accesscontextmanager.ServicePerimeter` resource,
// to enable them to be added separately.
//
// > **Note:** If this resource is used alongside a `accesscontextmanager.ServicePerimeter` resource,
// the service perimeter resource must have a `lifecycle` block with `ignoreChanges = [status[0].resources]` so
// they don't fight over which resources should be in the policy.
//
// To get more information about ServicePerimeterResource, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
// * How-to Guides
//     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Access Context Manager Service Perimeter Resource Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := accesscontextmanager.NewAccessPolicy(ctx, "access-policy", &accesscontextmanager.AccessPolicyArgs{
// 			Parent: pulumi.String("organizations/123456789"),
// 			Title:  pulumi.String("my policy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewServicePerimeter(ctx, "service-perimeter-resourceServicePerimeter", &accesscontextmanager.ServicePerimeterArgs{
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 			Title: pulumi.String("restrict_all"),
// 			Status: &accesscontextmanager.ServicePerimeterStatusArgs{
// 				RestrictedServices: pulumi.StringArray{
// 					pulumi.String("storage.googleapis.com"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewServicePerimeterResource(ctx, "service-perimeter-resourceServicePerimeterResource", &accesscontextmanager.ServicePerimeterResourceArgs{
// 			PerimeterName: service_perimeter_resourceServicePerimeter.Name,
// 			Resource:      pulumi.String("projects/987654321"),
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
// ServicePerimeterResource can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource default {{perimeter_name}}/{{resource}}
// ```
type ServicePerimeterResource struct {
	pulumi.CustomResourceState

	// The name of the Service Perimeter to add this resource to.
	PerimeterName pulumi.StringOutput `pulumi:"perimeterName"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource pulumi.StringOutput `pulumi:"resource"`
}

// NewServicePerimeterResource registers a new resource with the given unique name, arguments, and options.
func NewServicePerimeterResource(ctx *pulumi.Context,
	name string, args *ServicePerimeterResourceArgs, opts ...pulumi.ResourceOption) (*ServicePerimeterResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'PerimeterName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource ServicePerimeterResource
	err := ctx.RegisterResource("gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePerimeterResource gets an existing ServicePerimeterResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePerimeterResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePerimeterResourceState, opts ...pulumi.ResourceOption) (*ServicePerimeterResource, error) {
	var resource ServicePerimeterResource
	err := ctx.ReadResource("gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePerimeterResource resources.
type servicePerimeterResourceState struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName *string `pulumi:"perimeterName"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource *string `pulumi:"resource"`
}

type ServicePerimeterResourceState struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName pulumi.StringPtrInput
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource pulumi.StringPtrInput
}

func (ServicePerimeterResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterResourceState)(nil)).Elem()
}

type servicePerimeterResourceArgs struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName string `pulumi:"perimeterName"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource string `pulumi:"resource"`
}

// The set of arguments for constructing a ServicePerimeterResource resource.
type ServicePerimeterResourceArgs struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName pulumi.StringInput
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource pulumi.StringInput
}

func (ServicePerimeterResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterResourceArgs)(nil)).Elem()
}

type ServicePerimeterResourceInput interface {
	pulumi.Input

	ToServicePerimeterResourceOutput() ServicePerimeterResourceOutput
	ToServicePerimeterResourceOutputWithContext(ctx context.Context) ServicePerimeterResourceOutput
}

func (*ServicePerimeterResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePerimeterResource)(nil)).Elem()
}

func (i *ServicePerimeterResource) ToServicePerimeterResourceOutput() ServicePerimeterResourceOutput {
	return i.ToServicePerimeterResourceOutputWithContext(context.Background())
}

func (i *ServicePerimeterResource) ToServicePerimeterResourceOutputWithContext(ctx context.Context) ServicePerimeterResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimeterResourceOutput)
}

// ServicePerimeterResourceArrayInput is an input type that accepts ServicePerimeterResourceArray and ServicePerimeterResourceArrayOutput values.
// You can construct a concrete instance of `ServicePerimeterResourceArrayInput` via:
//
//          ServicePerimeterResourceArray{ ServicePerimeterResourceArgs{...} }
type ServicePerimeterResourceArrayInput interface {
	pulumi.Input

	ToServicePerimeterResourceArrayOutput() ServicePerimeterResourceArrayOutput
	ToServicePerimeterResourceArrayOutputWithContext(context.Context) ServicePerimeterResourceArrayOutput
}

type ServicePerimeterResourceArray []ServicePerimeterResourceInput

func (ServicePerimeterResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePerimeterResource)(nil)).Elem()
}

func (i ServicePerimeterResourceArray) ToServicePerimeterResourceArrayOutput() ServicePerimeterResourceArrayOutput {
	return i.ToServicePerimeterResourceArrayOutputWithContext(context.Background())
}

func (i ServicePerimeterResourceArray) ToServicePerimeterResourceArrayOutputWithContext(ctx context.Context) ServicePerimeterResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimeterResourceArrayOutput)
}

// ServicePerimeterResourceMapInput is an input type that accepts ServicePerimeterResourceMap and ServicePerimeterResourceMapOutput values.
// You can construct a concrete instance of `ServicePerimeterResourceMapInput` via:
//
//          ServicePerimeterResourceMap{ "key": ServicePerimeterResourceArgs{...} }
type ServicePerimeterResourceMapInput interface {
	pulumi.Input

	ToServicePerimeterResourceMapOutput() ServicePerimeterResourceMapOutput
	ToServicePerimeterResourceMapOutputWithContext(context.Context) ServicePerimeterResourceMapOutput
}

type ServicePerimeterResourceMap map[string]ServicePerimeterResourceInput

func (ServicePerimeterResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePerimeterResource)(nil)).Elem()
}

func (i ServicePerimeterResourceMap) ToServicePerimeterResourceMapOutput() ServicePerimeterResourceMapOutput {
	return i.ToServicePerimeterResourceMapOutputWithContext(context.Background())
}

func (i ServicePerimeterResourceMap) ToServicePerimeterResourceMapOutputWithContext(ctx context.Context) ServicePerimeterResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimeterResourceMapOutput)
}

type ServicePerimeterResourceOutput struct{ *pulumi.OutputState }

func (ServicePerimeterResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePerimeterResource)(nil)).Elem()
}

func (o ServicePerimeterResourceOutput) ToServicePerimeterResourceOutput() ServicePerimeterResourceOutput {
	return o
}

func (o ServicePerimeterResourceOutput) ToServicePerimeterResourceOutputWithContext(ctx context.Context) ServicePerimeterResourceOutput {
	return o
}

type ServicePerimeterResourceArrayOutput struct{ *pulumi.OutputState }

func (ServicePerimeterResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePerimeterResource)(nil)).Elem()
}

func (o ServicePerimeterResourceArrayOutput) ToServicePerimeterResourceArrayOutput() ServicePerimeterResourceArrayOutput {
	return o
}

func (o ServicePerimeterResourceArrayOutput) ToServicePerimeterResourceArrayOutputWithContext(ctx context.Context) ServicePerimeterResourceArrayOutput {
	return o
}

func (o ServicePerimeterResourceArrayOutput) Index(i pulumi.IntInput) ServicePerimeterResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePerimeterResource {
		return vs[0].([]*ServicePerimeterResource)[vs[1].(int)]
	}).(ServicePerimeterResourceOutput)
}

type ServicePerimeterResourceMapOutput struct{ *pulumi.OutputState }

func (ServicePerimeterResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePerimeterResource)(nil)).Elem()
}

func (o ServicePerimeterResourceMapOutput) ToServicePerimeterResourceMapOutput() ServicePerimeterResourceMapOutput {
	return o
}

func (o ServicePerimeterResourceMapOutput) ToServicePerimeterResourceMapOutputWithContext(ctx context.Context) ServicePerimeterResourceMapOutput {
	return o
}

func (o ServicePerimeterResourceMapOutput) MapIndex(k pulumi.StringInput) ServicePerimeterResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePerimeterResource {
		return vs[0].(map[string]*ServicePerimeterResource)[vs[1].(string)]
	}).(ServicePerimeterResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePerimeterResourceInput)(nil)).Elem(), &ServicePerimeterResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePerimeterResourceArrayInput)(nil)).Elem(), ServicePerimeterResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePerimeterResourceMapInput)(nil)).Elem(), ServicePerimeterResourceMap{})
	pulumi.RegisterOutputType(ServicePerimeterResourceOutput{})
	pulumi.RegisterOutputType(ServicePerimeterResourceArrayOutput{})
	pulumi.RegisterOutputType(ServicePerimeterResourceMapOutput{})
}
