// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager
{
    /// <summary>
    /// Allows configuring a single access level condition to be appended to an access level's conditions.
    /// This resource is intended to be used in cases where it is not possible to compile a full list
    /// of conditions to include in a `gcp.accesscontextmanager.AccessLevel` resource,
    /// to enable them to be added separately.
    /// 
    /// &gt; **Note:** If this resource is used alongside a `gcp.accesscontextmanager.AccessLevel` resource,
    /// the access level resource must have a `lifecycle` block with `ignore_changes = [basic[0].conditions]` so
    /// they don't fight over which service accounts should be included.
    /// 
    /// To get more information about AccessLevelCondition, see:
    /// 
    /// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
    /// * How-to Guides
    ///     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
    /// 
    /// &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
    /// you must specify a `billing_project` and set `user_project_override` to true
    /// in the provider configuration. Otherwise the ACM API will return a 403 error.
    /// Your account must have the `serviceusage.services.use` permission on the
    /// `billing_project` you defined.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:accesscontextmanager/accessLevelCondition:AccessLevelCondition")]
    public partial class AccessLevelCondition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Access Level to add this condition to.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// Device specific restrictions, all restrictions must hold for
        /// the Condition to be true. If not specified, all devices are
        /// allowed.
        /// Structure is documented below.
        /// </summary>
        [Output("devicePolicy")]
        public Output<Outputs.AccessLevelConditionDevicePolicy?> DevicePolicy { get; private set; } = null!;

        /// <summary>
        /// A list of CIDR block IP subnetwork specification. May be IPv4
        /// or IPv6.
        /// Note that for a CIDR IP address block, the specified IP address
        /// portion must be properly truncated (i.e. all the host bits must
        /// be zero) or the input is considered malformed. For example,
        /// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
        /// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
        /// is not. The originating IP of a request must be in one of the
        /// listed subnets in order for this Condition to be true.
        /// If empty, all IP addresses are allowed.
        /// </summary>
        [Output("ipSubnetworks")]
        public Output<ImmutableArray<string>> IpSubnetworks { get; private set; } = null!;

        /// <summary>
        /// An allowed list of members (users, service accounts).
        /// Using groups is not supported yet.
        /// The signed-in user originating the request must be a part of one
        /// of the provided members. If not specified, a request may come
        /// from any user (logged in/not logged in, not present in any
        /// groups, etc.).
        /// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Whether to negate the Condition. If true, the Condition becomes
        /// a NAND over its non-empty fields, each field must be false for
        /// the Condition overall to be satisfied. Defaults to false.
        /// </summary>
        [Output("negate")]
        public Output<bool?> Negate { get; private set; } = null!;

        /// <summary>
        /// The request must originate from one of the provided
        /// countries/regions.
        /// Format: A valid ISO 3166-1 alpha-2 code.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// A list of other access levels defined in the same Policy,
        /// referenced by resource name. Referencing an AccessLevel which
        /// does not exist is an error. All access levels listed must be
        /// granted for the Condition to be true.
        /// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Output("requiredAccessLevels")]
        public Output<ImmutableArray<string>> RequiredAccessLevels { get; private set; } = null!;


        /// <summary>
        /// Create a AccessLevelCondition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessLevelCondition(string name, AccessLevelConditionArgs args, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/accessLevelCondition:AccessLevelCondition", name, args ?? new AccessLevelConditionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessLevelCondition(string name, Input<string> id, AccessLevelConditionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/accessLevelCondition:AccessLevelCondition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessLevelCondition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessLevelCondition Get(string name, Input<string> id, AccessLevelConditionState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessLevelCondition(name, id, state, options);
        }
    }

    public sealed class AccessLevelConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Access Level to add this condition to.
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// Device specific restrictions, all restrictions must hold for
        /// the Condition to be true. If not specified, all devices are
        /// allowed.
        /// Structure is documented below.
        /// </summary>
        [Input("devicePolicy")]
        public Input<Inputs.AccessLevelConditionDevicePolicyArgs>? DevicePolicy { get; set; }

        [Input("ipSubnetworks")]
        private InputList<string>? _ipSubnetworks;

        /// <summary>
        /// A list of CIDR block IP subnetwork specification. May be IPv4
        /// or IPv6.
        /// Note that for a CIDR IP address block, the specified IP address
        /// portion must be properly truncated (i.e. all the host bits must
        /// be zero) or the input is considered malformed. For example,
        /// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
        /// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
        /// is not. The originating IP of a request must be in one of the
        /// listed subnets in order for this Condition to be true.
        /// If empty, all IP addresses are allowed.
        /// </summary>
        public InputList<string> IpSubnetworks
        {
            get => _ipSubnetworks ?? (_ipSubnetworks = new InputList<string>());
            set => _ipSubnetworks = value;
        }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// An allowed list of members (users, service accounts).
        /// Using groups is not supported yet.
        /// The signed-in user originating the request must be a part of one
        /// of the provided members. If not specified, a request may come
        /// from any user (logged in/not logged in, not present in any
        /// groups, etc.).
        /// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Whether to negate the Condition. If true, the Condition becomes
        /// a NAND over its non-empty fields, each field must be false for
        /// the Condition overall to be satisfied. Defaults to false.
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// The request must originate from one of the provided
        /// countries/regions.
        /// Format: A valid ISO 3166-1 alpha-2 code.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        [Input("requiredAccessLevels")]
        private InputList<string>? _requiredAccessLevels;

        /// <summary>
        /// A list of other access levels defined in the same Policy,
        /// referenced by resource name. Referencing an AccessLevel which
        /// does not exist is an error. All access levels listed must be
        /// granted for the Condition to be true.
        /// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        public InputList<string> RequiredAccessLevels
        {
            get => _requiredAccessLevels ?? (_requiredAccessLevels = new InputList<string>());
            set => _requiredAccessLevels = value;
        }

        public AccessLevelConditionArgs()
        {
        }
        public static new AccessLevelConditionArgs Empty => new AccessLevelConditionArgs();
    }

    public sealed class AccessLevelConditionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Access Level to add this condition to.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Device specific restrictions, all restrictions must hold for
        /// the Condition to be true. If not specified, all devices are
        /// allowed.
        /// Structure is documented below.
        /// </summary>
        [Input("devicePolicy")]
        public Input<Inputs.AccessLevelConditionDevicePolicyGetArgs>? DevicePolicy { get; set; }

        [Input("ipSubnetworks")]
        private InputList<string>? _ipSubnetworks;

        /// <summary>
        /// A list of CIDR block IP subnetwork specification. May be IPv4
        /// or IPv6.
        /// Note that for a CIDR IP address block, the specified IP address
        /// portion must be properly truncated (i.e. all the host bits must
        /// be zero) or the input is considered malformed. For example,
        /// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
        /// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
        /// is not. The originating IP of a request must be in one of the
        /// listed subnets in order for this Condition to be true.
        /// If empty, all IP addresses are allowed.
        /// </summary>
        public InputList<string> IpSubnetworks
        {
            get => _ipSubnetworks ?? (_ipSubnetworks = new InputList<string>());
            set => _ipSubnetworks = value;
        }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// An allowed list of members (users, service accounts).
        /// Using groups is not supported yet.
        /// The signed-in user originating the request must be a part of one
        /// of the provided members. If not specified, a request may come
        /// from any user (logged in/not logged in, not present in any
        /// groups, etc.).
        /// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Whether to negate the Condition. If true, the Condition becomes
        /// a NAND over its non-empty fields, each field must be false for
        /// the Condition overall to be satisfied. Defaults to false.
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// The request must originate from one of the provided
        /// countries/regions.
        /// Format: A valid ISO 3166-1 alpha-2 code.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        [Input("requiredAccessLevels")]
        private InputList<string>? _requiredAccessLevels;

        /// <summary>
        /// A list of other access levels defined in the same Policy,
        /// referenced by resource name. Referencing an AccessLevel which
        /// does not exist is an error. All access levels listed must be
        /// granted for the Condition to be true.
        /// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        public InputList<string> RequiredAccessLevels
        {
            get => _requiredAccessLevels ?? (_requiredAccessLevels = new InputList<string>());
            set => _requiredAccessLevels = value;
        }

        public AccessLevelConditionState()
        {
        }
        public static new AccessLevelConditionState Empty => new AccessLevelConditionState();
    }
}
