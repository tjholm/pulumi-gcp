// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificatemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Certificate Manager Certificate Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificatemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := certificatemanager.NewDnsAuthorization(ctx, "instance", &certificatemanager.DnsAuthorizationArgs{
//				Description: pulumi.String("The default dnss"),
//				Domain:      pulumi.String("subdomain.hashicorptest.com"),
//			})
//			if err != nil {
//				return err
//			}
//			instance2, err := certificatemanager.NewDnsAuthorization(ctx, "instance2", &certificatemanager.DnsAuthorizationArgs{
//				Description: pulumi.String("The default dnss"),
//				Domain:      pulumi.String("subdomain2.hashicorptest.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = certificatemanager.NewCertificate(ctx, "default", &certificatemanager.CertificateArgs{
//				Description: pulumi.String("The default cert"),
//				Scope:       pulumi.String("EDGE_CACHE"),
//				Managed: &certificatemanager.CertificateManagedArgs{
//					Domains: pulumi.StringArray{
//						instance.Domain,
//						instance2.Domain,
//					},
//					DnsAuthorizations: pulumi.StringArray{
//						instance.ID(),
//						instance2.ID(),
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
// # Certificate can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificate:Certificate default projects/{{project}}/locations/global/certificates/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificate:Certificate default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:certificatemanager/certificate:Certificate default {{name}}
//
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// A human-readable description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Set of label tags associated with the EdgeCache resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed CertificateManagedPtrOutput `pulumi:"managed"`
	// A user-defined name of the certificate. Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The scope of the certificate.
	// Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `EDGE_CACHE`.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged CertificateSelfManagedPtrOutput `pulumi:"selfManaged"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		args = &CertificateArgs{}
	}

	var resource Certificate
	err := ctx.RegisterResource("gcp:certificatemanager/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("gcp:certificatemanager/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the EdgeCache resource.
	Labels map[string]string `pulumi:"labels"`
	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed *CertificateManaged `pulumi:"managed"`
	// A user-defined name of the certificate. Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The scope of the certificate.
	// Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `EDGE_CACHE`.
	Scope *string `pulumi:"scope"`
	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged *CertificateSelfManaged `pulumi:"selfManaged"`
}

type CertificateState struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the EdgeCache resource.
	Labels pulumi.StringMapInput
	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed CertificateManagedPtrInput
	// A user-defined name of the certificate. Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The scope of the certificate.
	// Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `EDGE_CACHE`.
	Scope pulumi.StringPtrInput
	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged CertificateSelfManagedPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the EdgeCache resource.
	Labels map[string]string `pulumi:"labels"`
	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed *CertificateManaged `pulumi:"managed"`
	// A user-defined name of the certificate. Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The scope of the certificate.
	// Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `EDGE_CACHE`.
	Scope *string `pulumi:"scope"`
	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged *CertificateSelfManaged `pulumi:"selfManaged"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the EdgeCache resource.
	Labels pulumi.StringMapInput
	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed CertificateManagedPtrInput
	// A user-defined name of the certificate. Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The scope of the certificate.
	// Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `EDGE_CACHE`.
	Scope pulumi.StringPtrInput
	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged CertificateSelfManagedPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// A human-readable description of the resource.
func (o CertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Set of label tags associated with the EdgeCache resource.
func (o CertificateOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Configuration and state of a Managed Certificate.
// Certificate Manager provisions and renews Managed Certificates
// automatically, for as long as it's authorized to do so.
// Structure is documented below.
func (o CertificateOutput) Managed() CertificateManagedPtrOutput {
	return o.ApplyT(func(v *Certificate) CertificateManagedPtrOutput { return v.Managed }).(CertificateManagedPtrOutput)
}

// A user-defined name of the certificate. Certificate names must be unique
// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
// and all following characters must be a dash, underscore, letter or digit.
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o CertificateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The scope of the certificate.
// Certificates with default scope are served from core Google data centers.
// If unsure, choose this option.
// Certificates with scope EDGE_CACHE are special-purposed certificates,
// served from non-core Google data centers.
// Currently allowed only for managed certificates.
// Default value is `DEFAULT`.
// Possible values are `DEFAULT` and `EDGE_CACHE`.
func (o CertificateOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Certificate data for a SelfManaged Certificate.
// SelfManaged Certificates are uploaded by the user. Updating such
// certificates before they expire remains the user's responsibility.
// Structure is documented below.
func (o CertificateOutput) SelfManaged() CertificateSelfManagedPtrOutput {
	return o.ApplyT(func(v *Certificate) CertificateSelfManagedPtrOutput { return v.SelfManaged }).(CertificateSelfManagedPtrOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
