// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Four different resources help you manage your IAM policy for a folder. Each of these resources serves a different use case:
//
// * `folder.IAMPolicy`: Authoritative. Sets the IAM policy for the folder and replaces any existing policy already attached.
// * `folder.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the folder are preserved.
// * `folder.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the folder are preserved.
// * `folder.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `folder.IAMPolicy` **cannot** be used in conjunction with `folder.IAMBinding`, `folder.IAMMember`, or `folder.IamAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `folder.IAMBinding` resources **can be** used in conjunction with `folder.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:** The underlying API method `projects.setIamPolicy` has constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
//
//	IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning a 400 error code so please review these if you encounter errors with this resource.
//
// ## google\_folder\_iam\_policy
//
// !> **Be careful!** You can accidentally lock yourself out of your folder
//
//	using this resource. Deleting a `folder.IAMPolicy` removes access
//	from anyone without permissions on its parent folder/organization. Proceed with caution.
//	It's not recommended to use `folder.IAMPolicy` with your provider folder
//	to avoid locking yourself out, and it should generally only be used with folders
//	fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//	applying the change.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
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
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = folder.NewIAMPolicy(ctx, "folder", &folder.IAMPolicyArgs{
//				Folder:     pulumi.String("folders/1234567"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
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
//						Condition: {
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//							Title:       "expires_after_2019_12_31",
//						},
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Role: "roles/compute.admin",
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = folder.NewIAMPolicy(ctx, "folder", &folder.IAMPolicyArgs{
//				Folder:     pulumi.String("folders/1234567"),
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
// ## google\_folder\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := folder.NewIAMBinding(ctx, "folder", &folder.IAMBindingArgs{
//				Folder: pulumi.String("folders/1234567"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := folder.NewIAMBinding(ctx, "folder", &folder.IAMBindingArgs{
//				Condition: &folder.IAMBindingConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Folder: pulumi.String("folders/1234567"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/container.admin"),
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
// ## google\_folder\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := folder.NewIAMMember(ctx, "folder", &folder.IAMMemberArgs{
//				Folder: pulumi.String("folders/1234567"),
//				Member: pulumi.String("user:jane@example.com"),
//				Role:   pulumi.String("roles/editor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := folder.NewIAMMember(ctx, "folder", &folder.IAMMemberArgs{
//				Condition: &folder.IAMMemberConditionArgs{
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
//					Title:       pulumi.String("expires_after_2019_12_31"),
//				},
//				Folder: pulumi.String("folders/1234567"),
//				Member: pulumi.String("user:jane@example.com"),
//				Role:   pulumi.String("roles/firebase.admin"),
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
// ## google\_folder\_iam\_audit\_config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/folder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := folder.NewIamAuditConfig(ctx, "folder", &folder.IamAuditConfigArgs{
//				AuditLogConfigs: folder.IamAuditConfigAuditLogConfigArray{
//					&folder.IamAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&folder.IamAuditConfigAuditLogConfigArgs{
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@hashicorp.com"),
//						},
//						LogType: pulumi.String("DATA_READ"),
//					},
//				},
//				Folder:  pulumi.String("folders/1234567"),
//				Service: pulumi.String("allServices"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `folder`, role, and member e.g.
//
// ```sh
//
//	$ pulumi import gcp:folder/iamAuditConfig:IamAuditConfig my_folder "folder roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `folder` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:folder/iamAuditConfig:IamAuditConfig my_folder "folder roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `folder`.
//
// ```sh
//
//	$ pulumi import gcp:folder/iamAuditConfig:IamAuditConfig my_folder folder
//
// ```
//
//	IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//
//	$ pulumi import gcp:folder/iamAuditConfig:IamAuditConfig my_folder "folder foo.googleapis.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
//
// ```sh
//
//	$ pulumi import gcp:folder/iamAuditConfig:IamAuditConfig to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
//
// ```
type IamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// (Computed) The etag of the folder's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIamAuditConfig(ctx *pulumi.Context,
	name string, args *IamAuditConfigArgs, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource IamAuditConfig
	err := ctx.RegisterResource("gcp:folder/iamAuditConfig:IamAuditConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamAuditConfig gets an existing IamAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamAuditConfigState, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	var resource IamAuditConfig
	err := ctx.ReadResource("gcp:folder/iamAuditConfig:IamAuditConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamAuditConfig resources.
type iamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// (Computed) The etag of the folder's IAM policy.
	Etag *string `pulumi:"etag"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder *string `pulumi:"folder"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service *string `pulumi:"service"`
}

type IamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// (Computed) The etag of the folder's IAM policy.
	Etag pulumi.StringPtrInput
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringPtrInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringPtrInput
}

func (IamAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigState)(nil)).Elem()
}

type iamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `pulumi:"folder"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamAuditConfig resource.
type IamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringInput
}

func (IamAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigArgs)(nil)).Elem()
}

type IamAuditConfigInput interface {
	pulumi.Input

	ToIamAuditConfigOutput() IamAuditConfigOutput
	ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput
}

func (*IamAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAuditConfig)(nil)).Elem()
}

func (i *IamAuditConfig) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return i.ToIamAuditConfigOutputWithContext(context.Background())
}

func (i *IamAuditConfig) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigOutput)
}

// IamAuditConfigArrayInput is an input type that accepts IamAuditConfigArray and IamAuditConfigArrayOutput values.
// You can construct a concrete instance of `IamAuditConfigArrayInput` via:
//
//	IamAuditConfigArray{ IamAuditConfigArgs{...} }
type IamAuditConfigArrayInput interface {
	pulumi.Input

	ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput
	ToIamAuditConfigArrayOutputWithContext(context.Context) IamAuditConfigArrayOutput
}

type IamAuditConfigArray []IamAuditConfigInput

func (IamAuditConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamAuditConfig)(nil)).Elem()
}

func (i IamAuditConfigArray) ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput {
	return i.ToIamAuditConfigArrayOutputWithContext(context.Background())
}

func (i IamAuditConfigArray) ToIamAuditConfigArrayOutputWithContext(ctx context.Context) IamAuditConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigArrayOutput)
}

// IamAuditConfigMapInput is an input type that accepts IamAuditConfigMap and IamAuditConfigMapOutput values.
// You can construct a concrete instance of `IamAuditConfigMapInput` via:
//
//	IamAuditConfigMap{ "key": IamAuditConfigArgs{...} }
type IamAuditConfigMapInput interface {
	pulumi.Input

	ToIamAuditConfigMapOutput() IamAuditConfigMapOutput
	ToIamAuditConfigMapOutputWithContext(context.Context) IamAuditConfigMapOutput
}

type IamAuditConfigMap map[string]IamAuditConfigInput

func (IamAuditConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamAuditConfig)(nil)).Elem()
}

func (i IamAuditConfigMap) ToIamAuditConfigMapOutput() IamAuditConfigMapOutput {
	return i.ToIamAuditConfigMapOutputWithContext(context.Background())
}

func (i IamAuditConfigMap) ToIamAuditConfigMapOutputWithContext(ctx context.Context) IamAuditConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigMapOutput)
}

type IamAuditConfigOutput struct{ *pulumi.OutputState }

func (IamAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAuditConfig)(nil)).Elem()
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return o
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return o
}

// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
func (o IamAuditConfigOutput) AuditLogConfigs() IamAuditConfigAuditLogConfigArrayOutput {
	return o.ApplyT(func(v *IamAuditConfig) IamAuditConfigAuditLogConfigArrayOutput { return v.AuditLogConfigs }).(IamAuditConfigAuditLogConfigArrayOutput)
}

// (Computed) The etag of the folder's IAM policy.
func (o IamAuditConfigOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAuditConfig) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
func (o IamAuditConfigOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAuditConfig) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
func (o IamAuditConfigOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAuditConfig) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

type IamAuditConfigArrayOutput struct{ *pulumi.OutputState }

func (IamAuditConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamAuditConfig)(nil)).Elem()
}

func (o IamAuditConfigArrayOutput) ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput {
	return o
}

func (o IamAuditConfigArrayOutput) ToIamAuditConfigArrayOutputWithContext(ctx context.Context) IamAuditConfigArrayOutput {
	return o
}

func (o IamAuditConfigArrayOutput) Index(i pulumi.IntInput) IamAuditConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamAuditConfig {
		return vs[0].([]*IamAuditConfig)[vs[1].(int)]
	}).(IamAuditConfigOutput)
}

type IamAuditConfigMapOutput struct{ *pulumi.OutputState }

func (IamAuditConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamAuditConfig)(nil)).Elem()
}

func (o IamAuditConfigMapOutput) ToIamAuditConfigMapOutput() IamAuditConfigMapOutput {
	return o
}

func (o IamAuditConfigMapOutput) ToIamAuditConfigMapOutputWithContext(ctx context.Context) IamAuditConfigMapOutput {
	return o
}

func (o IamAuditConfigMapOutput) MapIndex(k pulumi.StringInput) IamAuditConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamAuditConfig {
		return vs[0].(map[string]*IamAuditConfig)[vs[1].(string)]
	}).(IamAuditConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamAuditConfigInput)(nil)).Elem(), &IamAuditConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamAuditConfigArrayInput)(nil)).Elem(), IamAuditConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamAuditConfigMapInput)(nil)).Elem(), IamAuditConfigMap{})
	pulumi.RegisterOutputType(IamAuditConfigOutput{})
	pulumi.RegisterOutputType(IamAuditConfigArrayOutput{})
	pulumi.RegisterOutputType(IamAuditConfigMapOutput{})
}
