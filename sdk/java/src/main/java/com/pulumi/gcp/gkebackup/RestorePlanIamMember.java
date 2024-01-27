// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkebackup;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkebackup.RestorePlanIamMemberArgs;
import com.pulumi.gcp.gkebackup.inputs.RestorePlanIamMemberState;
import com.pulumi.gcp.gkebackup.outputs.RestorePlanIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a Restore Plan instance.
 * 
 * To get more information about RestorePlan, see:
 * 
 * * [API documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/projects.locations.restorePlans)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke)
 * 
 * ## Example Usage
 * ### Gkebackup Restoreplan All Namespaces
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
 * import com.pulumi.gcp.gkebackup.BackupPlan;
 * import com.pulumi.gcp.gkebackup.BackupPlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.BackupPlanBackupConfigArgs;
 * import com.pulumi.gcp.gkebackup.RestorePlan;
 * import com.pulumi.gcp.gkebackup.RestorePlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool(&#34;my-project-name.svc.id.goog&#34;)
 *                 .build())
 *             .addonsConfig(ClusterAddonsConfigArgs.builder()
 *                 .gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;&#34;)
 *             .network(&#34;default&#34;)
 *             .subnetwork(&#34;default&#34;)
 *             .build());
 * 
 *         var basic = new BackupPlan(&#34;basic&#34;, BackupPlanArgs.builder()        
 *             .cluster(primary.id())
 *             .location(&#34;us-central1&#34;)
 *             .backupConfig(BackupPlanBackupConfigArgs.builder()
 *                 .includeVolumeData(true)
 *                 .includeSecrets(true)
 *                 .allNamespaces(true)
 *                 .build())
 *             .build());
 * 
 *         var allNs = new RestorePlan(&#34;allNs&#34;, RestorePlanArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupPlan(basic.id())
 *             .cluster(primary.id())
 *             .restoreConfig(RestorePlanRestoreConfigArgs.builder()
 *                 .allNamespaces(true)
 *                 .namespacedResourceRestoreMode(&#34;FAIL_ON_CONFLICT&#34;)
 *                 .volumeDataRestorePolicy(&#34;RESTORE_VOLUME_DATA_FROM_BACKUP&#34;)
 *                 .clusterResourceRestoreScope(RestorePlanRestoreConfigClusterResourceRestoreScopeArgs.builder()
 *                     .allGroupKinds(true)
 *                     .build())
 *                 .clusterResourceConflictPolicy(&#34;USE_EXISTING_VERSION&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkebackup Restoreplan Rollback Namespace
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
 * import com.pulumi.gcp.gkebackup.BackupPlan;
 * import com.pulumi.gcp.gkebackup.BackupPlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.BackupPlanBackupConfigArgs;
 * import com.pulumi.gcp.gkebackup.RestorePlan;
 * import com.pulumi.gcp.gkebackup.RestorePlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigSelectedNamespacesArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool(&#34;my-project-name.svc.id.goog&#34;)
 *                 .build())
 *             .addonsConfig(ClusterAddonsConfigArgs.builder()
 *                 .gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;&#34;)
 *             .network(&#34;default&#34;)
 *             .subnetwork(&#34;default&#34;)
 *             .build());
 * 
 *         var basic = new BackupPlan(&#34;basic&#34;, BackupPlanArgs.builder()        
 *             .cluster(primary.id())
 *             .location(&#34;us-central1&#34;)
 *             .backupConfig(BackupPlanBackupConfigArgs.builder()
 *                 .includeVolumeData(true)
 *                 .includeSecrets(true)
 *                 .allNamespaces(true)
 *                 .build())
 *             .build());
 * 
 *         var rollbackNs = new RestorePlan(&#34;rollbackNs&#34;, RestorePlanArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupPlan(basic.id())
 *             .cluster(primary.id())
 *             .restoreConfig(RestorePlanRestoreConfigArgs.builder()
 *                 .selectedNamespaces(RestorePlanRestoreConfigSelectedNamespacesArgs.builder()
 *                     .namespaces(&#34;my-ns&#34;)
 *                     .build())
 *                 .namespacedResourceRestoreMode(&#34;DELETE_AND_RESTORE&#34;)
 *                 .volumeDataRestorePolicy(&#34;RESTORE_VOLUME_DATA_FROM_BACKUP&#34;)
 *                 .clusterResourceRestoreScope(RestorePlanRestoreConfigClusterResourceRestoreScopeArgs.builder()
 *                     .selectedGroupKinds(                    
 *                         RestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindArgs.builder()
 *                             .resourceGroup(&#34;apiextension.k8s.io&#34;)
 *                             .resourceKind(&#34;CustomResourceDefinition&#34;)
 *                             .build(),
 *                         RestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindArgs.builder()
 *                             .resourceGroup(&#34;storage.k8s.io&#34;)
 *                             .resourceKind(&#34;StorageClass&#34;)
 *                             .build())
 *                     .build())
 *                 .clusterResourceConflictPolicy(&#34;USE_EXISTING_VERSION&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkebackup Restoreplan Protected Application
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
 * import com.pulumi.gcp.gkebackup.BackupPlan;
 * import com.pulumi.gcp.gkebackup.BackupPlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.BackupPlanBackupConfigArgs;
 * import com.pulumi.gcp.gkebackup.RestorePlan;
 * import com.pulumi.gcp.gkebackup.RestorePlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigSelectedApplicationsArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool(&#34;my-project-name.svc.id.goog&#34;)
 *                 .build())
 *             .addonsConfig(ClusterAddonsConfigArgs.builder()
 *                 .gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;&#34;)
 *             .network(&#34;default&#34;)
 *             .subnetwork(&#34;default&#34;)
 *             .build());
 * 
 *         var basic = new BackupPlan(&#34;basic&#34;, BackupPlanArgs.builder()        
 *             .cluster(primary.id())
 *             .location(&#34;us-central1&#34;)
 *             .backupConfig(BackupPlanBackupConfigArgs.builder()
 *                 .includeVolumeData(true)
 *                 .includeSecrets(true)
 *                 .allNamespaces(true)
 *                 .build())
 *             .build());
 * 
 *         var rollbackApp = new RestorePlan(&#34;rollbackApp&#34;, RestorePlanArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupPlan(basic.id())
 *             .cluster(primary.id())
 *             .restoreConfig(RestorePlanRestoreConfigArgs.builder()
 *                 .selectedApplications(RestorePlanRestoreConfigSelectedApplicationsArgs.builder()
 *                     .namespacedNames(RestorePlanRestoreConfigSelectedApplicationsNamespacedNameArgs.builder()
 *                         .name(&#34;my-app&#34;)
 *                         .namespace(&#34;my-ns&#34;)
 *                         .build())
 *                     .build())
 *                 .namespacedResourceRestoreMode(&#34;DELETE_AND_RESTORE&#34;)
 *                 .volumeDataRestorePolicy(&#34;REUSE_VOLUME_HANDLE_FROM_BACKUP&#34;)
 *                 .clusterResourceRestoreScope(RestorePlanRestoreConfigClusterResourceRestoreScopeArgs.builder()
 *                     .noGroupKinds(true)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkebackup Restoreplan All Cluster Resources
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
 * import com.pulumi.gcp.gkebackup.BackupPlan;
 * import com.pulumi.gcp.gkebackup.BackupPlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.BackupPlanBackupConfigArgs;
 * import com.pulumi.gcp.gkebackup.RestorePlan;
 * import com.pulumi.gcp.gkebackup.RestorePlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool(&#34;my-project-name.svc.id.goog&#34;)
 *                 .build())
 *             .addonsConfig(ClusterAddonsConfigArgs.builder()
 *                 .gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;&#34;)
 *             .network(&#34;default&#34;)
 *             .subnetwork(&#34;default&#34;)
 *             .build());
 * 
 *         var basic = new BackupPlan(&#34;basic&#34;, BackupPlanArgs.builder()        
 *             .cluster(primary.id())
 *             .location(&#34;us-central1&#34;)
 *             .backupConfig(BackupPlanBackupConfigArgs.builder()
 *                 .includeVolumeData(true)
 *                 .includeSecrets(true)
 *                 .allNamespaces(true)
 *                 .build())
 *             .build());
 * 
 *         var allClusterResources = new RestorePlan(&#34;allClusterResources&#34;, RestorePlanArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupPlan(basic.id())
 *             .cluster(primary.id())
 *             .restoreConfig(RestorePlanRestoreConfigArgs.builder()
 *                 .noNamespaces(true)
 *                 .namespacedResourceRestoreMode(&#34;FAIL_ON_CONFLICT&#34;)
 *                 .clusterResourceRestoreScope(RestorePlanRestoreConfigClusterResourceRestoreScopeArgs.builder()
 *                     .allGroupKinds(true)
 *                     .build())
 *                 .clusterResourceConflictPolicy(&#34;USE_EXISTING_VERSION&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkebackup Restoreplan Rename Namespace
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
 * import com.pulumi.gcp.gkebackup.BackupPlan;
 * import com.pulumi.gcp.gkebackup.BackupPlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.BackupPlanBackupConfigArgs;
 * import com.pulumi.gcp.gkebackup.RestorePlan;
 * import com.pulumi.gcp.gkebackup.RestorePlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigSelectedNamespacesArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool(&#34;my-project-name.svc.id.goog&#34;)
 *                 .build())
 *             .addonsConfig(ClusterAddonsConfigArgs.builder()
 *                 .gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;&#34;)
 *             .network(&#34;default&#34;)
 *             .subnetwork(&#34;default&#34;)
 *             .build());
 * 
 *         var basic = new BackupPlan(&#34;basic&#34;, BackupPlanArgs.builder()        
 *             .cluster(primary.id())
 *             .location(&#34;us-central1&#34;)
 *             .backupConfig(BackupPlanBackupConfigArgs.builder()
 *                 .includeVolumeData(true)
 *                 .includeSecrets(true)
 *                 .allNamespaces(true)
 *                 .build())
 *             .build());
 * 
 *         var renameNs = new RestorePlan(&#34;renameNs&#34;, RestorePlanArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupPlan(basic.id())
 *             .cluster(primary.id())
 *             .restoreConfig(RestorePlanRestoreConfigArgs.builder()
 *                 .selectedNamespaces(RestorePlanRestoreConfigSelectedNamespacesArgs.builder()
 *                     .namespaces(&#34;ns1&#34;)
 *                     .build())
 *                 .namespacedResourceRestoreMode(&#34;FAIL_ON_CONFLICT&#34;)
 *                 .volumeDataRestorePolicy(&#34;REUSE_VOLUME_HANDLE_FROM_BACKUP&#34;)
 *                 .clusterResourceRestoreScope(RestorePlanRestoreConfigClusterResourceRestoreScopeArgs.builder()
 *                     .noGroupKinds(true)
 *                     .build())
 *                 .transformationRules(                
 *                     RestorePlanRestoreConfigTransformationRuleArgs.builder()
 *                         .description(&#34;rename namespace from ns1 to ns2&#34;)
 *                         .resourceFilter(RestorePlanRestoreConfigTransformationRuleResourceFilterArgs.builder()
 *                             .groupKinds(RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKindArgs.builder()
 *                                 .resourceKind(&#34;Namespace&#34;)
 *                                 .build())
 *                             .jsonPath(&#34;.metadata[?(@.name == &#39;ns1&#39;)]&#34;)
 *                             .build())
 *                         .fieldActions(RestorePlanRestoreConfigTransformationRuleFieldActionArgs.builder()
 *                             .op(&#34;REPLACE&#34;)
 *                             .path(&#34;/metadata/name&#34;)
 *                             .value(&#34;ns2&#34;)
 *                             .build())
 *                         .build(),
 *                     RestorePlanRestoreConfigTransformationRuleArgs.builder()
 *                         .description(&#34;move all resources from ns1 to ns2&#34;)
 *                         .resourceFilter(RestorePlanRestoreConfigTransformationRuleResourceFilterArgs.builder()
 *                             .namespaces(&#34;ns1&#34;)
 *                             .build())
 *                         .fieldActions(RestorePlanRestoreConfigTransformationRuleFieldActionArgs.builder()
 *                             .op(&#34;REPLACE&#34;)
 *                             .path(&#34;/metadata/namespace&#34;)
 *                             .value(&#34;ns2&#34;)
 *                             .build())
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkebackup Restoreplan Second Transformation
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigArgs;
 * import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
 * import com.pulumi.gcp.gkebackup.BackupPlan;
 * import com.pulumi.gcp.gkebackup.BackupPlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.BackupPlanBackupConfigArgs;
 * import com.pulumi.gcp.gkebackup.RestorePlan;
 * import com.pulumi.gcp.gkebackup.RestorePlanArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigExcludedNamespacesArgs;
 * import com.pulumi.gcp.gkebackup.inputs.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool(&#34;my-project-name.svc.id.goog&#34;)
 *                 .build())
 *             .addonsConfig(ClusterAddonsConfigArgs.builder()
 *                 .gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .deletionProtection(&#34;&#34;)
 *             .network(&#34;default&#34;)
 *             .subnetwork(&#34;default&#34;)
 *             .build());
 * 
 *         var basic = new BackupPlan(&#34;basic&#34;, BackupPlanArgs.builder()        
 *             .cluster(primary.id())
 *             .location(&#34;us-central1&#34;)
 *             .backupConfig(BackupPlanBackupConfigArgs.builder()
 *                 .includeVolumeData(true)
 *                 .includeSecrets(true)
 *                 .allNamespaces(true)
 *                 .build())
 *             .build());
 * 
 *         var transformRule = new RestorePlan(&#34;transformRule&#34;, RestorePlanArgs.builder()        
 *             .description(&#34;copy nginx env variables&#34;)
 *             .labels(Map.of(&#34;app&#34;, &#34;nginx&#34;))
 *             .location(&#34;us-central1&#34;)
 *             .backupPlan(basic.id())
 *             .cluster(primary.id())
 *             .restoreConfig(RestorePlanRestoreConfigArgs.builder()
 *                 .excludedNamespaces(RestorePlanRestoreConfigExcludedNamespacesArgs.builder()
 *                     .namespaces(&#34;my-ns&#34;)
 *                     .build())
 *                 .namespacedResourceRestoreMode(&#34;DELETE_AND_RESTORE&#34;)
 *                 .volumeDataRestorePolicy(&#34;RESTORE_VOLUME_DATA_FROM_BACKUP&#34;)
 *                 .clusterResourceRestoreScope(RestorePlanRestoreConfigClusterResourceRestoreScopeArgs.builder()
 *                     .excludedGroupKinds(RestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKindArgs.builder()
 *                         .resourceGroup(&#34;apiextension.k8s.io&#34;)
 *                         .resourceKind(&#34;CustomResourceDefinition&#34;)
 *                         .build())
 *                     .build())
 *                 .clusterResourceConflictPolicy(&#34;USE_EXISTING_VERSION&#34;)
 *                 .transformationRules(RestorePlanRestoreConfigTransformationRuleArgs.builder()
 *                     .description(&#34;Copy environment variables from the nginx container to the install init container.&#34;)
 *                     .resourceFilter(RestorePlanRestoreConfigTransformationRuleResourceFilterArgs.builder()
 *                         .groupKinds(RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKindArgs.builder()
 *                             .resourceKind(&#34;Pod&#34;)
 *                             .resourceGroup(&#34;&#34;)
 *                             .build())
 *                         .jsonPath(&#34;.metadata[?(@.name == &#39;nginx&#39;)]&#34;)
 *                         .build())
 *                     .fieldActions(RestorePlanRestoreConfigTransformationRuleFieldActionArgs.builder()
 *                         .op(&#34;COPY&#34;)
 *                         .path(&#34;/spec/initContainers/0/env&#34;)
 *                         .fromPath(&#34;/spec/containers/0/env&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RestorePlan can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/restorePlans/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, RestorePlan can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:gkebackup/restorePlanIamMember:RestorePlanIamMember default projects/{{project}}/locations/{{location}}/restorePlans/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkebackup/restorePlanIamMember:RestorePlanIamMember default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkebackup/restorePlanIamMember:RestorePlanIamMember default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:gkebackup/restorePlanIamMember:RestorePlanIamMember")
public class RestorePlanIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", refs={RestorePlanIamMemberCondition.class}, tree="[0]")
    private Output</* @Nullable */ RestorePlanIamMemberCondition> condition;

    public Output<Optional<RestorePlanIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The region of the Restore Plan.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The region of the Restore Plan.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    @Export(name="member", refs={String.class}, tree="[0]")
    private Output<String> member;

    public Output<String> member() {
        return this.member;
    }
    /**
     * The full name of the BackupPlan Resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The full name of the BackupPlan Resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RestorePlanIamMember(String name) {
        this(name, RestorePlanIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RestorePlanIamMember(String name, RestorePlanIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RestorePlanIamMember(String name, RestorePlanIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkebackup/restorePlanIamMember:RestorePlanIamMember", name, args == null ? RestorePlanIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RestorePlanIamMember(String name, Output<String> id, @Nullable RestorePlanIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkebackup/restorePlanIamMember:RestorePlanIamMember", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RestorePlanIamMember get(String name, Output<String> id, @Nullable RestorePlanIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RestorePlanIamMember(name, id, state, options);
    }
}
