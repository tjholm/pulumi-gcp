// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An `Instance` is the runtime dataplane in Apigee.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances/create)
// * How-to Guides
//   - [Creating a runtime instance](https://cloud.google.com/apigee/docs/api-platform/get-started/create-instance)
//
// ## Example Usage
// ### Apigee Instance Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := organizations.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
//			if err != nil {
//				return err
//			}
//			apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(16),
//				Network:      apigeeNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
//				Network: apigeeNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					apigeeRange.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			apigeeOrg, err := apigee.NewOrganization(ctx, "apigeeOrg", &apigee.OrganizationArgs{
//				AnalyticsRegion:   pulumi.String("us-central1"),
//				ProjectId:         *pulumi.String(current.Project),
//				AuthorizedNetwork: apigeeNetwork.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				apigeeVpcConnection,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewInstance(ctx, "apigeeInstance", &apigee.InstanceArgs{
//				Location: pulumi.String("us-central1"),
//				OrgId:    apigeeOrg.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Apigee Instance Cidr Range
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := organizations.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
//			if err != nil {
//				return err
//			}
//			apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(22),
//				Network:      apigeeNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
//				Network: apigeeNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					apigeeRange.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			apigeeOrg, err := apigee.NewOrganization(ctx, "apigeeOrg", &apigee.OrganizationArgs{
//				AnalyticsRegion:   pulumi.String("us-central1"),
//				ProjectId:         *pulumi.String(current.Project),
//				AuthorizedNetwork: apigeeNetwork.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				apigeeVpcConnection,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewInstance(ctx, "apigeeInstance", &apigee.InstanceArgs{
//				Location:         pulumi.String("us-central1"),
//				OrgId:            apigeeOrg.ID(),
//				PeeringCidrRange: pulumi.String("SLASH_22"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Apigee Instance Ip Range
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := organizations.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
//			if err != nil {
//				return err
//			}
//			apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(22),
//				Network:      apigeeNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
//				Network: apigeeNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					apigeeRange.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			apigeeOrg, err := apigee.NewOrganization(ctx, "apigeeOrg", &apigee.OrganizationArgs{
//				AnalyticsRegion:   pulumi.String("us-central1"),
//				ProjectId:         *pulumi.String(current.Project),
//				AuthorizedNetwork: apigeeNetwork.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				apigeeVpcConnection,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewInstance(ctx, "apigeeInstance", &apigee.InstanceArgs{
//				Location: pulumi.String("us-central1"),
//				OrgId:    apigeeOrg.ID(),
//				IpRange:  pulumi.String("10.87.8.0/22"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Apigee Instance Full
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := organizations.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
//			if err != nil {
//				return err
//			}
//			apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(16),
//				Network:      apigeeNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
//				Network: apigeeNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					apigeeRange.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			apigeeKeyring, err := kms.NewKeyRing(ctx, "apigeeKeyring", &kms.KeyRingArgs{
//				Location: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeKey, err := kms.NewCryptoKey(ctx, "apigeeKey", &kms.CryptoKeyArgs{
//				KeyRing: apigeeKeyring.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeSa, err := projects.NewServiceIdentity(ctx, "apigeeSa", &projects.ServiceIdentityArgs{
//				Project: pulumi.Any(google_project.Project.Project_id),
//				Service: pulumi.Any(google_project_service.Apigee.Service),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			apigeeSaKeyuser, err := kms.NewCryptoKeyIAMBinding(ctx, "apigeeSaKeyuser", &kms.CryptoKeyIAMBindingArgs{
//				CryptoKeyId: apigeeKey.ID(),
//				Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
//				Members: pulumi.StringArray{
//					apigeeSa.Email.ApplyT(func(email string) (string, error) {
//						return fmt.Sprintf("serviceAccount:%v", email), nil
//					}).(pulumi.StringOutput),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			apigeeOrg, err := apigee.NewOrganization(ctx, "apigeeOrg", &apigee.OrganizationArgs{
//				AnalyticsRegion:                  pulumi.String("us-central1"),
//				DisplayName:                      pulumi.String("apigee-org"),
//				Description:                      pulumi.String("Auto-provisioned Apigee Org."),
//				ProjectId:                        *pulumi.String(current.Project),
//				AuthorizedNetwork:                apigeeNetwork.ID(),
//				RuntimeDatabaseEncryptionKeyName: apigeeKey.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				apigeeVpcConnection,
//				apigeeSaKeyuser,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewInstance(ctx, "apigeeInstance", &apigee.InstanceArgs{
//				Location:              pulumi.String("us-central1"),
//				Description:           pulumi.String("Auto-managed Apigee Runtime Instance"),
//				DisplayName:           pulumi.String("tf-test"),
//				OrgId:                 apigeeOrg.ID(),
//				DiskEncryptionKeyName: apigeeKey.ID(),
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
// # Instance can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:apigee/instance:Instance default {{org_id}}/instances/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigee/instance:Instance default {{org_id}}/{{name}}
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptLists pulumi.StringArrayOutput `pulumi:"consumerAcceptLists"`
	// Description of the instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringPtrOutput `pulumi:"diskEncryptionKeyName"`
	// Display name of the instance.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
	Host pulumi.StringOutput `pulumi:"host"`
	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IpRange pulumi.StringPtrOutput `pulumi:"ipRange"`
	// Required. Compute Engine location where the instance resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource ID of the instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Apigee Organization associated with the Apigee instance,
	// in the format `organizations/{{org_name}}`.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	PeeringCidrRange pulumi.StringOutput `pulumi:"peeringCidrRange"`
	// Output only. Port number of the exposed Apigee endpoint.
	Port pulumi.StringOutput `pulumi:"port"`
	// Output only. Resource name of the service attachment created for the instance in
	// the format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately
	// forward traffic to this service attachment using the PSC endpoints.
	ServiceAttachment pulumi.StringOutput `pulumi:"serviceAttachment"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:apigee/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:apigee/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptLists []string `pulumi:"consumerAcceptLists"`
	// Description of the instance.
	Description *string `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName *string `pulumi:"diskEncryptionKeyName"`
	// Display name of the instance.
	DisplayName *string `pulumi:"displayName"`
	// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
	Host *string `pulumi:"host"`
	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IpRange *string `pulumi:"ipRange"`
	// Required. Compute Engine location where the instance resides.
	Location *string `pulumi:"location"`
	// Resource ID of the instance.
	Name *string `pulumi:"name"`
	// The Apigee Organization associated with the Apigee instance,
	// in the format `organizations/{{org_name}}`.
	OrgId *string `pulumi:"orgId"`
	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	PeeringCidrRange *string `pulumi:"peeringCidrRange"`
	// Output only. Port number of the exposed Apigee endpoint.
	Port *string `pulumi:"port"`
	// Output only. Resource name of the service attachment created for the instance in
	// the format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately
	// forward traffic to this service attachment using the PSC endpoints.
	ServiceAttachment *string `pulumi:"serviceAttachment"`
}

