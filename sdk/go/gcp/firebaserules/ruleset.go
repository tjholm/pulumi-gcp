// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firebaserules

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// For more information, see:
// * [Get started with Firebase Security Rules](https://firebase.google.com/docs/rules/get-started)
// ## Example Usage
// ### Basic_ruleset
// Creates a basic Firestore ruleset
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebaserules"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firebaserules.NewRuleset(ctx, "primary", &firebaserules.RulesetArgs{
//				Project: pulumi.String("my-project-name"),
//				Source: &firebaserules.RulesetSourceArgs{
//					Files: firebaserules.RulesetSourceFileArray{
//						&firebaserules.RulesetSourceFileArgs{
//							Content:     pulumi.String("service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }"),
//							Fingerprint: pulumi.String(""),
//							Name:        pulumi.String("firestore.rules"),
//						},
//					},
//					Language: pulumi.String(""),
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
// ### Minimal_ruleset
// Creates a minimal Firestore ruleset
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebaserules"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firebaserules.NewRuleset(ctx, "primary", &firebaserules.RulesetArgs{
//				Project: pulumi.String("my-project-name"),
//				Source: &firebaserules.RulesetSourceArgs{
//					Files: firebaserules.RulesetSourceFileArray{
//						&firebaserules.RulesetSourceFileArgs{
//							Content: pulumi.String("service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }"),
//							Name:    pulumi.String("firestore.rules"),
//						},
//					},
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
// ## Import
//
// # Ruleset can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:firebaserules/ruleset:Ruleset default projects/{{project}}/rulesets/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firebaserules/ruleset:Ruleset default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firebaserules/ruleset:Ruleset default {{name}}
//
// ```
type Ruleset struct {
	pulumi.CustomResourceState

	// Output only. Time the `Ruleset` was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. The metadata for this ruleset.
	Metadatas RulesetMetadataArrayOutput `pulumi:"metadatas"`
	// File name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// `Source` for the `Ruleset`.
	Source RulesetSourceOutput `pulumi:"source"`
}

// NewRuleset registers a new resource with the given unique name, arguments, and options.
func NewRuleset(ctx *pulumi.Context,
	name string, args *RulesetArgs, opts ...pulumi.ResourceOption) (*Ruleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource Ruleset
	err := ctx.RegisterResource("gcp:firebaserules/ruleset:Ruleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleset gets an existing Ruleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RulesetState, opts ...pulumi.ResourceOption) (*Ruleset, error) {
	var resource Ruleset
	err := ctx.ReadResource("gcp:firebaserules/ruleset:Ruleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ruleset resources.
type rulesetState struct {
	// Output only. Time the `Ruleset` was created.
	CreateTime *string `pulumi:"createTime"`
	// Output only. The metadata for this ruleset.
	Metadatas []RulesetMetadata `pulumi:"metadatas"`
	// File name.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// `Source` for the `Ruleset`.
	Source *RulesetSource `pulumi:"source"`
}

type RulesetState struct {
	// Output only. Time the `Ruleset` was created.
	CreateTime pulumi.StringPtrInput
	// Output only. The metadata for this ruleset.
	Metadatas RulesetMetadataArrayInput
	// File name.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// `Source` for the `Ruleset`.
	Source RulesetSourcePtrInput
}

func (RulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesetState)(nil)).Elem()
}

type rulesetArgs struct {
	// The project for the resource
	Project *string `pulumi:"project"`
	// `Source` for the `Ruleset`.
	Source RulesetSource `pulumi:"source"`
}

// The set of arguments for constructing a Ruleset resource.
type RulesetArgs struct {
	// The project for the resource
	Project pulumi.StringPtrInput
	// `Source` for the `Ruleset`.
	Source RulesetSourceInput
}

func (RulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesetArgs)(nil)).Elem()
}

type RulesetInput interface {
	pulumi.Input

	ToRulesetOutput() RulesetOutput
	ToRulesetOutputWithContext(ctx context.Context) RulesetOutput
}

func (*Ruleset) ElementType() reflect.Type {
	return reflect.TypeOf((**Ruleset)(nil)).Elem()
}

func (i *Ruleset) ToRulesetOutput() RulesetOutput {
	return i.ToRulesetOutputWithContext(context.Background())
}

func (i *Ruleset) ToRulesetOutputWithContext(ctx context.Context) RulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesetOutput)
}

