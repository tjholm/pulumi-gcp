// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A CertificateAuthority represents an individual Certificate Authority. A
// CertificateAuthority can be used to create Certificates.
//
// > **Warning:** Please remember that all resources created during preview (via this provider)
// will be deleted when CA service transitions to General Availability (GA). Relying on these
// certificate authorities for production traffic is discouraged.
//
// To get more information about CertificateAuthority, see:
//
// * [API documentation](https://cloud.google.com/certificate-authority-service/docs/reference/rest)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/certificate-authority-service)
//
// ## Example Usage
// ### Privateca Certificate Authority Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewAuthority(ctx, "_default", &certificateauthority.AuthorityArgs{
// 			CertificateAuthorityId: pulumi.String("my-certificate-authority"),
// 			Location:               pulumi.String("us-central1"),
// 			Config: &certificateauthority.AuthorityConfigArgs{
// 				SubjectConfig: &certificateauthority.AuthorityConfigSubjectConfigArgs{
// 					Subject: &certificateauthority.AuthorityConfigSubjectConfigSubjectArgs{
// 						Organization: pulumi.String("HashiCorp"),
// 					},
// 					CommonName: pulumi.String("my-certificate-authority"),
// 					SubjectAltName: &certificateauthority.AuthorityConfigSubjectConfigSubjectAltNameArgs{
// 						DnsNames: pulumi.StringArray{
// 							pulumi.String("hashicorp.com"),
// 						},
// 					},
// 				},
// 				ReusableConfig: &certificateauthority.AuthorityConfigReusableConfigArgs{
// 					ReusableConfig: pulumi.String("projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"),
// 				},
// 			},
// 			KeySpec: &certificateauthority.AuthorityKeySpecArgs{
// 				Algorithm: pulumi.String("RSA_PKCS1_4096_SHA256"),
// 			},
// 			DisableOnDelete: pulumi.Bool(true),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Privateca Certificate Authority Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewAuthority(ctx, "_default", &certificateauthority.AuthorityArgs{
// 			CertificateAuthorityId: pulumi.String("my-certificate-authority"),
// 			Location:               pulumi.String("us-central1"),
// 			Tier:                   pulumi.String("DEVOPS"),
// 			Config: &certificateauthority.AuthorityConfigArgs{
// 				SubjectConfig: &certificateauthority.AuthorityConfigSubjectConfigArgs{
// 					Subject: &certificateauthority.AuthorityConfigSubjectConfigSubjectArgs{
// 						CountryCode:        pulumi.String("US"),
// 						Organization:       pulumi.String("HashiCorp"),
// 						OrganizationalUnit: pulumi.String("Terraform"),
// 						Locality:           pulumi.String("San Francisco"),
// 						Province:           pulumi.String("CA"),
// 						StreetAddress:      pulumi.String("101 2nd St #700"),
// 						PostalCode:         pulumi.String("94105"),
// 					},
// 					CommonName: pulumi.String("my-certificate-authority"),
// 					SubjectAltName: &certificateauthority.AuthorityConfigSubjectConfigSubjectAltNameArgs{
// 						DnsNames: pulumi.StringArray{
// 							pulumi.String("hashicorp.com"),
// 						},
// 						EmailAddresses: pulumi.StringArray{
// 							pulumi.String("email@example.com"),
// 						},
// 						IpAddresses: pulumi.StringArray{
// 							pulumi.String("127.0.0.1"),
// 						},
// 						Uris: pulumi.StringArray{
// 							pulumi.String("http://www.ietf.org/rfc/rfc3986.txt"),
// 						},
// 					},
// 				},
// 				ReusableConfig: &certificateauthority.AuthorityConfigReusableConfigArgs{
// 					ReusableConfig: pulumi.String("projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"),
// 				},
// 			},
// 			Lifetime: pulumi.String("86400s"),
// 			IssuingOptions: &certificateauthority.AuthorityIssuingOptionsArgs{
// 				IncludeCaCertUrl:    pulumi.Bool(true),
// 				IncludeCrlAccessUrl: pulumi.Bool(false),
// 			},
// 			KeySpec: &certificateauthority.AuthorityKeySpecArgs{
// 				Algorithm: pulumi.String("EC_P256_SHA256"),
// 			},
// 			DisableOnDelete: pulumi.Bool(true),
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
// CertificateAuthority can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:certificateauthority/authority:Authority default projects/{{project}}/locations/{{location}}/certificateAuthorities/{{certificate_authority_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:certificateauthority/authority:Authority default {{project}}/{{location}}/{{certificate_authority_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:certificateauthority/authority:Authority default {{location}}/{{certificate_authority_id}}
// ```
type Authority struct {
	pulumi.CustomResourceState

	// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
	AccessUrls AuthorityAccessUrlArrayOutput `pulumi:"accessUrls"`
	// The user provided Resource ID for this Certificate Authority.
	CertificateAuthorityId pulumi.StringOutput `pulumi:"certificateAuthorityId"`
	// The config used to create a self-signed X.509 certificate or CSR.
	// Structure is documented below.
	Config AuthorityConfigOutput `pulumi:"config"`
	// The time at which this CertificateAuthority was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// If set to `true`, the Certificate Authority will be disabled
	// on delete. If the Certitificate Authorities is not disabled,
	// it cannot be deleted. Use with care. Defaults to `false`.
	DisableOnDelete pulumi.BoolPtrOutput `pulumi:"disableOnDelete"`
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
	// such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
	// (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
	// my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
	// created.
	GcsBucket pulumi.StringPtrOutput `pulumi:"gcsBucket"`
	// Options that affect all certificates issued by a CertificateAuthority.
	// Structure is documented below.
	IssuingOptions AuthorityIssuingOptionsPtrOutput `pulumi:"issuingOptions"`
	// Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
	// is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
	// certificate. Otherwise, it is used to sign a CSR.
	// Structure is documented below.
	KeySpec AuthorityKeySpecOutput `pulumi:"keySpec"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	Lifetime pulumi.StringPtrOutput `pulumi:"lifetime"`
	// Location of the CertificateAuthority. A full list of valid locations can be found by
	// running `gcloud beta privateca locations list`.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for this CertificateAuthority in the format projects/*/locations/*/certificateAuthorities/*.
	Name pulumi.StringOutput `pulumi:"name"`
	// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such
	// that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the
	// current CertificateAuthority's certificate.
	PemCaCertificates pulumi.StringArrayOutput `pulumi:"pemCaCertificates"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The State for this CertificateAuthority.
	State pulumi.StringOutput `pulumi:"state"`
	// The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
	// server side certificates issued, and support certificate revocation. For more details,
	// please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
	// Default value is `ENTERPRISE`.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// The Type of this CertificateAuthority.
	// > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
	// be manually activated (via Cloud Console of `gcloud`) before they can
	// issue certificates.
	// Default value is `SELF_SIGNED`.
	// Possible values are `SELF_SIGNED` and `SUBORDINATE`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The time at which this CertificateAuthority was updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAuthority registers a new resource with the given unique name, arguments, and options.
func NewAuthority(ctx *pulumi.Context,
	name string, args *AuthorityArgs, opts ...pulumi.ResourceOption) (*Authority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthorityId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityId'")
	}
	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.KeySpec == nil {
		return nil, errors.New("invalid value for required argument 'KeySpec'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Authority
	err := ctx.RegisterResource("gcp:certificateauthority/authority:Authority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthority gets an existing Authority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorityState, opts ...pulumi.ResourceOption) (*Authority, error) {
	var resource Authority
	err := ctx.ReadResource("gcp:certificateauthority/authority:Authority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Authority resources.
type authorityState struct {
	// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
	AccessUrls []AuthorityAccessUrl `pulumi:"accessUrls"`
	// The user provided Resource ID for this Certificate Authority.
	CertificateAuthorityId *string `pulumi:"certificateAuthorityId"`
	// The config used to create a self-signed X.509 certificate or CSR.
	// Structure is documented below.
	Config *AuthorityConfig `pulumi:"config"`
	// The time at which this CertificateAuthority was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// If set to `true`, the Certificate Authority will be disabled
	// on delete. If the Certitificate Authorities is not disabled,
	// it cannot be deleted. Use with care. Defaults to `false`.
	DisableOnDelete *bool `pulumi:"disableOnDelete"`
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
	// such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
	// (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
	// my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
	// created.
	GcsBucket *string `pulumi:"gcsBucket"`
	// Options that affect all certificates issued by a CertificateAuthority.
	// Structure is documented below.
	IssuingOptions *AuthorityIssuingOptions `pulumi:"issuingOptions"`
	// Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
	// is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
	// certificate. Otherwise, it is used to sign a CSR.
	// Structure is documented below.
	KeySpec *AuthorityKeySpec `pulumi:"keySpec"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	Lifetime *string `pulumi:"lifetime"`
	// Location of the CertificateAuthority. A full list of valid locations can be found by
	// running `gcloud beta privateca locations list`.
	Location *string `pulumi:"location"`
	// The resource name for this CertificateAuthority in the format projects/*/locations/*/certificateAuthorities/*.
	Name *string `pulumi:"name"`
	// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such
	// that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the
	// current CertificateAuthority's certificate.
	PemCaCertificates []string `pulumi:"pemCaCertificates"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The State for this CertificateAuthority.
	State *string `pulumi:"state"`
	// The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
	// server side certificates issued, and support certificate revocation. For more details,
	// please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
	// Default value is `ENTERPRISE`.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
	Tier *string `pulumi:"tier"`
	// The Type of this CertificateAuthority.
	// > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
	// be manually activated (via Cloud Console of `gcloud`) before they can
	// issue certificates.
	// Default value is `SELF_SIGNED`.
	// Possible values are `SELF_SIGNED` and `SUBORDINATE`.
	Type *string `pulumi:"type"`
	// The time at which this CertificateAuthority was updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type AuthorityState struct {
	// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
	AccessUrls AuthorityAccessUrlArrayInput
	// The user provided Resource ID for this Certificate Authority.
	CertificateAuthorityId pulumi.StringPtrInput
	// The config used to create a self-signed X.509 certificate or CSR.
	// Structure is documented below.
	Config AuthorityConfigPtrInput
	// The time at which this CertificateAuthority was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// If set to `true`, the Certificate Authority will be disabled
	// on delete. If the Certitificate Authorities is not disabled,
	// it cannot be deleted. Use with care. Defaults to `false`.
	DisableOnDelete pulumi.BoolPtrInput
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
	// such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
	// (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
	// my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
	// created.
	GcsBucket pulumi.StringPtrInput
	// Options that affect all certificates issued by a CertificateAuthority.
	// Structure is documented below.
	IssuingOptions AuthorityIssuingOptionsPtrInput
	// Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
	// is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
	// certificate. Otherwise, it is used to sign a CSR.
	// Structure is documented below.
	KeySpec AuthorityKeySpecPtrInput
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	Lifetime pulumi.StringPtrInput
	// Location of the CertificateAuthority. A full list of valid locations can be found by
	// running `gcloud beta privateca locations list`.
	Location pulumi.StringPtrInput
	// The resource name for this CertificateAuthority in the format projects/*/locations/*/certificateAuthorities/*.
	Name pulumi.StringPtrInput
	// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such
	// that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the
	// current CertificateAuthority's certificate.
	PemCaCertificates pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The State for this CertificateAuthority.
	State pulumi.StringPtrInput
	// The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
	// server side certificates issued, and support certificate revocation. For more details,
	// please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
	// Default value is `ENTERPRISE`.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
	Tier pulumi.StringPtrInput
	// The Type of this CertificateAuthority.
	// > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
	// be manually activated (via Cloud Console of `gcloud`) before they can
	// issue certificates.
	// Default value is `SELF_SIGNED`.
	// Possible values are `SELF_SIGNED` and `SUBORDINATE`.
	Type pulumi.StringPtrInput
	// The time at which this CertificateAuthority was updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (AuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorityState)(nil)).Elem()
}

type authorityArgs struct {
	// The user provided Resource ID for this Certificate Authority.
	CertificateAuthorityId string `pulumi:"certificateAuthorityId"`
	// The config used to create a self-signed X.509 certificate or CSR.
	// Structure is documented below.
	Config AuthorityConfig `pulumi:"config"`
	// If set to `true`, the Certificate Authority will be disabled
	// on delete. If the Certitificate Authorities is not disabled,
	// it cannot be deleted. Use with care. Defaults to `false`.
	DisableOnDelete *bool `pulumi:"disableOnDelete"`
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
	// such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
	// (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
	// my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
	// created.
	GcsBucket *string `pulumi:"gcsBucket"`
	// Options that affect all certificates issued by a CertificateAuthority.
	// Structure is documented below.
	IssuingOptions *AuthorityIssuingOptions `pulumi:"issuingOptions"`
	// Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
	// is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
	// certificate. Otherwise, it is used to sign a CSR.
	// Structure is documented below.
	KeySpec AuthorityKeySpec `pulumi:"keySpec"`
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	Lifetime *string `pulumi:"lifetime"`
	// Location of the CertificateAuthority. A full list of valid locations can be found by
	// running `gcloud beta privateca locations list`.
	Location string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
	// server side certificates issued, and support certificate revocation. For more details,
	// please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
	// Default value is `ENTERPRISE`.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
	Tier *string `pulumi:"tier"`
	// The Type of this CertificateAuthority.
	// > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
	// be manually activated (via Cloud Console of `gcloud`) before they can
	// issue certificates.
	// Default value is `SELF_SIGNED`.
	// Possible values are `SELF_SIGNED` and `SUBORDINATE`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Authority resource.
type AuthorityArgs struct {
	// The user provided Resource ID for this Certificate Authority.
	CertificateAuthorityId pulumi.StringInput
	// The config used to create a self-signed X.509 certificate or CSR.
	// Structure is documented below.
	Config AuthorityConfigInput
	// If set to `true`, the Certificate Authority will be disabled
	// on delete. If the Certitificate Authorities is not disabled,
	// it cannot be deleted. Use with care. Defaults to `false`.
	DisableOnDelete pulumi.BoolPtrInput
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
	// such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
	// (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
	// my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
	// created.
	GcsBucket pulumi.StringPtrInput
	// Options that affect all certificates issued by a CertificateAuthority.
	// Structure is documented below.
	IssuingOptions AuthorityIssuingOptionsPtrInput
	// Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
	// is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
	// certificate. Otherwise, it is used to sign a CSR.
	// Structure is documented below.
	KeySpec AuthorityKeySpecInput
	// Labels with user-defined metadata.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	Lifetime pulumi.StringPtrInput
	// Location of the CertificateAuthority. A full list of valid locations can be found by
	// running `gcloud beta privateca locations list`.
	Location pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
	// server side certificates issued, and support certificate revocation. For more details,
	// please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
	// Default value is `ENTERPRISE`.
	// Possible values are `ENTERPRISE` and `DEVOPS`.
	Tier pulumi.StringPtrInput
	// The Type of this CertificateAuthority.
	// > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
	// be manually activated (via Cloud Console of `gcloud`) before they can
	// issue certificates.
	// Default value is `SELF_SIGNED`.
	// Possible values are `SELF_SIGNED` and `SUBORDINATE`.
	Type pulumi.StringPtrInput
}

func (AuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorityArgs)(nil)).Elem()
}

type AuthorityInput interface {
	pulumi.Input

	ToAuthorityOutput() AuthorityOutput
	ToAuthorityOutputWithContext(ctx context.Context) AuthorityOutput
}

func (*Authority) ElementType() reflect.Type {
	return reflect.TypeOf((*Authority)(nil))
}

func (i *Authority) ToAuthorityOutput() AuthorityOutput {
	return i.ToAuthorityOutputWithContext(context.Background())
}

func (i *Authority) ToAuthorityOutputWithContext(ctx context.Context) AuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityOutput)
}

func (i *Authority) ToAuthorityPtrOutput() AuthorityPtrOutput {
	return i.ToAuthorityPtrOutputWithContext(context.Background())
}

func (i *Authority) ToAuthorityPtrOutputWithContext(ctx context.Context) AuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityPtrOutput)
}

