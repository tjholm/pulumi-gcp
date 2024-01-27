// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Access Approval enables you to require your explicit approval whenever Google support and engineering need to access your customer content.
//
// To get more information about FolderSettings, see:
//
// * [API documentation](https://cloud.google.com/access-approval/docs/reference/rest/v1/folders)
//
// ## Example Usage
// ### Folder Access Approval Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/folder"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myFolder, err := organizations.NewFolder(ctx, "myFolder", &organizations.FolderArgs{
//				DisplayName: pulumi.String("my-folder"),
//				Parent:      pulumi.String("organizations/123456789"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = folder.NewAccessApprovalSettings(ctx, "folderAccessApproval", &folder.AccessApprovalSettingsArgs{
//				FolderId: myFolder.FolderId,
//				NotificationEmails: pulumi.StringArray{
//					pulumi.String("testuser@example.com"),
//					pulumi.String("example.user@example.com"),
//				},
//				EnrolledServices: folder.AccessApprovalSettingsEnrolledServiceArray{
//					&folder.AccessApprovalSettingsEnrolledServiceArgs{
//						CloudProduct: pulumi.String("all"),
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
// ### Folder Access Approval Active Key Version
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/accessapproval"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/folder"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myFolder, err := organizations.NewFolder(ctx, "myFolder", &organizations.FolderArgs{
//				DisplayName: pulumi.String("my-folder"),
//				Parent:      pulumi.String("organizations/123456789"),
//			})
//			if err != nil {
//				return err
//			}
//			myProject, err := organizations.NewProject(ctx, "myProject", &organizations.ProjectArgs{
//				FolderId: myFolder.Name,
//			})
//			if err != nil {
//				return err
//			}
//			keyRing, err := kms.NewKeyRing(ctx, "keyRing", &kms.KeyRingArgs{
//				Location: pulumi.String("global"),
//				Project:  myProject.ProjectId,
//			})
//			if err != nil {
//				return err
//			}
//			cryptoKey, err := kms.NewCryptoKey(ctx, "cryptoKey", &kms.CryptoKeyArgs{
//				KeyRing: keyRing.ID(),
//				Purpose: pulumi.String("ASYMMETRIC_SIGN"),
//				VersionTemplate: &kms.CryptoKeyVersionTemplateArgs{
//					Algorithm: pulumi.String("EC_SIGN_P384_SHA384"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			serviceAccount := accessapproval.GetFolderServiceAccountOutput(ctx, accessapproval.GetFolderServiceAccountOutputArgs{
//				FolderId: myFolder.FolderId,
//			}, nil)
//			iam, err := kms.NewCryptoKeyIAMMember(ctx, "iam", &kms.CryptoKeyIAMMemberArgs{
//				CryptoKeyId: cryptoKey.ID(),
//				Role:        pulumi.String("roles/cloudkms.signerVerifier"),
//				Member: serviceAccount.ApplyT(func(serviceAccount accessapproval.GetFolderServiceAccountResult) (string, error) {
//					return fmt.Sprintf("serviceAccount:%v", serviceAccount.AccountEmail), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			cryptoKeyVersion := kms.GetKMSCryptoKeyVersionOutput(ctx, kms.GetKMSCryptoKeyVersionOutputArgs{
//				CryptoKey: cryptoKey.ID(),
//			}, nil)
//			_, err = folder.NewAccessApprovalSettings(ctx, "folderAccessApproval", &folder.AccessApprovalSettingsArgs{
//				FolderId: myFolder.FolderId,
//				ActiveKeyVersion: cryptoKeyVersion.ApplyT(func(cryptoKeyVersion kms.GetKMSCryptoKeyVersionResult) (*string, error) {
//					return &cryptoKeyVersion.Name, nil
//				}).(pulumi.StringPtrOutput),
//				EnrolledServices: folder.AccessApprovalSettingsEnrolledServiceArray{
//					&folder.AccessApprovalSettingsEnrolledServiceArgs{
//						CloudProduct: pulumi.String("all"),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				iam,
//			}))
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
// FolderSettings can be imported using any of these accepted formats* `folders/{{folder_id}}/accessApprovalSettings` * `{{folder_id}}` When using the `pulumi import` command, FolderSettings can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:folder/accessApprovalSettings:AccessApprovalSettings default folders/{{folder_id}}/accessApprovalSettings
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:folder/accessApprovalSettings:AccessApprovalSettings default {{folder_id}}
//
// ```
type AccessApprovalSettings struct {
	pulumi.CustomResourceState

	// The asymmetric crypto key version to use for signing approval requests.
	// Empty activeKeyVersion indicates that a Google-managed key should be used for signing.
	// This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
	ActiveKeyVersion pulumi.StringPtrOutput `pulumi:"activeKeyVersion"`
	// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
	AncestorHasActiveKeyVersion pulumi.BoolOutput `pulumi:"ancestorHasActiveKeyVersion"`
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Folder.
	EnrolledAncestor pulumi.BoolOutput `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	// Structure is documented below.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayOutput `pulumi:"enrolledServices"`
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// If the field is true, that indicates that there is some configuration issue with the activeKeyVersion
	// configured on this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the
	// correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
	// as key versions are inherited top-down.
	InvalidKeyVersion pulumi.BoolOutput `pulumi:"invalidKeyVersion"`
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	NotificationEmails pulumi.StringArrayOutput `pulumi:"notificationEmails"`
}

// NewAccessApprovalSettings registers a new resource with the given unique name, arguments, and options.
func NewAccessApprovalSettings(ctx *pulumi.Context,
	name string, args *AccessApprovalSettingsArgs, opts ...pulumi.ResourceOption) (*AccessApprovalSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnrolledServices == nil {
		return nil, errors.New("invalid value for required argument 'EnrolledServices'")
	}
	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessApprovalSettings
	err := ctx.RegisterResource("gcp:folder/accessApprovalSettings:AccessApprovalSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessApprovalSettings gets an existing AccessApprovalSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessApprovalSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessApprovalSettingsState, opts ...pulumi.ResourceOption) (*AccessApprovalSettings, error) {
	var resource AccessApprovalSettings
	err := ctx.ReadResource("gcp:folder/accessApprovalSettings:AccessApprovalSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessApprovalSettings resources.
type accessApprovalSettingsState struct {
	// The asymmetric crypto key version to use for signing approval requests.
	// Empty activeKeyVersion indicates that a Google-managed key should be used for signing.
	// This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
	ActiveKeyVersion *string `pulumi:"activeKeyVersion"`
	// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
	AncestorHasActiveKeyVersion *bool `pulumi:"ancestorHasActiveKeyVersion"`
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Folder.
	EnrolledAncestor *bool `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	// Structure is documented below.
	EnrolledServices []AccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	// ID of the folder of the access approval settings.
	FolderId *string `pulumi:"folderId"`
	// If the field is true, that indicates that there is some configuration issue with the activeKeyVersion
	// configured on this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the
	// correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
	// as key versions are inherited top-down.
	InvalidKeyVersion *bool `pulumi:"invalidKeyVersion"`
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name *string `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	NotificationEmails []string `pulumi:"notificationEmails"`
}

type AccessApprovalSettingsState struct {
	// The asymmetric crypto key version to use for signing approval requests.
	// Empty activeKeyVersion indicates that a Google-managed key should be used for signing.
	// This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
	ActiveKeyVersion pulumi.StringPtrInput
	// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
	AncestorHasActiveKeyVersion pulumi.BoolPtrInput
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Folder.
	EnrolledAncestor pulumi.BoolPtrInput
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	// Structure is documented below.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayInput
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringPtrInput
	// If the field is true, that indicates that there is some configuration issue with the activeKeyVersion
	// configured on this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the
	// correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
	// as key versions are inherited top-down.
	InvalidKeyVersion pulumi.BoolPtrInput
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name pulumi.StringPtrInput
	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
}

func (AccessApprovalSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessApprovalSettingsState)(nil)).Elem()
}

type accessApprovalSettingsArgs struct {
	// The asymmetric crypto key version to use for signing approval requests.
	// Empty activeKeyVersion indicates that a Google-managed key should be used for signing.
	// This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
	ActiveKeyVersion *string `pulumi:"activeKeyVersion"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	// Structure is documented below.
	EnrolledServices []AccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	// ID of the folder of the access approval settings.
	FolderId string `pulumi:"folderId"`
	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	NotificationEmails []string `pulumi:"notificationEmails"`
}

// The set of arguments for constructing a AccessApprovalSettings resource.
type AccessApprovalSettingsArgs struct {
	// The asymmetric crypto key version to use for signing approval requests.
	// Empty activeKeyVersion indicates that a Google-managed key should be used for signing.
	// This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
	ActiveKeyVersion pulumi.StringPtrInput
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	// Structure is documented below.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayInput
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringInput
	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
}

func (AccessApprovalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessApprovalSettingsArgs)(nil)).Elem()
}

