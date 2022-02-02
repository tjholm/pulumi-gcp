// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Directory Service. Each of these resources serves a different use case:
//
// * `servicedirectory.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `servicedirectory.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `servicedirectory.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// > **Note:** `servicedirectory.ServiceIamPolicy` **cannot** be used in conjunction with `servicedirectory.ServiceIamBinding` and `servicedirectory.ServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `servicedirectory.ServiceIamBinding` resources **can be** used in conjunction with `servicedirectory.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_service\_directory\_service\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicedirectory.NewServiceIamPolicy(ctx, "policy", &servicedirectory.ServiceIamPolicyArgs{
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_directory\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicedirectory.NewServiceIamBinding(ctx, "binding", &servicedirectory.ServiceIamBindingArgs{
// 			Role: pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_directory\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicedirectory.NewServiceIamMember(ctx, "member", &servicedirectory.ServiceIamMemberArgs{
// 			Role:   pulumi.String("roles/viewer"),
// 			Member: pulumi.String("user:jane@example.com"),
// 		}, pulumi.Provider(google_beta))
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} * {{project}}/{{location}}/{{namespace_id}}/{{service_id}} * {{location}}/{{namespace_id}}/{{service_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/serviceIamBinding:ServiceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/serviceIamBinding:ServiceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/serviceIamBinding:ServiceIamBinding editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ServiceIamBinding struct {
	pulumi.CustomResourceState

	Condition ServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewServiceIamBinding(ctx *pulumi.Context,
	name string, args *ServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*ServiceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ServiceIamBinding
	err := ctx.RegisterResource("gcp:servicedirectory/serviceIamBinding:ServiceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIamBinding gets an existing ServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIamBindingState, opts ...pulumi.ResourceOption) (*ServiceIamBinding, error) {
	var resource ServiceIamBinding
	err := ctx.ReadResource("gcp:servicedirectory/serviceIamBinding:ServiceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamBinding resources.
type serviceIamBindingState struct {
	Condition *ServiceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ServiceIamBindingState struct {
	Condition ServiceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamBindingState)(nil)).Elem()
}

type serviceIamBindingArgs struct {
	Condition *ServiceIamBindingCondition `pulumi:"condition"`
	Members   []string                    `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ServiceIamBinding resource.
type ServiceIamBindingArgs struct {
	Condition ServiceIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamBindingArgs)(nil)).Elem()
}

type ServiceIamBindingInput interface {
	pulumi.Input

	ToServiceIamBindingOutput() ServiceIamBindingOutput
	ToServiceIamBindingOutputWithContext(ctx context.Context) ServiceIamBindingOutput
}

func (*ServiceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIamBinding)(nil)).Elem()
}

func (i *ServiceIamBinding) ToServiceIamBindingOutput() ServiceIamBindingOutput {
	return i.ToServiceIamBindingOutputWithContext(context.Background())
}

func (i *ServiceIamBinding) ToServiceIamBindingOutputWithContext(ctx context.Context) ServiceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamBindingOutput)
}

// ServiceIamBindingArrayInput is an input type that accepts ServiceIamBindingArray and ServiceIamBindingArrayOutput values.
// You can construct a concrete instance of `ServiceIamBindingArrayInput` via:
//
//          ServiceIamBindingArray{ ServiceIamBindingArgs{...} }
type ServiceIamBindingArrayInput interface {
	pulumi.Input

	ToServiceIamBindingArrayOutput() ServiceIamBindingArrayOutput
	ToServiceIamBindingArrayOutputWithContext(context.Context) ServiceIamBindingArrayOutput
}

type ServiceIamBindingArray []ServiceIamBindingInput

func (ServiceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceIamBinding)(nil)).Elem()
}

func (i ServiceIamBindingArray) ToServiceIamBindingArrayOutput() ServiceIamBindingArrayOutput {
	return i.ToServiceIamBindingArrayOutputWithContext(context.Background())
}

func (i ServiceIamBindingArray) ToServiceIamBindingArrayOutputWithContext(ctx context.Context) ServiceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamBindingArrayOutput)
}

// ServiceIamBindingMapInput is an input type that accepts ServiceIamBindingMap and ServiceIamBindingMapOutput values.
// You can construct a concrete instance of `ServiceIamBindingMapInput` via:
//
//          ServiceIamBindingMap{ "key": ServiceIamBindingArgs{...} }
type ServiceIamBindingMapInput interface {
	pulumi.Input

	ToServiceIamBindingMapOutput() ServiceIamBindingMapOutput
	ToServiceIamBindingMapOutputWithContext(context.Context) ServiceIamBindingMapOutput
}

type ServiceIamBindingMap map[string]ServiceIamBindingInput

func (ServiceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceIamBinding)(nil)).Elem()
}

func (i ServiceIamBindingMap) ToServiceIamBindingMapOutput() ServiceIamBindingMapOutput {
	return i.ToServiceIamBindingMapOutputWithContext(context.Background())
}

func (i ServiceIamBindingMap) ToServiceIamBindingMapOutputWithContext(ctx context.Context) ServiceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamBindingMapOutput)
}

type ServiceIamBindingOutput struct{ *pulumi.OutputState }

func (ServiceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIamBinding)(nil)).Elem()
}

func (o ServiceIamBindingOutput) ToServiceIamBindingOutput() ServiceIamBindingOutput {
	return o
}

func (o ServiceIamBindingOutput) ToServiceIamBindingOutputWithContext(ctx context.Context) ServiceIamBindingOutput {
	return o
}

type ServiceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ServiceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceIamBinding)(nil)).Elem()
}

func (o ServiceIamBindingArrayOutput) ToServiceIamBindingArrayOutput() ServiceIamBindingArrayOutput {
	return o
}

func (o ServiceIamBindingArrayOutput) ToServiceIamBindingArrayOutputWithContext(ctx context.Context) ServiceIamBindingArrayOutput {
	return o
}

func (o ServiceIamBindingArrayOutput) Index(i pulumi.IntInput) ServiceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceIamBinding {
		return vs[0].([]*ServiceIamBinding)[vs[1].(int)]
	}).(ServiceIamBindingOutput)
}

type ServiceIamBindingMapOutput struct{ *pulumi.OutputState }

func (ServiceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceIamBinding)(nil)).Elem()
}

func (o ServiceIamBindingMapOutput) ToServiceIamBindingMapOutput() ServiceIamBindingMapOutput {
	return o
}

func (o ServiceIamBindingMapOutput) ToServiceIamBindingMapOutputWithContext(ctx context.Context) ServiceIamBindingMapOutput {
	return o
}

func (o ServiceIamBindingMapOutput) MapIndex(k pulumi.StringInput) ServiceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceIamBinding {
		return vs[0].(map[string]*ServiceIamBinding)[vs[1].(string)]
	}).(ServiceIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceIamBindingInput)(nil)).Elem(), &ServiceIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceIamBindingArrayInput)(nil)).Elem(), ServiceIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceIamBindingMapInput)(nil)).Elem(), ServiceIamBindingMap{})
	pulumi.RegisterOutputType(ServiceIamBindingOutput{})
	pulumi.RegisterOutputType(ServiceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ServiceIamBindingMapOutput{})
}
