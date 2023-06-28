// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp;

import com.pulumi.core.TypeShape;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.config.inputs.Batching;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("gcp");
    public Optional<String> accessApprovalCustomEndpoint() {
        return Codegen.stringProp("accessApprovalCustomEndpoint").config(config).get();
    }
    public Optional<String> accessContextManagerCustomEndpoint() {
        return Codegen.stringProp("accessContextManagerCustomEndpoint").config(config).get();
    }
    public Optional<String> accessToken() {
        return Codegen.stringProp("accessToken").config(config).get();
    }
    public Optional<String> activeDirectoryCustomEndpoint() {
        return Codegen.stringProp("activeDirectoryCustomEndpoint").config(config).get();
    }
    public Optional<String> alloydbCustomEndpoint() {
        return Codegen.stringProp("alloydbCustomEndpoint").config(config).get();
    }
    public Optional<String> apiGatewayCustomEndpoint() {
        return Codegen.stringProp("apiGatewayCustomEndpoint").config(config).get();
    }
    public Optional<String> apigeeCustomEndpoint() {
        return Codegen.stringProp("apigeeCustomEndpoint").config(config).get();
    }
    public Optional<String> apikeysCustomEndpoint() {
        return Codegen.stringProp("apikeysCustomEndpoint").config(config).get();
    }
    public Optional<String> appEngineCustomEndpoint() {
        return Codegen.stringProp("appEngineCustomEndpoint").config(config).get();
    }
    public Optional<String> artifactRegistryCustomEndpoint() {
        return Codegen.stringProp("artifactRegistryCustomEndpoint").config(config).get();
    }
    public Optional<String> assuredWorkloadsCustomEndpoint() {
        return Codegen.stringProp("assuredWorkloadsCustomEndpoint").config(config).get();
    }
    public Optional<Batching> batching() {
        return Codegen.objectProp("batching", Batching.class).config(config).get();
    }
    public Optional<String> beyondcorpCustomEndpoint() {
        return Codegen.stringProp("beyondcorpCustomEndpoint").config(config).get();
    }
    public Optional<String> bigQueryCustomEndpoint() {
        return Codegen.stringProp("bigQueryCustomEndpoint").config(config).get();
    }
    public Optional<String> bigqueryAnalyticsHubCustomEndpoint() {
        return Codegen.stringProp("bigqueryAnalyticsHubCustomEndpoint").config(config).get();
    }
    public Optional<String> bigqueryConnectionCustomEndpoint() {
        return Codegen.stringProp("bigqueryConnectionCustomEndpoint").config(config).get();
    }
    public Optional<String> bigqueryDataTransferCustomEndpoint() {
        return Codegen.stringProp("bigqueryDataTransferCustomEndpoint").config(config).get();
    }
    public Optional<String> bigqueryDatapolicyCustomEndpoint() {
        return Codegen.stringProp("bigqueryDatapolicyCustomEndpoint").config(config).get();
    }
    public Optional<String> bigqueryReservationCustomEndpoint() {
        return Codegen.stringProp("bigqueryReservationCustomEndpoint").config(config).get();
    }
    public Optional<String> bigtableCustomEndpoint() {
        return Codegen.stringProp("bigtableCustomEndpoint").config(config).get();
    }
    public Optional<String> billingCustomEndpoint() {
        return Codegen.stringProp("billingCustomEndpoint").config(config).get();
    }
    public Optional<String> billingProject() {
        return Codegen.stringProp("billingProject").config(config).get();
    }
    public Optional<String> binaryAuthorizationCustomEndpoint() {
        return Codegen.stringProp("binaryAuthorizationCustomEndpoint").config(config).get();
    }
    public Optional<String> certificateManagerCustomEndpoint() {
        return Codegen.stringProp("certificateManagerCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudAssetCustomEndpoint() {
        return Codegen.stringProp("cloudAssetCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudBillingCustomEndpoint() {
        return Codegen.stringProp("cloudBillingCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudBuildCustomEndpoint() {
        return Codegen.stringProp("cloudBuildCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudBuildWorkerPoolCustomEndpoint() {
        return Codegen.stringProp("cloudBuildWorkerPoolCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudFunctionsCustomEndpoint() {
        return Codegen.stringProp("cloudFunctionsCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudIdentityCustomEndpoint() {
        return Codegen.stringProp("cloudIdentityCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudIdsCustomEndpoint() {
        return Codegen.stringProp("cloudIdsCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudIotCustomEndpoint() {
        return Codegen.stringProp("cloudIotCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudResourceManagerCustomEndpoint() {
        return Codegen.stringProp("cloudResourceManagerCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudRunCustomEndpoint() {
        return Codegen.stringProp("cloudRunCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudRunV2CustomEndpoint() {
        return Codegen.stringProp("cloudRunV2CustomEndpoint").config(config).get();
    }
    public Optional<String> cloudSchedulerCustomEndpoint() {
        return Codegen.stringProp("cloudSchedulerCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudTasksCustomEndpoint() {
        return Codegen.stringProp("cloudTasksCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudbuildv2CustomEndpoint() {
        return Codegen.stringProp("cloudbuildv2CustomEndpoint").config(config).get();
    }
    public Optional<String> clouddeployCustomEndpoint() {
        return Codegen.stringProp("clouddeployCustomEndpoint").config(config).get();
    }
    public Optional<String> cloudfunctions2CustomEndpoint() {
        return Codegen.stringProp("cloudfunctions2CustomEndpoint").config(config).get();
    }
    public Optional<String> composerCustomEndpoint() {
        return Codegen.stringProp("composerCustomEndpoint").config(config).get();
    }
    public Optional<String> computeCustomEndpoint() {
        return Codegen.stringProp("computeCustomEndpoint").config(config).get();
    }
    public Optional<String> containerAnalysisCustomEndpoint() {
        return Codegen.stringProp("containerAnalysisCustomEndpoint").config(config).get();
    }
    public Optional<String> containerAttachedCustomEndpoint() {
        return Codegen.stringProp("containerAttachedCustomEndpoint").config(config).get();
    }
    public Optional<String> containerAwsCustomEndpoint() {
        return Codegen.stringProp("containerAwsCustomEndpoint").config(config).get();
    }
    public Optional<String> containerAzureCustomEndpoint() {
        return Codegen.stringProp("containerAzureCustomEndpoint").config(config).get();
    }
    public Optional<String> containerCustomEndpoint() {
        return Codegen.stringProp("containerCustomEndpoint").config(config).get();
    }
    public Optional<String> credentials() {
        return Codegen.stringProp("credentials").config(config).get();
    }
    public Optional<String> dataCatalogCustomEndpoint() {
        return Codegen.stringProp("dataCatalogCustomEndpoint").config(config).get();
    }
    public Optional<String> dataFusionCustomEndpoint() {
        return Codegen.stringProp("dataFusionCustomEndpoint").config(config).get();
    }
    public Optional<String> dataLossPreventionCustomEndpoint() {
        return Codegen.stringProp("dataLossPreventionCustomEndpoint").config(config).get();
    }
    public Optional<String> databaseMigrationServiceCustomEndpoint() {
        return Codegen.stringProp("databaseMigrationServiceCustomEndpoint").config(config).get();
    }
    public Optional<String> dataflowCustomEndpoint() {
        return Codegen.stringProp("dataflowCustomEndpoint").config(config).get();
    }
    public Optional<String> dataformCustomEndpoint() {
        return Codegen.stringProp("dataformCustomEndpoint").config(config).get();
    }
    public Optional<String> dataplexCustomEndpoint() {
        return Codegen.stringProp("dataplexCustomEndpoint").config(config).get();
    }
    public Optional<String> dataprocCustomEndpoint() {
        return Codegen.stringProp("dataprocCustomEndpoint").config(config).get();
    }
    public Optional<String> dataprocMetastoreCustomEndpoint() {
        return Codegen.stringProp("dataprocMetastoreCustomEndpoint").config(config).get();
    }
    public Optional<String> datastoreCustomEndpoint() {
        return Codegen.stringProp("datastoreCustomEndpoint").config(config).get();
    }
    public Optional<String> datastreamCustomEndpoint() {
        return Codegen.stringProp("datastreamCustomEndpoint").config(config).get();
    }
    public Optional<String> deploymentManagerCustomEndpoint() {
        return Codegen.stringProp("deploymentManagerCustomEndpoint").config(config).get();
    }
    public Optional<String> dialogflowCustomEndpoint() {
        return Codegen.stringProp("dialogflowCustomEndpoint").config(config).get();
    }
    public Optional<String> dialogflowCxCustomEndpoint() {
        return Codegen.stringProp("dialogflowCxCustomEndpoint").config(config).get();
    }
    public Optional<Boolean> disableGooglePartnerName() {
        return Codegen.booleanProp("disableGooglePartnerName").config(config).get();
    }
    public Optional<String> dnsCustomEndpoint() {
        return Codegen.stringProp("dnsCustomEndpoint").config(config).get();
    }
    public Optional<String> documentAiCustomEndpoint() {
        return Codegen.stringProp("documentAiCustomEndpoint").config(config).get();
    }
    public Optional<String> essentialContactsCustomEndpoint() {
        return Codegen.stringProp("essentialContactsCustomEndpoint").config(config).get();
    }
    public Optional<String> eventarcCustomEndpoint() {
        return Codegen.stringProp("eventarcCustomEndpoint").config(config).get();
    }
    public Optional<String> filestoreCustomEndpoint() {
        return Codegen.stringProp("filestoreCustomEndpoint").config(config).get();
    }
    public Optional<String> firebaseCustomEndpoint() {
        return Codegen.stringProp("firebaseCustomEndpoint").config(config).get();
    }
    public Optional<String> firebaseDatabaseCustomEndpoint() {
        return Codegen.stringProp("firebaseDatabaseCustomEndpoint").config(config).get();
    }
    public Optional<String> firebaseHostingCustomEndpoint() {
        return Codegen.stringProp("firebaseHostingCustomEndpoint").config(config).get();
    }
    public Optional<String> firebaseStorageCustomEndpoint() {
        return Codegen.stringProp("firebaseStorageCustomEndpoint").config(config).get();
    }
    public Optional<String> firebaserulesCustomEndpoint() {
        return Codegen.stringProp("firebaserulesCustomEndpoint").config(config).get();
    }
    public Optional<String> firestoreCustomEndpoint() {
        return Codegen.stringProp("firestoreCustomEndpoint").config(config).get();
    }
    public Optional<String> gameServicesCustomEndpoint() {
        return Codegen.stringProp("gameServicesCustomEndpoint").config(config).get();
    }
    public Optional<String> gkeBackupCustomEndpoint() {
        return Codegen.stringProp("gkeBackupCustomEndpoint").config(config).get();
    }
    public Optional<String> gkeHub2CustomEndpoint() {
        return Codegen.stringProp("gkeHub2CustomEndpoint").config(config).get();
    }
    public Optional<String> gkeHubCustomEndpoint() {
        return Codegen.stringProp("gkeHubCustomEndpoint").config(config).get();
    }
    public Optional<String> gkehubFeatureCustomEndpoint() {
        return Codegen.stringProp("gkehubFeatureCustomEndpoint").config(config).get();
    }
    public Optional<String> gkeonpremCustomEndpoint() {
        return Codegen.stringProp("gkeonpremCustomEndpoint").config(config).get();
    }
    public Optional<String> googlePartnerName() {
        return Codegen.stringProp("googlePartnerName").config(config).get();
    }
    public Optional<String> healthcareCustomEndpoint() {
        return Codegen.stringProp("healthcareCustomEndpoint").config(config).get();
    }
    public Optional<String> iam2CustomEndpoint() {
        return Codegen.stringProp("iam2CustomEndpoint").config(config).get();
    }
    public Optional<String> iamBetaCustomEndpoint() {
        return Codegen.stringProp("iamBetaCustomEndpoint").config(config).get();
    }
    public Optional<String> iamCredentialsCustomEndpoint() {
        return Codegen.stringProp("iamCredentialsCustomEndpoint").config(config).get();
    }
    public Optional<String> iamCustomEndpoint() {
        return Codegen.stringProp("iamCustomEndpoint").config(config).get();
    }
    public Optional<String> iamWorkforcePoolCustomEndpoint() {
        return Codegen.stringProp("iamWorkforcePoolCustomEndpoint").config(config).get();
    }
    public Optional<String> iapCustomEndpoint() {
        return Codegen.stringProp("iapCustomEndpoint").config(config).get();
    }
    public Optional<String> identityPlatformCustomEndpoint() {
        return Codegen.stringProp("identityPlatformCustomEndpoint").config(config).get();
    }
    public Optional<String> impersonateServiceAccount() {
        return Codegen.stringProp("impersonateServiceAccount").config(config).get();
    }
    public Optional<List<String>> impersonateServiceAccountDelegates() {
        return Codegen.objectProp("impersonateServiceAccountDelegates", TypeShape.<List<String>>builder(List.class).addParameter(String.class).build()).config(config).get();
    }
    public Optional<String> kmsCustomEndpoint() {
        return Codegen.stringProp("kmsCustomEndpoint").config(config).get();
    }
    public Optional<String> loggingCustomEndpoint() {
        return Codegen.stringProp("loggingCustomEndpoint").config(config).get();
    }
    public Optional<String> memcacheCustomEndpoint() {
        return Codegen.stringProp("memcacheCustomEndpoint").config(config).get();
    }
    public Optional<String> mlEngineCustomEndpoint() {
        return Codegen.stringProp("mlEngineCustomEndpoint").config(config).get();
    }
    public Optional<String> monitoringCustomEndpoint() {
        return Codegen.stringProp("monitoringCustomEndpoint").config(config).get();
    }
    public Optional<String> networkConnectivityCustomEndpoint() {
        return Codegen.stringProp("networkConnectivityCustomEndpoint").config(config).get();
    }
    public Optional<String> networkManagementCustomEndpoint() {
        return Codegen.stringProp("networkManagementCustomEndpoint").config(config).get();
    }
    public Optional<String> networkSecurityCustomEndpoint() {
        return Codegen.stringProp("networkSecurityCustomEndpoint").config(config).get();
    }
    public Optional<String> networkServicesCustomEndpoint() {
        return Codegen.stringProp("networkServicesCustomEndpoint").config(config).get();
    }
    public Optional<String> notebooksCustomEndpoint() {
        return Codegen.stringProp("notebooksCustomEndpoint").config(config).get();
    }
    public Optional<String> orgPolicyCustomEndpoint() {
        return Codegen.stringProp("orgPolicyCustomEndpoint").config(config).get();
    }
    public Optional<String> osConfigCustomEndpoint() {
        return Codegen.stringProp("osConfigCustomEndpoint").config(config).get();
    }
    public Optional<String> osLoginCustomEndpoint() {
        return Codegen.stringProp("osLoginCustomEndpoint").config(config).get();
    }
    public Optional<String> privatecaCustomEndpoint() {
        return Codegen.stringProp("privatecaCustomEndpoint").config(config).get();
    }
    public Optional<String> project() {
        return Codegen.stringProp("project").config(config).env("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT").get();
    }
    public Optional<String> pubsubCustomEndpoint() {
        return Codegen.stringProp("pubsubCustomEndpoint").config(config).get();
    }
    public Optional<String> pubsubLiteCustomEndpoint() {
        return Codegen.stringProp("pubsubLiteCustomEndpoint").config(config).get();
    }
    public Optional<String> recaptchaEnterpriseCustomEndpoint() {
        return Codegen.stringProp("recaptchaEnterpriseCustomEndpoint").config(config).get();
    }
    public Optional<String> redisCustomEndpoint() {
        return Codegen.stringProp("redisCustomEndpoint").config(config).get();
    }
    public Optional<String> region() {
        return Codegen.stringProp("region").config(config).env("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION").get();
    }
    public Optional<String> requestReason() {
        return Codegen.stringProp("requestReason").config(config).get();
    }
    public Optional<String> requestTimeout() {
        return Codegen.stringProp("requestTimeout").config(config).get();
    }
    public Optional<String> resourceManagerCustomEndpoint() {
        return Codegen.stringProp("resourceManagerCustomEndpoint").config(config).get();
    }
    public Optional<String> resourceManagerV3CustomEndpoint() {
        return Codegen.stringProp("resourceManagerV3CustomEndpoint").config(config).get();
    }
    public Optional<String> runtimeConfigCustomEndpoint() {
        return Codegen.stringProp("runtimeConfigCustomEndpoint").config(config).get();
    }
    public Optional<String> runtimeconfigCustomEndpoint() {
        return Codegen.stringProp("runtimeconfigCustomEndpoint").config(config).get();
    }
    public Optional<List<String>> scopes() {
        return Codegen.objectProp("scopes", TypeShape.<List<String>>builder(List.class).addParameter(String.class).build()).config(config).get();
    }
    public Optional<String> secretManagerCustomEndpoint() {
        return Codegen.stringProp("secretManagerCustomEndpoint").config(config).get();
    }
    public Optional<String> securityCenterCustomEndpoint() {
        return Codegen.stringProp("securityCenterCustomEndpoint").config(config).get();
    }
    public Optional<String> securityScannerCustomEndpoint() {
        return Codegen.stringProp("securityScannerCustomEndpoint").config(config).get();
    }
    public Optional<String> serviceDirectoryCustomEndpoint() {
        return Codegen.stringProp("serviceDirectoryCustomEndpoint").config(config).get();
    }
    public Optional<String> serviceManagementCustomEndpoint() {
        return Codegen.stringProp("serviceManagementCustomEndpoint").config(config).get();
    }
    public Optional<String> serviceNetworkingCustomEndpoint() {
        return Codegen.stringProp("serviceNetworkingCustomEndpoint").config(config).get();
    }
    public Optional<String> serviceUsageCustomEndpoint() {
        return Codegen.stringProp("serviceUsageCustomEndpoint").config(config).get();
    }
    public Optional<String> sourceRepoCustomEndpoint() {
        return Codegen.stringProp("sourceRepoCustomEndpoint").config(config).get();
    }
    public Optional<String> spannerCustomEndpoint() {
        return Codegen.stringProp("spannerCustomEndpoint").config(config).get();
    }
    public Optional<String> sqlCustomEndpoint() {
        return Codegen.stringProp("sqlCustomEndpoint").config(config).get();
    }
    public Optional<String> storageCustomEndpoint() {
        return Codegen.stringProp("storageCustomEndpoint").config(config).get();
    }
    public Optional<String> storageTransferCustomEndpoint() {
        return Codegen.stringProp("storageTransferCustomEndpoint").config(config).get();
    }
    public Optional<String> tagsCustomEndpoint() {
        return Codegen.stringProp("tagsCustomEndpoint").config(config).get();
    }
    public Optional<String> tagsLocationCustomEndpoint() {
        return Codegen.stringProp("tagsLocationCustomEndpoint").config(config).get();
    }
    public Optional<String> tpuCustomEndpoint() {
        return Codegen.stringProp("tpuCustomEndpoint").config(config).get();
    }
    public Optional<Boolean> userProjectOverride() {
        return Codegen.booleanProp("userProjectOverride").config(config).get();
    }
    public Optional<String> vertexAiCustomEndpoint() {
        return Codegen.stringProp("vertexAiCustomEndpoint").config(config).get();
    }
    public Optional<String> vmwareengineCustomEndpoint() {
        return Codegen.stringProp("vmwareengineCustomEndpoint").config(config).get();
    }
    public Optional<String> vpcAccessCustomEndpoint() {
        return Codegen.stringProp("vpcAccessCustomEndpoint").config(config).get();
    }
    public Optional<String> workflowsCustomEndpoint() {
        return Codegen.stringProp("workflowsCustomEndpoint").config(config).get();
    }
    public Optional<String> workstationsCustomEndpoint() {
        return Codegen.stringProp("workstationsCustomEndpoint").config(config).get();
    }
    public Optional<String> zone() {
        return Codegen.stringProp("zone").config(config).env("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE").get();
    }
}