type AccessApprovalSettingsInput interface {
	pulumi.Input

	ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput
	ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput
}

func (*AccessApprovalSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessApprovalSettings)(nil)).Elem()
}

func (i *AccessApprovalSettings) ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput {
	return i.ToAccessApprovalSettingsOutputWithContext(context.Background())
}

func (i *AccessApprovalSettings) ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsOutput)
}

// AccessApprovalSettingsArrayInput is an input type that accepts AccessApprovalSettingsArray and AccessApprovalSettingsArrayOutput values.
// You can construct a concrete instance of `AccessApprovalSettingsArrayInput` via:
//
//	AccessApprovalSettingsArray{ AccessApprovalSettingsArgs{...} }
type AccessApprovalSettingsArrayInput interface {
	pulumi.Input

	ToAccessApprovalSettingsArrayOutput() AccessApprovalSettingsArrayOutput
	ToAccessApprovalSettingsArrayOutputWithContext(context.Context) AccessApprovalSettingsArrayOutput
}

type AccessApprovalSettingsArray []AccessApprovalSettingsInput

func (AccessApprovalSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessApprovalSettings)(nil)).Elem()
}

func (i AccessApprovalSettingsArray) ToAccessApprovalSettingsArrayOutput() AccessApprovalSettingsArrayOutput {
	return i.ToAccessApprovalSettingsArrayOutputWithContext(context.Background())
}