// RulesetArrayInput is an input type that accepts RulesetArray and RulesetArrayOutput values.
// You can construct a concrete instance of `RulesetArrayInput` via:
//
//	RulesetArray{ RulesetArgs{...} }
type RulesetArrayInput interface {
	pulumi.Input

	ToRulesetArrayOutput() RulesetArrayOutput
	ToRulesetArrayOutputWithContext(context.Context) RulesetArrayOutput
}

type RulesetArray []RulesetInput

func (RulesetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ruleset)(nil)).Elem()
}

func (i RulesetArray) ToRulesetArrayOutput() RulesetArrayOutput {
	return i.ToRulesetArrayOutputWithContext(context.Background())
}

func (i RulesetArray) ToRulesetArrayOutputWithContext(ctx context.Context) RulesetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesetArrayOutput)
}

// RulesetMapInput is an input type that accepts RulesetMap and RulesetMapOutput values.
// You can construct a concrete instance of `RulesetMapInput` via:
//
//	RulesetMap{ "key": RulesetArgs{...} }
type RulesetMapInput interface {
	pulumi.Input

	ToRulesetMapOutput() RulesetMapOutput
	ToRulesetMapOutputWithContext(context.Context) RulesetMapOutput
}

type RulesetMap map[string]RulesetInput

func (RulesetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ruleset)(nil)).Elem()
}

func (i RulesetMap) ToRulesetMapOutput() RulesetMapOutput {
	return i.ToRulesetMapOutputWithContext(context.Background())
}

func (i RulesetMap) ToRulesetMapOutputWithContext(ctx context.Context) RulesetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesetMapOutput)
}

type RulesetOutput struct{ *pulumi.OutputState }

func (RulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ruleset)(nil)).Elem()
}

func (o RulesetOutput) ToRulesetOutput() RulesetOutput {
	return o
}

func (o RulesetOutput) ToRulesetOutputWithContext(ctx context.Context) RulesetOutput {
	return o
}

// Output only. Time the `Ruleset` was created.
func (o RulesetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Ruleset) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. The metadata for this ruleset.
func (o RulesetOutput) Metadatas() RulesetMetadataArrayOutput {
	return o.ApplyT(func(v *Ruleset) RulesetMetadataArrayOutput { return v.Metadatas }).(RulesetMetadataArrayOutput)
}

// File name.
func (o RulesetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ruleset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o RulesetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Ruleset) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// `Source` for the `Ruleset`.
func (o RulesetOutput) Source() RulesetSourceOutput {
	return o.ApplyT(func(v *Ruleset) RulesetSourceOutput { return v.Source }).(RulesetSourceOutput)
}

type RulesetArrayOutput struct{ *pulumi.OutputState }

func (RulesetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ruleset)(nil)).Elem()
}

func (o RulesetArrayOutput) ToRulesetArrayOutput() RulesetArrayOutput {
	return o
}

func (o RulesetArrayOutput) ToRulesetArrayOutputWithContext(ctx context.Context) RulesetArrayOutput {
	return o
}

func (o RulesetArrayOutput) Index(i pulumi.IntInput) RulesetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ruleset {
		return vs[0].([]*Ruleset)[vs[1].(int)]
	}).(RulesetOutput)
}

type RulesetMapOutput struct{ *pulumi.OutputState }

func (RulesetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ruleset)(nil)).Elem()
}

func (o RulesetMapOutput) ToRulesetMapOutput() RulesetMapOutput {
	return o
}

func (o RulesetMapOutput) ToRulesetMapOutputWithContext(ctx context.Context) RulesetMapOutput {
	return o
}

func (o RulesetMapOutput) MapIndex(k pulumi.StringInput) RulesetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ruleset {
		return vs[0].(map[string]*Ruleset)[vs[1].(string)]
	}).(RulesetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RulesetInput)(nil)).Elem(), &Ruleset{})
	pulumi.RegisterInputType(reflect.TypeOf((*RulesetArrayInput)(nil)).Elem(), RulesetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RulesetMapInput)(nil)).Elem(), RulesetMap{})
	pulumi.RegisterOutputType(RulesetOutput{})
	pulumi.RegisterOutputType(RulesetArrayOutput{})
	pulumi.RegisterOutputType(RulesetMapOutput{})
}
