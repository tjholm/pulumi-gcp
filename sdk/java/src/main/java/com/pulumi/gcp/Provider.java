// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.ProviderArgs;
import com.pulumi.gcp.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the google-beta package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:gcp")
public class Provider extends com.pulumi.resources.ProviderResource {
    @Export(name="accessApprovalCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessApprovalCustomEndpoint;

    public Output<Optional<String>> accessApprovalCustomEndpoint() {
        return Codegen.optional(this.accessApprovalCustomEndpoint);
    }
    @Export(name="accessContextManagerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessContextManagerCustomEndpoint;

    public Output<Optional<String>> accessContextManagerCustomEndpoint() {
        return Codegen.optional(this.accessContextManagerCustomEndpoint);
    }
    @Export(name="accessToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessToken;

    public Output<Optional<String>> accessToken() {
        return Codegen.optional(this.accessToken);
    }
    @Export(name="activeDirectoryCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> activeDirectoryCustomEndpoint;

    public Output<Optional<String>> activeDirectoryCustomEndpoint() {
        return Codegen.optional(this.activeDirectoryCustomEndpoint);
    }
    @Export(name="apiGatewayCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> apiGatewayCustomEndpoint;

    public Output<Optional<String>> apiGatewayCustomEndpoint() {
        return Codegen.optional(this.apiGatewayCustomEndpoint);
    }
    @Export(name="apigeeCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> apigeeCustomEndpoint;

    public Output<Optional<String>> apigeeCustomEndpoint() {
        return Codegen.optional(this.apigeeCustomEndpoint);
    }
    @Export(name="apikeysCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> apikeysCustomEndpoint;

    public Output<Optional<String>> apikeysCustomEndpoint() {
        return Codegen.optional(this.apikeysCustomEndpoint);
    }
    @Export(name="appEngineCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> appEngineCustomEndpoint;

    public Output<Optional<String>> appEngineCustomEndpoint() {
        return Codegen.optional(this.appEngineCustomEndpoint);
    }
    @Export(name="artifactRegistryCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> artifactRegistryCustomEndpoint;

    public Output<Optional<String>> artifactRegistryCustomEndpoint() {
        return Codegen.optional(this.artifactRegistryCustomEndpoint);
    }
    @Export(name="assuredWorkloadsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> assuredWorkloadsCustomEndpoint;

    public Output<Optional<String>> assuredWorkloadsCustomEndpoint() {
        return Codegen.optional(this.assuredWorkloadsCustomEndpoint);
    }
    @Export(name="bigQueryCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> bigQueryCustomEndpoint;

    public Output<Optional<String>> bigQueryCustomEndpoint() {
        return Codegen.optional(this.bigQueryCustomEndpoint);
    }
    @Export(name="bigqueryConnectionCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> bigqueryConnectionCustomEndpoint;

    public Output<Optional<String>> bigqueryConnectionCustomEndpoint() {
        return Codegen.optional(this.bigqueryConnectionCustomEndpoint);
    }
    @Export(name="bigqueryDataTransferCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> bigqueryDataTransferCustomEndpoint;

    public Output<Optional<String>> bigqueryDataTransferCustomEndpoint() {
        return Codegen.optional(this.bigqueryDataTransferCustomEndpoint);
    }
    @Export(name="bigqueryReservationCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> bigqueryReservationCustomEndpoint;

    public Output<Optional<String>> bigqueryReservationCustomEndpoint() {
        return Codegen.optional(this.bigqueryReservationCustomEndpoint);
    }
    @Export(name="bigtableCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> bigtableCustomEndpoint;

    public Output<Optional<String>> bigtableCustomEndpoint() {
        return Codegen.optional(this.bigtableCustomEndpoint);
    }
    @Export(name="billingCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> billingCustomEndpoint;

    public Output<Optional<String>> billingCustomEndpoint() {
        return Codegen.optional(this.billingCustomEndpoint);
    }
    @Export(name="billingProject", type=String.class, parameters={})
    private Output</* @Nullable */ String> billingProject;

    public Output<Optional<String>> billingProject() {
        return Codegen.optional(this.billingProject);
    }
    @Export(name="binaryAuthorizationCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> binaryAuthorizationCustomEndpoint;

    public Output<Optional<String>> binaryAuthorizationCustomEndpoint() {
        return Codegen.optional(this.binaryAuthorizationCustomEndpoint);
    }
    @Export(name="certificateManagerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> certificateManagerCustomEndpoint;

    public Output<Optional<String>> certificateManagerCustomEndpoint() {
        return Codegen.optional(this.certificateManagerCustomEndpoint);
    }
    @Export(name="cloudAssetCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudAssetCustomEndpoint;

    public Output<Optional<String>> cloudAssetCustomEndpoint() {
        return Codegen.optional(this.cloudAssetCustomEndpoint);
    }
    @Export(name="cloudBillingCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudBillingCustomEndpoint;

    public Output<Optional<String>> cloudBillingCustomEndpoint() {
        return Codegen.optional(this.cloudBillingCustomEndpoint);
    }
    @Export(name="cloudBuildCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudBuildCustomEndpoint;

    public Output<Optional<String>> cloudBuildCustomEndpoint() {
        return Codegen.optional(this.cloudBuildCustomEndpoint);
    }
    @Export(name="cloudBuildWorkerPoolCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudBuildWorkerPoolCustomEndpoint;

    public Output<Optional<String>> cloudBuildWorkerPoolCustomEndpoint() {
        return Codegen.optional(this.cloudBuildWorkerPoolCustomEndpoint);
    }
    @Export(name="cloudFunctionsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudFunctionsCustomEndpoint;

    public Output<Optional<String>> cloudFunctionsCustomEndpoint() {
        return Codegen.optional(this.cloudFunctionsCustomEndpoint);
    }
    @Export(name="cloudIdentityCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudIdentityCustomEndpoint;

    public Output<Optional<String>> cloudIdentityCustomEndpoint() {
        return Codegen.optional(this.cloudIdentityCustomEndpoint);
    }
    @Export(name="cloudIotCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudIotCustomEndpoint;

    public Output<Optional<String>> cloudIotCustomEndpoint() {
        return Codegen.optional(this.cloudIotCustomEndpoint);
    }
    @Export(name="cloudResourceManagerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudResourceManagerCustomEndpoint;

    public Output<Optional<String>> cloudResourceManagerCustomEndpoint() {
        return Codegen.optional(this.cloudResourceManagerCustomEndpoint);
    }
    @Export(name="cloudRunCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudRunCustomEndpoint;

    public Output<Optional<String>> cloudRunCustomEndpoint() {
        return Codegen.optional(this.cloudRunCustomEndpoint);
    }
    @Export(name="cloudSchedulerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudSchedulerCustomEndpoint;

    public Output<Optional<String>> cloudSchedulerCustomEndpoint() {
        return Codegen.optional(this.cloudSchedulerCustomEndpoint);
    }
    @Export(name="cloudTasksCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudTasksCustomEndpoint;

    public Output<Optional<String>> cloudTasksCustomEndpoint() {
        return Codegen.optional(this.cloudTasksCustomEndpoint);
    }
    @Export(name="clouddeployCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> clouddeployCustomEndpoint;

    public Output<Optional<String>> clouddeployCustomEndpoint() {
        return Codegen.optional(this.clouddeployCustomEndpoint);
    }
    @Export(name="cloudfunctions2CustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloudfunctions2CustomEndpoint;

    public Output<Optional<String>> cloudfunctions2CustomEndpoint() {
        return Codegen.optional(this.cloudfunctions2CustomEndpoint);
    }
    @Export(name="composerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> composerCustomEndpoint;

    public Output<Optional<String>> composerCustomEndpoint() {
        return Codegen.optional(this.composerCustomEndpoint);
    }
    @Export(name="computeCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> computeCustomEndpoint;

    public Output<Optional<String>> computeCustomEndpoint() {
        return Codegen.optional(this.computeCustomEndpoint);
    }
    @Export(name="containerAnalysisCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> containerAnalysisCustomEndpoint;

    public Output<Optional<String>> containerAnalysisCustomEndpoint() {
        return Codegen.optional(this.containerAnalysisCustomEndpoint);
    }
    @Export(name="containerAwsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> containerAwsCustomEndpoint;

    public Output<Optional<String>> containerAwsCustomEndpoint() {
        return Codegen.optional(this.containerAwsCustomEndpoint);
    }
    @Export(name="containerAzureCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> containerAzureCustomEndpoint;

    public Output<Optional<String>> containerAzureCustomEndpoint() {
        return Codegen.optional(this.containerAzureCustomEndpoint);
    }
    @Export(name="containerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> containerCustomEndpoint;

    public Output<Optional<String>> containerCustomEndpoint() {
        return Codegen.optional(this.containerCustomEndpoint);
    }
    @Export(name="credentials", type=String.class, parameters={})
    private Output</* @Nullable */ String> credentials;

    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    @Export(name="dataCatalogCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataCatalogCustomEndpoint;

    public Output<Optional<String>> dataCatalogCustomEndpoint() {
        return Codegen.optional(this.dataCatalogCustomEndpoint);
    }
    @Export(name="dataFusionCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataFusionCustomEndpoint;

    public Output<Optional<String>> dataFusionCustomEndpoint() {
        return Codegen.optional(this.dataFusionCustomEndpoint);
    }
    @Export(name="dataLossPreventionCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataLossPreventionCustomEndpoint;

    public Output<Optional<String>> dataLossPreventionCustomEndpoint() {
        return Codegen.optional(this.dataLossPreventionCustomEndpoint);
    }
    @Export(name="dataflowCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataflowCustomEndpoint;

    public Output<Optional<String>> dataflowCustomEndpoint() {
        return Codegen.optional(this.dataflowCustomEndpoint);
    }
    @Export(name="dataplexCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataplexCustomEndpoint;

    public Output<Optional<String>> dataplexCustomEndpoint() {
        return Codegen.optional(this.dataplexCustomEndpoint);
    }
    @Export(name="dataprocCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataprocCustomEndpoint;

    public Output<Optional<String>> dataprocCustomEndpoint() {
        return Codegen.optional(this.dataprocCustomEndpoint);
    }
    @Export(name="dataprocMetastoreCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dataprocMetastoreCustomEndpoint;

    public Output<Optional<String>> dataprocMetastoreCustomEndpoint() {
        return Codegen.optional(this.dataprocMetastoreCustomEndpoint);
    }
    @Export(name="datastoreCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> datastoreCustomEndpoint;

    public Output<Optional<String>> datastoreCustomEndpoint() {
        return Codegen.optional(this.datastoreCustomEndpoint);
    }
    @Export(name="deploymentManagerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> deploymentManagerCustomEndpoint;

    public Output<Optional<String>> deploymentManagerCustomEndpoint() {
        return Codegen.optional(this.deploymentManagerCustomEndpoint);
    }
    @Export(name="dialogflowCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dialogflowCustomEndpoint;

    public Output<Optional<String>> dialogflowCustomEndpoint() {
        return Codegen.optional(this.dialogflowCustomEndpoint);
    }
    @Export(name="dialogflowCxCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dialogflowCxCustomEndpoint;

    public Output<Optional<String>> dialogflowCxCustomEndpoint() {
        return Codegen.optional(this.dialogflowCxCustomEndpoint);
    }
    @Export(name="dnsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> dnsCustomEndpoint;

    public Output<Optional<String>> dnsCustomEndpoint() {
        return Codegen.optional(this.dnsCustomEndpoint);
    }
    @Export(name="documentAiCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> documentAiCustomEndpoint;

    public Output<Optional<String>> documentAiCustomEndpoint() {
        return Codegen.optional(this.documentAiCustomEndpoint);
    }
    @Export(name="essentialContactsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> essentialContactsCustomEndpoint;

    public Output<Optional<String>> essentialContactsCustomEndpoint() {
        return Codegen.optional(this.essentialContactsCustomEndpoint);
    }
    @Export(name="eventarcCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> eventarcCustomEndpoint;

    public Output<Optional<String>> eventarcCustomEndpoint() {
        return Codegen.optional(this.eventarcCustomEndpoint);
    }
    @Export(name="filestoreCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> filestoreCustomEndpoint;

    public Output<Optional<String>> filestoreCustomEndpoint() {
        return Codegen.optional(this.filestoreCustomEndpoint);
    }
    @Export(name="firebaseCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> firebaseCustomEndpoint;

    public Output<Optional<String>> firebaseCustomEndpoint() {
        return Codegen.optional(this.firebaseCustomEndpoint);
    }
    @Export(name="firebaserulesCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> firebaserulesCustomEndpoint;

    public Output<Optional<String>> firebaserulesCustomEndpoint() {
        return Codegen.optional(this.firebaserulesCustomEndpoint);
    }
    @Export(name="firestoreCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> firestoreCustomEndpoint;

    public Output<Optional<String>> firestoreCustomEndpoint() {
        return Codegen.optional(this.firestoreCustomEndpoint);
    }
    @Export(name="gameServicesCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> gameServicesCustomEndpoint;

    public Output<Optional<String>> gameServicesCustomEndpoint() {
        return Codegen.optional(this.gameServicesCustomEndpoint);
    }
    @Export(name="gkeHubCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> gkeHubCustomEndpoint;

    public Output<Optional<String>> gkeHubCustomEndpoint() {
        return Codegen.optional(this.gkeHubCustomEndpoint);
    }
    @Export(name="gkehubFeatureCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> gkehubFeatureCustomEndpoint;

    public Output<Optional<String>> gkehubFeatureCustomEndpoint() {
        return Codegen.optional(this.gkehubFeatureCustomEndpoint);
    }
    @Export(name="googlePartnerName", type=String.class, parameters={})
    private Output</* @Nullable */ String> googlePartnerName;

    public Output<Optional<String>> googlePartnerName() {
        return Codegen.optional(this.googlePartnerName);
    }
    @Export(name="healthcareCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> healthcareCustomEndpoint;

    public Output<Optional<String>> healthcareCustomEndpoint() {
        return Codegen.optional(this.healthcareCustomEndpoint);
    }
    @Export(name="iam2CustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> iam2CustomEndpoint;

    public Output<Optional<String>> iam2CustomEndpoint() {
        return Codegen.optional(this.iam2CustomEndpoint);
    }
    @Export(name="iamBetaCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamBetaCustomEndpoint;

    public Output<Optional<String>> iamBetaCustomEndpoint() {
        return Codegen.optional(this.iamBetaCustomEndpoint);
    }
    @Export(name="iamCredentialsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamCredentialsCustomEndpoint;

    public Output<Optional<String>> iamCredentialsCustomEndpoint() {
        return Codegen.optional(this.iamCredentialsCustomEndpoint);
    }
    @Export(name="iamCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamCustomEndpoint;

    public Output<Optional<String>> iamCustomEndpoint() {
        return Codegen.optional(this.iamCustomEndpoint);
    }
    @Export(name="iapCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> iapCustomEndpoint;

    public Output<Optional<String>> iapCustomEndpoint() {
        return Codegen.optional(this.iapCustomEndpoint);
    }
    @Export(name="identityPlatformCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> identityPlatformCustomEndpoint;

    public Output<Optional<String>> identityPlatformCustomEndpoint() {
        return Codegen.optional(this.identityPlatformCustomEndpoint);
    }
    @Export(name="impersonateServiceAccount", type=String.class, parameters={})
    private Output</* @Nullable */ String> impersonateServiceAccount;

    public Output<Optional<String>> impersonateServiceAccount() {
        return Codegen.optional(this.impersonateServiceAccount);
    }
    @Export(name="kmsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsCustomEndpoint;

    public Output<Optional<String>> kmsCustomEndpoint() {
        return Codegen.optional(this.kmsCustomEndpoint);
    }
    @Export(name="loggingCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> loggingCustomEndpoint;

    public Output<Optional<String>> loggingCustomEndpoint() {
        return Codegen.optional(this.loggingCustomEndpoint);
    }
    @Export(name="memcacheCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> memcacheCustomEndpoint;

    public Output<Optional<String>> memcacheCustomEndpoint() {
        return Codegen.optional(this.memcacheCustomEndpoint);
    }
    @Export(name="mlEngineCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> mlEngineCustomEndpoint;

    public Output<Optional<String>> mlEngineCustomEndpoint() {
        return Codegen.optional(this.mlEngineCustomEndpoint);
    }
    @Export(name="monitoringCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> monitoringCustomEndpoint;

    public Output<Optional<String>> monitoringCustomEndpoint() {
        return Codegen.optional(this.monitoringCustomEndpoint);
    }
    @Export(name="networkConnectivityCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> networkConnectivityCustomEndpoint;

    public Output<Optional<String>> networkConnectivityCustomEndpoint() {
        return Codegen.optional(this.networkConnectivityCustomEndpoint);
    }
    @Export(name="networkManagementCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> networkManagementCustomEndpoint;

    public Output<Optional<String>> networkManagementCustomEndpoint() {
        return Codegen.optional(this.networkManagementCustomEndpoint);
    }
    @Export(name="networkServicesCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> networkServicesCustomEndpoint;

    public Output<Optional<String>> networkServicesCustomEndpoint() {
        return Codegen.optional(this.networkServicesCustomEndpoint);
    }
    @Export(name="notebooksCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> notebooksCustomEndpoint;

    public Output<Optional<String>> notebooksCustomEndpoint() {
        return Codegen.optional(this.notebooksCustomEndpoint);
    }
    @Export(name="orgPolicyCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> orgPolicyCustomEndpoint;

    public Output<Optional<String>> orgPolicyCustomEndpoint() {
        return Codegen.optional(this.orgPolicyCustomEndpoint);
    }
    @Export(name="osConfigCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> osConfigCustomEndpoint;

    public Output<Optional<String>> osConfigCustomEndpoint() {
        return Codegen.optional(this.osConfigCustomEndpoint);
    }
    @Export(name="osLoginCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> osLoginCustomEndpoint;

    public Output<Optional<String>> osLoginCustomEndpoint() {
        return Codegen.optional(this.osLoginCustomEndpoint);
    }
    @Export(name="privatecaCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> privatecaCustomEndpoint;

    public Output<Optional<String>> privatecaCustomEndpoint() {
        return Codegen.optional(this.privatecaCustomEndpoint);
    }
    @Export(name="project", type=String.class, parameters={})
    private Output</* @Nullable */ String> project;

    public Output<Optional<String>> project() {
        return Codegen.optional(this.project);
    }
    @Export(name="pubsubCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> pubsubCustomEndpoint;

    public Output<Optional<String>> pubsubCustomEndpoint() {
        return Codegen.optional(this.pubsubCustomEndpoint);
    }
    @Export(name="pubsubLiteCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> pubsubLiteCustomEndpoint;

    public Output<Optional<String>> pubsubLiteCustomEndpoint() {
        return Codegen.optional(this.pubsubLiteCustomEndpoint);
    }
    @Export(name="recaptchaEnterpriseCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> recaptchaEnterpriseCustomEndpoint;

    public Output<Optional<String>> recaptchaEnterpriseCustomEndpoint() {
        return Codegen.optional(this.recaptchaEnterpriseCustomEndpoint);
    }
    @Export(name="redisCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> redisCustomEndpoint;

    public Output<Optional<String>> redisCustomEndpoint() {
        return Codegen.optional(this.redisCustomEndpoint);
    }
    @Export(name="region", type=String.class, parameters={})
    private Output</* @Nullable */ String> region;

    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    @Export(name="requestReason", type=String.class, parameters={})
    private Output</* @Nullable */ String> requestReason;

    public Output<Optional<String>> requestReason() {
        return Codegen.optional(this.requestReason);
    }
    @Export(name="requestTimeout", type=String.class, parameters={})
    private Output</* @Nullable */ String> requestTimeout;

    public Output<Optional<String>> requestTimeout() {
        return Codegen.optional(this.requestTimeout);
    }
    @Export(name="resourceManagerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceManagerCustomEndpoint;

    public Output<Optional<String>> resourceManagerCustomEndpoint() {
        return Codegen.optional(this.resourceManagerCustomEndpoint);
    }
    @Export(name="resourceManagerV3CustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceManagerV3CustomEndpoint;

    public Output<Optional<String>> resourceManagerV3CustomEndpoint() {
        return Codegen.optional(this.resourceManagerV3CustomEndpoint);
    }
    @Export(name="runtimeConfigCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> runtimeConfigCustomEndpoint;

    public Output<Optional<String>> runtimeConfigCustomEndpoint() {
        return Codegen.optional(this.runtimeConfigCustomEndpoint);
    }
    @Export(name="runtimeconfigCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> runtimeconfigCustomEndpoint;

    public Output<Optional<String>> runtimeconfigCustomEndpoint() {
        return Codegen.optional(this.runtimeconfigCustomEndpoint);
    }
    @Export(name="secretManagerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> secretManagerCustomEndpoint;

    public Output<Optional<String>> secretManagerCustomEndpoint() {
        return Codegen.optional(this.secretManagerCustomEndpoint);
    }
    @Export(name="securityCenterCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityCenterCustomEndpoint;

    public Output<Optional<String>> securityCenterCustomEndpoint() {
        return Codegen.optional(this.securityCenterCustomEndpoint);
    }
    @Export(name="securityScannerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityScannerCustomEndpoint;

    public Output<Optional<String>> securityScannerCustomEndpoint() {
        return Codegen.optional(this.securityScannerCustomEndpoint);
    }
    @Export(name="serviceDirectoryCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceDirectoryCustomEndpoint;

    public Output<Optional<String>> serviceDirectoryCustomEndpoint() {
        return Codegen.optional(this.serviceDirectoryCustomEndpoint);
    }
    @Export(name="serviceManagementCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceManagementCustomEndpoint;

    public Output<Optional<String>> serviceManagementCustomEndpoint() {
        return Codegen.optional(this.serviceManagementCustomEndpoint);
    }
    @Export(name="serviceNetworkingCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceNetworkingCustomEndpoint;

    public Output<Optional<String>> serviceNetworkingCustomEndpoint() {
        return Codegen.optional(this.serviceNetworkingCustomEndpoint);
    }
    @Export(name="serviceUsageCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceUsageCustomEndpoint;

    public Output<Optional<String>> serviceUsageCustomEndpoint() {
        return Codegen.optional(this.serviceUsageCustomEndpoint);
    }
    @Export(name="sourceRepoCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceRepoCustomEndpoint;

    public Output<Optional<String>> sourceRepoCustomEndpoint() {
        return Codegen.optional(this.sourceRepoCustomEndpoint);
    }
    @Export(name="spannerCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> spannerCustomEndpoint;

    public Output<Optional<String>> spannerCustomEndpoint() {
        return Codegen.optional(this.spannerCustomEndpoint);
    }
    @Export(name="sqlCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> sqlCustomEndpoint;

    public Output<Optional<String>> sqlCustomEndpoint() {
        return Codegen.optional(this.sqlCustomEndpoint);
    }
    @Export(name="storageCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> storageCustomEndpoint;

    public Output<Optional<String>> storageCustomEndpoint() {
        return Codegen.optional(this.storageCustomEndpoint);
    }
    @Export(name="storageTransferCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> storageTransferCustomEndpoint;

    public Output<Optional<String>> storageTransferCustomEndpoint() {
        return Codegen.optional(this.storageTransferCustomEndpoint);
    }
    @Export(name="tagsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> tagsCustomEndpoint;

    public Output<Optional<String>> tagsCustomEndpoint() {
        return Codegen.optional(this.tagsCustomEndpoint);
    }
    @Export(name="tpuCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> tpuCustomEndpoint;

    public Output<Optional<String>> tpuCustomEndpoint() {
        return Codegen.optional(this.tpuCustomEndpoint);
    }
    @Export(name="vertexAiCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> vertexAiCustomEndpoint;

    public Output<Optional<String>> vertexAiCustomEndpoint() {
        return Codegen.optional(this.vertexAiCustomEndpoint);
    }
    @Export(name="vpcAccessCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> vpcAccessCustomEndpoint;

    public Output<Optional<String>> vpcAccessCustomEndpoint() {
        return Codegen.optional(this.vpcAccessCustomEndpoint);
    }
    @Export(name="workflowsCustomEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> workflowsCustomEndpoint;

    public Output<Optional<String>> workflowsCustomEndpoint() {
        return Codegen.optional(this.workflowsCustomEndpoint);
    }
    @Export(name="zone", type=String.class, parameters={})
    private Output</* @Nullable */ String> zone;

    public Output<Optional<String>> zone() {
        return Codegen.optional(this.zone);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