type AuthorityPtrInput interface {
	pulumi.Input

	ToAuthorityPtrOutput() AuthorityPtrOutput
	ToAuthorityPtrOutputWithContext(ctx context.Context) AuthorityPtrOutput
}

type authorityPtrType AuthorityArgs

func (*authorityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Authority)(nil))
}

func (i *authorityPtrType) ToAuthorityPtrOutput() AuthorityPtrOutput {
	return i.ToAuthorityPtrOutputWithContext(context.Background())
}

func (i *authorityPtrType) ToAuthorityPtrOutputWithContext(ctx context.Context) AuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityPtrOutput)
}

// AuthorityArrayInput is an input type that accepts AuthorityArray and AuthorityArrayOutput values.
// You can construct a concrete instance of `AuthorityArrayInput` via:
//
//          AuthorityArray{ AuthorityArgs{...} }
type AuthorityArrayInput interface {
	pulumi.Input

	ToAuthorityArrayOutput() AuthorityArrayOutput
	ToAuthorityArrayOutputWithContext(context.Context) AuthorityArrayOutput
}

type AuthorityArray []AuthorityInput

func (AuthorityArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Authority)(nil))
}

func (i AuthorityArray) ToAuthorityArrayOutput() AuthorityArrayOutput {
	return i.ToAuthorityArrayOutputWithContext(context.Background())
}

