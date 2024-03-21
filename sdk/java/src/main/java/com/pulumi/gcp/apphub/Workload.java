// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apphub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apphub.WorkloadArgs;
import com.pulumi.gcp.apphub.inputs.WorkloadState;
import com.pulumi.gcp.apphub.outputs.WorkloadAttributes;
import com.pulumi.gcp.apphub.outputs.WorkloadWorkloadProperty;
import com.pulumi.gcp.apphub.outputs.WorkloadWorkloadReference;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Workload represents a binary deployment (such as Managed Instance Groups (MIGs), GKE deployments, etc.) that performs the smallest logical subset of business functionality. It registers identified workload to the Application.
 * 
 * ## Example Usage
 * 
 * ### Apphub Workload Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apphub.Application;
 * import com.pulumi.gcp.apphub.ApplicationArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationScopeArgs;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.time.sleep;
 * import com.pulumi.time.SleepArgs;
 * import com.pulumi.gcp.apphub.ServiceProjectAttachment;
 * import com.pulumi.gcp.apphub.ServiceProjectAttachmentArgs;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
 * import com.pulumi.gcp.compute.RegionInstanceGroupManager;
 * import com.pulumi.gcp.compute.RegionInstanceGroupManagerArgs;
 * import com.pulumi.gcp.compute.inputs.RegionInstanceGroupManagerVersionArgs;
 * import com.pulumi.gcp.apphub.ApphubFunctions;
 * import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
 * import com.pulumi.gcp.apphub.Workload;
 * import com.pulumi.gcp.apphub.WorkloadArgs;
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
 *         var application = new Application(&#34;application&#34;, ApplicationArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .applicationId(&#34;example-application-1&#34;)
 *             .scope(ApplicationScopeArgs.builder()
 *                 .type(&#34;REGIONAL&#34;)
 *                 .build())
 *             .build());
 * 
 *         var serviceProject = new Project(&#34;serviceProject&#34;, ProjectArgs.builder()        
 *             .projectId(&#34;project-1&#34;)
 *             .name(&#34;Service Project&#34;)
 *             .orgId(&#34;123456789&#34;)
 *             .billingAccount(&#34;000000-0000000-0000000-000000&#34;)
 *             .build());
 * 
 *         var computeServiceProject = new Service(&#34;computeServiceProject&#34;, ServiceArgs.builder()        
 *             .project(serviceProject.projectId())
 *             .service(&#34;compute.googleapis.com&#34;)
 *             .build());
 * 
 *         var wait120s = new Sleep(&#34;wait120s&#34;, SleepArgs.builder()        
 *             .createDuration(&#34;120s&#34;)
 *             .build());
 * 
 *         var serviceProjectAttachment = new ServiceProjectAttachment(&#34;serviceProjectAttachment&#34;, ServiceProjectAttachmentArgs.builder()        
 *             .serviceProjectAttachmentId(serviceProject.projectId())
 *             .build());
 * 
 *         var ilbNetwork = new Network(&#34;ilbNetwork&#34;, NetworkArgs.builder()        
 *             .name(&#34;l7-ilb-network&#34;)
 *             .project(serviceProject.projectId())
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var ilbSubnet = new Subnetwork(&#34;ilbSubnet&#34;, SubnetworkArgs.builder()        
 *             .name(&#34;l7-ilb-subnet&#34;)
 *             .project(serviceProject.projectId())
 *             .ipCidrRange(&#34;10.0.1.0/24&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .network(ilbNetwork.id())
 *             .build());
 * 
 *         var instanceTemplate = new InstanceTemplate(&#34;instanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .accessConfigs()
 *                 .network(ilbNetwork.id())
 *                 .subnetwork(ilbSubnet.id())
 *                 .build())
 *             .name(&#34;l7-ilb-mig-template&#34;)
 *             .project(serviceProject.projectId())
 *             .machineType(&#34;e2-small&#34;)
 *             .tags(&#34;http-server&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(&#34;debian-cloud/debian-10&#34;)
 *                 .autoDelete(true)
 *                 .boot(true)
 *                 .build())
 *             .metadata(Map.of(&#34;startup-script&#34;, &#34;&#34;&#34;
 * #! /bin/bash
 * set -euo pipefail
 * export DEBIAN_FRONTEND=noninteractive
 * apt-get update
 * apt-get install -y nginx-light jq
 * NAME=$(curl -H &#34;Metadata-Flavor: Google&#34; &#34;http://metadata.google.internal/computeMetadata/v1/instance/hostname&#34;)
 * IP=$(curl -H &#34;Metadata-Flavor: Google&#34; &#34;http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip&#34;)
 * METADATA=$(curl -f -H &#34;Metadata-Flavor: Google&#34; &#34;http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=True&#34; | jq &#39;del(.[&#34;startup-script&#34;])&#39;)
 * cat &lt;&lt;EOF &gt; /var/www/html/index.html
 * &lt;pre&gt;
 * Name: $NAME
 * IP: $IP
 * Metadata: $METADATA
 * &lt;/pre&gt;
 * EOF
 *             &#34;&#34;&#34;))
 *             .build());
 * 
 *         var mig = new RegionInstanceGroupManager(&#34;mig&#34;, RegionInstanceGroupManagerArgs.builder()        
 *             .name(&#34;l7-ilb-mig1&#34;)
 *             .project(serviceProject.projectId())
 *             .region(&#34;us-central1&#34;)
 *             .versions(RegionInstanceGroupManagerVersionArgs.builder()
 *                 .instanceTemplate(instanceTemplate.id())
 *                 .name(&#34;primary&#34;)
 *                 .build())
 *             .baseInstanceName(&#34;vm&#34;)
 *             .targetSize(2)
 *             .build());
 * 
 *         final var catalog-workload = ApphubFunctions.getDiscoveredWorkload(GetDiscoveredWorkloadArgs.builder()
 *             .location(&#34;us-central1&#34;)
 *             .workloadUri(StdFunctions.replace().applyValue(invoke -&gt; invoke.result()))
 *             .build());
 * 
 *         var wait120sForResourceIngestion = new Sleep(&#34;wait120sForResourceIngestion&#34;, SleepArgs.builder()        
 *             .createDuration(&#34;120s&#34;)
 *             .build());
 * 
 *         var example = new Workload(&#34;example&#34;, WorkloadArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .applicationId(application.applicationId())
 *             .workloadId(mig.name())
 *             .discoveredWorkload(catalog_workload.applyValue(catalog_workload -&gt; catalog_workload.name()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Apphub Workload Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apphub.Application;
 * import com.pulumi.gcp.apphub.ApplicationArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationScopeArgs;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.time.sleep;
 * import com.pulumi.time.SleepArgs;
 * import com.pulumi.gcp.apphub.ServiceProjectAttachment;
 * import com.pulumi.gcp.apphub.ServiceProjectAttachmentArgs;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
 * import com.pulumi.gcp.compute.RegionInstanceGroupManager;
 * import com.pulumi.gcp.compute.RegionInstanceGroupManagerArgs;
 * import com.pulumi.gcp.compute.inputs.RegionInstanceGroupManagerVersionArgs;
 * import com.pulumi.gcp.apphub.ApphubFunctions;
 * import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
 * import com.pulumi.gcp.apphub.Workload;
 * import com.pulumi.gcp.apphub.WorkloadArgs;
 * import com.pulumi.gcp.apphub.inputs.WorkloadAttributesArgs;
 * import com.pulumi.gcp.apphub.inputs.WorkloadAttributesEnvironmentArgs;
 * import com.pulumi.gcp.apphub.inputs.WorkloadAttributesCriticalityArgs;
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
 *         var application = new Application(&#34;application&#34;, ApplicationArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .applicationId(&#34;example-application-1&#34;)
 *             .scope(ApplicationScopeArgs.builder()
 *                 .type(&#34;REGIONAL&#34;)
 *                 .build())
 *             .build());
 * 
 *         var serviceProject = new Project(&#34;serviceProject&#34;, ProjectArgs.builder()        
 *             .projectId(&#34;project-1&#34;)
 *             .name(&#34;Service Project&#34;)
 *             .orgId(&#34;123456789&#34;)
 *             .billingAccount(&#34;000000-0000000-0000000-000000&#34;)
 *             .build());
 * 
 *         var computeServiceProject = new Service(&#34;computeServiceProject&#34;, ServiceArgs.builder()        
 *             .project(serviceProject.projectId())
 *             .service(&#34;compute.googleapis.com&#34;)
 *             .build());
 * 
 *         var wait120s = new Sleep(&#34;wait120s&#34;, SleepArgs.builder()        
 *             .createDuration(&#34;120s&#34;)
 *             .build());
 * 
 *         var serviceProjectAttachment = new ServiceProjectAttachment(&#34;serviceProjectAttachment&#34;, ServiceProjectAttachmentArgs.builder()        
 *             .serviceProjectAttachmentId(serviceProject.projectId())
 *             .build());
 * 
 *         var ilbNetwork = new Network(&#34;ilbNetwork&#34;, NetworkArgs.builder()        
 *             .name(&#34;l7-ilb-network&#34;)
 *             .project(serviceProject.projectId())
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var ilbSubnet = new Subnetwork(&#34;ilbSubnet&#34;, SubnetworkArgs.builder()        
 *             .name(&#34;l7-ilb-subnet&#34;)
 *             .project(serviceProject.projectId())
 *             .ipCidrRange(&#34;10.0.1.0/24&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .network(ilbNetwork.id())
 *             .build());
 * 
 *         var instanceTemplate = new InstanceTemplate(&#34;instanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .accessConfigs()
 *                 .network(ilbNetwork.id())
 *                 .subnetwork(ilbSubnet.id())
 *                 .build())
 *             .name(&#34;l7-ilb-mig-template&#34;)
 *             .project(serviceProject.projectId())
 *             .machineType(&#34;e2-small&#34;)
 *             .tags(&#34;http-server&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(&#34;debian-cloud/debian-10&#34;)
 *                 .autoDelete(true)
 *                 .boot(true)
 *                 .build())
 *             .metadata(Map.of(&#34;startup-script&#34;, &#34;&#34;&#34;
 * #! /bin/bash
 * set -euo pipefail
 * export DEBIAN_FRONTEND=noninteractive
 * apt-get update
 * apt-get install -y nginx-light jq
 * NAME=$(curl -H &#34;Metadata-Flavor: Google&#34; &#34;http://metadata.google.internal/computeMetadata/v1/instance/hostname&#34;)
 * IP=$(curl -H &#34;Metadata-Flavor: Google&#34; &#34;http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip&#34;)
 * METADATA=$(curl -f -H &#34;Metadata-Flavor: Google&#34; &#34;http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=True&#34; | jq &#39;del(.[&#34;startup-script&#34;])&#39;)
 * cat &lt;&lt;EOF &gt; /var/www/html/index.html
 * &lt;pre&gt;
 * Name: $NAME
 * IP: $IP
 * Metadata: $METADATA
 * &lt;/pre&gt;
 * EOF
 *             &#34;&#34;&#34;))
 *             .build());
 * 
 *         var mig = new RegionInstanceGroupManager(&#34;mig&#34;, RegionInstanceGroupManagerArgs.builder()        
 *             .name(&#34;l7-ilb-mig1&#34;)
 *             .project(serviceProject.projectId())
 *             .region(&#34;us-central1&#34;)
 *             .versions(RegionInstanceGroupManagerVersionArgs.builder()
 *                 .instanceTemplate(instanceTemplate.id())
 *                 .name(&#34;primary&#34;)
 *                 .build())
 *             .baseInstanceName(&#34;vm&#34;)
 *             .targetSize(2)
 *             .build());
 * 
 *         final var catalog-workload = ApphubFunctions.getDiscoveredWorkload(GetDiscoveredWorkloadArgs.builder()
 *             .location(&#34;us-central1&#34;)
 *             .workloadUri(StdFunctions.replace().applyValue(invoke -&gt; invoke.result()))
 *             .build());
 * 
 *         var wait120sForResourceIngestion = new Sleep(&#34;wait120sForResourceIngestion&#34;, SleepArgs.builder()        
 *             .createDuration(&#34;120s&#34;)
 *             .build());
 * 
 *         var example = new Workload(&#34;example&#34;, WorkloadArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .applicationId(application.applicationId())
 *             .workloadId(mig.name())
 *             .discoveredWorkload(catalog_workload.applyValue(catalog_workload -&gt; catalog_workload.name()))
 *             .displayName(&#34;Example Service Full&#34;)
 *             .description(&#34;Register service for testing&#34;)
 *             .attributes(WorkloadAttributesArgs.builder()
 *                 .environment(WorkloadAttributesEnvironmentArgs.builder()
 *                     .type(&#34;STAGING&#34;)
 *                     .build())
 *                 .criticality(WorkloadAttributesCriticalityArgs.builder()
 *                     .type(&#34;MISSION_CRITICAL&#34;)
 *                     .build())
 *                 .businessOwners(WorkloadAttributesBusinessOwnerArgs.builder()
 *                     .displayName(&#34;Alice&#34;)
 *                     .email(&#34;alice@google.com&#34;)
 *                     .build())
 *                 .developerOwners(WorkloadAttributesDeveloperOwnerArgs.builder()
 *                     .displayName(&#34;Bob&#34;)
 *                     .email(&#34;bob@google.com&#34;)
 *                     .build())
 *                 .operatorOwners(WorkloadAttributesOperatorOwnerArgs.builder()
 *                     .displayName(&#34;Charlie&#34;)
 *                     .email(&#34;charlie@google.com&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Workload can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/applications/{{application_id}}/workloads/{{workload_id}}`
 * 
 * * `{{project}}/{{location}}/{{application_id}}/{{workload_id}}`
 * 
 * * `{{location}}/{{application_id}}/{{workload_id}}`
 * 
 * When using the `pulumi import` command, Workload can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:apphub/workload:Workload default projects/{{project}}/locations/{{location}}/applications/{{application_id}}/workloads/{{workload_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apphub/workload:Workload default {{project}}/{{location}}/{{application_id}}/{{workload_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apphub/workload:Workload default {{location}}/{{application_id}}/{{workload_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:apphub/workload:Workload")
public class Workload extends com.pulumi.resources.CustomResource {
    /**
     * Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * Consumer provided attributes.
     * Structure is documented below.
     * 
     */
    @Export(name="attributes", refs={WorkloadAttributes.class}, tree="[0]")
    private Output</* @Nullable */ WorkloadAttributes> attributes;

    /**
     * @return Consumer provided attributes.
     * Structure is documented below.
     * 
     */
    public Output<Optional<WorkloadAttributes>> attributes() {
        return Codegen.optional(this.attributes);
    }
    /**
     * Output only. Create time.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Create time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * User-defined description of a Workload.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User-defined description of a Workload.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Immutable. The resource name of the original discovered workload.
     * 
     */
    @Export(name="discoveredWorkload", refs={String.class}, tree="[0]")
    private Output<String> discoveredWorkload;

    /**
     * @return Immutable. The resource name of the original discovered workload.
     * 
     */
    public Output<String> discoveredWorkload() {
        return this.discoveredWorkload;
    }
    /**
     * User-defined name for the Workload.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User-defined name for the Workload.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Identifier. The resource name of the Workload. Format:&#34;projects/{host-project-id}/locations/{location}/applications/{application-id}/workloads/{workload-id}&#34;
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Identifier. The resource name of the Workload. Format:&#34;projects/{host-project-id}/locations/{location}/applications/{application-id}/workloads/{workload-id}&#34;
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
    /**
     * Output only. Workload state. Possible values:  STATE_UNSPECIFIED CREATING ACTIVE DELETING DETACHED
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. Workload state. Possible values:  STATE_UNSPECIFIED CREATING ACTIVE DELETING DETACHED
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Output only. A universally unique identifier (UUID) for the `Workload` in the UUID4 format.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. A universally unique identifier (UUID) for the `Workload` in the UUID4 format.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. Update time.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. Update time.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * The Workload identifier.
     * 
     * ***
     * 
     */
    @Export(name="workloadId", refs={String.class}, tree="[0]")
    private Output<String> workloadId;

    /**
     * @return The Workload identifier.
     * 
     * ***
     * 
     */
    public Output<String> workloadId() {
        return this.workloadId;
    }
    /**
     * Properties of an underlying compute resource represented by the Workload.
     * Structure is documented below.
     * 
     */
    @Export(name="workloadProperties", refs={List.class,WorkloadWorkloadProperty.class}, tree="[0,1]")
    private Output<List<WorkloadWorkloadProperty>> workloadProperties;

    /**
     * @return Properties of an underlying compute resource represented by the Workload.
     * Structure is documented below.
     * 
     */
    public Output<List<WorkloadWorkloadProperty>> workloadProperties() {
        return this.workloadProperties;
    }
    /**
     * Reference of an underlying compute resource represented by the Workload.
     * Structure is documented below.
     * 
     */
    @Export(name="workloadReferences", refs={List.class,WorkloadWorkloadReference.class}, tree="[0,1]")
    private Output<List<WorkloadWorkloadReference>> workloadReferences;

    /**
     * @return Reference of an underlying compute resource represented by the Workload.
     * Structure is documented below.
     * 
     */
    public Output<List<WorkloadWorkloadReference>> workloadReferences() {
        return this.workloadReferences;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Workload(String name) {
        this(name, WorkloadArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Workload(String name, WorkloadArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Workload(String name, WorkloadArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apphub/workload:Workload", name, args == null ? WorkloadArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Workload(String name, Output<String> id, @Nullable WorkloadState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apphub/workload:Workload", name, state, makeResourceOptions(options, id));
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
    public static Workload get(String name, Output<String> id, @Nullable WorkloadState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Workload(name, id, state, options);
    }
}
