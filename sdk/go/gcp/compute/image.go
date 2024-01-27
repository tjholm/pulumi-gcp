// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an Image resource.
//
// Google Compute Engine uses operating system images to create the root
// persistent disks for your instances. You specify an image when you create
// an instance. Images contain a boot loader, an operating system, and a
// root file system. Linux operating system images are also capable of
// running containers on Compute Engine.
//
// Images can be either public or custom.
//
// Public images are provided and maintained by Google, open-source
// communities, and third-party vendors. By default, all projects have
// access to these images and can use them to create instances.  Custom
// images are available only to your project. You can create a custom image
// from root persistent disks and other images. Then, use the custom image
// to create an instance.
//
// To get more information about Image, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/images)
//
// ## Example Usage
// ### Image Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewImage(ctx, "example", &compute.ImageArgs{
//				RawDisk: &compute.ImageRawDiskArgs{
//					Source: pulumi.String("https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"),
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
// ### Image Guest Os
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewImage(ctx, "example", &compute.ImageArgs{
//				GuestOsFeatures: compute.ImageGuestOsFeatureArray{
//					&compute.ImageGuestOsFeatureArgs{
//						Type: pulumi.String("SECURE_BOOT"),
//					},
//					&compute.ImageGuestOsFeatureArgs{
//						Type: pulumi.String("MULTI_IP_SUBNET"),
//					},
//				},
//				RawDisk: &compute.ImageRawDiskArgs{
//					Source: pulumi.String("https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"),
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
// ### Image Basic Storage Location
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewImage(ctx, "example", &compute.ImageArgs{
//				RawDisk: &compute.ImageRawDiskArgs{
//					Source: pulumi.String("https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"),
//				},
//				StorageLocations: pulumi.StringArray{
//					pulumi.String("us-central1"),
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
// Image can be imported using any of these accepted formats* `projects/{{project}}/global/images/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Image can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/image:Image default projects/{{project}}/global/images/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/image:Image default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/image:Image default {{name}}
//
// ```
type Image struct {
	pulumi.CustomResourceState

	// Size of the image tar.gz archive stored in Google Cloud Storage (in
	// bytes).
	ArchiveSizeBytes pulumi.IntOutput `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.IntOutput `pulumi:"diskSizeGb"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family pulumi.StringPtrOutput `pulumi:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures ImageGuestOsFeatureArrayOutput `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key.
	// After you encrypt an image with a customer-supplied key, you must
	// provide the same key if you use the image later (e.g. to create a
	// disk from the image)
	// Structure is documented below.
	ImageEncryptionKey ImageImageEncryptionKeyPtrOutput `pulumi:"imageEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this Image.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Any applicable license URI.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk ImageRawDiskPtrOutput `pulumi:"rawDisk"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk pulumi.StringPtrOutput `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage pulumi.StringPtrOutput `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot pulumi.StringPtrOutput `pulumi:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image
	// (regional or multi-regional).
	// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		args = &ImageArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("gcp:compute/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("gcp:compute/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in
	// bytes).
	ArchiveSizeBytes *int `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures []ImageGuestOsFeature `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key.
	// After you encrypt an image with a customer-supplied key, you must
	// provide the same key if you use the image later (e.g. to create a
	// disk from the image)
	// Structure is documented below.
	ImageEncryptionKey *ImageImageEncryptionKey `pulumi:"imageEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this Image.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk *ImageRawDisk `pulumi:"rawDisk"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk *string `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image
	// (regional or multi-regional).
	// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations []string `pulumi:"storageLocations"`
}

type ImageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in
	// bytes).
	ArchiveSizeBytes pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.IntPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures ImageGuestOsFeatureArrayInput
	// Encrypts the image using a customer-supplied encryption key.
	// After you encrypt an image with a customer-supplied key, you must
	// provide the same key if you use the image later (e.g. to create a
	// disk from the image)
	// Structure is documented below.
	ImageEncryptionKey ImageImageEncryptionKeyPtrInput
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this Image.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk ImageRawDiskPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk pulumi.StringPtrInput
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
	// Cloud Storage bucket storage location of the image
	// (regional or multi-regional).
	// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations pulumi.StringArrayInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures []ImageGuestOsFeature `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key.
	// After you encrypt an image with a customer-supplied key, you must
	// provide the same key if you use the image later (e.g. to create a
	// disk from the image)
	// Structure is documented below.
	ImageEncryptionKey *ImageImageEncryptionKey `pulumi:"imageEncryptionKey"`
	// Labels to apply to this Image.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk *ImageRawDisk `pulumi:"rawDisk"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk *string `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image
	// (regional or multi-regional).
	// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations []string `pulumi:"storageLocations"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.IntPtrInput
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures ImageGuestOsFeatureArrayInput
	// Encrypts the image using a customer-supplied encryption key.
	// After you encrypt an image with a customer-supplied key, you must
	// provide the same key if you use the image later (e.g. to create a
	// disk from the image)
	// Structure is documented below.
	ImageEncryptionKey ImageImageEncryptionKeyPtrInput
	// Labels to apply to this Image.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk ImageRawDiskPtrInput
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk pulumi.StringPtrInput
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
	// Cloud Storage bucket storage location of the image
	// (regional or multi-regional).
	// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations pulumi.StringArrayInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//	ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//	ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// Size of the image tar.gz archive stored in Google Cloud Storage (in
// bytes).
func (o ImageOutput) ArchiveSizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *Image) pulumi.IntOutput { return v.ArchiveSizeBytes }).(pulumi.IntOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ImageOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when
// you create the resource.
func (o ImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Size of the image when restored onto a persistent disk (in GB).
func (o ImageOutput) DiskSizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Image) pulumi.IntOutput { return v.DiskSizeGb }).(pulumi.IntOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o ImageOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the image family to which this image belongs. You can
// create disks by specifying an image family instead of a specific
// image name. The image family always returns its latest image that is
// not deprecated. The name of the image family must comply with
// RFC1035.
func (o ImageOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Family }).(pulumi.StringPtrOutput)
}

// A list of features to enable on the guest operating system.
// Applicable only for bootable images.
// Structure is documented below.
func (o ImageOutput) GuestOsFeatures() ImageGuestOsFeatureArrayOutput {
	return o.ApplyT(func(v *Image) ImageGuestOsFeatureArrayOutput { return v.GuestOsFeatures }).(ImageGuestOsFeatureArrayOutput)
}

// Encrypts the image using a customer-supplied encryption key.
// After you encrypt an image with a customer-supplied key, you must
// provide the same key if you use the image later (e.g. to create a
// disk from the image)
// Structure is documented below.
func (o ImageOutput) ImageEncryptionKey() ImageImageEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *Image) ImageImageEncryptionKeyPtrOutput { return v.ImageEncryptionKey }).(ImageImageEncryptionKeyPtrOutput)
}

