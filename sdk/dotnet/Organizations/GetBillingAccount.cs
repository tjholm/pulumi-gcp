// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static class GetBillingAccount
    {
        /// <summary>
        /// Use this data source to get information about a Google Billing Account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var acct = Gcp.Organizations.GetBillingAccount.Invoke(new()
        ///     {
        ///         DisplayName = "My Billing Account",
        ///         Open = true,
        ///     });
        /// 
        ///     var myProject = new Gcp.Organizations.Project("myProject", new()
        ///     {
        ///         ProjectId = "your-project-id",
        ///         OrgId = "1234567",
        ///         BillingAccount = acct.Apply(getBillingAccountResult =&gt; getBillingAccountResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBillingAccountResult> InvokeAsync(GetBillingAccountArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBillingAccountResult>("gcp:organizations/getBillingAccount:getBillingAccount", args ?? new GetBillingAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Google Billing Account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var acct = Gcp.Organizations.GetBillingAccount.Invoke(new()
        ///     {
        ///         DisplayName = "My Billing Account",
        ///         Open = true,
        ///     });
        /// 
        ///     var myProject = new Gcp.Organizations.Project("myProject", new()
        ///     {
        ///         ProjectId = "your-project-id",
        ///         OrgId = "1234567",
        ///         BillingAccount = acct.Apply(getBillingAccountResult =&gt; getBillingAccountResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBillingAccountResult> Invoke(GetBillingAccountInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingAccountResult>("gcp:organizations/getBillingAccount:getBillingAccount", args ?? new GetBillingAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBillingAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the billing account in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        /// </summary>
        [Input("billingAccount")]
        public string? BillingAccount { get; set; }

        /// <summary>
        /// The display name of the billing account.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// `true` if projects associated with the billing account should be read, `false` if this step
        /// should be skipped. Setting `false` may be useful if the user permissions do not allow listing projects. Defaults to `true`.
        /// 
        /// &gt; **NOTE:** One of `billing_account` or `display_name` must be specified.
        /// </summary>
        [Input("lookupProjects")]
        public bool? LookupProjects { get; set; }

        /// <summary>
        /// `true` if the billing account is open, `false` if the billing account is closed.
        /// </summary>
        [Input("open")]
        public bool? Open { get; set; }

        public GetBillingAccountArgs()
        {
        }
        public static new GetBillingAccountArgs Empty => new GetBillingAccountArgs();
    }

    public sealed class GetBillingAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the billing account in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// The display name of the billing account.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// `true` if projects associated with the billing account should be read, `false` if this step
        /// should be skipped. Setting `false` may be useful if the user permissions do not allow listing projects. Defaults to `true`.
        /// 
        /// &gt; **NOTE:** One of `billing_account` or `display_name` must be specified.
        /// </summary>
        [Input("lookupProjects")]
        public Input<bool>? LookupProjects { get; set; }

        /// <summary>
        /// `true` if the billing account is open, `false` if the billing account is closed.
        /// </summary>
        [Input("open")]
        public Input<bool>? Open { get; set; }

        public GetBillingAccountInvokeArgs()
        {
        }
        public static new GetBillingAccountInvokeArgs Empty => new GetBillingAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetBillingAccountResult
    {
        public readonly string? BillingAccount;
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? LookupProjects;
        /// <summary>
        /// The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
        /// </summary>
        public readonly string Name;
        public readonly bool Open;
        /// <summary>
        /// The IDs of any projects associated with the billing account. `lookup_projects` must not be false
        /// for this to be populated.
        /// </summary>
        public readonly ImmutableArray<string> ProjectIds;

        [OutputConstructor]
        private GetBillingAccountResult(
            string? billingAccount,

            string displayName,

            string id,

            bool? lookupProjects,

            string name,

            bool open,

            ImmutableArray<string> projectIds)
        {
            BillingAccount = billingAccount;
            DisplayName = displayName;
            Id = id;
            LookupProjects = lookupProjects;
            Name = name;
            Open = open;
            ProjectIds = projectIds;
        }
    }
}
