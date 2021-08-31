// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The provider type for the google-beta package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'gcp';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    public readonly accessApprovalCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly accessContextManagerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly accessToken!: pulumi.Output<string | undefined>;
    public readonly activeDirectoryCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly apiGatewayCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly apigeeCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly appEngineCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly artifactRegistryCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly assuredWorkloadsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly bigQueryCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly bigqueryConnectionCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly bigqueryDataTransferCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly bigqueryReservationCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly bigtableCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly billingCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly billingProject!: pulumi.Output<string | undefined>;
    public readonly binaryAuthorizationCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudAssetCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudBillingCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudBuildCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudBuildWorkerPoolCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudFunctionsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudIdentityCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudIotCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudRunCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudSchedulerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly cloudTasksCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly composerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly computeBetaCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly computeCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly containerAnalysisCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly containerBetaCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly containerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly credentials!: pulumi.Output<string | undefined>;
    public readonly dataCatalogCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dataFusionCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dataLossPreventionCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dataflowCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dataprocBetaCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dataprocCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dataprocMetastoreCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly datastoreCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly deploymentManagerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dialogflowCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dialogflowCxCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly dnsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly essentialContactsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly eventarcCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly filestoreCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly firebaseCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly firestoreCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly gameServicesCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly gkeHubCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly gkehubFeatureCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly googlePartnerName!: pulumi.Output<string | undefined>;
    public readonly healthcareCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly iamBetaCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly iamCredentialsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly iamCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly iapCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly identityPlatformCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly impersonateServiceAccount!: pulumi.Output<string | undefined>;
    public readonly kmsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly loggingCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly memcacheCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly mlEngineCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly monitoringCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly networkManagementCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly networkServicesCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly notebooksCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly osConfigCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly osLoginCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly privatecaCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly project!: pulumi.Output<string | undefined>;
    public readonly pubsubCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly pubsubLiteCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly redisCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string | undefined>;
    public readonly requestReason!: pulumi.Output<string | undefined>;
    public readonly requestTimeout!: pulumi.Output<string | undefined>;
    public readonly resourceManagerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly resourceManagerV2CustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly runtimeConfigCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly runtimeconfigCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly secretManagerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly securityCenterCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly securityScannerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly serviceDirectoryCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly serviceManagementCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly serviceNetworkingCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly serviceUsageCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly sourceRepoCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly spannerCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly sqlCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly storageCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly storageTransferCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly tagsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly tpuCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly vertexAiCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly vpcAccessCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly workflowsCustomEndpoint!: pulumi.Output<string | undefined>;
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            inputs["accessApprovalCustomEndpoint"] = args ? args.accessApprovalCustomEndpoint : undefined;
            inputs["accessContextManagerCustomEndpoint"] = args ? args.accessContextManagerCustomEndpoint : undefined;
            inputs["accessToken"] = args ? args.accessToken : undefined;
            inputs["activeDirectoryCustomEndpoint"] = args ? args.activeDirectoryCustomEndpoint : undefined;
            inputs["apiGatewayCustomEndpoint"] = args ? args.apiGatewayCustomEndpoint : undefined;
            inputs["apigeeCustomEndpoint"] = args ? args.apigeeCustomEndpoint : undefined;
            inputs["appEngineCustomEndpoint"] = args ? args.appEngineCustomEndpoint : undefined;
            inputs["artifactRegistryCustomEndpoint"] = args ? args.artifactRegistryCustomEndpoint : undefined;
            inputs["assuredWorkloadsCustomEndpoint"] = args ? args.assuredWorkloadsCustomEndpoint : undefined;
            inputs["batching"] = pulumi.output(args ? args.batching : undefined).apply(JSON.stringify);
            inputs["bigQueryCustomEndpoint"] = args ? args.bigQueryCustomEndpoint : undefined;
            inputs["bigqueryConnectionCustomEndpoint"] = args ? args.bigqueryConnectionCustomEndpoint : undefined;
            inputs["bigqueryDataTransferCustomEndpoint"] = args ? args.bigqueryDataTransferCustomEndpoint : undefined;
            inputs["bigqueryReservationCustomEndpoint"] = args ? args.bigqueryReservationCustomEndpoint : undefined;
            inputs["bigtableCustomEndpoint"] = args ? args.bigtableCustomEndpoint : undefined;
            inputs["billingCustomEndpoint"] = args ? args.billingCustomEndpoint : undefined;
            inputs["billingProject"] = args ? args.billingProject : undefined;
            inputs["binaryAuthorizationCustomEndpoint"] = args ? args.binaryAuthorizationCustomEndpoint : undefined;
            inputs["cloudAssetCustomEndpoint"] = args ? args.cloudAssetCustomEndpoint : undefined;
            inputs["cloudBillingCustomEndpoint"] = args ? args.cloudBillingCustomEndpoint : undefined;
            inputs["cloudBuildCustomEndpoint"] = args ? args.cloudBuildCustomEndpoint : undefined;
            inputs["cloudBuildWorkerPoolCustomEndpoint"] = args ? args.cloudBuildWorkerPoolCustomEndpoint : undefined;
            inputs["cloudFunctionsCustomEndpoint"] = args ? args.cloudFunctionsCustomEndpoint : undefined;
            inputs["cloudIdentityCustomEndpoint"] = args ? args.cloudIdentityCustomEndpoint : undefined;
            inputs["cloudIotCustomEndpoint"] = args ? args.cloudIotCustomEndpoint : undefined;
            inputs["cloudRunCustomEndpoint"] = args ? args.cloudRunCustomEndpoint : undefined;
            inputs["cloudSchedulerCustomEndpoint"] = args ? args.cloudSchedulerCustomEndpoint : undefined;
            inputs["cloudTasksCustomEndpoint"] = args ? args.cloudTasksCustomEndpoint : undefined;
            inputs["composerCustomEndpoint"] = args ? args.composerCustomEndpoint : undefined;
            inputs["computeBetaCustomEndpoint"] = args ? args.computeBetaCustomEndpoint : undefined;
            inputs["computeCustomEndpoint"] = args ? args.computeCustomEndpoint : undefined;
            inputs["containerAnalysisCustomEndpoint"] = args ? args.containerAnalysisCustomEndpoint : undefined;
            inputs["containerBetaCustomEndpoint"] = args ? args.containerBetaCustomEndpoint : undefined;
            inputs["containerCustomEndpoint"] = args ? args.containerCustomEndpoint : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["dataCatalogCustomEndpoint"] = args ? args.dataCatalogCustomEndpoint : undefined;
            inputs["dataFusionCustomEndpoint"] = args ? args.dataFusionCustomEndpoint : undefined;
            inputs["dataLossPreventionCustomEndpoint"] = args ? args.dataLossPreventionCustomEndpoint : undefined;
            inputs["dataflowCustomEndpoint"] = args ? args.dataflowCustomEndpoint : undefined;
            inputs["dataprocBetaCustomEndpoint"] = args ? args.dataprocBetaCustomEndpoint : undefined;
            inputs["dataprocCustomEndpoint"] = args ? args.dataprocCustomEndpoint : undefined;
            inputs["dataprocMetastoreCustomEndpoint"] = args ? args.dataprocMetastoreCustomEndpoint : undefined;
            inputs["datastoreCustomEndpoint"] = args ? args.datastoreCustomEndpoint : undefined;
            inputs["deploymentManagerCustomEndpoint"] = args ? args.deploymentManagerCustomEndpoint : undefined;
            inputs["dialogflowCustomEndpoint"] = args ? args.dialogflowCustomEndpoint : undefined;
            inputs["dialogflowCxCustomEndpoint"] = args ? args.dialogflowCxCustomEndpoint : undefined;
            inputs["disableGooglePartnerName"] = pulumi.output(args ? args.disableGooglePartnerName : undefined).apply(JSON.stringify);
            inputs["dnsCustomEndpoint"] = args ? args.dnsCustomEndpoint : undefined;
            inputs["essentialContactsCustomEndpoint"] = args ? args.essentialContactsCustomEndpoint : undefined;
            inputs["eventarcCustomEndpoint"] = args ? args.eventarcCustomEndpoint : undefined;
            inputs["filestoreCustomEndpoint"] = args ? args.filestoreCustomEndpoint : undefined;
            inputs["firebaseCustomEndpoint"] = args ? args.firebaseCustomEndpoint : undefined;
            inputs["firestoreCustomEndpoint"] = args ? args.firestoreCustomEndpoint : undefined;
            inputs["gameServicesCustomEndpoint"] = args ? args.gameServicesCustomEndpoint : undefined;
            inputs["gkeHubCustomEndpoint"] = args ? args.gkeHubCustomEndpoint : undefined;
            inputs["gkehubFeatureCustomEndpoint"] = args ? args.gkehubFeatureCustomEndpoint : undefined;
            inputs["googlePartnerName"] = args ? args.googlePartnerName : undefined;
            inputs["healthcareCustomEndpoint"] = args ? args.healthcareCustomEndpoint : undefined;
            inputs["iamBetaCustomEndpoint"] = args ? args.iamBetaCustomEndpoint : undefined;
            inputs["iamCredentialsCustomEndpoint"] = args ? args.iamCredentialsCustomEndpoint : undefined;
            inputs["iamCustomEndpoint"] = args ? args.iamCustomEndpoint : undefined;
            inputs["iapCustomEndpoint"] = args ? args.iapCustomEndpoint : undefined;
            inputs["identityPlatformCustomEndpoint"] = args ? args.identityPlatformCustomEndpoint : undefined;
            inputs["impersonateServiceAccount"] = args ? args.impersonateServiceAccount : undefined;
            inputs["impersonateServiceAccountDelegates"] = pulumi.output(args ? args.impersonateServiceAccountDelegates : undefined).apply(JSON.stringify);
            inputs["kmsCustomEndpoint"] = args ? args.kmsCustomEndpoint : undefined;
            inputs["loggingCustomEndpoint"] = args ? args.loggingCustomEndpoint : undefined;
            inputs["memcacheCustomEndpoint"] = args ? args.memcacheCustomEndpoint : undefined;
            inputs["mlEngineCustomEndpoint"] = args ? args.mlEngineCustomEndpoint : undefined;
            inputs["monitoringCustomEndpoint"] = args ? args.monitoringCustomEndpoint : undefined;
            inputs["networkManagementCustomEndpoint"] = args ? args.networkManagementCustomEndpoint : undefined;
            inputs["networkServicesCustomEndpoint"] = args ? args.networkServicesCustomEndpoint : undefined;
            inputs["notebooksCustomEndpoint"] = args ? args.notebooksCustomEndpoint : undefined;
            inputs["osConfigCustomEndpoint"] = args ? args.osConfigCustomEndpoint : undefined;
            inputs["osLoginCustomEndpoint"] = args ? args.osLoginCustomEndpoint : undefined;
            inputs["privatecaCustomEndpoint"] = args ? args.privatecaCustomEndpoint : undefined;
            inputs["project"] = (args ? args.project : undefined) ?? utilities.getEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");
            inputs["pubsubCustomEndpoint"] = args ? args.pubsubCustomEndpoint : undefined;
            inputs["pubsubLiteCustomEndpoint"] = args ? args.pubsubLiteCustomEndpoint : undefined;
            inputs["redisCustomEndpoint"] = args ? args.redisCustomEndpoint : undefined;
            inputs["region"] = (args ? args.region : undefined) ?? utilities.getEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");
            inputs["requestReason"] = args ? args.requestReason : undefined;
            inputs["requestTimeout"] = args ? args.requestTimeout : undefined;
            inputs["resourceManagerCustomEndpoint"] = args ? args.resourceManagerCustomEndpoint : undefined;
            inputs["resourceManagerV2CustomEndpoint"] = args ? args.resourceManagerV2CustomEndpoint : undefined;
            inputs["runtimeConfigCustomEndpoint"] = args ? args.runtimeConfigCustomEndpoint : undefined;
            inputs["runtimeconfigCustomEndpoint"] = args ? args.runtimeconfigCustomEndpoint : undefined;
            inputs["scopes"] = pulumi.output(args ? args.scopes : undefined).apply(JSON.stringify);
            inputs["secretManagerCustomEndpoint"] = args ? args.secretManagerCustomEndpoint : undefined;
            inputs["securityCenterCustomEndpoint"] = args ? args.securityCenterCustomEndpoint : undefined;
            inputs["securityScannerCustomEndpoint"] = args ? args.securityScannerCustomEndpoint : undefined;
            inputs["serviceDirectoryCustomEndpoint"] = args ? args.serviceDirectoryCustomEndpoint : undefined;
            inputs["serviceManagementCustomEndpoint"] = args ? args.serviceManagementCustomEndpoint : undefined;
            inputs["serviceNetworkingCustomEndpoint"] = args ? args.serviceNetworkingCustomEndpoint : undefined;
            inputs["serviceUsageCustomEndpoint"] = args ? args.serviceUsageCustomEndpoint : undefined;
            inputs["sourceRepoCustomEndpoint"] = args ? args.sourceRepoCustomEndpoint : undefined;
            inputs["spannerCustomEndpoint"] = args ? args.spannerCustomEndpoint : undefined;
            inputs["sqlCustomEndpoint"] = args ? args.sqlCustomEndpoint : undefined;
            inputs["storageCustomEndpoint"] = args ? args.storageCustomEndpoint : undefined;
            inputs["storageTransferCustomEndpoint"] = args ? args.storageTransferCustomEndpoint : undefined;
            inputs["tagsCustomEndpoint"] = args ? args.tagsCustomEndpoint : undefined;
            inputs["tpuCustomEndpoint"] = args ? args.tpuCustomEndpoint : undefined;
            inputs["userProjectOverride"] = pulumi.output(args ? args.userProjectOverride : undefined).apply(JSON.stringify);
            inputs["vertexAiCustomEndpoint"] = args ? args.vertexAiCustomEndpoint : undefined;
            inputs["vpcAccessCustomEndpoint"] = args ? args.vpcAccessCustomEndpoint : undefined;
            inputs["workflowsCustomEndpoint"] = args ? args.workflowsCustomEndpoint : undefined;
            inputs["zone"] = (args ? args.zone : undefined) ?? utilities.getEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    accessApprovalCustomEndpoint?: pulumi.Input<string>;
    accessContextManagerCustomEndpoint?: pulumi.Input<string>;
    accessToken?: pulumi.Input<string>;
    activeDirectoryCustomEndpoint?: pulumi.Input<string>;
    apiGatewayCustomEndpoint?: pulumi.Input<string>;
    apigeeCustomEndpoint?: pulumi.Input<string>;
    appEngineCustomEndpoint?: pulumi.Input<string>;
    artifactRegistryCustomEndpoint?: pulumi.Input<string>;
    assuredWorkloadsCustomEndpoint?: pulumi.Input<string>;
    batching?: pulumi.Input<inputs.ProviderBatching>;
    bigQueryCustomEndpoint?: pulumi.Input<string>;
    bigqueryConnectionCustomEndpoint?: pulumi.Input<string>;
    bigqueryDataTransferCustomEndpoint?: pulumi.Input<string>;
    bigqueryReservationCustomEndpoint?: pulumi.Input<string>;
    bigtableCustomEndpoint?: pulumi.Input<string>;
    billingCustomEndpoint?: pulumi.Input<string>;
    billingProject?: pulumi.Input<string>;
    binaryAuthorizationCustomEndpoint?: pulumi.Input<string>;
    cloudAssetCustomEndpoint?: pulumi.Input<string>;
    cloudBillingCustomEndpoint?: pulumi.Input<string>;
    cloudBuildCustomEndpoint?: pulumi.Input<string>;
    cloudBuildWorkerPoolCustomEndpoint?: pulumi.Input<string>;
    cloudFunctionsCustomEndpoint?: pulumi.Input<string>;
    cloudIdentityCustomEndpoint?: pulumi.Input<string>;
    cloudIotCustomEndpoint?: pulumi.Input<string>;
    cloudRunCustomEndpoint?: pulumi.Input<string>;
    cloudSchedulerCustomEndpoint?: pulumi.Input<string>;
    cloudTasksCustomEndpoint?: pulumi.Input<string>;
    composerCustomEndpoint?: pulumi.Input<string>;
    computeBetaCustomEndpoint?: pulumi.Input<string>;
    computeCustomEndpoint?: pulumi.Input<string>;
    containerAnalysisCustomEndpoint?: pulumi.Input<string>;
    containerBetaCustomEndpoint?: pulumi.Input<string>;
    containerCustomEndpoint?: pulumi.Input<string>;
    credentials?: pulumi.Input<string>;
    dataCatalogCustomEndpoint?: pulumi.Input<string>;
    dataFusionCustomEndpoint?: pulumi.Input<string>;
    dataLossPreventionCustomEndpoint?: pulumi.Input<string>;
    dataflowCustomEndpoint?: pulumi.Input<string>;
    dataprocBetaCustomEndpoint?: pulumi.Input<string>;
    dataprocCustomEndpoint?: pulumi.Input<string>;
    dataprocMetastoreCustomEndpoint?: pulumi.Input<string>;
    datastoreCustomEndpoint?: pulumi.Input<string>;
    deploymentManagerCustomEndpoint?: pulumi.Input<string>;
    dialogflowCustomEndpoint?: pulumi.Input<string>;
    dialogflowCxCustomEndpoint?: pulumi.Input<string>;
    disableGooglePartnerName?: pulumi.Input<boolean>;
    dnsCustomEndpoint?: pulumi.Input<string>;
    essentialContactsCustomEndpoint?: pulumi.Input<string>;
    eventarcCustomEndpoint?: pulumi.Input<string>;
    filestoreCustomEndpoint?: pulumi.Input<string>;
    firebaseCustomEndpoint?: pulumi.Input<string>;
    firestoreCustomEndpoint?: pulumi.Input<string>;
    gameServicesCustomEndpoint?: pulumi.Input<string>;
    gkeHubCustomEndpoint?: pulumi.Input<string>;
    gkehubFeatureCustomEndpoint?: pulumi.Input<string>;
    googlePartnerName?: pulumi.Input<string>;
    healthcareCustomEndpoint?: pulumi.Input<string>;
    iamBetaCustomEndpoint?: pulumi.Input<string>;
    iamCredentialsCustomEndpoint?: pulumi.Input<string>;
    iamCustomEndpoint?: pulumi.Input<string>;
    iapCustomEndpoint?: pulumi.Input<string>;
    identityPlatformCustomEndpoint?: pulumi.Input<string>;
    impersonateServiceAccount?: pulumi.Input<string>;
    impersonateServiceAccountDelegates?: pulumi.Input<pulumi.Input<string>[]>;
    kmsCustomEndpoint?: pulumi.Input<string>;
    loggingCustomEndpoint?: pulumi.Input<string>;
    memcacheCustomEndpoint?: pulumi.Input<string>;
    mlEngineCustomEndpoint?: pulumi.Input<string>;
    monitoringCustomEndpoint?: pulumi.Input<string>;
    networkManagementCustomEndpoint?: pulumi.Input<string>;
    networkServicesCustomEndpoint?: pulumi.Input<string>;
    notebooksCustomEndpoint?: pulumi.Input<string>;
    osConfigCustomEndpoint?: pulumi.Input<string>;
    osLoginCustomEndpoint?: pulumi.Input<string>;
    privatecaCustomEndpoint?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    pubsubCustomEndpoint?: pulumi.Input<string>;
    pubsubLiteCustomEndpoint?: pulumi.Input<string>;
    redisCustomEndpoint?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    requestReason?: pulumi.Input<string>;
    requestTimeout?: pulumi.Input<string>;
    resourceManagerCustomEndpoint?: pulumi.Input<string>;
    resourceManagerV2CustomEndpoint?: pulumi.Input<string>;
    runtimeConfigCustomEndpoint?: pulumi.Input<string>;
    runtimeconfigCustomEndpoint?: pulumi.Input<string>;
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    secretManagerCustomEndpoint?: pulumi.Input<string>;
    securityCenterCustomEndpoint?: pulumi.Input<string>;
    securityScannerCustomEndpoint?: pulumi.Input<string>;
    serviceDirectoryCustomEndpoint?: pulumi.Input<string>;
    serviceManagementCustomEndpoint?: pulumi.Input<string>;
    serviceNetworkingCustomEndpoint?: pulumi.Input<string>;
    serviceUsageCustomEndpoint?: pulumi.Input<string>;
    sourceRepoCustomEndpoint?: pulumi.Input<string>;
    spannerCustomEndpoint?: pulumi.Input<string>;
    sqlCustomEndpoint?: pulumi.Input<string>;
    storageCustomEndpoint?: pulumi.Input<string>;
    storageTransferCustomEndpoint?: pulumi.Input<string>;
    tagsCustomEndpoint?: pulumi.Input<string>;
    tpuCustomEndpoint?: pulumi.Input<string>;
    userProjectOverride?: pulumi.Input<boolean>;
    vertexAiCustomEndpoint?: pulumi.Input<string>;
    vpcAccessCustomEndpoint?: pulumi.Input<string>;
    workflowsCustomEndpoint?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}
