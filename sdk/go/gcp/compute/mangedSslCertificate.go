// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An SslCertificate resource, used for HTTPS load balancing.  This resource
// represents a certificate for which the certificate secrets are created and
// managed by Google.
//
// For a resource where you provide the key, see the
// SSL Certificate resource.
//
// To get more information about ManagedSslCertificate, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
//
// > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
// certificate is complex.  Ensure that you understand the lifecycle of a
// certificate before attempting complex tasks like cert rotation automatically.
// This resource will "return" as soon as the certificate object is created,
// but post-creation the certificate object will go through a "provisioning"
// process.  The provisioning process can complete only when the domain name
// for which the certificate is created points to a target pool which, itself,
// points at the certificate.  Depending on your DNS provider, this may take
// some time, and migrating from self-managed certificates to Google-managed
// certificates may entail some downtime while the certificate provisions.
//
// In conclusion: Be extremely cautious.
//
// ## Example Usage
//
// ## Import
//
// ManagedSslCertificate can be imported using any of these accepted formats* `projects/{{project}}/global/sslCertificates/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, ManagedSslCertificate can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/mangedSslCertificate:MangedSslCertificate default projects/{{project}}/global/sslCertificates/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/mangedSslCertificate:MangedSslCertificate default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/mangedSslCertificate:MangedSslCertificate default {{name}}
//
// ```
//
// Deprecated: gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate
type MangedSslCertificate struct {
	pulumi.CustomResourceState

	// The unique identifier for the resource.
	CertificateId pulumi.IntOutput `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Expire time of the certificate in RFC3339 text format.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed MangedSslCertificateManagedPtrOutput `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// These are in the same namespace as the managed SSL certificates.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are: `MANAGED`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewMangedSslCertificate registers a new resource with the given unique name, arguments, and options.
func NewMangedSslCertificate(ctx *pulumi.Context,
	name string, args *MangedSslCertificateArgs, opts ...pulumi.ResourceOption) (*MangedSslCertificate, error) {
	if args == nil {
		args = &MangedSslCertificateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MangedSslCertificate
	err := ctx.RegisterResource("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMangedSslCertificate gets an existing MangedSslCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMangedSslCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MangedSslCertificateState, opts ...pulumi.ResourceOption) (*MangedSslCertificate, error) {
	var resource MangedSslCertificate
	err := ctx.ReadResource("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MangedSslCertificate resources.
type mangedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Expire time of the certificate in RFC3339 text format.
	ExpireTime *string `pulumi:"expireTime"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed *MangedSslCertificateManaged `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// These are in the same namespace as the managed SSL certificates.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are: `MANAGED`.
	Type *string `pulumi:"type"`
}

type MangedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Expire time of the certificate in RFC3339 text format.
	ExpireTime pulumi.StringPtrInput
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed MangedSslCertificateManagedPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// These are in the same namespace as the managed SSL certificates.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayInput
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are: `MANAGED`.
	Type pulumi.StringPtrInput
}

func (MangedSslCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mangedSslCertificateState)(nil)).Elem()
}

type mangedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed *MangedSslCertificateManaged `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// These are in the same namespace as the managed SSL certificates.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are: `MANAGED`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a MangedSslCertificate resource.
type MangedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed MangedSslCertificateManagedPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// These are in the same namespace as the managed SSL certificates.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are: `MANAGED`.
	Type pulumi.StringPtrInput
}

func (MangedSslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mangedSslCertificateArgs)(nil)).Elem()
}

type MangedSslCertificateInput interface {
	pulumi.Input

	ToMangedSslCertificateOutput() MangedSslCertificateOutput
	ToMangedSslCertificateOutputWithContext(ctx context.Context) MangedSslCertificateOutput
}

func (*MangedSslCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**MangedSslCertificate)(nil)).Elem()
}

func (i *MangedSslCertificate) ToMangedSslCertificateOutput() MangedSslCertificateOutput {
	return i.ToMangedSslCertificateOutputWithContext(context.Background())
}

func (i *MangedSslCertificate) ToMangedSslCertificateOutputWithContext(ctx context.Context) MangedSslCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MangedSslCertificateOutput)
}

// MangedSslCertificateArrayInput is an input type that accepts MangedSslCertificateArray and MangedSslCertificateArrayOutput values.
// You can construct a concrete instance of `MangedSslCertificateArrayInput` via:
//
//	MangedSslCertificateArray{ MangedSslCertificateArgs{...} }
type MangedSslCertificateArrayInput interface {
	pulumi.Input

	ToMangedSslCertificateArrayOutput() MangedSslCertificateArrayOutput
	ToMangedSslCertificateArrayOutputWithContext(context.Context) MangedSslCertificateArrayOutput
}

type MangedSslCertificateArray []MangedSslCertificateInput

func (MangedSslCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MangedSslCertificate)(nil)).Elem()
}

func (i MangedSslCertificateArray) ToMangedSslCertificateArrayOutput() MangedSslCertificateArrayOutput {
	return i.ToMangedSslCertificateArrayOutputWithContext(context.Background())
}

func (i MangedSslCertificateArray) ToMangedSslCertificateArrayOutputWithContext(ctx context.Context) MangedSslCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MangedSslCertificateArrayOutput)
}

// MangedSslCertificateMapInput is an input type that accepts MangedSslCertificateMap and MangedSslCertificateMapOutput values.
// You can construct a concrete instance of `MangedSslCertificateMapInput` via:
//
//	MangedSslCertificateMap{ "key": MangedSslCertificateArgs{...} }
type MangedSslCertificateMapInput interface {
	pulumi.Input

	ToMangedSslCertificateMapOutput() MangedSslCertificateMapOutput
	ToMangedSslCertificateMapOutputWithContext(context.Context) MangedSslCertificateMapOutput
}

type MangedSslCertificateMap map[string]MangedSslCertificateInput

func (MangedSslCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MangedSslCertificate)(nil)).Elem()
}

func (i MangedSslCertificateMap) ToMangedSslCertificateMapOutput() MangedSslCertificateMapOutput {
	return i.ToMangedSslCertificateMapOutputWithContext(context.Background())
}

func (i MangedSslCertificateMap) ToMangedSslCertificateMapOutputWithContext(ctx context.Context) MangedSslCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MangedSslCertificateMapOutput)
}

type MangedSslCertificateOutput struct{ *pulumi.OutputState }

func (MangedSslCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MangedSslCertificate)(nil)).Elem()
}

func (o MangedSslCertificateOutput) ToMangedSslCertificateOutput() MangedSslCertificateOutput {
	return o
}

func (o MangedSslCertificateOutput) ToMangedSslCertificateOutputWithContext(ctx context.Context) MangedSslCertificateOutput {
	return o
}

// The unique identifier for the resource.
func (o MangedSslCertificateOutput) CertificateId() pulumi.IntOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.IntOutput { return v.CertificateId }).(pulumi.IntOutput)
}

// Creation timestamp in RFC3339 text format.
func (o MangedSslCertificateOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o MangedSslCertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Expire time of the certificate in RFC3339 text format.
func (o MangedSslCertificateOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// Properties relevant to a managed certificate.  These will be used if the
// certificate is managed (as indicated by a value of `MANAGED` in `type`).
// Structure is documented below.
func (o MangedSslCertificateOutput) Managed() MangedSslCertificateManagedPtrOutput {
	return o.ApplyT(func(v *MangedSslCertificate) MangedSslCertificateManagedPtrOutput { return v.Managed }).(MangedSslCertificateManagedPtrOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
//
// These are in the same namespace as the managed SSL certificates.
func (o MangedSslCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o MangedSslCertificateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o MangedSslCertificateOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Domains associated with the certificate via Subject Alternative Name.
func (o MangedSslCertificateOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringArrayOutput { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

// Enum field whose value is always `MANAGED` - used to signal to the API
// which type this is.
// Default value is `MANAGED`.
// Possible values are: `MANAGED`.
func (o MangedSslCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MangedSslCertificate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type MangedSslCertificateArrayOutput struct{ *pulumi.OutputState }

func (MangedSslCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MangedSslCertificate)(nil)).Elem()
}

func (o MangedSslCertificateArrayOutput) ToMangedSslCertificateArrayOutput() MangedSslCertificateArrayOutput {
	return o
}

func (o MangedSslCertificateArrayOutput) ToMangedSslCertificateArrayOutputWithContext(ctx context.Context) MangedSslCertificateArrayOutput {
	return o
}

func (o MangedSslCertificateArrayOutput) Index(i pulumi.IntInput) MangedSslCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MangedSslCertificate {
		return vs[0].([]*MangedSslCertificate)[vs[1].(int)]
	}).(MangedSslCertificateOutput)
}

type MangedSslCertificateMapOutput struct{ *pulumi.OutputState }

func (MangedSslCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MangedSslCertificate)(nil)).Elem()
}

func (o MangedSslCertificateMapOutput) ToMangedSslCertificateMapOutput() MangedSslCertificateMapOutput {
	return o
}

func (o MangedSslCertificateMapOutput) ToMangedSslCertificateMapOutputWithContext(ctx context.Context) MangedSslCertificateMapOutput {
	return o
}

func (o MangedSslCertificateMapOutput) MapIndex(k pulumi.StringInput) MangedSslCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MangedSslCertificate {
		return vs[0].(map[string]*MangedSslCertificate)[vs[1].(string)]
	}).(MangedSslCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MangedSslCertificateInput)(nil)).Elem(), &MangedSslCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MangedSslCertificateArrayInput)(nil)).Elem(), MangedSslCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MangedSslCertificateMapInput)(nil)).Elem(), MangedSslCertificateMap{})
	pulumi.RegisterOutputType(MangedSslCertificateOutput{})
	pulumi.RegisterOutputType(MangedSslCertificateArrayOutput{})
	pulumi.RegisterOutputType(MangedSslCertificateMapOutput{})
}
