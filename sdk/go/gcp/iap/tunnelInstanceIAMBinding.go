// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelInstance. Each of these resources serves a different use case:
//
// * `iap.TunnelInstanceIAMPolicy`: Authoritative. Sets the IAM policy for the tunnelinstance and replaces any existing policy already attached.
// * `iap.TunnelInstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnelinstance are preserved.
// * `iap.TunnelInstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnelinstance are preserved.
//
// > **Note:** `iap.TunnelInstanceIAMPolicy` **cannot** be used in conjunction with `iap.TunnelInstanceIAMBinding` and `iap.TunnelInstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.TunnelInstanceIAMBinding` resources **can be** used in conjunction with `iap.TunnelInstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_tunnel\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.tunnelResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:       pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance:   pulumi.Any(google_compute_instance.Tunnelvm.Name),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.tunnelResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:       pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance:   pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_tunnel\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.TunnelInstanceIAMBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_tunnel\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Condition: &iap.TunnelInstanceIAMMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{name}} * projects/{{project}}/zones/{{zone}}/instances/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy tunnelinstance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding editor projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TunnelInstanceIAMBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTunnelInstanceIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewTunnelInstanceIAMBinding(ctx *pulumi.Context,
	name string, args *TunnelInstanceIAMBindingArgs, opts ...pulumi.ResourceOption) (*TunnelInstanceIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource TunnelInstanceIAMBinding
	err := ctx.RegisterResource("gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTunnelInstanceIAMBinding gets an existing TunnelInstanceIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunnelInstanceIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TunnelInstanceIAMBindingState, opts ...pulumi.ResourceOption) (*TunnelInstanceIAMBinding, error) {
	var resource TunnelInstanceIAMBinding
	err := ctx.ReadResource("gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TunnelInstanceIAMBinding resources.
type tunnelInstanceIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *TunnelInstanceIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	Zone *string `pulumi:"zone"`
}

type TunnelInstanceIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	Zone pulumi.StringPtrInput
}

func (TunnelInstanceIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelInstanceIAMBindingState)(nil)).Elem()
}

type tunnelInstanceIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *TunnelInstanceIAMBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string  `pulumi:"role"`
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TunnelInstanceIAMBinding resource.
type TunnelInstanceIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	Zone pulumi.StringPtrInput
}

func (TunnelInstanceIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelInstanceIAMBindingArgs)(nil)).Elem()
}

type TunnelInstanceIAMBindingInput interface {
	pulumi.Input

	ToTunnelInstanceIAMBindingOutput() TunnelInstanceIAMBindingOutput
	ToTunnelInstanceIAMBindingOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingOutput
}

func (*TunnelInstanceIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelInstanceIAMBinding)(nil))
}

func (i *TunnelInstanceIAMBinding) ToTunnelInstanceIAMBindingOutput() TunnelInstanceIAMBindingOutput {
	return i.ToTunnelInstanceIAMBindingOutputWithContext(context.Background())
}

func (i *TunnelInstanceIAMBinding) ToTunnelInstanceIAMBindingOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMBindingOutput)
}

func (i *TunnelInstanceIAMBinding) ToTunnelInstanceIAMBindingPtrOutput() TunnelInstanceIAMBindingPtrOutput {
	return i.ToTunnelInstanceIAMBindingPtrOutputWithContext(context.Background())
}

func (i *TunnelInstanceIAMBinding) ToTunnelInstanceIAMBindingPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMBindingPtrOutput)
}

type TunnelInstanceIAMBindingPtrInput interface {
	pulumi.Input

	ToTunnelInstanceIAMBindingPtrOutput() TunnelInstanceIAMBindingPtrOutput
	ToTunnelInstanceIAMBindingPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingPtrOutput
}

type tunnelInstanceIAMBindingPtrType TunnelInstanceIAMBindingArgs

func (*tunnelInstanceIAMBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelInstanceIAMBinding)(nil))
}

func (i *tunnelInstanceIAMBindingPtrType) ToTunnelInstanceIAMBindingPtrOutput() TunnelInstanceIAMBindingPtrOutput {
	return i.ToTunnelInstanceIAMBindingPtrOutputWithContext(context.Background())
}

func (i *tunnelInstanceIAMBindingPtrType) ToTunnelInstanceIAMBindingPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMBindingPtrOutput)
}

// TunnelInstanceIAMBindingArrayInput is an input type that accepts TunnelInstanceIAMBindingArray and TunnelInstanceIAMBindingArrayOutput values.
// You can construct a concrete instance of `TunnelInstanceIAMBindingArrayInput` via:
//
//          TunnelInstanceIAMBindingArray{ TunnelInstanceIAMBindingArgs{...} }
type TunnelInstanceIAMBindingArrayInput interface {
	pulumi.Input

	ToTunnelInstanceIAMBindingArrayOutput() TunnelInstanceIAMBindingArrayOutput
	ToTunnelInstanceIAMBindingArrayOutputWithContext(context.Context) TunnelInstanceIAMBindingArrayOutput
}

