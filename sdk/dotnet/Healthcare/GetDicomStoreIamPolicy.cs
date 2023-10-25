// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    public static class GetDicomStoreIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for a Google Cloud Healthcare DICOM store.
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Gcp.Healthcare.GetDicomStoreIamPolicy.Invoke(new()
        ///     {
        ///         DicomStoreId = google_healthcare_dicom_store.Dicom_store.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDicomStoreIamPolicyResult> InvokeAsync(GetDicomStoreIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDicomStoreIamPolicyResult>("gcp:healthcare/getDicomStoreIamPolicy:getDicomStoreIamPolicy", args ?? new GetDicomStoreIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for a Google Cloud Healthcare DICOM store.
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Gcp.Healthcare.GetDicomStoreIamPolicy.Invoke(new()
        ///     {
        ///         DicomStoreId = google_healthcare_dicom_store.Dicom_store.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDicomStoreIamPolicyResult> Invoke(GetDicomStoreIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDicomStoreIamPolicyResult>("gcp:healthcare/getDicomStoreIamPolicy:getDicomStoreIamPolicy", args ?? new GetDicomStoreIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDicomStoreIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DICOM store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        /// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("dicomStoreId", required: true)]
        public string DicomStoreId { get; set; } = null!;

        public GetDicomStoreIamPolicyArgs()
        {
        }
        public static new GetDicomStoreIamPolicyArgs Empty => new GetDicomStoreIamPolicyArgs();
    }

    public sealed class GetDicomStoreIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DICOM store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        /// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("dicomStoreId", required: true)]
        public Input<string> DicomStoreId { get; set; } = null!;

        public GetDicomStoreIamPolicyInvokeArgs()
        {
        }
        public static new GetDicomStoreIamPolicyInvokeArgs Empty => new GetDicomStoreIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetDicomStoreIamPolicyResult
    {
        public readonly string DicomStoreId;
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The policy data
        /// </summary>
        public readonly string PolicyData;

        [OutputConstructor]
        private GetDicomStoreIamPolicyResult(
            string dicomStoreId,

            string etag,

            string id,

            string policyData)
        {
            DicomStoreId = dicomStoreId;
            Etag = etag;
            Id = id;
            PolicyData = policyData;
        }
    }
}
