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
    /// ServicePerimeter describes a set of GCP resources which can freely import
    /// and export data amongst themselves, but not export outside of the
    /// ServicePerimeter. If a request with a source within this ServicePerimeter
    /// has a target outside of the ServicePerimeter, the request will be blocked.
    /// Otherwise the request is allowed. There are two types of Service Perimeter
    /// - Regular and Bridge. Regular Service Perimeters cannot overlap, a single
    ///   GCP project can only belong to a single regular Service Perimeter. Service
    ///   Perimeter Bridges can contain only GCP projects as members, a single GCP
    ///   project may belong to multiple Service Perimeter Bridges.
    /// 
    /// To get more information about ServicePerimeter, see:
    /// 
    /// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
    /// * How-to Guides
    ///     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
    /// 
    /// &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
    /// you must specify a `billing_project` and set `user_project_override` to true
    /// in the provider configuration. Otherwise the ACM API will return a 403 error.
    /// Your account must have the `serviceusage.services.use` permission on the
    /// `billing_project` you defined.
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class ServicePerimeter : Pulumi.CustomResource
    {
        /// <summary>
        /// Time the AccessPolicy was created in UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the ServicePerimeter and its use. Does not affect
        /// behavior.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Resource name for the ServicePerimeter. The short_name component must
        /// begin with a letter and only include alphanumeric and '_'.
        /// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AccessPolicy this ServicePerimeter lives in.
        /// Format: accessPolicies/{policy_id}
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of the Perimeter. There are two types: regular and
        /// bridge. Regular Service Perimeter contains resources, access levels,
        /// and restricted services. Every resource can be in at most
        /// ONE regular Service Perimeter.
        /// In addition to being in a regular service perimeter, a resource can also
        /// be in zero or more perimeter bridges. A perimeter bridge only contains
        /// resources. Cross project operations are permitted if all effected
        /// resources share some perimeter (whether bridge or regular). Perimeter
        /// Bridge does not contain access levels or services: those are governed
        /// entirely by the regular perimeter that resource is in.
        /// Perimeter Bridges are typically useful when building more complex
        /// topologies with many independent perimeters that need to share some data
        /// with a common perimeter, but should not be able to share data among
        /// themselves.
        /// Default value is `PERIMETER_TYPE_REGULAR`.
        /// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
        /// </summary>
        [Output("perimeterType")]
        public Output<string?> PerimeterType { get; private set; } = null!;

        /// <summary>
        /// Proposed (or dry run) ServicePerimeter configuration.
        /// This configuration allows to specify and test ServicePerimeter configuration
        /// without enforcing actual access restrictions. Only allowed to be set when
        /// the `useExplicitDryRunSpec` flag is set.
        /// Structure is documented below.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.ServicePerimeterSpec?> Spec { get; private set; } = null!;

        /// <summary>
        /// ServicePerimeter configuration. Specifies sets of resources,
        /// restricted services and access levels that determine
        /// perimeter content and boundaries.
        /// Structure is documented below.
        /// </summary>
        [Output("status")]
        public Output<Outputs.ServicePerimeterStatus?> Status { get; private set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Time the AccessPolicy was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
        /// for all Service Perimeters, and that spec is identical to the status for those
        /// Service Perimeters. When this flag is set, it inhibits the generation of the
        /// implicit spec, thereby allowing the user to explicitly provide a
        /// configuration ("spec") to use in a dry-run version of the Service Perimeter.
        /// This allows the user to test changes to the enforced config ("status") without
        /// actually enforcing them. This testing is done through analyzing the differences
        /// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
        /// bet set to True if any of the fields in the spec are set to non-default values.
        /// </summary>
        [Output("useExplicitDryRunSpec")]
        public Output<bool?> UseExplicitDryRunSpec { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePerimeter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePerimeter(string name, ServicePerimeterArgs args, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/servicePerimeter:ServicePerimeter", name, args ?? new ServicePerimeterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePerimeter(string name, Input<string> id, ServicePerimeterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/servicePerimeter:ServicePerimeter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServicePerimeter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePerimeter Get(string name, Input<string> id, ServicePerimeterState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePerimeter(name, id, state, options);
        }
    }

    public sealed class ServicePerimeterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the ServicePerimeter and its use. Does not affect
        /// behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the ServicePerimeter. The short_name component must
        /// begin with a letter and only include alphanumeric and '_'.
        /// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AccessPolicy this ServicePerimeter lives in.
        /// Format: accessPolicies/{policy_id}
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Specifies the type of the Perimeter. There are two types: regular and
        /// bridge. Regular Service Perimeter contains resources, access levels,
        /// and restricted services. Every resource can be in at most
        /// ONE regular Service Perimeter.
        /// In addition to being in a regular service perimeter, a resource can also
        /// be in zero or more perimeter bridges. A perimeter bridge only contains
        /// resources. Cross project operations are permitted if all effected
        /// resources share some perimeter (whether bridge or regular). Perimeter
        /// Bridge does not contain access levels or services: those are governed
        /// entirely by the regular perimeter that resource is in.
        /// Perimeter Bridges are typically useful when building more complex
        /// topologies with many independent perimeters that need to share some data
        /// with a common perimeter, but should not be able to share data among
        /// themselves.
        /// Default value is `PERIMETER_TYPE_REGULAR`.
        /// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
        /// </summary>
        [Input("perimeterType")]
        public Input<string>? PerimeterType { get; set; }

        /// <summary>
        /// Proposed (or dry run) ServicePerimeter configuration.
        /// This configuration allows to specify and test ServicePerimeter configuration
        /// without enforcing actual access restrictions. Only allowed to be set when
        /// the `useExplicitDryRunSpec` flag is set.
        /// Structure is documented below.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.ServicePerimeterSpecArgs>? Spec { get; set; }

        /// <summary>
        /// ServicePerimeter configuration. Specifies sets of resources,
        /// restricted services and access levels that determine
        /// perimeter content and boundaries.
        /// Structure is documented below.
        /// </summary>
        [Input("status")]
        public Input<Inputs.ServicePerimeterStatusArgs>? Status { get; set; }

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
        /// for all Service Perimeters, and that spec is identical to the status for those
        /// Service Perimeters. When this flag is set, it inhibits the generation of the
        /// implicit spec, thereby allowing the user to explicitly provide a
        /// configuration ("spec") to use in a dry-run version of the Service Perimeter.
        /// This allows the user to test changes to the enforced config ("status") without
        /// actually enforcing them. This testing is done through analyzing the differences
        /// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
        /// bet set to True if any of the fields in the spec are set to non-default values.
        /// </summary>
        [Input("useExplicitDryRunSpec")]
        public Input<bool>? UseExplicitDryRunSpec { get; set; }

        public ServicePerimeterArgs()
        {
        }
    }

    public sealed class ServicePerimeterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time the AccessPolicy was created in UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the ServicePerimeter and its use. Does not affect
        /// behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the ServicePerimeter. The short_name component must
        /// begin with a letter and only include alphanumeric and '_'.
        /// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AccessPolicy this ServicePerimeter lives in.
        /// Format: accessPolicies/{policy_id}
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Specifies the type of the Perimeter. There are two types: regular and
        /// bridge. Regular Service Perimeter contains resources, access levels,
        /// and restricted services. Every resource can be in at most
        /// ONE regular Service Perimeter.
        /// In addition to being in a regular service perimeter, a resource can also
        /// be in zero or more perimeter bridges. A perimeter bridge only contains
        /// resources. Cross project operations are permitted if all effected
        /// resources share some perimeter (whether bridge or regular). Perimeter
        /// Bridge does not contain access levels or services: those are governed
        /// entirely by the regular perimeter that resource is in.
        /// Perimeter Bridges are typically useful when building more complex
        /// topologies with many independent perimeters that need to share some data
        /// with a common perimeter, but should not be able to share data among
        /// themselves.
        /// Default value is `PERIMETER_TYPE_REGULAR`.
        /// Possible values are `PERIMETER_TYPE_REGULAR` and `PERIMETER_TYPE_BRIDGE`.
        /// </summary>
        [Input("perimeterType")]
        public Input<string>? PerimeterType { get; set; }

        /// <summary>
        /// Proposed (or dry run) ServicePerimeter configuration.
        /// This configuration allows to specify and test ServicePerimeter configuration
        /// without enforcing actual access restrictions. Only allowed to be set when
        /// the `useExplicitDryRunSpec` flag is set.
        /// Structure is documented below.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.ServicePerimeterSpecGetArgs>? Spec { get; set; }

        /// <summary>
        /// ServicePerimeter configuration. Specifies sets of resources,
        /// restricted services and access levels that determine
        /// perimeter content and boundaries.
        /// Structure is documented below.
        /// </summary>
        [Input("status")]
        public Input<Inputs.ServicePerimeterStatusGetArgs>? Status { get; set; }

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Time the AccessPolicy was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
        /// for all Service Perimeters, and that spec is identical to the status for those
        /// Service Perimeters. When this flag is set, it inhibits the generation of the
        /// implicit spec, thereby allowing the user to explicitly provide a
        /// configuration ("spec") to use in a dry-run version of the Service Perimeter.
        /// This allows the user to test changes to the enforced config ("status") without
        /// actually enforcing them. This testing is done through analyzing the differences
        /// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
        /// bet set to True if any of the fields in the spec are set to non-default values.
        /// </summary>
        [Input("useExplicitDryRunSpec")]
        public Input<bool>? UseExplicitDryRunSpec { get; set; }

        public ServicePerimeterState()
        {
        }
    }
}