func (i AuthorityArray) ToAuthorityArrayOutputWithContext(ctx context.Context) AuthorityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityArrayOutput)
}

// AuthorityMapInput is an input type that accepts AuthorityMap and AuthorityMapOutput values.
// You can construct a concrete instance of `AuthorityMapInput` via:
//
//          AuthorityMap{ "key": AuthorityArgs{...} }
type AuthorityMapInput interface {
	pulumi.Input

	ToAuthorityMapOutput() AuthorityMapOutput
	ToAuthorityMapOutputWithContext(context.Context) AuthorityMapOutput
}

type AuthorityMap map[string]AuthorityInput

func (AuthorityMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Authority)(nil))
}

func (i AuthorityMap) ToAuthorityMapOutput() AuthorityMapOutput {
	return i.ToAuthorityMapOutputWithContext(context.Background())
}

func (i AuthorityMap) ToAuthorityMapOutputWithContext(ctx context.Context) AuthorityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityMapOutput)
}

type AuthorityOutput struct {
	*pulumi.OutputState
}

func (AuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Authority)(nil))
}

func (o AuthorityOutput) ToAuthorityOutput() AuthorityOutput {
	return o
}

func (o AuthorityOutput) ToAuthorityOutputWithContext(ctx context.Context) AuthorityOutput {
	return o
}

