// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp
{
    /// <summary>
    /// The provider type for the google-beta package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [GcpResourceType("pulumi:providers:gcp")]
    public partial class Provider : Pulumi.ProviderResource
    {
        [Output("accessApprovalCustomEndpoint")]
        public Output<string?> AccessApprovalCustomEndpoint { get; private set; } = null!;

        [Output("accessContextManagerCustomEndpoint")]
        public Output<string?> AccessContextManagerCustomEndpoint { get; private set; } = null!;

        [Output("accessToken")]
        public Output<string?> AccessToken { get; private set; } = null!;

        [Output("activeDirectoryCustomEndpoint")]
        public Output<string?> ActiveDirectoryCustomEndpoint { get; private set; } = null!;

        [Output("apiGatewayCustomEndpoint")]
        public Output<string?> ApiGatewayCustomEndpoint { get; private set; } = null!;

        [Output("apigeeCustomEndpoint")]
        public Output<string?> ApigeeCustomEndpoint { get; private set; } = null!;

        [Output("apikeysCustomEndpoint")]
        public Output<string?> ApikeysCustomEndpoint { get; private set; } = null!;

        [Output("appEngineCustomEndpoint")]
        public Output<string?> AppEngineCustomEndpoint { get; private set; } = null!;

        [Output("artifactRegistryCustomEndpoint")]
        public Output<string?> ArtifactRegistryCustomEndpoint { get; private set; } = null!;

        [Output("assuredWorkloadsCustomEndpoint")]
        public Output<string?> AssuredWorkloadsCustomEndpoint { get; private set; } = null!;

        [Output("bigQueryCustomEndpoint")]
        public Output<string?> BigQueryCustomEndpoint { get; private set; } = null!;

        [Output("bigqueryConnectionCustomEndpoint")]
        public Output<string?> BigqueryConnectionCustomEndpoint { get; private set; } = null!;

        [Output("bigqueryDataTransferCustomEndpoint")]
        public Output<string?> BigqueryDataTransferCustomEndpoint { get; private set; } = null!;

        [Output("bigqueryReservationCustomEndpoint")]
        public Output<string?> BigqueryReservationCustomEndpoint { get; private set; } = null!;

        [Output("bigtableCustomEndpoint")]
        public Output<string?> BigtableCustomEndpoint { get; private set; } = null!;

        [Output("billingCustomEndpoint")]
        public Output<string?> BillingCustomEndpoint { get; private set; } = null!;

        [Output("billingProject")]
        public Output<string?> BillingProject { get; private set; } = null!;

        [Output("binaryAuthorizationCustomEndpoint")]
        public Output<string?> BinaryAuthorizationCustomEndpoint { get; private set; } = null!;

        [Output("certificateManagerCustomEndpoint")]
        public Output<string?> CertificateManagerCustomEndpoint { get; private set; } = null!;

        [Output("cloudAssetCustomEndpoint")]
        public Output<string?> CloudAssetCustomEndpoint { get; private set; } = null!;

        [Output("cloudBillingCustomEndpoint")]
        public Output<string?> CloudBillingCustomEndpoint { get; private set; } = null!;

        [Output("cloudBuildCustomEndpoint")]
        public Output<string?> CloudBuildCustomEndpoint { get; private set; } = null!;

        [Output("cloudBuildWorkerPoolCustomEndpoint")]
        public Output<string?> CloudBuildWorkerPoolCustomEndpoint { get; private set; } = null!;

        [Output("cloudFunctionsCustomEndpoint")]
        public Output<string?> CloudFunctionsCustomEndpoint { get; private set; } = null!;

        [Output("cloudIdentityCustomEndpoint")]
        public Output<string?> CloudIdentityCustomEndpoint { get; private set; } = null!;

        [Output("cloudIotCustomEndpoint")]
        public Output<string?> CloudIotCustomEndpoint { get; private set; } = null!;

        [Output("cloudResourceManagerCustomEndpoint")]
        public Output<string?> CloudResourceManagerCustomEndpoint { get; private set; } = null!;

        [Output("cloudRunCustomEndpoint")]
        public Output<string?> CloudRunCustomEndpoint { get; private set; } = null!;

        [Output("cloudSchedulerCustomEndpoint")]
        public Output<string?> CloudSchedulerCustomEndpoint { get; private set; } = null!;

        [Output("cloudTasksCustomEndpoint")]
        public Output<string?> CloudTasksCustomEndpoint { get; private set; } = null!;

        [Output("clouddeployCustomEndpoint")]
        public Output<string?> ClouddeployCustomEndpoint { get; private set; } = null!;

        [Output("cloudfunctions2CustomEndpoint")]
        public Output<string?> Cloudfunctions2CustomEndpoint { get; private set; } = null!;

        [Output("composerCustomEndpoint")]
        public Output<string?> ComposerCustomEndpoint { get; private set; } = null!;

        [Output("computeCustomEndpoint")]
        public Output<string?> ComputeCustomEndpoint { get; private set; } = null!;

        [Output("containerAnalysisCustomEndpoint")]
        public Output<string?> ContainerAnalysisCustomEndpoint { get; private set; } = null!;

        [Output("containerAwsCustomEndpoint")]
        public Output<string?> ContainerAwsCustomEndpoint { get; private set; } = null!;

        [Output("containerAzureCustomEndpoint")]
        public Output<string?> ContainerAzureCustomEndpoint { get; private set; } = null!;

        [Output("containerCustomEndpoint")]
        public Output<string?> ContainerCustomEndpoint { get; private set; } = null!;

        [Output("credentials")]
        public Output<string?> Credentials { get; private set; } = null!;

        [Output("dataCatalogCustomEndpoint")]
        public Output<string?> DataCatalogCustomEndpoint { get; private set; } = null!;

        [Output("dataFusionCustomEndpoint")]
        public Output<string?> DataFusionCustomEndpoint { get; private set; } = null!;

        [Output("dataLossPreventionCustomEndpoint")]
        public Output<string?> DataLossPreventionCustomEndpoint { get; private set; } = null!;

        [Output("dataflowCustomEndpoint")]
        public Output<string?> DataflowCustomEndpoint { get; private set; } = null!;

        [Output("dataplexCustomEndpoint")]
        public Output<string?> DataplexCustomEndpoint { get; private set; } = null!;

        [Output("dataprocCustomEndpoint")]
        public Output<string?> DataprocCustomEndpoint { get; private set; } = null!;

        [Output("dataprocMetastoreCustomEndpoint")]
        public Output<string?> DataprocMetastoreCustomEndpoint { get; private set; } = null!;

        [Output("datastoreCustomEndpoint")]
        public Output<string?> DatastoreCustomEndpoint { get; private set; } = null!;

        [Output("deploymentManagerCustomEndpoint")]
        public Output<string?> DeploymentManagerCustomEndpoint { get; private set; } = null!;

        [Output("dialogflowCustomEndpoint")]
        public Output<string?> DialogflowCustomEndpoint { get; private set; } = null!;

        [Output("dialogflowCxCustomEndpoint")]
        public Output<string?> DialogflowCxCustomEndpoint { get; private set; } = null!;

        [Output("dnsCustomEndpoint")]
        public Output<string?> DnsCustomEndpoint { get; private set; } = null!;

        [Output("essentialContactsCustomEndpoint")]
        public Output<string?> EssentialContactsCustomEndpoint { get; private set; } = null!;

        [Output("eventarcCustomEndpoint")]
        public Output<string?> EventarcCustomEndpoint { get; private set; } = null!;

        [Output("filestoreCustomEndpoint")]
        public Output<string?> FilestoreCustomEndpoint { get; private set; } = null!;

        [Output("firebaseCustomEndpoint")]
        public Output<string?> FirebaseCustomEndpoint { get; private set; } = null!;

        [Output("firebaserulesCustomEndpoint")]
        public Output<string?> FirebaserulesCustomEndpoint { get; private set; } = null!;

        [Output("firestoreCustomEndpoint")]
        public Output<string?> FirestoreCustomEndpoint { get; private set; } = null!;

        [Output("gameServicesCustomEndpoint")]
        public Output<string?> GameServicesCustomEndpoint { get; private set; } = null!;

        [Output("gkeHubCustomEndpoint")]
        public Output<string?> GkeHubCustomEndpoint { get; private set; } = null!;

        [Output("gkehubFeatureCustomEndpoint")]
        public Output<string?> GkehubFeatureCustomEndpoint { get; private set; } = null!;

        [Output("googlePartnerName")]
        public Output<string?> GooglePartnerName { get; private set; } = null!;

        [Output("healthcareCustomEndpoint")]
        public Output<string?> HealthcareCustomEndpoint { get; private set; } = null!;

        [Output("iam2CustomEndpoint")]
        public Output<string?> Iam2CustomEndpoint { get; private set; } = null!;

        [Output("iamBetaCustomEndpoint")]
        public Output<string?> IamBetaCustomEndpoint { get; private set; } = null!;

        [Output("iamCredentialsCustomEndpoint")]
        public Output<string?> IamCredentialsCustomEndpoint { get; private set; } = null!;

        [Output("iamCustomEndpoint")]
        public Output<string?> IamCustomEndpoint { get; private set; } = null!;

        [Output("iapCustomEndpoint")]
        public Output<string?> IapCustomEndpoint { get; private set; } = null!;

        [Output("identityPlatformCustomEndpoint")]
        public Output<string?> IdentityPlatformCustomEndpoint { get; private set; } = null!;

        [Output("impersonateServiceAccount")]
        public Output<string?> ImpersonateServiceAccount { get; private set; } = null!;

        [Output("kmsCustomEndpoint")]
        public Output<string?> KmsCustomEndpoint { get; private set; } = null!;

        [Output("loggingCustomEndpoint")]
        public Output<string?> LoggingCustomEndpoint { get; private set; } = null!;

        [Output("memcacheCustomEndpoint")]
        public Output<string?> MemcacheCustomEndpoint { get; private set; } = null!;

        [Output("mlEngineCustomEndpoint")]
        public Output<string?> MlEngineCustomEndpoint { get; private set; } = null!;

        [Output("monitoringCustomEndpoint")]
        public Output<string?> MonitoringCustomEndpoint { get; private set; } = null!;

        [Output("networkConnectivityCustomEndpoint")]
        public Output<string?> NetworkConnectivityCustomEndpoint { get; private set; } = null!;

        [Output("networkManagementCustomEndpoint")]
        public Output<string?> NetworkManagementCustomEndpoint { get; private set; } = null!;

        [Output("networkServicesCustomEndpoint")]
        public Output<string?> NetworkServicesCustomEndpoint { get; private set; } = null!;

        [Output("notebooksCustomEndpoint")]
        public Output<string?> NotebooksCustomEndpoint { get; private set; } = null!;

        [Output("orgPolicyCustomEndpoint")]
        public Output<string?> OrgPolicyCustomEndpoint { get; private set; } = null!;

        [Output("osConfigCustomEndpoint")]
        public Output<string?> OsConfigCustomEndpoint { get; private set; } = null!;

        [Output("osLoginCustomEndpoint")]
        public Output<string?> OsLoginCustomEndpoint { get; private set; } = null!;

        [Output("privatecaCustomEndpoint")]
        public Output<string?> PrivatecaCustomEndpoint { get; private set; } = null!;

        [Output("project")]
        public Output<string?> Project { get; private set; } = null!;

        [Output("pubsubCustomEndpoint")]
        public Output<string?> PubsubCustomEndpoint { get; private set; } = null!;

        [Output("pubsubLiteCustomEndpoint")]
        public Output<string?> PubsubLiteCustomEndpoint { get; private set; } = null!;

        [Output("recaptchaEnterpriseCustomEndpoint")]
        public Output<string?> RecaptchaEnterpriseCustomEndpoint { get; private set; } = null!;

        [Output("redisCustomEndpoint")]
        public Output<string?> RedisCustomEndpoint { get; private set; } = null!;

        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        [Output("requestReason")]
        public Output<string?> RequestReason { get; private set; } = null!;

        [Output("requestTimeout")]
        public Output<string?> RequestTimeout { get; private set; } = null!;

        [Output("resourceManagerCustomEndpoint")]
        public Output<string?> ResourceManagerCustomEndpoint { get; private set; } = null!;

        [Output("resourceManagerV3CustomEndpoint")]
        public Output<string?> ResourceManagerV3CustomEndpoint { get; private set; } = null!;

        [Output("runtimeConfigCustomEndpoint")]
        public Output<string?> RuntimeConfigCustomEndpoint { get; private set; } = null!;

        [Output("runtimeconfigCustomEndpoint")]
        public Output<string?> RuntimeconfigCustomEndpoint { get; private set; } = null!;

        [Output("secretManagerCustomEndpoint")]
        public Output<string?> SecretManagerCustomEndpoint { get; private set; } = null!;

        [Output("securityCenterCustomEndpoint")]
        public Output<string?> SecurityCenterCustomEndpoint { get; private set; } = null!;

        [Output("securityScannerCustomEndpoint")]
        public Output<string?> SecurityScannerCustomEndpoint { get; private set; } = null!;

        [Output("serviceDirectoryCustomEndpoint")]
        public Output<string?> ServiceDirectoryCustomEndpoint { get; private set; } = null!;

        [Output("serviceManagementCustomEndpoint")]
        public Output<string?> ServiceManagementCustomEndpoint { get; private set; } = null!;

        [Output("serviceNetworkingCustomEndpoint")]
        public Output<string?> ServiceNetworkingCustomEndpoint { get; private set; } = null!;

        [Output("serviceUsageCustomEndpoint")]
        public Output<string?> ServiceUsageCustomEndpoint { get; private set; } = null!;

        [Output("sourceRepoCustomEndpoint")]
        public Output<string?> SourceRepoCustomEndpoint { get; private set; } = null!;

        [Output("spannerCustomEndpoint")]
        public Output<string?> SpannerCustomEndpoint { get; private set; } = null!;

        [Output("sqlCustomEndpoint")]
        public Output<string?> SqlCustomEndpoint { get; private set; } = null!;

        [Output("storageCustomEndpoint")]
        public Output<string?> StorageCustomEndpoint { get; private set; } = null!;

        [Output("storageTransferCustomEndpoint")]
        public Output<string?> StorageTransferCustomEndpoint { get; private set; } = null!;

        [Output("tagsCustomEndpoint")]
        public Output<string?> TagsCustomEndpoint { get; private set; } = null!;

        [Output("tpuCustomEndpoint")]
        public Output<string?> TpuCustomEndpoint { get; private set; } = null!;

        [Output("vertexAiCustomEndpoint")]
        public Output<string?> VertexAiCustomEndpoint { get; private set; } = null!;

        [Output("vpcAccessCustomEndpoint")]
        public Output<string?> VpcAccessCustomEndpoint { get; private set; } = null!;

        [Output("workflowsCustomEndpoint")]
        public Output<string?> WorkflowsCustomEndpoint { get; private set; } = null!;

        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        [Input("accessApprovalCustomEndpoint")]
        public Input<string>? AccessApprovalCustomEndpoint { get; set; }

        [Input("accessContextManagerCustomEndpoint")]
        public Input<string>? AccessContextManagerCustomEndpoint { get; set; }

        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        [Input("activeDirectoryCustomEndpoint")]
        public Input<string>? ActiveDirectoryCustomEndpoint { get; set; }

        [Input("apiGatewayCustomEndpoint")]
        public Input<string>? ApiGatewayCustomEndpoint { get; set; }

        [Input("apigeeCustomEndpoint")]
        public Input<string>? ApigeeCustomEndpoint { get; set; }

        [Input("apikeysCustomEndpoint")]
        public Input<string>? ApikeysCustomEndpoint { get; set; }

        [Input("appEngineCustomEndpoint")]
        public Input<string>? AppEngineCustomEndpoint { get; set; }

        [Input("artifactRegistryCustomEndpoint")]
        public Input<string>? ArtifactRegistryCustomEndpoint { get; set; }

        [Input("assuredWorkloadsCustomEndpoint")]
        public Input<string>? AssuredWorkloadsCustomEndpoint { get; set; }

        [Input("batching", json: true)]
        public Input<Inputs.ProviderBatchingArgs>? Batching { get; set; }

        [Input("bigQueryCustomEndpoint")]
        public Input<string>? BigQueryCustomEndpoint { get; set; }

        [Input("bigqueryConnectionCustomEndpoint")]
        public Input<string>? BigqueryConnectionCustomEndpoint { get; set; }

        [Input("bigqueryDataTransferCustomEndpoint")]
        public Input<string>? BigqueryDataTransferCustomEndpoint { get; set; }

        [Input("bigqueryReservationCustomEndpoint")]
        public Input<string>? BigqueryReservationCustomEndpoint { get; set; }

        [Input("bigtableCustomEndpoint")]
        public Input<string>? BigtableCustomEndpoint { get; set; }

        [Input("billingCustomEndpoint")]
        public Input<string>? BillingCustomEndpoint { get; set; }

        [Input("billingProject")]
        public Input<string>? BillingProject { get; set; }

        [Input("binaryAuthorizationCustomEndpoint")]
        public Input<string>? BinaryAuthorizationCustomEndpoint { get; set; }

        [Input("certificateManagerCustomEndpoint")]
        public Input<string>? CertificateManagerCustomEndpoint { get; set; }

        [Input("cloudAssetCustomEndpoint")]
        public Input<string>? CloudAssetCustomEndpoint { get; set; }

        [Input("cloudBillingCustomEndpoint")]
        public Input<string>? CloudBillingCustomEndpoint { get; set; }

        [Input("cloudBuildCustomEndpoint")]
        public Input<string>? CloudBuildCustomEndpoint { get; set; }

        [Input("cloudBuildWorkerPoolCustomEndpoint")]
        public Input<string>? CloudBuildWorkerPoolCustomEndpoint { get; set; }

        [Input("cloudFunctionsCustomEndpoint")]
        public Input<string>? CloudFunctionsCustomEndpoint { get; set; }

        [Input("cloudIdentityCustomEndpoint")]
        public Input<string>? CloudIdentityCustomEndpoint { get; set; }

        [Input("cloudIotCustomEndpoint")]
        public Input<string>? CloudIotCustomEndpoint { get; set; }

        [Input("cloudResourceManagerCustomEndpoint")]
        public Input<string>? CloudResourceManagerCustomEndpoint { get; set; }

        [Input("cloudRunCustomEndpoint")]
        public Input<string>? CloudRunCustomEndpoint { get; set; }

        [Input("cloudSchedulerCustomEndpoint")]
        public Input<string>? CloudSchedulerCustomEndpoint { get; set; }

        [Input("cloudTasksCustomEndpoint")]
        public Input<string>? CloudTasksCustomEndpoint { get; set; }

        [Input("clouddeployCustomEndpoint")]
        public Input<string>? ClouddeployCustomEndpoint { get; set; }

        [Input("cloudfunctions2CustomEndpoint")]
        public Input<string>? Cloudfunctions2CustomEndpoint { get; set; }

        [Input("composerCustomEndpoint")]
        public Input<string>? ComposerCustomEndpoint { get; set; }

        [Input("computeCustomEndpoint")]
        public Input<string>? ComputeCustomEndpoint { get; set; }

        [Input("containerAnalysisCustomEndpoint")]
        public Input<string>? ContainerAnalysisCustomEndpoint { get; set; }

        [Input("containerAwsCustomEndpoint")]
        public Input<string>? ContainerAwsCustomEndpoint { get; set; }

        [Input("containerAzureCustomEndpoint")]
        public Input<string>? ContainerAzureCustomEndpoint { get; set; }

        [Input("containerCustomEndpoint")]
        public Input<string>? ContainerCustomEndpoint { get; set; }

        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        [Input("dataCatalogCustomEndpoint")]
        public Input<string>? DataCatalogCustomEndpoint { get; set; }

        [Input("dataFusionCustomEndpoint")]
        public Input<string>? DataFusionCustomEndpoint { get; set; }

        [Input("dataLossPreventionCustomEndpoint")]
        public Input<string>? DataLossPreventionCustomEndpoint { get; set; }

        [Input("dataflowCustomEndpoint")]
        public Input<string>? DataflowCustomEndpoint { get; set; }

        [Input("dataplexCustomEndpoint")]
        public Input<string>? DataplexCustomEndpoint { get; set; }

        [Input("dataprocCustomEndpoint")]
        public Input<string>? DataprocCustomEndpoint { get; set; }

        [Input("dataprocMetastoreCustomEndpoint")]
        public Input<string>? DataprocMetastoreCustomEndpoint { get; set; }

        [Input("datastoreCustomEndpoint")]
        public Input<string>? DatastoreCustomEndpoint { get; set; }

        [Input("deploymentManagerCustomEndpoint")]
        public Input<string>? DeploymentManagerCustomEndpoint { get; set; }

        [Input("dialogflowCustomEndpoint")]
        public Input<string>? DialogflowCustomEndpoint { get; set; }

        [Input("dialogflowCxCustomEndpoint")]
        public Input<string>? DialogflowCxCustomEndpoint { get; set; }

        [Input("disableGooglePartnerName", json: true)]
        public Input<bool>? DisableGooglePartnerName { get; set; }

        [Input("dnsCustomEndpoint")]
        public Input<string>? DnsCustomEndpoint { get; set; }

        [Input("essentialContactsCustomEndpoint")]
        public Input<string>? EssentialContactsCustomEndpoint { get; set; }

        [Input("eventarcCustomEndpoint")]
        public Input<string>? EventarcCustomEndpoint { get; set; }

        [Input("filestoreCustomEndpoint")]
        public Input<string>? FilestoreCustomEndpoint { get; set; }

        [Input("firebaseCustomEndpoint")]
        public Input<string>? FirebaseCustomEndpoint { get; set; }

        [Input("firebaserulesCustomEndpoint")]
        public Input<string>? FirebaserulesCustomEndpoint { get; set; }

        [Input("firestoreCustomEndpoint")]
        public Input<string>? FirestoreCustomEndpoint { get; set; }

        [Input("gameServicesCustomEndpoint")]
        public Input<string>? GameServicesCustomEndpoint { get; set; }

        [Input("gkeHubCustomEndpoint")]
        public Input<string>? GkeHubCustomEndpoint { get; set; }

        [Input("gkehubFeatureCustomEndpoint")]
        public Input<string>? GkehubFeatureCustomEndpoint { get; set; }

        [Input("googlePartnerName")]
        public Input<string>? GooglePartnerName { get; set; }

        [Input("healthcareCustomEndpoint")]
        public Input<string>? HealthcareCustomEndpoint { get; set; }

        [Input("iam2CustomEndpoint")]
        public Input<string>? Iam2CustomEndpoint { get; set; }

        [Input("iamBetaCustomEndpoint")]
        public Input<string>? IamBetaCustomEndpoint { get; set; }

        [Input("iamCredentialsCustomEndpoint")]
        public Input<string>? IamCredentialsCustomEndpoint { get; set; }

        [Input("iamCustomEndpoint")]
        public Input<string>? IamCustomEndpoint { get; set; }

        [Input("iapCustomEndpoint")]
        public Input<string>? IapCustomEndpoint { get; set; }

        [Input("identityPlatformCustomEndpoint")]
        public Input<string>? IdentityPlatformCustomEndpoint { get; set; }

        [Input("impersonateServiceAccount")]
        public Input<string>? ImpersonateServiceAccount { get; set; }

        [Input("impersonateServiceAccountDelegates", json: true)]
        private InputList<string>? _impersonateServiceAccountDelegates;
        public InputList<string> ImpersonateServiceAccountDelegates
        {
            get => _impersonateServiceAccountDelegates ?? (_impersonateServiceAccountDelegates = new InputList<string>());
            set => _impersonateServiceAccountDelegates = value;
        }

        [Input("kmsCustomEndpoint")]
        public Input<string>? KmsCustomEndpoint { get; set; }

        [Input("loggingCustomEndpoint")]
        public Input<string>? LoggingCustomEndpoint { get; set; }

        [Input("memcacheCustomEndpoint")]
        public Input<string>? MemcacheCustomEndpoint { get; set; }

        [Input("mlEngineCustomEndpoint")]
        public Input<string>? MlEngineCustomEndpoint { get; set; }

        [Input("monitoringCustomEndpoint")]
        public Input<string>? MonitoringCustomEndpoint { get; set; }

        [Input("networkConnectivityCustomEndpoint")]
        public Input<string>? NetworkConnectivityCustomEndpoint { get; set; }

        [Input("networkManagementCustomEndpoint")]
        public Input<string>? NetworkManagementCustomEndpoint { get; set; }

        [Input("networkServicesCustomEndpoint")]
        public Input<string>? NetworkServicesCustomEndpoint { get; set; }

        [Input("notebooksCustomEndpoint")]
        public Input<string>? NotebooksCustomEndpoint { get; set; }

        [Input("orgPolicyCustomEndpoint")]
        public Input<string>? OrgPolicyCustomEndpoint { get; set; }

        [Input("osConfigCustomEndpoint")]
        public Input<string>? OsConfigCustomEndpoint { get; set; }

        [Input("osLoginCustomEndpoint")]
        public Input<string>? OsLoginCustomEndpoint { get; set; }

        [Input("privatecaCustomEndpoint")]
        public Input<string>? PrivatecaCustomEndpoint { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pubsubCustomEndpoint")]
        public Input<string>? PubsubCustomEndpoint { get; set; }

        [Input("pubsubLiteCustomEndpoint")]
        public Input<string>? PubsubLiteCustomEndpoint { get; set; }

        [Input("recaptchaEnterpriseCustomEndpoint")]
        public Input<string>? RecaptchaEnterpriseCustomEndpoint { get; set; }

        [Input("redisCustomEndpoint")]
        public Input<string>? RedisCustomEndpoint { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("requestReason")]
        public Input<string>? RequestReason { get; set; }

        [Input("requestTimeout")]
        public Input<string>? RequestTimeout { get; set; }

        [Input("resourceManagerCustomEndpoint")]
        public Input<string>? ResourceManagerCustomEndpoint { get; set; }

        [Input("resourceManagerV3CustomEndpoint")]
        public Input<string>? ResourceManagerV3CustomEndpoint { get; set; }

        [Input("runtimeConfigCustomEndpoint")]
        public Input<string>? RuntimeConfigCustomEndpoint { get; set; }

        [Input("runtimeconfigCustomEndpoint")]
        public Input<string>? RuntimeconfigCustomEndpoint { get; set; }

        [Input("scopes", json: true)]
        private InputList<string>? _scopes;
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("secretManagerCustomEndpoint")]
        public Input<string>? SecretManagerCustomEndpoint { get; set; }

        [Input("securityCenterCustomEndpoint")]
        public Input<string>? SecurityCenterCustomEndpoint { get; set; }

        [Input("securityScannerCustomEndpoint")]
        public Input<string>? SecurityScannerCustomEndpoint { get; set; }

        [Input("serviceDirectoryCustomEndpoint")]
        public Input<string>? ServiceDirectoryCustomEndpoint { get; set; }

        [Input("serviceManagementCustomEndpoint")]
        public Input<string>? ServiceManagementCustomEndpoint { get; set; }

        [Input("serviceNetworkingCustomEndpoint")]
        public Input<string>? ServiceNetworkingCustomEndpoint { get; set; }

        [Input("serviceUsageCustomEndpoint")]
        public Input<string>? ServiceUsageCustomEndpoint { get; set; }

        [Input("sourceRepoCustomEndpoint")]
        public Input<string>? SourceRepoCustomEndpoint { get; set; }

        [Input("spannerCustomEndpoint")]
        public Input<string>? SpannerCustomEndpoint { get; set; }

        [Input("sqlCustomEndpoint")]
        public Input<string>? SqlCustomEndpoint { get; set; }

        [Input("storageCustomEndpoint")]
        public Input<string>? StorageCustomEndpoint { get; set; }

        [Input("storageTransferCustomEndpoint")]
        public Input<string>? StorageTransferCustomEndpoint { get; set; }

        [Input("tagsCustomEndpoint")]
        public Input<string>? TagsCustomEndpoint { get; set; }

        [Input("tpuCustomEndpoint")]
        public Input<string>? TpuCustomEndpoint { get; set; }

        [Input("userProjectOverride", json: true)]
        public Input<bool>? UserProjectOverride { get; set; }

        [Input("vertexAiCustomEndpoint")]
        public Input<string>? VertexAiCustomEndpoint { get; set; }

        [Input("vpcAccessCustomEndpoint")]
        public Input<string>? VpcAccessCustomEndpoint { get; set; }

        [Input("workflowsCustomEndpoint")]
        public Input<string>? WorkflowsCustomEndpoint { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ProviderArgs()
        {
            Project = Utilities.GetEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");
            Region = Utilities.GetEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");
            Zone = Utilities.GetEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");
        }
    }
}