func (i AccessApprovalSettingsArray) ToAccessApprovalSettingsArrayOutputWithContext(ctx context.Context) AccessApprovalSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsArrayOutput)
}

// AccessApprovalSettingsMapInput is an input type that accepts AccessApprovalSettingsMap and AccessApprovalSettingsMapOutput values.
// You can construct a concrete instance of `AccessApprovalSettingsMapInput` via:
//
//	AccessApprovalSettingsMap{ "key": AccessApprovalSettingsArgs{...} }
type AccessApprovalSettingsMapInput interface {
	pulumi.Input

	ToAccessApprovalSettingsMapOutput() AccessApprovalSettingsMapOutput
	ToAccessApprovalSettingsMapOutputWithContext(context.Context) AccessApprovalSettingsMapOutput
}

type AccessApprovalSettingsMap map[string]AccessApprovalSettingsInput

func (AccessApprovalSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessApprovalSettings)(nil)).Elem()
}

func (i AccessApprovalSettingsMap) ToAccessApprovalSettingsMapOutput() AccessApprovalSettingsMapOutput {
	return i.ToAccessApprovalSettingsMapOutputWithContext(context.Background())
}

func (i AccessApprovalSettingsMap) ToAccessApprovalSettingsMapOutputWithContext(ctx context.Context) AccessApprovalSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsMapOutput)
}

type AccessApprovalSettingsOutput struct{ *pulumi.OutputState }

func (AccessApprovalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessApprovalSettings)(nil)).Elem()
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput {
	return o
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput {
	return o
}

// The asymmetric crypto key version to use for signing approval requests.
// Empty activeKeyVersion indicates that a Google-managed key should be used for signing.
// This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
func (o AccessApprovalSettingsOutput) ActiveKeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.StringPtrOutput { return v.ActiveKeyVersion }).(pulumi.StringPtrOutput)
}