// The fingerprint used for optimistic locking of this resource. Used
// internally during updates.
func (o ImageOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this Image.
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o ImageOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Any applicable license URI.
func (o ImageOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.Licenses }).(pulumi.StringArrayOutput)
}

// Name of the resource; provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and
// match the regular expression `a-z?` which means
// the first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the
// last character, which cannot be a dash.
//
// ***
func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ImageOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o ImageOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The parameters of the raw disk image.
// Structure is documented below.
func (o ImageOutput) RawDisk() ImageRawDiskPtrOutput {
	return o.ApplyT(func(v *Image) ImageRawDiskPtrOutput { return v.RawDisk }).(ImageRawDiskPtrOutput)
}

// The URI of the created resource.
func (o ImageOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The source disk to create this image based on.
// You must provide either this property or the
// rawDisk.source property but not both to create an image.
func (o ImageOutput) SourceDisk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.SourceDisk }).(pulumi.StringPtrOutput)
}

// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
// URL of one of the following:
// * The selfLink URL
// * This property
// * The rawDisk.source URL
// * The sourceDisk URL
func (o ImageOutput) SourceImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.SourceImage }).(pulumi.StringPtrOutput)
}

// URL of the source snapshot used to create this image.
// In order to create an image, you must provide the full or partial URL of one of the following:
// * The selfLink URL
// * This property
// * The sourceImage URL
// * The rawDisk.source URL
// * The sourceDisk URL
func (o ImageOutput) SourceSnapshot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.SourceSnapshot }).(pulumi.StringPtrOutput)
}

// Cloud Storage bucket storage location of the image
// (regional or multi-regional).
// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
func (o ImageOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
