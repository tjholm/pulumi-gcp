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
// 		})
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
// 		})
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} * {{project}}/{{location}}/{{namespace_id}}/{{service_id}} * {{location}}/{{namespace_id}}/{{service_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ServiceIamMember struct {
	pulumi.CustomResourceState

	Condition ServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewServiceIamMember(ctx *pulumi.Context,
	name string, args *ServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*ServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ServiceIamMember
	err := ctx.RegisterResource("gcp:servicedirectory/serviceIamMember:ServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIamMember gets an existing ServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIamMemberState, opts ...pulumi.ResourceOption) (*ServiceIamMember, error) {
	var resource ServiceIamMember
	err := ctx.ReadResource("gcp:servicedirectory/serviceIamMember:ServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamMember resources.
type serviceIamMemberState struct {
	Condition *ServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ServiceIamMemberState struct {
	Condition ServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamMemberState)(nil)).Elem()
}

type serviceIamMemberArgs struct {
	Condition *ServiceIamMemberCondition `pulumi:"condition"`
	Member    string                     `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ServiceIamMember resource.
type ServiceIamMemberArgs struct {
	Condition ServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamMemberArgs)(nil)).Elem()
}

type ServiceIamMemberInput interface {
	pulumi.Input

	ToServiceIamMemberOutput() ServiceIamMemberOutput
	ToServiceIamMemberOutputWithContext(ctx context.Context) ServiceIamMemberOutput
}

func (*ServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIamMember)(nil))
}

func (i *ServiceIamMember) ToServiceIamMemberOutput() ServiceIamMemberOutput {
	return i.ToServiceIamMemberOutputWithContext(context.Background())
}

func (i *ServiceIamMember) ToServiceIamMemberOutputWithContext(ctx context.Context) ServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamMemberOutput)
}

func (i *ServiceIamMember) ToServiceIamMemberPtrOutput() ServiceIamMemberPtrOutput {
	return i.ToServiceIamMemberPtrOutputWithContext(context.Background())
}

func (i *ServiceIamMember) ToServiceIamMemberPtrOutputWithContext(ctx context.Context) ServiceIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamMemberPtrOutput)
}

type ServiceIamMemberPtrInput interface {
	pulumi.Input

	ToServiceIamMemberPtrOutput() ServiceIamMemberPtrOutput
	ToServiceIamMemberPtrOutputWithContext(ctx context.Context) ServiceIamMemberPtrOutput
}

type serviceIamMemberPtrType ServiceIamMemberArgs

func (*serviceIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIamMember)(nil))
}

func (i *serviceIamMemberPtrType) ToServiceIamMemberPtrOutput() ServiceIamMemberPtrOutput {
	return i.ToServiceIamMemberPtrOutputWithContext(context.Background())
}

func (i *serviceIamMemberPtrType) ToServiceIamMemberPtrOutputWithContext(ctx context.Context) ServiceIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamMemberPtrOutput)
}

// ServiceIamMemberArrayInput is an input type that accepts ServiceIamMemberArray and ServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `ServiceIamMemberArrayInput` via:
//
//          ServiceIamMemberArray{ ServiceIamMemberArgs{...} }
type ServiceIamMemberArrayInput interface {
	pulumi.Input

	ToServiceIamMemberArrayOutput() ServiceIamMemberArrayOutput
	ToServiceIamMemberArrayOutputWithContext(context.Context) ServiceIamMemberArrayOutput
}

type ServiceIamMemberArray []ServiceIamMemberInput

func (ServiceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceIamMember)(nil)).Elem()
}

func (i ServiceIamMemberArray) ToServiceIamMemberArrayOutput() ServiceIamMemberArrayOutput {
	return i.ToServiceIamMemberArrayOutputWithContext(context.Background())
}

func (i ServiceIamMemberArray) ToServiceIamMemberArrayOutputWithContext(ctx context.Context) ServiceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamMemberArrayOutput)
}

// ServiceIamMemberMapInput is an input type that accepts ServiceIamMemberMap and ServiceIamMemberMapOutput values.
// You can construct a concrete instance of `ServiceIamMemberMapInput` via:
//
//          ServiceIamMemberMap{ "key": ServiceIamMemberArgs{...} }
type ServiceIamMemberMapInput interface {
	pulumi.Input

	ToServiceIamMemberMapOutput() ServiceIamMemberMapOutput
	ToServiceIamMemberMapOutputWithContext(context.Context) ServiceIamMemberMapOutput
}

type ServiceIamMemberMap map[string]ServiceIamMemberInput

func (ServiceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceIamMember)(nil)).Elem()
}

func (i ServiceIamMemberMap) ToServiceIamMemberMapOutput() ServiceIamMemberMapOutput {
	return i.ToServiceIamMemberMapOutputWithContext(context.Background())
}

func (i ServiceIamMemberMap) ToServiceIamMemberMapOutputWithContext(ctx context.Context) ServiceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamMemberMapOutput)
}

type ServiceIamMemberOutput struct{ *pulumi.OutputState }

func (ServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIamMember)(nil))
}

func (o ServiceIamMemberOutput) ToServiceIamMemberOutput() ServiceIamMemberOutput {
	return o
}

func (o ServiceIamMemberOutput) ToServiceIamMemberOutputWithContext(ctx context.Context) ServiceIamMemberOutput {
	return o
}

func (o ServiceIamMemberOutput) ToServiceIamMemberPtrOutput() ServiceIamMemberPtrOutput {
	return o.ToServiceIamMemberPtrOutputWithContext(context.Background())
}

func (o ServiceIamMemberOutput) ToServiceIamMemberPtrOutputWithContext(ctx context.Context) ServiceIamMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceIamMember) *ServiceIamMember {
		return &v
	}).(ServiceIamMemberPtrOutput)
}

type ServiceIamMemberPtrOutput struct{ *pulumi.OutputState }

func (ServiceIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIamMember)(nil))
}

func (o ServiceIamMemberPtrOutput) ToServiceIamMemberPtrOutput() ServiceIamMemberPtrOutput {
	return o
}

func (o ServiceIamMemberPtrOutput) ToServiceIamMemberPtrOutputWithContext(ctx context.Context) ServiceIamMemberPtrOutput {
	return o
}

func (o ServiceIamMemberPtrOutput) Elem() ServiceIamMemberOutput {
	return o.ApplyT(func(v *ServiceIamMember) ServiceIamMember {
		if v != nil {
			return *v
		}
		var ret ServiceIamMember
		return ret
	}).(ServiceIamMemberOutput)
}

type ServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceIamMember)(nil))
}

func (o ServiceIamMemberArrayOutput) ToServiceIamMemberArrayOutput() ServiceIamMemberArrayOutput {
	return o
}

func (o ServiceIamMemberArrayOutput) ToServiceIamMemberArrayOutputWithContext(ctx context.Context) ServiceIamMemberArrayOutput {
	return o
}

func (o ServiceIamMemberArrayOutput) Index(i pulumi.IntInput) ServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceIamMember {
		return vs[0].([]ServiceIamMember)[vs[1].(int)]
	}).(ServiceIamMemberOutput)
}

type ServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (ServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceIamMember)(nil))
}

func (o ServiceIamMemberMapOutput) ToServiceIamMemberMapOutput() ServiceIamMemberMapOutput {
	return o
}

func (o ServiceIamMemberMapOutput) ToServiceIamMemberMapOutputWithContext(ctx context.Context) ServiceIamMemberMapOutput {
	return o
}

func (o ServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) ServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceIamMember {
		return vs[0].(map[string]ServiceIamMember)[vs[1].(string)]
	}).(ServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceIamMemberOutput{})
	pulumi.RegisterOutputType(ServiceIamMemberPtrOutput{})
	pulumi.RegisterOutputType(ServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ServiceIamMemberMapOutput{})
}
