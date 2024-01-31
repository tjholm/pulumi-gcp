# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

accessApprovalCustomEndpoint: Optional[str]

accessContextManagerCustomEndpoint: Optional[str]

accessToken: Optional[str]

activeDirectoryCustomEndpoint: Optional[str]

alloydbCustomEndpoint: Optional[str]

apiGatewayCustomEndpoint: Optional[str]

apigeeCustomEndpoint: Optional[str]

apikeysCustomEndpoint: Optional[str]

appEngineCustomEndpoint: Optional[str]

artifactRegistryCustomEndpoint: Optional[str]

assuredWorkloadsCustomEndpoint: Optional[str]

backupDrCustomEndpoint: Optional[str]

batching: Optional[str]

beyondcorpCustomEndpoint: Optional[str]

bigQueryCustomEndpoint: Optional[str]

biglakeCustomEndpoint: Optional[str]

bigqueryAnalyticsHubCustomEndpoint: Optional[str]

bigqueryConnectionCustomEndpoint: Optional[str]

bigqueryDataTransferCustomEndpoint: Optional[str]

bigqueryDatapolicyCustomEndpoint: Optional[str]

bigqueryReservationCustomEndpoint: Optional[str]

bigtableCustomEndpoint: Optional[str]

billingCustomEndpoint: Optional[str]

billingProject: Optional[str]

binaryAuthorizationCustomEndpoint: Optional[str]

certificateManagerCustomEndpoint: Optional[str]

cloudAssetCustomEndpoint: Optional[str]

cloudBillingCustomEndpoint: Optional[str]

cloudBuildCustomEndpoint: Optional[str]

cloudBuildWorkerPoolCustomEndpoint: Optional[str]

cloudFunctionsCustomEndpoint: Optional[str]

cloudIdentityCustomEndpoint: Optional[str]

cloudIdsCustomEndpoint: Optional[str]

cloudResourceManagerCustomEndpoint: Optional[str]

cloudRunCustomEndpoint: Optional[str]

cloudRunV2CustomEndpoint: Optional[str]

cloudSchedulerCustomEndpoint: Optional[str]

cloudTasksCustomEndpoint: Optional[str]

cloudbuildv2CustomEndpoint: Optional[str]

clouddeployCustomEndpoint: Optional[str]

clouddomainsCustomEndpoint: Optional[str]

cloudfunctions2CustomEndpoint: Optional[str]

composerCustomEndpoint: Optional[str]

computeCustomEndpoint: Optional[str]

containerAnalysisCustomEndpoint: Optional[str]

containerAttachedCustomEndpoint: Optional[str]

containerAwsCustomEndpoint: Optional[str]

containerAzureCustomEndpoint: Optional[str]

containerCustomEndpoint: Optional[str]

coreBillingCustomEndpoint: Optional[str]

credentials: Optional[str]

dataCatalogCustomEndpoint: Optional[str]

dataFusionCustomEndpoint: Optional[str]

dataLossPreventionCustomEndpoint: Optional[str]

dataPipelineCustomEndpoint: Optional[str]

databaseMigrationServiceCustomEndpoint: Optional[str]

dataflowCustomEndpoint: Optional[str]

dataformCustomEndpoint: Optional[str]

dataplexCustomEndpoint: Optional[str]

dataprocCustomEndpoint: Optional[str]

dataprocMetastoreCustomEndpoint: Optional[str]

datastoreCustomEndpoint: Optional[str]

datastreamCustomEndpoint: Optional[str]

defaultLabels: Optional[str]

deploymentManagerCustomEndpoint: Optional[str]

dialogflowCustomEndpoint: Optional[str]

dialogflowCxCustomEndpoint: Optional[str]

disableGooglePartnerName: Optional[bool]

dnsCustomEndpoint: Optional[str]

documentAiCustomEndpoint: Optional[str]

documentAiWarehouseCustomEndpoint: Optional[str]

edgecontainerCustomEndpoint: Optional[str]

edgenetworkCustomEndpoint: Optional[str]

essentialContactsCustomEndpoint: Optional[str]

eventarcCustomEndpoint: Optional[str]

filestoreCustomEndpoint: Optional[str]

firebaseCustomEndpoint: Optional[str]

