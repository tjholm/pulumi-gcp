// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("gcp");

export declare const accessApprovalCustomEndpoint: string | undefined;
Object.defineProperty(exports, "accessApprovalCustomEndpoint", {
    get() {
        return __config.get("accessApprovalCustomEndpoint");
    },
    enumerable: true,
});

export declare const accessContextManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "accessContextManagerCustomEndpoint", {
    get() {
        return __config.get("accessContextManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const accessToken: string | undefined;
Object.defineProperty(exports, "accessToken", {
    get() {
        return __config.get("accessToken");
    },
    enumerable: true,
});

export declare const activeDirectoryCustomEndpoint: string | undefined;
Object.defineProperty(exports, "activeDirectoryCustomEndpoint", {
    get() {
        return __config.get("activeDirectoryCustomEndpoint");
    },
    enumerable: true,
});

export declare const addTerraformAttributionLabel: boolean | undefined;
Object.defineProperty(exports, "addTerraformAttributionLabel", {
    get() {
        return __config.getObject<boolean>("addTerraformAttributionLabel");
    },
    enumerable: true,
});

export declare const alloydbCustomEndpoint: string | undefined;
Object.defineProperty(exports, "alloydbCustomEndpoint", {
    get() {
        return __config.get("alloydbCustomEndpoint");
    },
    enumerable: true,
});

export declare const apiGatewayCustomEndpoint: string | undefined;
Object.defineProperty(exports, "apiGatewayCustomEndpoint", {
    get() {
        return __config.get("apiGatewayCustomEndpoint");
    },
    enumerable: true,
});

export declare const apigeeCustomEndpoint: string | undefined;
Object.defineProperty(exports, "apigeeCustomEndpoint", {
    get() {
        return __config.get("apigeeCustomEndpoint");
    },
    enumerable: true,
});

export declare const apikeysCustomEndpoint: string | undefined;
Object.defineProperty(exports, "apikeysCustomEndpoint", {
    get() {
        return __config.get("apikeysCustomEndpoint");
    },
    enumerable: true,
});

export declare const appEngineCustomEndpoint: string | undefined;
Object.defineProperty(exports, "appEngineCustomEndpoint", {
    get() {
        return __config.get("appEngineCustomEndpoint");
    },
    enumerable: true,
});

export declare const apphubCustomEndpoint: string | undefined;
Object.defineProperty(exports, "apphubCustomEndpoint", {
    get() {
        return __config.get("apphubCustomEndpoint");
    },
    enumerable: true,
});

export declare const artifactRegistryCustomEndpoint: string | undefined;
Object.defineProperty(exports, "artifactRegistryCustomEndpoint", {
    get() {
        return __config.get("artifactRegistryCustomEndpoint");
    },
    enumerable: true,
});

export declare const assuredWorkloadsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "assuredWorkloadsCustomEndpoint", {
    get() {
        return __config.get("assuredWorkloadsCustomEndpoint");
    },
    enumerable: true,
});

export declare const backupDrCustomEndpoint: string | undefined;
Object.defineProperty(exports, "backupDrCustomEndpoint", {
    get() {
        return __config.get("backupDrCustomEndpoint");
    },
    enumerable: true,
});

export declare const batching: outputs.config.Batching | undefined;
Object.defineProperty(exports, "batching", {
    get() {
        return __config.getObject<outputs.config.Batching>("batching");
    },
    enumerable: true,
});

export declare const beyondcorpCustomEndpoint: string | undefined;
Object.defineProperty(exports, "beyondcorpCustomEndpoint", {
    get() {
        return __config.get("beyondcorpCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigQueryCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigQueryCustomEndpoint", {
    get() {
        return __config.get("bigQueryCustomEndpoint");
    },
    enumerable: true,
});

export declare const biglakeCustomEndpoint: string | undefined;
Object.defineProperty(exports, "biglakeCustomEndpoint", {
    get() {
        return __config.get("biglakeCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigqueryAnalyticsHubCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigqueryAnalyticsHubCustomEndpoint", {
    get() {
        return __config.get("bigqueryAnalyticsHubCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigqueryConnectionCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigqueryConnectionCustomEndpoint", {
    get() {
        return __config.get("bigqueryConnectionCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigqueryDataTransferCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigqueryDataTransferCustomEndpoint", {
    get() {
        return __config.get("bigqueryDataTransferCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigqueryDatapolicyCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigqueryDatapolicyCustomEndpoint", {
    get() {
        return __config.get("bigqueryDatapolicyCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigqueryReservationCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigqueryReservationCustomEndpoint", {
    get() {
        return __config.get("bigqueryReservationCustomEndpoint");
    },
    enumerable: true,
});

export declare const bigtableCustomEndpoint: string | undefined;
Object.defineProperty(exports, "bigtableCustomEndpoint", {
    get() {
        return __config.get("bigtableCustomEndpoint");
    },
    enumerable: true,
});

export declare const billingCustomEndpoint: string | undefined;
Object.defineProperty(exports, "billingCustomEndpoint", {
    get() {
        return __config.get("billingCustomEndpoint");
    },
    enumerable: true,
});

export declare const billingProject: string | undefined;
Object.defineProperty(exports, "billingProject", {
    get() {
        return __config.get("billingProject");
    },
    enumerable: true,
});

export declare const binaryAuthorizationCustomEndpoint: string | undefined;
Object.defineProperty(exports, "binaryAuthorizationCustomEndpoint", {
    get() {
        return __config.get("binaryAuthorizationCustomEndpoint");
    },
    enumerable: true,
});

export declare const blockchainNodeEngineCustomEndpoint: string | undefined;
Object.defineProperty(exports, "blockchainNodeEngineCustomEndpoint", {
    get() {
        return __config.get("blockchainNodeEngineCustomEndpoint");
    },
    enumerable: true,
});

export declare const certificateManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "certificateManagerCustomEndpoint", {
    get() {
        return __config.get("certificateManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudAssetCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudAssetCustomEndpoint", {
    get() {
        return __config.get("cloudAssetCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudBillingCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudBillingCustomEndpoint", {
    get() {
        return __config.get("cloudBillingCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudBuildCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudBuildCustomEndpoint", {
    get() {
        return __config.get("cloudBuildCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudBuildWorkerPoolCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudBuildWorkerPoolCustomEndpoint", {
    get() {
        return __config.get("cloudBuildWorkerPoolCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudFunctionsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudFunctionsCustomEndpoint", {
    get() {
        return __config.get("cloudFunctionsCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudIdentityCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudIdentityCustomEndpoint", {
    get() {
        return __config.get("cloudIdentityCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudIdsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudIdsCustomEndpoint", {
    get() {
        return __config.get("cloudIdsCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudQuotasCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudQuotasCustomEndpoint", {
    get() {
        return __config.get("cloudQuotasCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudResourceManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudResourceManagerCustomEndpoint", {
    get() {
        return __config.get("cloudResourceManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudRunCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudRunCustomEndpoint", {
    get() {
        return __config.get("cloudRunCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudRunV2CustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudRunV2CustomEndpoint", {
    get() {
        return __config.get("cloudRunV2CustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudSchedulerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudSchedulerCustomEndpoint", {
    get() {
        return __config.get("cloudSchedulerCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudTasksCustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudTasksCustomEndpoint", {
    get() {
        return __config.get("cloudTasksCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudbuildv2CustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudbuildv2CustomEndpoint", {
    get() {
        return __config.get("cloudbuildv2CustomEndpoint");
    },
    enumerable: true,
});

export declare const clouddeployCustomEndpoint: string | undefined;
Object.defineProperty(exports, "clouddeployCustomEndpoint", {
    get() {
        return __config.get("clouddeployCustomEndpoint");
    },
    enumerable: true,
});

export declare const clouddomainsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "clouddomainsCustomEndpoint", {
    get() {
        return __config.get("clouddomainsCustomEndpoint");
    },
    enumerable: true,
});

export declare const cloudfunctions2CustomEndpoint: string | undefined;
Object.defineProperty(exports, "cloudfunctions2CustomEndpoint", {
    get() {
        return __config.get("cloudfunctions2CustomEndpoint");
    },
    enumerable: true,
});

export declare const composerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "composerCustomEndpoint", {
    get() {
        return __config.get("composerCustomEndpoint");
    },
    enumerable: true,
});

export declare const computeCustomEndpoint: string | undefined;
Object.defineProperty(exports, "computeCustomEndpoint", {
    get() {
        return __config.get("computeCustomEndpoint");
    },
    enumerable: true,
});

export declare const containerAnalysisCustomEndpoint: string | undefined;
Object.defineProperty(exports, "containerAnalysisCustomEndpoint", {
    get() {
        return __config.get("containerAnalysisCustomEndpoint");
    },
    enumerable: true,
});

export declare const containerAttachedCustomEndpoint: string | undefined;
Object.defineProperty(exports, "containerAttachedCustomEndpoint", {
    get() {
        return __config.get("containerAttachedCustomEndpoint");
    },
    enumerable: true,
});

export declare const containerAwsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "containerAwsCustomEndpoint", {
    get() {
        return __config.get("containerAwsCustomEndpoint");
    },
    enumerable: true,
});

export declare const containerAzureCustomEndpoint: string | undefined;
Object.defineProperty(exports, "containerAzureCustomEndpoint", {
    get() {
        return __config.get("containerAzureCustomEndpoint");
    },
    enumerable: true,
});

export declare const containerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "containerCustomEndpoint", {
    get() {
        return __config.get("containerCustomEndpoint");
    },
    enumerable: true,
});

export declare const coreBillingCustomEndpoint: string | undefined;
Object.defineProperty(exports, "coreBillingCustomEndpoint", {
    get() {
        return __config.get("coreBillingCustomEndpoint");
    },
    enumerable: true,
});

export declare const credentials: string | undefined;
Object.defineProperty(exports, "credentials", {
    get() {
        return __config.get("credentials");
    },
    enumerable: true,
});

export declare const dataCatalogCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataCatalogCustomEndpoint", {
    get() {
        return __config.get("dataCatalogCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataFusionCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataFusionCustomEndpoint", {
    get() {
        return __config.get("dataFusionCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataLossPreventionCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataLossPreventionCustomEndpoint", {
    get() {
        return __config.get("dataLossPreventionCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataPipelineCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataPipelineCustomEndpoint", {
    get() {
        return __config.get("dataPipelineCustomEndpoint");
    },
    enumerable: true,
});

export declare const databaseMigrationServiceCustomEndpoint: string | undefined;
Object.defineProperty(exports, "databaseMigrationServiceCustomEndpoint", {
    get() {
        return __config.get("databaseMigrationServiceCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataflowCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataflowCustomEndpoint", {
    get() {
        return __config.get("dataflowCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataformCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataformCustomEndpoint", {
    get() {
        return __config.get("dataformCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataplexCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataplexCustomEndpoint", {
    get() {
        return __config.get("dataplexCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataprocCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataprocCustomEndpoint", {
    get() {
        return __config.get("dataprocCustomEndpoint");
    },
    enumerable: true,
});

export declare const dataprocMetastoreCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dataprocMetastoreCustomEndpoint", {
    get() {
        return __config.get("dataprocMetastoreCustomEndpoint");
    },
    enumerable: true,
});

export declare const datastoreCustomEndpoint: string | undefined;
Object.defineProperty(exports, "datastoreCustomEndpoint", {
    get() {
        return __config.get("datastoreCustomEndpoint");
    },
    enumerable: true,
});

export declare const datastreamCustomEndpoint: string | undefined;
Object.defineProperty(exports, "datastreamCustomEndpoint", {
    get() {
        return __config.get("datastreamCustomEndpoint");
    },
    enumerable: true,
});

export declare const defaultLabels: {[key: string]: string} | undefined;
Object.defineProperty(exports, "defaultLabels", {
    get() {
        return __config.getObject<{[key: string]: string}>("defaultLabels");
    },
    enumerable: true,
});

export declare const deploymentManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "deploymentManagerCustomEndpoint", {
    get() {
        return __config.get("deploymentManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const dialogflowCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dialogflowCustomEndpoint", {
    get() {
        return __config.get("dialogflowCustomEndpoint");
    },
    enumerable: true,
});

export declare const dialogflowCxCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dialogflowCxCustomEndpoint", {
    get() {
        return __config.get("dialogflowCxCustomEndpoint");
    },
    enumerable: true,
});

export declare const disableGooglePartnerName: boolean | undefined;
Object.defineProperty(exports, "disableGooglePartnerName", {
    get() {
        return __config.getObject<boolean>("disableGooglePartnerName");
    },
    enumerable: true,
});

export declare const discoveryEngineCustomEndpoint: string | undefined;
Object.defineProperty(exports, "discoveryEngineCustomEndpoint", {
    get() {
        return __config.get("discoveryEngineCustomEndpoint");
    },
    enumerable: true,
});

export declare const dnsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "dnsCustomEndpoint", {
    get() {
        return __config.get("dnsCustomEndpoint");
    },
    enumerable: true,
});

export declare const documentAiCustomEndpoint: string | undefined;
Object.defineProperty(exports, "documentAiCustomEndpoint", {
    get() {
        return __config.get("documentAiCustomEndpoint");
    },
    enumerable: true,
});

export declare const documentAiWarehouseCustomEndpoint: string | undefined;
Object.defineProperty(exports, "documentAiWarehouseCustomEndpoint", {
    get() {
        return __config.get("documentAiWarehouseCustomEndpoint");
    },
    enumerable: true,
});

export declare const edgecontainerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "edgecontainerCustomEndpoint", {
    get() {
        return __config.get("edgecontainerCustomEndpoint");
    },
    enumerable: true,
});

export declare const edgenetworkCustomEndpoint: string | undefined;
Object.defineProperty(exports, "edgenetworkCustomEndpoint", {
    get() {
        return __config.get("edgenetworkCustomEndpoint");
    },
    enumerable: true,
});

export declare const essentialContactsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "essentialContactsCustomEndpoint", {
    get() {
        return __config.get("essentialContactsCustomEndpoint");
    },
    enumerable: true,
});

export declare const eventarcCustomEndpoint: string | undefined;
Object.defineProperty(exports, "eventarcCustomEndpoint", {
    get() {
        return __config.get("eventarcCustomEndpoint");
    },
    enumerable: true,
});

export declare const filestoreCustomEndpoint: string | undefined;
Object.defineProperty(exports, "filestoreCustomEndpoint", {
    get() {
        return __config.get("filestoreCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaseAppCheckCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaseAppCheckCustomEndpoint", {
    get() {
        return __config.get("firebaseAppCheckCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaseCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaseCustomEndpoint", {
    get() {
        return __config.get("firebaseCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaseDatabaseCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaseDatabaseCustomEndpoint", {
    get() {
        return __config.get("firebaseDatabaseCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaseExtensionsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaseExtensionsCustomEndpoint", {
    get() {
        return __config.get("firebaseExtensionsCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaseHostingCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaseHostingCustomEndpoint", {
    get() {
        return __config.get("firebaseHostingCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaseStorageCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaseStorageCustomEndpoint", {
    get() {
        return __config.get("firebaseStorageCustomEndpoint");
    },
    enumerable: true,
});

export declare const firebaserulesCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firebaserulesCustomEndpoint", {
    get() {
        return __config.get("firebaserulesCustomEndpoint");
    },
    enumerable: true,
});

export declare const firestoreCustomEndpoint: string | undefined;
Object.defineProperty(exports, "firestoreCustomEndpoint", {
    get() {
        return __config.get("firestoreCustomEndpoint");
    },
    enumerable: true,
});

export declare const gkeBackupCustomEndpoint: string | undefined;
Object.defineProperty(exports, "gkeBackupCustomEndpoint", {
    get() {
        return __config.get("gkeBackupCustomEndpoint");
    },
    enumerable: true,
});

export declare const gkeHub2CustomEndpoint: string | undefined;
Object.defineProperty(exports, "gkeHub2CustomEndpoint", {
    get() {
        return __config.get("gkeHub2CustomEndpoint");
    },
    enumerable: true,
});

export declare const gkeHubCustomEndpoint: string | undefined;
Object.defineProperty(exports, "gkeHubCustomEndpoint", {
    get() {
        return __config.get("gkeHubCustomEndpoint");
    },
    enumerable: true,
});

export declare const gkehubFeatureCustomEndpoint: string | undefined;
Object.defineProperty(exports, "gkehubFeatureCustomEndpoint", {
    get() {
        return __config.get("gkehubFeatureCustomEndpoint");
    },
    enumerable: true,
});

export declare const gkeonpremCustomEndpoint: string | undefined;
Object.defineProperty(exports, "gkeonpremCustomEndpoint", {
    get() {
        return __config.get("gkeonpremCustomEndpoint");
    },
    enumerable: true,
});

export declare const googlePartnerName: string | undefined;
Object.defineProperty(exports, "googlePartnerName", {
    get() {
        return __config.get("googlePartnerName");
    },
    enumerable: true,
});

export declare const healthcareCustomEndpoint: string | undefined;
Object.defineProperty(exports, "healthcareCustomEndpoint", {
    get() {
        return __config.get("healthcareCustomEndpoint");
    },
    enumerable: true,
});

export declare const iam2CustomEndpoint: string | undefined;
Object.defineProperty(exports, "iam2CustomEndpoint", {
    get() {
        return __config.get("iam2CustomEndpoint");
    },
    enumerable: true,
});

export declare const iamBetaCustomEndpoint: string | undefined;
Object.defineProperty(exports, "iamBetaCustomEndpoint", {
    get() {
        return __config.get("iamBetaCustomEndpoint");
    },
    enumerable: true,
});

export declare const iamCredentialsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "iamCredentialsCustomEndpoint", {
    get() {
        return __config.get("iamCredentialsCustomEndpoint");
    },
    enumerable: true,
});

export declare const iamCustomEndpoint: string | undefined;
Object.defineProperty(exports, "iamCustomEndpoint", {
    get() {
        return __config.get("iamCustomEndpoint");
    },
    enumerable: true,
});

export declare const iamWorkforcePoolCustomEndpoint: string | undefined;
Object.defineProperty(exports, "iamWorkforcePoolCustomEndpoint", {
    get() {
        return __config.get("iamWorkforcePoolCustomEndpoint");
    },
    enumerable: true,
});

export declare const iapCustomEndpoint: string | undefined;
Object.defineProperty(exports, "iapCustomEndpoint", {
    get() {
        return __config.get("iapCustomEndpoint");
    },
    enumerable: true,
});

export declare const identityPlatformCustomEndpoint: string | undefined;
Object.defineProperty(exports, "identityPlatformCustomEndpoint", {
    get() {
        return __config.get("identityPlatformCustomEndpoint");
    },
    enumerable: true,
});

export declare const impersonateServiceAccount: string | undefined;
Object.defineProperty(exports, "impersonateServiceAccount", {
    get() {
        return __config.get("impersonateServiceAccount");
    },
    enumerable: true,
});

export declare const impersonateServiceAccountDelegates: string[] | undefined;
Object.defineProperty(exports, "impersonateServiceAccountDelegates", {
    get() {
        return __config.getObject<string[]>("impersonateServiceAccountDelegates");
    },
    enumerable: true,
});

export declare const integrationConnectorsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "integrationConnectorsCustomEndpoint", {
    get() {
        return __config.get("integrationConnectorsCustomEndpoint");
    },
    enumerable: true,
});

export declare const kmsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "kmsCustomEndpoint", {
    get() {
        return __config.get("kmsCustomEndpoint");
    },
    enumerable: true,
});

export declare const loggingCustomEndpoint: string | undefined;
Object.defineProperty(exports, "loggingCustomEndpoint", {
    get() {
        return __config.get("loggingCustomEndpoint");
    },
    enumerable: true,
});

export declare const lookerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "lookerCustomEndpoint", {
    get() {
        return __config.get("lookerCustomEndpoint");
    },
    enumerable: true,
});

export declare const memcacheCustomEndpoint: string | undefined;
Object.defineProperty(exports, "memcacheCustomEndpoint", {
    get() {
        return __config.get("memcacheCustomEndpoint");
    },
    enumerable: true,
});

export declare const migrationCenterCustomEndpoint: string | undefined;
Object.defineProperty(exports, "migrationCenterCustomEndpoint", {
    get() {
        return __config.get("migrationCenterCustomEndpoint");
    },
    enumerable: true,
});

export declare const mlEngineCustomEndpoint: string | undefined;
Object.defineProperty(exports, "mlEngineCustomEndpoint", {
    get() {
        return __config.get("mlEngineCustomEndpoint");
    },
    enumerable: true,
});

export declare const monitoringCustomEndpoint: string | undefined;
Object.defineProperty(exports, "monitoringCustomEndpoint", {
    get() {
        return __config.get("monitoringCustomEndpoint");
    },
    enumerable: true,
});

export declare const netappCustomEndpoint: string | undefined;
Object.defineProperty(exports, "netappCustomEndpoint", {
    get() {
        return __config.get("netappCustomEndpoint");
    },
    enumerable: true,
});

export declare const networkConnectivityCustomEndpoint: string | undefined;
Object.defineProperty(exports, "networkConnectivityCustomEndpoint", {
    get() {
        return __config.get("networkConnectivityCustomEndpoint");
    },
    enumerable: true,
});

export declare const networkManagementCustomEndpoint: string | undefined;
Object.defineProperty(exports, "networkManagementCustomEndpoint", {
    get() {
        return __config.get("networkManagementCustomEndpoint");
    },
    enumerable: true,
});

export declare const networkSecurityCustomEndpoint: string | undefined;
Object.defineProperty(exports, "networkSecurityCustomEndpoint", {
    get() {
        return __config.get("networkSecurityCustomEndpoint");
    },
    enumerable: true,
});

export declare const networkServicesCustomEndpoint: string | undefined;
Object.defineProperty(exports, "networkServicesCustomEndpoint", {
    get() {
        return __config.get("networkServicesCustomEndpoint");
    },
    enumerable: true,
});

export declare const notebooksCustomEndpoint: string | undefined;
Object.defineProperty(exports, "notebooksCustomEndpoint", {
    get() {
        return __config.get("notebooksCustomEndpoint");
    },
    enumerable: true,
});

export declare const orgPolicyCustomEndpoint: string | undefined;
Object.defineProperty(exports, "orgPolicyCustomEndpoint", {
    get() {
        return __config.get("orgPolicyCustomEndpoint");
    },
    enumerable: true,
});

export declare const osConfigCustomEndpoint: string | undefined;
Object.defineProperty(exports, "osConfigCustomEndpoint", {
    get() {
        return __config.get("osConfigCustomEndpoint");
    },
    enumerable: true,
});

export declare const osLoginCustomEndpoint: string | undefined;
Object.defineProperty(exports, "osLoginCustomEndpoint", {
    get() {
        return __config.get("osLoginCustomEndpoint");
    },
    enumerable: true,
});

export declare const privatecaCustomEndpoint: string | undefined;
Object.defineProperty(exports, "privatecaCustomEndpoint", {
    get() {
        return __config.get("privatecaCustomEndpoint");
    },
    enumerable: true,
});

export declare const project: string | undefined;
Object.defineProperty(exports, "project", {
    get() {
        return __config.get("project") ?? utilities.getEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");
    },
    enumerable: true,
});

export declare const publicCaCustomEndpoint: string | undefined;
Object.defineProperty(exports, "publicCaCustomEndpoint", {
    get() {
        return __config.get("publicCaCustomEndpoint");
    },
    enumerable: true,
});

export declare const pubsubCustomEndpoint: string | undefined;
Object.defineProperty(exports, "pubsubCustomEndpoint", {
    get() {
        return __config.get("pubsubCustomEndpoint");
    },
    enumerable: true,
});

export declare const pubsubLiteCustomEndpoint: string | undefined;
Object.defineProperty(exports, "pubsubLiteCustomEndpoint", {
    get() {
        return __config.get("pubsubLiteCustomEndpoint");
    },
    enumerable: true,
});

export declare const recaptchaEnterpriseCustomEndpoint: string | undefined;
Object.defineProperty(exports, "recaptchaEnterpriseCustomEndpoint", {
    get() {
        return __config.get("recaptchaEnterpriseCustomEndpoint");
    },
    enumerable: true,
});

export declare const redisCustomEndpoint: string | undefined;
Object.defineProperty(exports, "redisCustomEndpoint", {
    get() {
        return __config.get("redisCustomEndpoint");
    },
    enumerable: true,
});

export declare const region: string | undefined;
Object.defineProperty(exports, "region", {
    get() {
        return __config.get("region") ?? utilities.getEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");
    },
    enumerable: true,
});

export declare const requestReason: string | undefined;
Object.defineProperty(exports, "requestReason", {
    get() {
        return __config.get("requestReason");
    },
    enumerable: true,
});

export declare const requestTimeout: string | undefined;
Object.defineProperty(exports, "requestTimeout", {
    get() {
        return __config.get("requestTimeout");
    },
    enumerable: true,
});

export declare const resourceManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "resourceManagerCustomEndpoint", {
    get() {
        return __config.get("resourceManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const resourceManagerV3CustomEndpoint: string | undefined;
Object.defineProperty(exports, "resourceManagerV3CustomEndpoint", {
    get() {
        return __config.get("resourceManagerV3CustomEndpoint");
    },
    enumerable: true,
});

export declare const runtimeConfigCustomEndpoint: string | undefined;
Object.defineProperty(exports, "runtimeConfigCustomEndpoint", {
    get() {
        return __config.get("runtimeConfigCustomEndpoint");
    },
    enumerable: true,
});

export declare const runtimeconfigCustomEndpoint: string | undefined;
Object.defineProperty(exports, "runtimeconfigCustomEndpoint", {
    get() {
        return __config.get("runtimeconfigCustomEndpoint");
    },
    enumerable: true,
});

export declare const scopes: string[] | undefined;
Object.defineProperty(exports, "scopes", {
    get() {
        return __config.getObject<string[]>("scopes");
    },
    enumerable: true,
});

export declare const secretManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "secretManagerCustomEndpoint", {
    get() {
        return __config.get("secretManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const secureSourceManagerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "secureSourceManagerCustomEndpoint", {
    get() {
        return __config.get("secureSourceManagerCustomEndpoint");
    },
    enumerable: true,
});

export declare const securityCenterCustomEndpoint: string | undefined;
Object.defineProperty(exports, "securityCenterCustomEndpoint", {
    get() {
        return __config.get("securityCenterCustomEndpoint");
    },
    enumerable: true,
});

export declare const securityScannerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "securityScannerCustomEndpoint", {
    get() {
        return __config.get("securityScannerCustomEndpoint");
    },
    enumerable: true,
});

export declare const securitypostureCustomEndpoint: string | undefined;
Object.defineProperty(exports, "securitypostureCustomEndpoint", {
    get() {
        return __config.get("securitypostureCustomEndpoint");
    },
    enumerable: true,
});

export declare const serviceDirectoryCustomEndpoint: string | undefined;
Object.defineProperty(exports, "serviceDirectoryCustomEndpoint", {
    get() {
        return __config.get("serviceDirectoryCustomEndpoint");
    },
    enumerable: true,
});

export declare const serviceManagementCustomEndpoint: string | undefined;
Object.defineProperty(exports, "serviceManagementCustomEndpoint", {
    get() {
        return __config.get("serviceManagementCustomEndpoint");
    },
    enumerable: true,
});

export declare const serviceNetworkingCustomEndpoint: string | undefined;
Object.defineProperty(exports, "serviceNetworkingCustomEndpoint", {
    get() {
        return __config.get("serviceNetworkingCustomEndpoint");
    },
    enumerable: true,
});

export declare const serviceUsageCustomEndpoint: string | undefined;
Object.defineProperty(exports, "serviceUsageCustomEndpoint", {
    get() {
        return __config.get("serviceUsageCustomEndpoint");
    },
    enumerable: true,
});

export declare const skipRegionValidation: boolean;
Object.defineProperty(exports, "skipRegionValidation", {
    get() {
        return __config.getObject<boolean>("skipRegionValidation") ?? (utilities.getEnvBoolean("PULUMI_GCP_SKIP_REGION_VALIDATION") || false);
    },
    enumerable: true,
});

export declare const sourceRepoCustomEndpoint: string | undefined;
Object.defineProperty(exports, "sourceRepoCustomEndpoint", {
    get() {
        return __config.get("sourceRepoCustomEndpoint");
    },
    enumerable: true,
});

export declare const spannerCustomEndpoint: string | undefined;
Object.defineProperty(exports, "spannerCustomEndpoint", {
    get() {
        return __config.get("spannerCustomEndpoint");
    },
    enumerable: true,
});

export declare const sqlCustomEndpoint: string | undefined;
Object.defineProperty(exports, "sqlCustomEndpoint", {
    get() {
        return __config.get("sqlCustomEndpoint");
    },
    enumerable: true,
});

export declare const storageCustomEndpoint: string | undefined;
Object.defineProperty(exports, "storageCustomEndpoint", {
    get() {
        return __config.get("storageCustomEndpoint");
    },
    enumerable: true,
});

export declare const storageInsightsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "storageInsightsCustomEndpoint", {
    get() {
        return __config.get("storageInsightsCustomEndpoint");
    },
    enumerable: true,
});

export declare const storageTransferCustomEndpoint: string | undefined;
Object.defineProperty(exports, "storageTransferCustomEndpoint", {
    get() {
        return __config.get("storageTransferCustomEndpoint");
    },
    enumerable: true,
});

export declare const tagsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "tagsCustomEndpoint", {
    get() {
        return __config.get("tagsCustomEndpoint");
    },
    enumerable: true,
});

export declare const tagsLocationCustomEndpoint: string | undefined;
Object.defineProperty(exports, "tagsLocationCustomEndpoint", {
    get() {
        return __config.get("tagsLocationCustomEndpoint");
    },
    enumerable: true,
});

export declare const terraformAttributionLabelAdditionStrategy: string | undefined;
Object.defineProperty(exports, "terraformAttributionLabelAdditionStrategy", {
    get() {
        return __config.get("terraformAttributionLabelAdditionStrategy");
    },
    enumerable: true,
});

export declare const tpuCustomEndpoint: string | undefined;
Object.defineProperty(exports, "tpuCustomEndpoint", {
    get() {
        return __config.get("tpuCustomEndpoint");
    },
    enumerable: true,
});

export declare const tpuV2CustomEndpoint: string | undefined;
Object.defineProperty(exports, "tpuV2CustomEndpoint", {
    get() {
        return __config.get("tpuV2CustomEndpoint");
    },
    enumerable: true,
});

export declare const universeDomain: string | undefined;
Object.defineProperty(exports, "universeDomain", {
    get() {
        return __config.get("universeDomain");
    },
    enumerable: true,
});

export declare const userProjectOverride: boolean | undefined;
Object.defineProperty(exports, "userProjectOverride", {
    get() {
        return __config.getObject<boolean>("userProjectOverride");
    },
    enumerable: true,
});

export declare const vertexAiCustomEndpoint: string | undefined;
Object.defineProperty(exports, "vertexAiCustomEndpoint", {
    get() {
        return __config.get("vertexAiCustomEndpoint");
    },
    enumerable: true,
});

export declare const vmwareengineCustomEndpoint: string | undefined;
Object.defineProperty(exports, "vmwareengineCustomEndpoint", {
    get() {
        return __config.get("vmwareengineCustomEndpoint");
    },
    enumerable: true,
});

export declare const vpcAccessCustomEndpoint: string | undefined;
Object.defineProperty(exports, "vpcAccessCustomEndpoint", {
    get() {
        return __config.get("vpcAccessCustomEndpoint");
    },
    enumerable: true,
});

export declare const workbenchCustomEndpoint: string | undefined;
Object.defineProperty(exports, "workbenchCustomEndpoint", {
    get() {
        return __config.get("workbenchCustomEndpoint");
    },
    enumerable: true,
});

export declare const workflowsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "workflowsCustomEndpoint", {
    get() {
        return __config.get("workflowsCustomEndpoint");
    },
    enumerable: true,
});

export declare const workstationsCustomEndpoint: string | undefined;
Object.defineProperty(exports, "workstationsCustomEndpoint", {
    get() {
        return __config.get("workstationsCustomEndpoint");
    },
    enumerable: true,
});

export declare const zone: string | undefined;
Object.defineProperty(exports, "zone", {
    get() {
        return __config.get("zone") ?? utilities.getEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");
    },
    enumerable: true,
});

