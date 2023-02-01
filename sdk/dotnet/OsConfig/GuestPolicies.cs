// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig
{
    /// <summary>
    /// ## Example Usage
    /// ### Os Config Guest Policies Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myImage = Gcp.Compute.GetImage.Invoke(new()
    ///     {
    ///         Family = "debian-11",
    ///         Project = "debian-cloud",
    ///     });
    /// 
    ///     var foobar = new Gcp.Compute.Instance("foobar", new()
    ///     {
    ///         MachineType = "e2-medium",
    ///         Zone = "us-central1-a",
    ///         CanIpForward = false,
    ///         Tags = new[]
    ///         {
    ///             "foo",
    ///             "bar",
    ///         },
    ///         BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
    ///         {
    ///             InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
    ///             {
    ///                 Image = myImage.Apply(getImageResult =&gt; getImageResult.SelfLink),
    ///             },
    ///         },
    ///         NetworkInterfaces = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.InstanceNetworkInterfaceArgs
    ///             {
    ///                 Network = "default",
    ///             },
    ///         },
    ///         Metadata = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var guestPolicies = new Gcp.OsConfig.GuestPolicies("guestPolicies", new()
    ///     {
    ///         GuestPolicyId = "guest-policy",
    ///         Assignment = new Gcp.OsConfig.Inputs.GuestPoliciesAssignmentArgs
    ///         {
    ///             Instances = new[]
    ///             {
    ///                 foobar.Id,
    ///             },
    ///         },
    ///         Packages = new[]
    ///         {
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesPackageArgs
    ///             {
    ///                 Name = "my-package",
    ///                 DesiredState = "UPDATED",
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Os Config Guest Policies Packages
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var guestPolicies = new Gcp.OsConfig.GuestPolicies("guestPolicies", new()
    ///     {
    ///         GuestPolicyId = "guest-policy",
    ///         Assignment = new Gcp.OsConfig.Inputs.GuestPoliciesAssignmentArgs
    ///         {
    ///             GroupLabels = new[]
    ///             {
    ///                 new Gcp.OsConfig.Inputs.GuestPoliciesAssignmentGroupLabelArgs
    ///                 {
    ///                     Labels = 
    ///                     {
    ///                         { "color", "red" },
    ///                         { "env", "test" },
    ///                     },
    ///                 },
    ///                 new Gcp.OsConfig.Inputs.GuestPoliciesAssignmentGroupLabelArgs
    ///                 {
    ///                     Labels = 
    ///                     {
    ///                         { "color", "blue" },
    ///                         { "env", "test" },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Packages = new[]
    ///         {
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesPackageArgs
    ///             {
    ///                 Name = "my-package",
    ///                 DesiredState = "INSTALLED",
    ///             },
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesPackageArgs
    ///             {
    ///                 Name = "bad-package-1",
    ///                 DesiredState = "REMOVED",
    ///             },
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesPackageArgs
    ///             {
    ///                 Name = "bad-package-2",
    ///                 DesiredState = "REMOVED",
    ///                 Manager = "APT",
    ///             },
    ///         },
    ///         PackageRepositories = new[]
    ///         {
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesPackageRepositoryArgs
    ///             {
    ///                 Apt = new Gcp.OsConfig.Inputs.GuestPoliciesPackageRepositoryAptArgs
    ///                 {
    ///                     Uri = "https://packages.cloud.google.com/apt",
    ///                     ArchiveType = "DEB",
    ///                     Distribution = "cloud-sdk-stretch",
    ///                     Components = new[]
    ///                     {
    ///                         "main",
    ///                     },
    ///                 },
    ///             },
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesPackageRepositoryArgs
    ///             {
    ///                 Yum = new Gcp.OsConfig.Inputs.GuestPoliciesPackageRepositoryYumArgs
    ///                 {
    ///                     Id = "google-cloud-sdk",
    ///                     DisplayName = "Google Cloud SDK",
    ///                     BaseUrl = "https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64",
    ///                     GpgKeys = new[]
    ///                     {
    ///                         "https://packages.cloud.google.com/yum/doc/yum-key.gpg",
    ///                         "https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Os Config Guest Policies Recipes
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var guestPolicies = new Gcp.OsConfig.GuestPolicies("guestPolicies", new()
    ///     {
    ///         GuestPolicyId = "guest-policy",
    ///         Assignment = new Gcp.OsConfig.Inputs.GuestPoliciesAssignmentArgs
    ///         {
    ///             Zones = new[]
    ///             {
    ///                 "us-east1-b",
    ///                 "us-east1-d",
    ///             },
    ///         },
    ///         Recipes = new[]
    ///         {
    ///             new Gcp.OsConfig.Inputs.GuestPoliciesRecipeArgs
    ///             {
    ///                 Name = "guest-policy-recipe",
    ///                 DesiredState = "INSTALLED",
    ///                 Artifacts = new[]
    ///                 {
    ///                     new Gcp.OsConfig.Inputs.GuestPoliciesRecipeArtifactArgs
    ///                     {
    ///                         Id = "guest-policy-artifact-id",
    ///                         Gcs = new Gcp.OsConfig.Inputs.GuestPoliciesRecipeArtifactGcsArgs
    ///                         {
    ///                             Bucket = "my-bucket",
    ///                             Object = "executable.msi",
    ///                             Generation = 1546030865175603,
    ///                         },
    ///                     },
    ///                 },
    ///                 InstallSteps = new[]
    ///                 {
    ///                     new Gcp.OsConfig.Inputs.GuestPoliciesRecipeInstallStepArgs
    ///                     {
    ///                         MsiInstallation = new Gcp.OsConfig.Inputs.GuestPoliciesRecipeInstallStepMsiInstallationArgs
    ///                         {
    ///                             ArtifactId = "guest-policy-artifact-id",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GuestPolicies can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default projects/{{project}}/guestPolicies/{{guest_policy_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default {{project}}/{{guest_policy_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default {{guest_policy_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:osconfig/guestPolicies:GuestPolicies")]
    public partial class GuestPolicies : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the VM instances that are assigned to this policy. This allows you to target sets
        /// or groups of VM instances by different parameters such as labels, names, OS, or zones.
        /// If left empty, all VM instances underneath this policy are targeted.
        /// At the same level in the resource hierarchy (that is within a project), the service prevents
        /// the creation of multiple policies that conflict with each other.
        /// For more information, see how the service
        /// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        /// Structure is documented below.
        /// </summary>
        [Output("assignment")]
        public Output<Outputs.GuestPoliciesAssignment> Assignment { get; private set; } = null!;

        /// <summary>
        /// Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the guest policy. Length of the description is limited to 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The etag for this guest policy. If this is provided on update, it must match the server's etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The logical name of the guest policy in the project with the following restrictions:
        /// * Must contain only lowercase letters, numbers, and hyphens.
        /// * Must start with a letter.
        /// * Must be between 1-63 characters.
        /// * Must end with a number or a letter.
        /// * Must be unique within the project.
        /// </summary>
        [Output("guestPolicyId")]
        public Output<string> GuestPolicyId { get; private set; } = null!;

        /// <summary>
        /// The name of the package. A package is uniquely identified for conflict validation
        /// by checking the package name and the manager(s) that the package targets.
        /// (Required)
        /// The name of the repository.
        /// (Required)
        /// Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
        /// Names are also used to identify resources which helps to determine whether guest policies have conflicts.
        /// This means that requests to create multiple recipes with the same name and version are rejected since they
        /// could potentially have conflicting assignments.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of package repositories to configure on the VM instance.
        /// This is done before any other configs are applied so they can use these repos.
        /// Package repositories are only configured if the corresponding package manager(s) are available.
        /// Structure is documented below.
        /// </summary>
        [Output("packageRepositories")]
        public Output<ImmutableArray<Outputs.GuestPoliciesPackageRepository>> PackageRepositories { get; private set; } = null!;

        /// <summary>
        /// The software packages to be managed by this policy.
        /// Structure is documented below.
        /// </summary>
        [Output("packages")]
        public Output<ImmutableArray<Outputs.GuestPoliciesPackage>> Packages { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A list of Recipes to install on the VM instance.
        /// Structure is documented below.
        /// </summary>
        [Output("recipes")]
        public Output<ImmutableArray<Outputs.GuestPoliciesRecipe>> Recipes { get; private set; } = null!;

        /// <summary>
        /// Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GuestPolicies resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GuestPolicies(string name, GuestPoliciesArgs args, CustomResourceOptions? options = null)
            : base("gcp:osconfig/guestPolicies:GuestPolicies", name, args ?? new GuestPoliciesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GuestPolicies(string name, Input<string> id, GuestPoliciesState? state = null, CustomResourceOptions? options = null)
            : base("gcp:osconfig/guestPolicies:GuestPolicies", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GuestPolicies resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GuestPolicies Get(string name, Input<string> id, GuestPoliciesState? state = null, CustomResourceOptions? options = null)
        {
            return new GuestPolicies(name, id, state, options);
        }
    }

    public sealed class GuestPoliciesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the VM instances that are assigned to this policy. This allows you to target sets
        /// or groups of VM instances by different parameters such as labels, names, OS, or zones.
        /// If left empty, all VM instances underneath this policy are targeted.
        /// At the same level in the resource hierarchy (that is within a project), the service prevents
        /// the creation of multiple policies that conflict with each other.
        /// For more information, see how the service
        /// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        /// Structure is documented below.
        /// </summary>
        [Input("assignment", required: true)]
        public Input<Inputs.GuestPoliciesAssignmentArgs> Assignment { get; set; } = null!;

        /// <summary>
        /// Description of the guest policy. Length of the description is limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The etag for this guest policy. If this is provided on update, it must match the server's etag.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The logical name of the guest policy in the project with the following restrictions:
        /// * Must contain only lowercase letters, numbers, and hyphens.
        /// * Must start with a letter.
        /// * Must be between 1-63 characters.
        /// * Must end with a number or a letter.
        /// * Must be unique within the project.
        /// </summary>
        [Input("guestPolicyId", required: true)]
        public Input<string> GuestPolicyId { get; set; } = null!;

        [Input("packageRepositories")]
        private InputList<Inputs.GuestPoliciesPackageRepositoryArgs>? _packageRepositories;

        /// <summary>
        /// A list of package repositories to configure on the VM instance.
        /// This is done before any other configs are applied so they can use these repos.
        /// Package repositories are only configured if the corresponding package manager(s) are available.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GuestPoliciesPackageRepositoryArgs> PackageRepositories
        {
            get => _packageRepositories ?? (_packageRepositories = new InputList<Inputs.GuestPoliciesPackageRepositoryArgs>());
            set => _packageRepositories = value;
        }

        [Input("packages")]
        private InputList<Inputs.GuestPoliciesPackageArgs>? _packages;

        /// <summary>
        /// The software packages to be managed by this policy.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GuestPoliciesPackageArgs> Packages
        {
            get => _packages ?? (_packages = new InputList<Inputs.GuestPoliciesPackageArgs>());
            set => _packages = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("recipes")]
        private InputList<Inputs.GuestPoliciesRecipeArgs>? _recipes;

        /// <summary>
        /// A list of Recipes to install on the VM instance.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GuestPoliciesRecipeArgs> Recipes
        {
            get => _recipes ?? (_recipes = new InputList<Inputs.GuestPoliciesRecipeArgs>());
            set => _recipes = value;
        }

        public GuestPoliciesArgs()
        {
        }
        public static new GuestPoliciesArgs Empty => new GuestPoliciesArgs();
    }

    public sealed class GuestPoliciesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the VM instances that are assigned to this policy. This allows you to target sets
        /// or groups of VM instances by different parameters such as labels, names, OS, or zones.
        /// If left empty, all VM instances underneath this policy are targeted.
        /// At the same level in the resource hierarchy (that is within a project), the service prevents
        /// the creation of multiple policies that conflict with each other.
        /// For more information, see how the service
        /// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        /// Structure is documented below.
        /// </summary>
        [Input("assignment")]
        public Input<Inputs.GuestPoliciesAssignmentGetArgs>? Assignment { get; set; }

        /// <summary>
        /// Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the guest policy. Length of the description is limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The etag for this guest policy. If this is provided on update, it must match the server's etag.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The logical name of the guest policy in the project with the following restrictions:
        /// * Must contain only lowercase letters, numbers, and hyphens.
        /// * Must start with a letter.
        /// * Must be between 1-63 characters.
        /// * Must end with a number or a letter.
        /// * Must be unique within the project.
        /// </summary>
        [Input("guestPolicyId")]
        public Input<string>? GuestPolicyId { get; set; }

        /// <summary>
        /// The name of the package. A package is uniquely identified for conflict validation
        /// by checking the package name and the manager(s) that the package targets.
        /// (Required)
        /// The name of the repository.
        /// (Required)
        /// Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
        /// Names are also used to identify resources which helps to determine whether guest policies have conflicts.
        /// This means that requests to create multiple recipes with the same name and version are rejected since they
        /// could potentially have conflicting assignments.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("packageRepositories")]
        private InputList<Inputs.GuestPoliciesPackageRepositoryGetArgs>? _packageRepositories;

        /// <summary>
        /// A list of package repositories to configure on the VM instance.
        /// This is done before any other configs are applied so they can use these repos.
        /// Package repositories are only configured if the corresponding package manager(s) are available.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GuestPoliciesPackageRepositoryGetArgs> PackageRepositories
        {
            get => _packageRepositories ?? (_packageRepositories = new InputList<Inputs.GuestPoliciesPackageRepositoryGetArgs>());
            set => _packageRepositories = value;
        }

        [Input("packages")]
        private InputList<Inputs.GuestPoliciesPackageGetArgs>? _packages;

        /// <summary>
        /// The software packages to be managed by this policy.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GuestPoliciesPackageGetArgs> Packages
        {
            get => _packages ?? (_packages = new InputList<Inputs.GuestPoliciesPackageGetArgs>());
            set => _packages = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("recipes")]
        private InputList<Inputs.GuestPoliciesRecipeGetArgs>? _recipes;

        /// <summary>
        /// A list of Recipes to install on the VM instance.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GuestPoliciesRecipeGetArgs> Recipes
        {
            get => _recipes ?? (_recipes = new InputList<Inputs.GuestPoliciesRecipeGetArgs>());
            set => _recipes = value;
        }

        /// <summary>
        /// Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public GuestPoliciesState()
        {
        }
        public static new GuestPoliciesState Empty => new GuestPoliciesState();
    }
}
