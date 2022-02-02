// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Recaptcha.Inputs
{

    public sealed class EnterpriseKeyAndroidSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, it means allowed_package_names will not be enforced.
        /// </summary>
        [Input("allowAllPackageNames")]
        public Input<bool>? AllowAllPackageNames { get; set; }

        [Input("allowedPackageNames")]
        private InputList<string>? _allowedPackageNames;

        /// <summary>
        /// Android package names of apps allowed to use the key. Example: 'com.companyname.appname'
        /// </summary>
        public InputList<string> AllowedPackageNames
        {
            get => _allowedPackageNames ?? (_allowedPackageNames = new InputList<string>());
            set => _allowedPackageNames = value;
        }

        public EnterpriseKeyAndroidSettingsGetArgs()
        {
        }
    }
}
