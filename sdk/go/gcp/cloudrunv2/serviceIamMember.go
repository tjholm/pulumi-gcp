// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrunv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Run (v2 API) Service. Each of these resources serves a different use case:
//
// * `cloudrunv2.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `cloudrunv2.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `cloudrunv2.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// > **Note:** `cloudrunv2.ServiceIamPolicy` **cannot** be used in conjunction with `cloudrunv2.ServiceIamBinding` and `cloudrunv2.ServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudrunv2.ServiceIamBinding` resources **can be** used in conjunction with `cloudrunv2.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloud\_run\_v2\_service\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudrunv2.NewServiceIamPolicy(ctx, "policy", &cloudrunv2.ServiceIamPolicyArgs{
//				Project:    pulumi.Any(google_cloud_run_v2_service.Default.Project),
//				Location:   pulumi.Any(google_cloud_run_v2_service.Default.Location),
//				PolicyData: *pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_cloud\_run\_v2\_service\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrunv2.NewServiceIamBinding(ctx, "binding", &cloudrunv2.ServiceIamBindingArgs{
//				Project:  pulumi.Any(google_cloud_run_v2_service.Default.Project),
//				Location: pulumi.Any(google_cloud_run_v2_service.Default.Location),
//				Role:     pulumi.String("roles/viewer"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_cloud\_run\_v2\_service\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrunv2.NewServiceIamMember(ctx, "member", &cloudrunv2.ServiceIamMemberArgs{
//				Project:  pulumi.Any(google_cloud_run_v2_service.Default.Project),
//				Location: pulumi.Any(google_cloud_run_v2_service.Default.Location),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
//			})
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/services/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Run (v2 API) service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor projects/{{project}}/locations/{{location}}/services/{{service}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ServiceIamMember struct {
	pulumi.CustomResourceState

	Condition ServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the cloud run service Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
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
	err := ctx.RegisterResource("gcp:cloudrunv2/serviceIamMember:ServiceIamMember", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:cloudrunv2/serviceIamMember:ServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamMember resources.
type serviceIamMemberState struct {
	Condition *ServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of the cloud run service Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ServiceIamMemberState struct {
	Condition ServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of the cloud run service Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamMemberState)(nil)).Elem()
}

type serviceIamMemberArgs struct {
	Condition *ServiceIamMemberCondition `pulumi:"condition"`
	// The location of the cloud run service Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   string  `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ServiceIamMember resource.
type ServiceIamMemberArgs struct {
	Condition ServiceIamMemberConditionPtrInput
	// The location of the cloud run service Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
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
	return reflect.TypeOf((**ServiceIamMember)(nil)).Elem()
}

func (i *ServiceIamMember) ToServiceIamMemberOutput() ServiceIamMemberOutput {
	return i.ToServiceIamMemberOutputWithContext(context.Background())
}

func (i *ServiceIamMember) ToServiceIamMemberOutputWithContext(ctx context.Context) ServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamMemberOutput)
}

// ServiceIamMemberArrayInput is an input type that accepts ServiceIamMemberArray and ServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `ServiceIamMemberArrayInput` via:
//
//	ServiceIamMemberArray{ ServiceIamMemberArgs{...} }
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
//	ServiceIamMemberMap{ "key": ServiceIamMemberArgs{...} }
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
	return reflect.TypeOf((**ServiceIamMember)(nil)).Elem()
}

func (o ServiceIamMemberOutput) ToServiceIamMemberOutput() ServiceIamMemberOutput {
	return o
}

func (o ServiceIamMemberOutput) ToServiceIamMemberOutputWithContext(ctx context.Context) ServiceIamMemberOutput {
	return o
}

func (o ServiceIamMemberOutput) Condition() ServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *ServiceIamMember) ServiceIamMemberConditionPtrOutput { return v.Condition }).(ServiceIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o ServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location of the cloud run service Used to find the parent resource to bind the IAM policy to
func (o ServiceIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o ServiceIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o ServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o ServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceIamMember)(nil)).Elem()
}

func (o ServiceIamMemberArrayOutput) ToServiceIamMemberArrayOutput() ServiceIamMemberArrayOutput {
	return o
}

func (o ServiceIamMemberArrayOutput) ToServiceIamMemberArrayOutputWithContext(ctx context.Context) ServiceIamMemberArrayOutput {
	return o
}

func (o ServiceIamMemberArrayOutput) Index(i pulumi.IntInput) ServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceIamMember {
		return vs[0].([]*ServiceIamMember)[vs[1].(int)]
	}).(ServiceIamMemberOutput)
}

type ServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (ServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceIamMember)(nil)).Elem()
}

func (o ServiceIamMemberMapOutput) ToServiceIamMemberMapOutput() ServiceIamMemberMapOutput {
	return o
}

func (o ServiceIamMemberMapOutput) ToServiceIamMemberMapOutputWithContext(ctx context.Context) ServiceIamMemberMapOutput {
	return o
}

func (o ServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) ServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceIamMember {
		return vs[0].(map[string]*ServiceIamMember)[vs[1].(string)]
	}).(ServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceIamMemberInput)(nil)).Elem(), &ServiceIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceIamMemberArrayInput)(nil)).Elem(), ServiceIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceIamMemberMapInput)(nil)).Elem(), ServiceIamMemberMap{})
	pulumi.RegisterOutputType(ServiceIamMemberOutput{})
	pulumi.RegisterOutputType(ServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ServiceIamMemberMapOutput{})
}
