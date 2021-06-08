// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// The Google Compute Engine Instance Group Manager API creates and manages pools
    /// of homogeneous Compute Engine virtual machine instances from a common instance
    /// template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
    /// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)
    /// 
    /// &gt; **Note:** Use [gcp.compute.RegionInstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.
    /// 
    /// ## Example Usage
    /// ### With Top Level Instance Template (`Google` Provider)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var autohealing = new Gcp.Compute.HealthCheck("autohealing", new Gcp.Compute.HealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 5,
    ///             TimeoutSec = 5,
    ///             HealthyThreshold = 2,
    ///             UnhealthyThreshold = 10,
    ///             HttpHealthCheck = new Gcp.Compute.Inputs.HealthCheckHttpHealthCheckArgs
    ///             {
    ///                 RequestPath = "/healthz",
    ///                 Port = 8080,
    ///             },
    ///         });
    ///         var appserver = new Gcp.Compute.InstanceGroupManager("appserver", new Gcp.Compute.InstanceGroupManagerArgs
    ///         {
    ///             BaseInstanceName = "app",
    ///             Zone = "us-central1-a",
    ///             Versions = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceGroupManagerVersionArgs
    ///                 {
    ///                     InstanceTemplate = google_compute_instance_template.Appserver.Id,
    ///                 },
    ///             },
    ///             TargetPools = 
    ///             {
    ///                 google_compute_target_pool.Appserver.Id,
    ///             },
    ///             TargetSize = 2,
    ///             NamedPorts = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceGroupManagerNamedPortArgs
    ///                 {
    ///                     Name = "customHTTP",
    ///                     Port = 8888,
    ///                 },
    ///             },
    ///             AutoHealingPolicies = new Gcp.Compute.Inputs.InstanceGroupManagerAutoHealingPoliciesArgs
    ///             {
    ///                 HealthCheck = autohealing.Id,
    ///                 InitialDelaySec = 300,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With Multiple Versions (`Google-Beta` Provider)
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var appserver = new Gcp.Compute.InstanceGroupManager("appserver", new Gcp.Compute.InstanceGroupManagerArgs
    ///         {
    ///             BaseInstanceName = "app",
    ///             Zone = "us-central1-a",
    ///             TargetSize = 5,
    ///             Versions = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceGroupManagerVersionArgs
    ///                 {
    ///                     Name = "appserver",
    ///                     InstanceTemplate = google_compute_instance_template.Appserver.Id,
    ///                 },
    ///                 new Gcp.Compute.Inputs.InstanceGroupManagerVersionArgs
    ///                 {
    ///                     Name = "appserver-canary",
    ///                     InstanceTemplate = google_compute_instance_template.Appserver_canary.Id,
    ///                     TargetSize = new Gcp.Compute.Inputs.InstanceGroupManagerVersionTargetSizeArgs
    ///                     {
    ///                         Fixed = 1,
    ///                     },
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance group managers can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{zone}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/instanceGroupManager:InstanceGroupManager")]
    public partial class InstanceGroupManager : Pulumi.CustomResource
    {
        /// <summary>
        /// The autohealing policies for this managed instance
        /// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        /// </summary>
        [Output("autoHealingPolicies")]
        public Output<Outputs.InstanceGroupManagerAutoHealingPolicies?> AutoHealingPolicies { get; private set; } = null!;

        /// <summary>
        /// The base instance name to use for
        /// instances in this group. The value must be a valid
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        /// are lowercase letters, numbers, and hyphens (-). Instances are named by
        /// appending a hyphen and a random four-character string to the base instance
        /// name.
        /// </summary>
        [Output("baseInstanceName")]
        public Output<string> BaseInstanceName { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the instance
        /// group manager.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the instance group manager.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The full URL of the instance group created by the manager.
        /// </summary>
        [Output("instanceGroup")]
        public Output<string> InstanceGroup { get; private set; } = null!;

        /// <summary>
        /// - Version name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The named port configuration. See the section below
        /// for details on configuration.
        /// </summary>
        [Output("namedPorts")]
        public Output<ImmutableArray<Outputs.InstanceGroupManagerNamedPort>> NamedPorts { get; private set; } = null!;

        [Output("operation")]
        public Output<string> Operation { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URL of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        /// </summary>
        [Output("statefulDisks")]
        public Output<ImmutableArray<Outputs.InstanceGroupManagerStatefulDisk>> StatefulDisks { get; private set; } = null!;

        /// <summary>
        /// The status of this managed instance group.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.InstanceGroupManagerStatus>> Statuses { get; private set; } = null!;

        /// <summary>
        /// The full URL of all target pools to which new
        /// instances in the group are added. Updating the target pools attribute does
        /// not affect existing instances.
        /// </summary>
        [Output("targetPools")]
        public Output<ImmutableArray<string>> TargetPools { get; private set; } = null!;

        /// <summary>
        /// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        /// </summary>
        [Output("targetSize")]
        public Output<int> TargetSize { get; private set; } = null!;

        /// <summary>
        /// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
        /// </summary>
        [Output("updatePolicy")]
        public Output<Outputs.InstanceGroupManagerUpdatePolicy> UpdatePolicy { get; private set; } = null!;

        /// <summary>
        /// Application versions managed by this instance group. Each
        /// version deals with a specific instance template, allowing canary release scenarios.
        /// Structure is documented below.
        /// </summary>
        [Output("versions")]
        public Output<ImmutableArray<Outputs.InstanceGroupManagerVersion>> Versions { get; private set; } = null!;

        /// <summary>
        /// Whether to wait for all instances to be created/updated before
        /// returning. Note that if this is set to true and the operation does not succeed, this provider will
        /// continue trying until it times out.
        /// </summary>
        [Output("waitForInstances")]
        public Output<bool?> WaitForInstances { get; private set; } = null!;

        /// <summary>
        /// When used with `wait_for_instances` it specifies the status to wait for.
        /// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
        /// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
        /// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
        /// </summary>
        [Output("waitForInstancesStatus")]
        public Output<string?> WaitForInstancesStatus { get; private set; } = null!;

        /// <summary>
        /// The zone that instances in this group should be created
        /// in.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceGroupManager resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceGroupManager(string name, InstanceGroupManagerArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceGroupManager:InstanceGroupManager", name, args ?? new InstanceGroupManagerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceGroupManager(string name, Input<string> id, InstanceGroupManagerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceGroupManager:InstanceGroupManager", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceGroupManager resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceGroupManager Get(string name, Input<string> id, InstanceGroupManagerState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceGroupManager(name, id, state, options);
        }
    }

    public sealed class InstanceGroupManagerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autohealing policies for this managed instance
        /// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        /// </summary>
        [Input("autoHealingPolicies")]
        public Input<Inputs.InstanceGroupManagerAutoHealingPoliciesArgs>? AutoHealingPolicies { get; set; }

        /// <summary>
        /// The base instance name to use for
        /// instances in this group. The value must be a valid
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        /// are lowercase letters, numbers, and hyphens (-). Instances are named by
        /// appending a hyphen and a random four-character string to the base instance
        /// name.
        /// </summary>
        [Input("baseInstanceName", required: true)]
        public Input<string> BaseInstanceName { get; set; } = null!;

        /// <summary>
        /// An optional textual description of the instance
        /// group manager.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// - Version name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namedPorts")]
        private InputList<Inputs.InstanceGroupManagerNamedPortArgs>? _namedPorts;

        /// <summary>
        /// The named port configuration. See the section below
        /// for details on configuration.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerNamedPortArgs> NamedPorts
        {
            get => _namedPorts ?? (_namedPorts = new InputList<Inputs.InstanceGroupManagerNamedPortArgs>());
            set => _namedPorts = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("statefulDisks")]
        private InputList<Inputs.InstanceGroupManagerStatefulDiskArgs>? _statefulDisks;

        /// <summary>
        /// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerStatefulDiskArgs> StatefulDisks
        {
            get => _statefulDisks ?? (_statefulDisks = new InputList<Inputs.InstanceGroupManagerStatefulDiskArgs>());
            set => _statefulDisks = value;
        }

        [Input("targetPools")]
        private InputList<string>? _targetPools;

        /// <summary>
        /// The full URL of all target pools to which new
        /// instances in the group are added. Updating the target pools attribute does
        /// not affect existing instances.
        /// </summary>
        public InputList<string> TargetPools
        {
            get => _targetPools ?? (_targetPools = new InputList<string>());
            set => _targetPools = value;
        }

        /// <summary>
        /// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        /// </summary>
        [Input("targetSize")]
        public Input<int>? TargetSize { get; set; }

        /// <summary>
        /// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
        /// </summary>
        [Input("updatePolicy")]
        public Input<Inputs.InstanceGroupManagerUpdatePolicyArgs>? UpdatePolicy { get; set; }

        [Input("versions", required: true)]
        private InputList<Inputs.InstanceGroupManagerVersionArgs>? _versions;

        /// <summary>
        /// Application versions managed by this instance group. Each
        /// version deals with a specific instance template, allowing canary release scenarios.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerVersionArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.InstanceGroupManagerVersionArgs>());
            set => _versions = value;
        }

        /// <summary>
        /// Whether to wait for all instances to be created/updated before
        /// returning. Note that if this is set to true and the operation does not succeed, this provider will
        /// continue trying until it times out.
        /// </summary>
        [Input("waitForInstances")]
        public Input<bool>? WaitForInstances { get; set; }

        /// <summary>
        /// When used with `wait_for_instances` it specifies the status to wait for.
        /// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
        /// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
        /// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
        /// </summary>
        [Input("waitForInstancesStatus")]
        public Input<string>? WaitForInstancesStatus { get; set; }

        /// <summary>
        /// The zone that instances in this group should be created
        /// in.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceGroupManagerArgs()
        {
        }
    }

    public sealed class InstanceGroupManagerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autohealing policies for this managed instance
        /// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        /// </summary>
        [Input("autoHealingPolicies")]
        public Input<Inputs.InstanceGroupManagerAutoHealingPoliciesGetArgs>? AutoHealingPolicies { get; set; }

        /// <summary>
        /// The base instance name to use for
        /// instances in this group. The value must be a valid
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        /// are lowercase letters, numbers, and hyphens (-). Instances are named by
        /// appending a hyphen and a random four-character string to the base instance
        /// name.
        /// </summary>
        [Input("baseInstanceName")]
        public Input<string>? BaseInstanceName { get; set; }

        /// <summary>
        /// An optional textual description of the instance
        /// group manager.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The fingerprint of the instance group manager.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The full URL of the instance group created by the manager.
        /// </summary>
        [Input("instanceGroup")]
        public Input<string>? InstanceGroup { get; set; }

        /// <summary>
        /// - Version name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namedPorts")]
        private InputList<Inputs.InstanceGroupManagerNamedPortGetArgs>? _namedPorts;

        /// <summary>
        /// The named port configuration. See the section below
        /// for details on configuration.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerNamedPortGetArgs> NamedPorts
        {
            get => _namedPorts ?? (_namedPorts = new InputList<Inputs.InstanceGroupManagerNamedPortGetArgs>());
            set => _namedPorts = value;
        }

        [Input("operation")]
        public Input<string>? Operation { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URL of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("statefulDisks")]
        private InputList<Inputs.InstanceGroupManagerStatefulDiskGetArgs>? _statefulDisks;

        /// <summary>
        /// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerStatefulDiskGetArgs> StatefulDisks
        {
            get => _statefulDisks ?? (_statefulDisks = new InputList<Inputs.InstanceGroupManagerStatefulDiskGetArgs>());
            set => _statefulDisks = value;
        }

        [Input("statuses")]
        private InputList<Inputs.InstanceGroupManagerStatusGetArgs>? _statuses;

        /// <summary>
        /// The status of this managed instance group.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerStatusGetArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.InstanceGroupManagerStatusGetArgs>());
            set => _statuses = value;
        }

        [Input("targetPools")]
        private InputList<string>? _targetPools;

        /// <summary>
        /// The full URL of all target pools to which new
        /// instances in the group are added. Updating the target pools attribute does
        /// not affect existing instances.
        /// </summary>
        public InputList<string> TargetPools
        {
            get => _targetPools ?? (_targetPools = new InputList<string>());
            set => _targetPools = value;
        }

        /// <summary>
        /// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        /// </summary>
        [Input("targetSize")]
        public Input<int>? TargetSize { get; set; }

        /// <summary>
        /// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
        /// </summary>
        [Input("updatePolicy")]
        public Input<Inputs.InstanceGroupManagerUpdatePolicyGetArgs>? UpdatePolicy { get; set; }

        [Input("versions")]
        private InputList<Inputs.InstanceGroupManagerVersionGetArgs>? _versions;

        /// <summary>
        /// Application versions managed by this instance group. Each
        /// version deals with a specific instance template, allowing canary release scenarios.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerVersionGetArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.InstanceGroupManagerVersionGetArgs>());
            set => _versions = value;
        }

        /// <summary>
        /// Whether to wait for all instances to be created/updated before
        /// returning. Note that if this is set to true and the operation does not succeed, this provider will
        /// continue trying until it times out.
        /// </summary>
        [Input("waitForInstances")]
        public Input<bool>? WaitForInstances { get; set; }

        /// <summary>
        /// When used with `wait_for_instances` it specifies the status to wait for.
        /// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
        /// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
        /// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
        /// </summary>
        [Input("waitForInstancesStatus")]
        public Input<string>? WaitForInstancesStatus { get; set; }

        /// <summary>
        /// The zone that instances in this group should be created
        /// in.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceGroupManagerState()
        {
        }
    }
}