firebaseDatabaseCustomEndpoint: Optional[str]

firebaseExtensionsCustomEndpoint: Optional[str]

firebaseHostingCustomEndpoint: Optional[str]

firebaseStorageCustomEndpoint: Optional[str]

firebaserulesCustomEndpoint: Optional[str]

firestoreCustomEndpoint: Optional[str]

gkeBackupCustomEndpoint: Optional[str]

gkeHub2CustomEndpoint: Optional[str]

gkeHubCustomEndpoint: Optional[str]

gkehubFeatureCustomEndpoint: Optional[str]

gkeonpremCustomEndpoint: Optional[str]

googlePartnerName: Optional[str]

healthcareCustomEndpoint: Optional[str]

iam2CustomEndpoint: Optional[str]

iamBetaCustomEndpoint: Optional[str]

iamCredentialsCustomEndpoint: Optional[str]

iamCustomEndpoint: Optional[str]

iamWorkforcePoolCustomEndpoint: Optional[str]

iapCustomEndpoint: Optional[str]

identityPlatformCustomEndpoint: Optional[str]

impersonateServiceAccount: Optional[str]

impersonateServiceAccountDelegates: Optional[str]

integrationConnectorsCustomEndpoint: Optional[str]

kmsCustomEndpoint: Optional[str]

loggingCustomEndpoint: Optional[str]

lookerCustomEndpoint: Optional[str]

memcacheCustomEndpoint: Optional[str]

migrationCenterCustomEndpoint: Optional[str]

mlEngineCustomEndpoint: Optional[str]

monitoringCustomEndpoint: Optional[str]

netappCustomEndpoint: Optional[str]

networkConnectivityCustomEndpoint: Optional[str]

networkManagementCustomEndpoint: Optional[str]

networkSecurityCustomEndpoint: Optional[str]

networkServicesCustomEndpoint: Optional[str]

notebooksCustomEndpoint: Optional[str]

orgPolicyCustomEndpoint: Optional[str]

osConfigCustomEndpoint: Optional[str]

osLoginCustomEndpoint: Optional[str]

privatecaCustomEndpoint: Optional[str]

project: Optional[str]

publicCaCustomEndpoint: Optional[str]

pubsubCustomEndpoint: Optional[str]

pubsubLiteCustomEndpoint: Optional[str]

recaptchaEnterpriseCustomEndpoint: Optional[str]

redisCustomEndpoint: Optional[str]

region: Optional[str]

requestReason: Optional[str]

requestTimeout: Optional[str]

resourceManagerCustomEndpoint: Optional[str]

resourceManagerV3CustomEndpoint: Optional[str]

runtimeConfigCustomEndpoint: Optional[str]

runtimeconfigCustomEndpoint: Optional[str]

scopes: Optional[str]

secretManagerCustomEndpoint: Optional[str]

secureSourceManagerCustomEndpoint: Optional[str]

securityCenterCustomEndpoint: Optional[str]

securityScannerCustomEndpoint: Optional[str]

serviceDirectoryCustomEndpoint: Optional[str]

serviceManagementCustomEndpoint: Optional[str]

serviceNetworkingCustomEndpoint: Optional[str]

serviceUsageCustomEndpoint: Optional[str]

skipRegionValidation: bool

sourceRepoCustomEndpoint: Optional[str]

spannerCustomEndpoint: Optional[str]

sqlCustomEndpoint: Optional[str]

storageCustomEndpoint: Optional[str]

storageInsightsCustomEndpoint: Optional[str]

storageTransferCustomEndpoint: Optional[str]

tagsCustomEndpoint: Optional[str]

tagsLocationCustomEndpoint: Optional[str]

tpuCustomEndpoint: Optional[str]

tpuV2CustomEndpoint: Optional[str]

universeDomain: Optional[str]

userProjectOverride: Optional[bool]

vertexAiCustomEndpoint: Optional[str]

vmwareengineCustomEndpoint: Optional[str]

vpcAccessCustomEndpoint: Optional[str]

workbenchCustomEndpoint: Optional[str]

workflowsCustomEndpoint: Optional[str]

workstationsCustomEndpoint: Optional[str]

zone: Optional[str]

