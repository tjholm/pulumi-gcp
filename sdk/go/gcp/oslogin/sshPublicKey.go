// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oslogin

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The SSH public key information associated with a Google account.
//
// To get more information about SSHPublicKey, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest/v1/users.sshPublicKeys)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/oslogin)
//
// ## Example Usage
// ### Os Login Ssh Key Basic
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/oslogin"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			me, err := organizations.GetClientOpenIdUserInfo(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = oslogin.NewSshPublicKey(ctx, "cache", &oslogin.SshPublicKeyArgs{
//				User: *pulumi.String(me.Email),
//				Key:  readFileOrPanic("path/to/id_rsa.pub"),
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
// # SSHPublicKey can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default users/{{user}}/sshPublicKeys/{{fingerprint}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default {{user}}/{{fingerprint}}
//
// ```
type SshPublicKey struct {
	pulumi.CustomResourceState

	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrOutput `pulumi:"expirationTimeUsec"`
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringOutput `pulumi:"key"`
	// The project ID of the Google Cloud Platform project.
	Project pulumi.StringPtrOutput `pulumi:"project"`
	// The user email.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewSshPublicKey registers a new resource with the given unique name, arguments, and options.
func NewSshPublicKey(ctx *pulumi.Context,
	name string, args *SshPublicKeyArgs, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource SshPublicKey
	err := ctx.RegisterResource("gcp:oslogin/sshPublicKey:SshPublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshPublicKey gets an existing SshPublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshPublicKeyState, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	var resource SshPublicKey
	err := ctx.ReadResource("gcp:oslogin/sshPublicKey:SshPublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshPublicKey resources.
type sshPublicKeyState struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec *string `pulumi:"expirationTimeUsec"`
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint *string `pulumi:"fingerprint"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key *string `pulumi:"key"`
	// The project ID of the Google Cloud Platform project.
	Project *string `pulumi:"project"`
	// The user email.
	User *string `pulumi:"user"`
}

type SshPublicKeyState struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrInput
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint pulumi.StringPtrInput
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringPtrInput
	// The project ID of the Google Cloud Platform project.
	Project pulumi.StringPtrInput
	// The user email.
	User pulumi.StringPtrInput
}

func (SshPublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyState)(nil)).Elem()
}

type sshPublicKeyArgs struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec *string `pulumi:"expirationTimeUsec"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `pulumi:"key"`
	// The project ID of the Google Cloud Platform project.
	Project *string `pulumi:"project"`
	// The user email.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a SshPublicKey resource.
type SshPublicKeyArgs struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrInput
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringInput
	// The project ID of the Google Cloud Platform project.
	Project pulumi.StringPtrInput
	// The user email.
	User pulumi.StringInput
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyArgs)(nil)).Elem()
}

type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput
}

func (*SshPublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (i *SshPublicKey) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i *SshPublicKey) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}

// SshPublicKeyArrayInput is an input type that accepts SshPublicKeyArray and SshPublicKeyArrayOutput values.
// You can construct a concrete instance of `SshPublicKeyArrayInput` via:
//
//	SshPublicKeyArray{ SshPublicKeyArgs{...} }
type SshPublicKeyArrayInput interface {
	pulumi.Input

	ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput
	ToSshPublicKeyArrayOutputWithContext(context.Context) SshPublicKeyArrayOutput
}

type SshPublicKeyArray []SshPublicKeyInput

func (SshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return i.ToSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyArrayOutput)
}

// SshPublicKeyMapInput is an input type that accepts SshPublicKeyMap and SshPublicKeyMapOutput values.
// You can construct a concrete instance of `SshPublicKeyMapInput` via:
//
//	SshPublicKeyMap{ "key": SshPublicKeyArgs{...} }
type SshPublicKeyMapInput interface {
	pulumi.Input

	ToSshPublicKeyMapOutput() SshPublicKeyMapOutput
	ToSshPublicKeyMapOutputWithContext(context.Context) SshPublicKeyMapOutput
}

type SshPublicKeyMap map[string]SshPublicKeyInput

func (SshPublicKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyMap) ToSshPublicKeyMapOutput() SshPublicKeyMapOutput {
	return i.ToSshPublicKeyMapOutputWithContext(context.Background())
}

func (i SshPublicKeyMap) ToSshPublicKeyMapOutputWithContext(ctx context.Context) SshPublicKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyMapOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

// An expiration time in microseconds since epoch.
func (o SshPublicKeyOutput) ExpirationTimeUsec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringPtrOutput { return v.ExpirationTimeUsec }).(pulumi.StringPtrOutput)
}

// The SHA-256 fingerprint of the SSH public key.
func (o SshPublicKeyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Public key text in SSH format, defined by RFC4253 section 6.6.
func (o SshPublicKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The project ID of the Google Cloud Platform project.
func (o SshPublicKeyOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringPtrOutput { return v.Project }).(pulumi.StringPtrOutput)
}

// The user email.
func (o SshPublicKeyOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type SshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) Index(i pulumi.IntInput) SshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SshPublicKey {
		return vs[0].([]*SshPublicKey)[vs[1].(int)]
	}).(SshPublicKeyOutput)
}

type SshPublicKeyMapOutput struct{ *pulumi.OutputState }

func (SshPublicKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyMapOutput) ToSshPublicKeyMapOutput() SshPublicKeyMapOutput {
	return o
}

func (o SshPublicKeyMapOutput) ToSshPublicKeyMapOutputWithContext(ctx context.Context) SshPublicKeyMapOutput {
	return o
}

func (o SshPublicKeyMapOutput) MapIndex(k pulumi.StringInput) SshPublicKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SshPublicKey {
		return vs[0].(map[string]*SshPublicKey)[vs[1].(string)]
	}).(SshPublicKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SshPublicKeyInput)(nil)).Elem(), &SshPublicKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshPublicKeyArrayInput)(nil)).Elem(), SshPublicKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshPublicKeyMapInput)(nil)).Elem(), SshPublicKeyMap{})
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
	pulumi.RegisterOutputType(SshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(SshPublicKeyMapOutput{})
}
