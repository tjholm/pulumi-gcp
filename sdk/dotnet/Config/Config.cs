// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Gcp
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("gcp");
        public static string? AccessApprovalCustomEndpoint { get; set; } = __config.Get("accessApprovalCustomEndpoint");

        public static string? AccessContextManagerCustomEndpoint { get; set; } = __config.Get("accessContextManagerCustomEndpoint");

        public static string? AccessToken { get; set; } = __config.Get("accessToken");

        public static string? ActiveDirectoryCustomEndpoint { get; set; } = __config.Get("activeDirectoryCustomEndpoint");

        public static string? ApiGatewayCustomEndpoint { get; set; } = __config.Get("apiGatewayCustomEndpoint");

        public static string? ApigeeCustomEndpoint { get; set; } = __config.Get("apigeeCustomEndpoint");

        public static string? AppEngineCustomEndpoint { get; set; } = __config.Get("appEngineCustomEndpoint");

        public static string? ArtifactRegistryCustomEndpoint { get; set; } = __config.Get("artifactRegistryCustomEndpoint");

        public static Pulumi.Gcp.Config.Types.Batching? Batching { get; set; } = __config.GetObject<Pulumi.Gcp.Config.Types.Batching>("batching");

        public static string? BigQueryCustomEndpoint { get; set; } = __config.Get("bigQueryCustomEndpoint");

        public static string? BigqueryConnectionCustomEndpoint { get; set; } = __config.Get("bigqueryConnectionCustomEndpoint");

        public static string? BigqueryDataTransferCustomEndpoint { get; set; } = __config.Get("bigqueryDataTransferCustomEndpoint");

        public static string? BigqueryReservationCustomEndpoint { get; set; } = __config.Get("bigqueryReservationCustomEndpoint");

        public static string? BigtableCustomEndpoint { get; set; } = __config.Get("bigtableCustomEndpoint");

        public static string? BillingCustomEndpoint { get; set; } = __config.Get("billingCustomEndpoint");

        public static string? BillingProject { get; set; } = __config.Get("billingProject");

        public static string? BinaryAuthorizationCustomEndpoint { get; set; } = __config.Get("binaryAuthorizationCustomEndpoint");

        public static string? CloudAssetCustomEndpoint { get; set; } = __config.Get("cloudAssetCustomEndpoint");

        public static string? CloudBillingCustomEndpoint { get; set; } = __config.Get("cloudBillingCustomEndpoint");

        public static string? CloudBuildCustomEndpoint { get; set; } = __config.Get("cloudBuildCustomEndpoint");

        public static string? CloudFunctionsCustomEndpoint { get; set; } = __config.Get("cloudFunctionsCustomEndpoint");

        public static string? CloudIdentityCustomEndpoint { get; set; } = __config.Get("cloudIdentityCustomEndpoint");

        public static string? CloudIotCustomEndpoint { get; set; } = __config.Get("cloudIotCustomEndpoint");

        public static string? CloudRunCustomEndpoint { get; set; } = __config.Get("cloudRunCustomEndpoint");

        public static string? CloudSchedulerCustomEndpoint { get; set; } = __config.Get("cloudSchedulerCustomEndpoint");

        public static string? CloudTasksCustomEndpoint { get; set; } = __config.Get("cloudTasksCustomEndpoint");

        public static string? ComposerCustomEndpoint { get; set; } = __config.Get("composerCustomEndpoint");

        public static string? ComputeBetaCustomEndpoint { get; set; } = __config.Get("computeBetaCustomEndpoint");

        public static string? ComputeCustomEndpoint { get; set; } = __config.Get("computeCustomEndpoint");

        public static string? ContainerAnalysisCustomEndpoint { get; set; } = __config.Get("containerAnalysisCustomEndpoint");

        public static string? ContainerBetaCustomEndpoint { get; set; } = __config.Get("containerBetaCustomEndpoint");

        public static string? ContainerCustomEndpoint { get; set; } = __config.Get("containerCustomEndpoint");

        public static string? Credentials { get; set; } = __config.Get("credentials");

        public static string? DataCatalogCustomEndpoint { get; set; } = __config.Get("dataCatalogCustomEndpoint");

        public static string? DataFusionCustomEndpoint { get; set; } = __config.Get("dataFusionCustomEndpoint");

        public static string? DataLossPreventionCustomEndpoint { get; set; } = __config.Get("dataLossPreventionCustomEndpoint");

        public static string? DataflowCustomEndpoint { get; set; } = __config.Get("dataflowCustomEndpoint");

        public static string? DataprocBetaCustomEndpoint { get; set; } = __config.Get("dataprocBetaCustomEndpoint");

        public static string? DataprocCustomEndpoint { get; set; } = __config.Get("dataprocCustomEndpoint");

        public static string? DataprocMetastoreCustomEndpoint { get; set; } = __config.Get("dataprocMetastoreCustomEndpoint");

        public static string? DatastoreCustomEndpoint { get; set; } = __config.Get("datastoreCustomEndpoint");

        public static string? DeploymentManagerCustomEndpoint { get; set; } = __config.Get("deploymentManagerCustomEndpoint");

        public static string? DialogflowCustomEndpoint { get; set; } = __config.Get("dialogflowCustomEndpoint");

        public static string? DialogflowCxCustomEndpoint { get; set; } = __config.Get("dialogflowCxCustomEndpoint");

        public static string? DnsCustomEndpoint { get; set; } = __config.Get("dnsCustomEndpoint");

        public static string? EssentialContactsCustomEndpoint { get; set; } = __config.Get("essentialContactsCustomEndpoint");

        public static string? EventarcCustomEndpoint { get; set; } = __config.Get("eventarcCustomEndpoint");

        public static string? FilestoreCustomEndpoint { get; set; } = __config.Get("filestoreCustomEndpoint");

        public static string? FirebaseCustomEndpoint { get; set; } = __config.Get("firebaseCustomEndpoint");

        public static string? FirestoreCustomEndpoint { get; set; } = __config.Get("firestoreCustomEndpoint");

        public static string? GameServicesCustomEndpoint { get; set; } = __config.Get("gameServicesCustomEndpoint");

        public static string? GkeHubCustomEndpoint { get; set; } = __config.Get("gkeHubCustomEndpoint");

        public static string? GkehubFeatureCustomEndpoint { get; set; } = __config.Get("gkehubFeatureCustomEndpoint");

        public static string? HealthcareCustomEndpoint { get; set; } = __config.Get("healthcareCustomEndpoint");

        public static string? IamBetaCustomEndpoint { get; set; } = __config.Get("iamBetaCustomEndpoint");

        public static string? IamCredentialsCustomEndpoint { get; set; } = __config.Get("iamCredentialsCustomEndpoint");

        public static string? IamCustomEndpoint { get; set; } = __config.Get("iamCustomEndpoint");

        public static string? IapCustomEndpoint { get; set; } = __config.Get("iapCustomEndpoint");

        public static string? IdentityPlatformCustomEndpoint { get; set; } = __config.Get("identityPlatformCustomEndpoint");

        public static string? ImpersonateServiceAccount { get; set; } = __config.Get("impersonateServiceAccount");

        public static ImmutableArray<string> ImpersonateServiceAccountDelegates { get; set; } = __config.GetObject<ImmutableArray<string>>("impersonateServiceAccountDelegates");

        public static string? KmsCustomEndpoint { get; set; } = __config.Get("kmsCustomEndpoint");

        public static string? LoggingCustomEndpoint { get; set; } = __config.Get("loggingCustomEndpoint");

        public static string? MemcacheCustomEndpoint { get; set; } = __config.Get("memcacheCustomEndpoint");

        public static string? MlEngineCustomEndpoint { get; set; } = __config.Get("mlEngineCustomEndpoint");

        public static string? MonitoringCustomEndpoint { get; set; } = __config.Get("monitoringCustomEndpoint");

        public static string? NetworkManagementCustomEndpoint { get; set; } = __config.Get("networkManagementCustomEndpoint");

        public static string? NotebooksCustomEndpoint { get; set; } = __config.Get("notebooksCustomEndpoint");

        public static string? OsConfigCustomEndpoint { get; set; } = __config.Get("osConfigCustomEndpoint");

        public static string? OsLoginCustomEndpoint { get; set; } = __config.Get("osLoginCustomEndpoint");

        public static string? PrivatecaCustomEndpoint { get; set; } = __config.Get("privatecaCustomEndpoint");

        public static string? Project { get; set; } = __config.Get("project") ?? Utilities.GetEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");

        public static string? PubsubCustomEndpoint { get; set; } = __config.Get("pubsubCustomEndpoint");

        public static string? PubsubLiteCustomEndpoint { get; set; } = __config.Get("pubsubLiteCustomEndpoint");

        public static string? RedisCustomEndpoint { get; set; } = __config.Get("redisCustomEndpoint");

        public static string? Region { get; set; } = __config.Get("region") ?? Utilities.GetEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");

        public static string? RequestTimeout { get; set; } = __config.Get("requestTimeout");

        public static string? ResourceManagerCustomEndpoint { get; set; } = __config.Get("resourceManagerCustomEndpoint");

        public static string? ResourceManagerV2CustomEndpoint { get; set; } = __config.Get("resourceManagerV2CustomEndpoint");

        public static string? RuntimeConfigCustomEndpoint { get; set; } = __config.Get("runtimeConfigCustomEndpoint");

        public static string? RuntimeconfigCustomEndpoint { get; set; } = __config.Get("runtimeconfigCustomEndpoint");

        public static ImmutableArray<string> Scopes { get; set; } = __config.GetObject<ImmutableArray<string>>("scopes");

        public static string? SecretManagerCustomEndpoint { get; set; } = __config.Get("secretManagerCustomEndpoint");

        public static string? SecurityCenterCustomEndpoint { get; set; } = __config.Get("securityCenterCustomEndpoint");

        public static string? SecurityScannerCustomEndpoint { get; set; } = __config.Get("securityScannerCustomEndpoint");

        public static string? ServiceDirectoryCustomEndpoint { get; set; } = __config.Get("serviceDirectoryCustomEndpoint");

        public static string? ServiceManagementCustomEndpoint { get; set; } = __config.Get("serviceManagementCustomEndpoint");

        public static string? ServiceNetworkingCustomEndpoint { get; set; } = __config.Get("serviceNetworkingCustomEndpoint");

        public static string? ServiceUsageCustomEndpoint { get; set; } = __config.Get("serviceUsageCustomEndpoint");

        public static string? SourceRepoCustomEndpoint { get; set; } = __config.Get("sourceRepoCustomEndpoint");

        public static string? SpannerCustomEndpoint { get; set; } = __config.Get("spannerCustomEndpoint");

        public static string? SqlCustomEndpoint { get; set; } = __config.Get("sqlCustomEndpoint");

        public static string? StorageCustomEndpoint { get; set; } = __config.Get("storageCustomEndpoint");

        public static string? StorageTransferCustomEndpoint { get; set; } = __config.Get("storageTransferCustomEndpoint");

        public static string? TagsCustomEndpoint { get; set; } = __config.Get("tagsCustomEndpoint");

        public static string? TpuCustomEndpoint { get; set; } = __config.Get("tpuCustomEndpoint");

        public static bool? UserProjectOverride { get; set; } = __config.GetBoolean("userProjectOverride");

        public static string? VpcAccessCustomEndpoint { get; set; } = __config.Get("vpcAccessCustomEndpoint");

        public static string? WorkflowsCustomEndpoint { get; set; } = __config.Get("workflowsCustomEndpoint");

        public static string? Zone { get; set; } = __config.Get("zone") ?? Utilities.GetEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");

        public static class Types
        {

             public class Batching
             {
                public bool? EnableBatching { get; set; }
                public string? SendAfter { get; set; } = null!;
            }
        }
    }
}