func (o AuthorityOutput) ToAuthorityPtrOutput() AuthorityPtrOutput {
	return o.ToAuthorityPtrOutputWithContext(context.Background())
}

func (o AuthorityOutput) ToAuthorityPtrOutputWithContext(ctx context.Context) AuthorityPtrOutput {
	return o.ApplyT(func(v Authority) *Authority {
		return &v
	}).(AuthorityPtrOutput)
}

type AuthorityPtrOutput struct {
	*pulumi.OutputState
}

func (AuthorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authority)(nil))
}

func (o AuthorityPtrOutput) ToAuthorityPtrOutput() AuthorityPtrOutput {
	return o
}

func (o AuthorityPtrOutput) ToAuthorityPtrOutputWithContext(ctx context.Context) AuthorityPtrOutput {
	return o
}

type AuthorityArrayOutput struct{ *pulumi.OutputState }

func (AuthorityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Authority)(nil))
}

func (o AuthorityArrayOutput) ToAuthorityArrayOutput() AuthorityArrayOutput {
	return o
}

func (o AuthorityArrayOutput) ToAuthorityArrayOutputWithContext(ctx context.Context) AuthorityArrayOutput {
	return o
}

func (o AuthorityArrayOutput) Index(i pulumi.IntInput) AuthorityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Authority {
		return vs[0].([]Authority)[vs[1].(int)]
	}).(AuthorityOutput)
}

type AuthorityMapOutput struct{ *pulumi.OutputState }

func (AuthorityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Authority)(nil))
}

func (o AuthorityMapOutput) ToAuthorityMapOutput() AuthorityMapOutput {
	return o
}

func (o AuthorityMapOutput) ToAuthorityMapOutputWithContext(ctx context.Context) AuthorityMapOutput {
	return o
}

func (o AuthorityMapOutput) MapIndex(k pulumi.StringInput) AuthorityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Authority {
		return vs[0].(map[string]Authority)[vs[1].(string)]
	}).(AuthorityOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorityOutput{})
	pulumi.RegisterOutputType(AuthorityPtrOutput{})
	pulumi.RegisterOutputType(AuthorityArrayOutput{})
	pulumi.RegisterOutputType(AuthorityMapOutput{})
}