type InstanceState struct {
	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptLists pulumi.StringArrayInput
	// Description of the instance.
	Description pulumi.StringPtrInput
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringPtrInput
	// Display name of the instance.
	DisplayName pulumi.StringPtrInput
	// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
	Host pulumi.StringPtrInput
	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IpRange pulumi.StringPtrInput
	// Required. Compute Engine location where the instance resides.
	Location pulumi.StringPtrInput
	// Resource ID of the instance.
	Name pulumi.StringPtrInput
	// The Apigee Organization associated with the Apigee instance,
	// in the format `organizations/{{org_name}}`.
	OrgId pulumi.StringPtrInput
	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	PeeringCidrRange pulumi.StringPtrInput
	// Output only. Port number of the exposed Apigee endpoint.
	Port pulumi.StringPtrInput
	// Output only. Resource name of the service attachment created for the instance in
	// the format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately
	// forward traffic to this service attachment using the PSC endpoints.
	ServiceAttachment pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptLists []string `pulumi:"consumerAcceptLists"`
	// Description of the instance.
	Description *string `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName *string `pulumi:"diskEncryptionKeyName"`
	// Display name of the instance.
	DisplayName *string `pulumi:"displayName"`
	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IpRange *string `pulumi:"ipRange"`
	// Required. Compute Engine location where the instance resides.
	Location string `pulumi:"location"`
	// Resource ID of the instance.
	Name *string `pulumi:"name"`
	// The Apigee Organization associated with the Apigee instance,
	// in the format `organizations/{{org_name}}`.
	OrgId string `pulumi:"orgId"`
	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	PeeringCidrRange *string `pulumi:"peeringCidrRange"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptLists pulumi.StringArrayInput
	// Description of the instance.
	Description pulumi.StringPtrInput
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringPtrInput
	// Display name of the instance.
	DisplayName pulumi.StringPtrInput
	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IpRange pulumi.StringPtrInput
	// Required. Compute Engine location where the instance resides.
	Location pulumi.StringInput
	// Resource ID of the instance.
	Name pulumi.StringPtrInput
	// The Apigee Organization associated with the Apigee instance,
	// in the format `organizations/{{org_name}}`.
	OrgId pulumi.StringInput
	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	PeeringCidrRange pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Optional. Customer accept list represents the list of projects (id/number) on customer
// side that can privately connect to the service attachment. It is an optional field
// which the customers can provide during the instance creation. By default, the customer
// project associated with the Apigee organization will be included to the list.
func (o InstanceOutput) ConsumerAcceptLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.ConsumerAcceptLists }).(pulumi.StringArrayOutput)
}

// Description of the instance.
func (o InstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
func (o InstanceOutput) DiskEncryptionKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.DiskEncryptionKeyName }).(pulumi.StringPtrOutput)
}

// Display name of the instance.
func (o InstanceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
func (o InstanceOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// IP range represents the customer-provided CIDR block of length 22 that will be used for
// the Apigee instance creation. This optional range, if provided, should be freely
// available as part of larger named range the customer has allocated to the Service
// Networking peering. If this is not provided, Apigee will automatically request for any
// available /22 CIDR block from Service Networking. The customer should use this CIDR block
// for configuring their firewall needs to allow traffic from Apigee.
// Input format: "a.b.c.d/22"
func (o InstanceOutput) IpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.IpRange }).(pulumi.StringPtrOutput)
}

// Required. Compute Engine location where the instance resides.
func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource ID of the instance.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Apigee Organization associated with the Apigee instance,
// in the format `organizations/{{org_name}}`.
func (o InstanceOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// The size of the CIDR block range that will be reserved by the instance. For valid values,
// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
func (o InstanceOutput) PeeringCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PeeringCidrRange }).(pulumi.StringOutput)
}

// Output only. Port number of the exposed Apigee endpoint.
func (o InstanceOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Output only. Resource name of the service attachment created for the instance in
// the format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately
// forward traffic to this service attachment using the PSC endpoints.
func (o InstanceOutput) ServiceAttachment() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceAttachment }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
