// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager.Inputs
{

    public sealed class ServicePerimeterStatusIngressPolicyIngressFromGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("identities")]
        private InputList<string>? _identities;

        /// <summary>
        /// 'A list of identities that are allowed access through this `IngressPolicy`.
        /// To specify an identity or identity group, use the IAM v1
        /// format specified [here](https://cloud.google.com/iam/docs/principal-identifiers.md#v1).
        /// The following prefixes are supprted: user, group, serviceAccount, principal, and principalSet.'
        /// </summary>
        public InputList<string> Identities
        {
            get => _identities ?? (_identities = new InputList<string>());
            set => _identities = value;
        }

        /// <summary>
        /// Specifies the type of identities that are allowed access from outside the
        /// perimeter. If left unspecified, then members of `identities` field will be
        /// allowed access.
        /// Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.
        /// </summary>
        [Input("identityType")]
        public Input<string>? IdentityType { get; set; }

        [Input("sources")]
        private InputList<Inputs.ServicePerimeterStatusIngressPolicyIngressFromSourceGetArgs>? _sources;

        /// <summary>
        /// Sources that this `IngressPolicy` authorizes access from.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServicePerimeterStatusIngressPolicyIngressFromSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.ServicePerimeterStatusIngressPolicyIngressFromSourceGetArgs>());
            set => _sources = value;
        }

        public ServicePerimeterStatusIngressPolicyIngressFromGetArgs()
        {
        }
        public static new ServicePerimeterStatusIngressPolicyIngressFromGetArgs Empty => new ServicePerimeterStatusIngressPolicyIngressFromGetArgs();
    }
}