type TunnelInstanceIAMBindingArray []TunnelInstanceIAMBindingInput

func (TunnelInstanceIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TunnelInstanceIAMBinding)(nil))
}

func (i TunnelInstanceIAMBindingArray) ToTunnelInstanceIAMBindingArrayOutput() TunnelInstanceIAMBindingArrayOutput {
	return i.ToTunnelInstanceIAMBindingArrayOutputWithContext(context.Background())
}

func (i TunnelInstanceIAMBindingArray) ToTunnelInstanceIAMBindingArrayOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMBindingArrayOutput)
}

// TunnelInstanceIAMBindingMapInput is an input type that accepts TunnelInstanceIAMBindingMap and TunnelInstanceIAMBindingMapOutput values.
// You can construct a concrete instance of `TunnelInstanceIAMBindingMapInput` via:
//
//          TunnelInstanceIAMBindingMap{ "key": TunnelInstanceIAMBindingArgs{...} }
type TunnelInstanceIAMBindingMapInput interface {
	pulumi.Input

	ToTunnelInstanceIAMBindingMapOutput() TunnelInstanceIAMBindingMapOutput
	ToTunnelInstanceIAMBindingMapOutputWithContext(context.Context) TunnelInstanceIAMBindingMapOutput
}

type TunnelInstanceIAMBindingMap map[string]TunnelInstanceIAMBindingInput

func (TunnelInstanceIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TunnelInstanceIAMBinding)(nil))
}

func (i TunnelInstanceIAMBindingMap) ToTunnelInstanceIAMBindingMapOutput() TunnelInstanceIAMBindingMapOutput {
	return i.ToTunnelInstanceIAMBindingMapOutputWithContext(context.Background())
}

func (i TunnelInstanceIAMBindingMap) ToTunnelInstanceIAMBindingMapOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMBindingMapOutput)
}

type TunnelInstanceIAMBindingOutput struct {
	*pulumi.OutputState
}

func (TunnelInstanceIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelInstanceIAMBinding)(nil))
}

func (o TunnelInstanceIAMBindingOutput) ToTunnelInstanceIAMBindingOutput() TunnelInstanceIAMBindingOutput {
	return o
}

func (o TunnelInstanceIAMBindingOutput) ToTunnelInstanceIAMBindingOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingOutput {
	return o
}

func (o TunnelInstanceIAMBindingOutput) ToTunnelInstanceIAMBindingPtrOutput() TunnelInstanceIAMBindingPtrOutput {
	return o.ToTunnelInstanceIAMBindingPtrOutputWithContext(context.Background())
}

func (o TunnelInstanceIAMBindingOutput) ToTunnelInstanceIAMBindingPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingPtrOutput {
	return o.ApplyT(func(v TunnelInstanceIAMBinding) *TunnelInstanceIAMBinding {
		return &v
	}).(TunnelInstanceIAMBindingPtrOutput)
}

type TunnelInstanceIAMBindingPtrOutput struct {
	*pulumi.OutputState
}

func (TunnelInstanceIAMBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelInstanceIAMBinding)(nil))
}

func (o TunnelInstanceIAMBindingPtrOutput) ToTunnelInstanceIAMBindingPtrOutput() TunnelInstanceIAMBindingPtrOutput {
	return o
}

func (o TunnelInstanceIAMBindingPtrOutput) ToTunnelInstanceIAMBindingPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingPtrOutput {
	return o
}

type TunnelInstanceIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TunnelInstanceIAMBinding)(nil))
}

func (o TunnelInstanceIAMBindingArrayOutput) ToTunnelInstanceIAMBindingArrayOutput() TunnelInstanceIAMBindingArrayOutput {
	return o
}

func (o TunnelInstanceIAMBindingArrayOutput) ToTunnelInstanceIAMBindingArrayOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingArrayOutput {
	return o
}

func (o TunnelInstanceIAMBindingArrayOutput) Index(i pulumi.IntInput) TunnelInstanceIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TunnelInstanceIAMBinding {
		return vs[0].([]TunnelInstanceIAMBinding)[vs[1].(int)]
	}).(TunnelInstanceIAMBindingOutput)
}

type TunnelInstanceIAMBindingMapOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TunnelInstanceIAMBinding)(nil))
}

func (o TunnelInstanceIAMBindingMapOutput) ToTunnelInstanceIAMBindingMapOutput() TunnelInstanceIAMBindingMapOutput {
	return o
}

func (o TunnelInstanceIAMBindingMapOutput) ToTunnelInstanceIAMBindingMapOutputWithContext(ctx context.Context) TunnelInstanceIAMBindingMapOutput {
	return o
}

func (o TunnelInstanceIAMBindingMapOutput) MapIndex(k pulumi.StringInput) TunnelInstanceIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TunnelInstanceIAMBinding {
		return vs[0].(map[string]TunnelInstanceIAMBinding)[vs[1].(string)]
	}).(TunnelInstanceIAMBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(TunnelInstanceIAMBindingOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMBindingPtrOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMBindingMapOutput{})
}