// If the field is true, that indicates that an ancestor of this Folder has set active_key_version.
func (o AccessApprovalSettingsOutput) AncestorHasActiveKeyVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.BoolOutput { return v.AncestorHasActiveKeyVersion }).(pulumi.BoolOutput)
}

// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Folder.
func (o AccessApprovalSettingsOutput) EnrolledAncestor() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.BoolOutput { return v.EnrolledAncestor }).(pulumi.BoolOutput)
}

// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
// Access requests for the resource given by name against any of these services contained here will be required
// to have explicit approval. Enrollment can only be done on an all or nothing basis.
// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
// Structure is documented below.
func (o AccessApprovalSettingsOutput) EnrolledServices() AccessApprovalSettingsEnrolledServiceArrayOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) AccessApprovalSettingsEnrolledServiceArrayOutput {
		return v.EnrolledServices
	}).(AccessApprovalSettingsEnrolledServiceArrayOutput)
}

// ID of the folder of the access approval settings.
func (o AccessApprovalSettingsOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// If the field is true, that indicates that there is some configuration issue with the activeKeyVersion
// configured on this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the
// correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
// as key versions are inherited top-down.
func (o AccessApprovalSettingsOutput) InvalidKeyVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.BoolOutput { return v.InvalidKeyVersion }).(pulumi.BoolOutput)
}

// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
func (o AccessApprovalSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of email addresses to which notifications relating to approval requests should be sent.
// Notifications relating to a resource will be sent to all emails in the settings of ancestor
// resources of that resource. A maximum of 50 email addresses are allowed.
func (o AccessApprovalSettingsOutput) NotificationEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessApprovalSettings) pulumi.StringArrayOutput { return v.NotificationEmails }).(pulumi.StringArrayOutput)
}

type AccessApprovalSettingsArrayOutput struct{ *pulumi.OutputState }

func (AccessApprovalSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessApprovalSettings)(nil)).Elem()
}

func (o AccessApprovalSettingsArrayOutput) ToAccessApprovalSettingsArrayOutput() AccessApprovalSettingsArrayOutput {
	return o
}

func (o AccessApprovalSettingsArrayOutput) ToAccessApprovalSettingsArrayOutputWithContext(ctx context.Context) AccessApprovalSettingsArrayOutput {
	return o
}

func (o AccessApprovalSettingsArrayOutput) Index(i pulumi.IntInput) AccessApprovalSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessApprovalSettings {
		return vs[0].([]*AccessApprovalSettings)[vs[1].(int)]
	}).(AccessApprovalSettingsOutput)
}

type AccessApprovalSettingsMapOutput struct{ *pulumi.OutputState }

func (AccessApprovalSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessApprovalSettings)(nil)).Elem()
}

func (o AccessApprovalSettingsMapOutput) ToAccessApprovalSettingsMapOutput() AccessApprovalSettingsMapOutput {
	return o
}

func (o AccessApprovalSettingsMapOutput) ToAccessApprovalSettingsMapOutputWithContext(ctx context.Context) AccessApprovalSettingsMapOutput {
	return o
}

func (o AccessApprovalSettingsMapOutput) MapIndex(k pulumi.StringInput) AccessApprovalSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessApprovalSettings {
		return vs[0].(map[string]*AccessApprovalSettings)[vs[1].(string)]
	}).(AccessApprovalSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessApprovalSettingsInput)(nil)).Elem(), &AccessApprovalSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessApprovalSettingsArrayInput)(nil)).Elem(), AccessApprovalSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessApprovalSettingsMapInput)(nil)).Elem(), AccessApprovalSettingsMap{})
	pulumi.RegisterOutputType(AccessApprovalSettingsOutput{})
	pulumi.RegisterOutputType(AccessApprovalSettingsArrayOutput{})
	pulumi.RegisterOutputType(AccessApprovalSettingsMapOutput{})
}
