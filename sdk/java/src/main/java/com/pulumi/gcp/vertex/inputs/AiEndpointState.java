// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.vertex.inputs.AiEndpointDeployedModelArgs;
import com.pulumi.gcp.vertex.inputs.AiEndpointEncryptionSpecArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiEndpointState extends com.pulumi.resources.ResourceArgs {

    public static final AiEndpointState Empty = new AiEndpointState();

    /**
     * (Output)
     * Output only. Timestamp when the DeployedModel was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return (Output)
     * Output only. Timestamp when the DeployedModel was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
     * Structure is documented below.
     * 
     */
    @Import(name="deployedModels")
    private @Nullable Output<List<AiEndpointDeployedModelArgs>> deployedModels;

    /**
     * @return Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<AiEndpointDeployedModelArgs>>> deployedModels() {
        return Optional.ofNullable(this.deployedModels);
    }

    /**
     * The description of the Endpoint.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the Endpoint.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,String>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
     * Structure is documented below.
     * 
     */
    @Import(name="encryptionSpec")
    private @Nullable Output<AiEndpointEncryptionSpecArgs> encryptionSpec;

    /**
     * @return Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AiEndpointEncryptionSpecArgs>> encryptionSpec() {
        return Optional.ofNullable(this.encryptionSpec);
    }

    /**
     * Used to perform consistent read-modify-write updates. If not set, a blind &#34;overwrite&#34; update happens.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return Used to perform consistent read-modify-write updates. If not set, a blind &#34;overwrite&#34; update happens.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location for the resource
     * 
     * ***
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location for the resource
     * 
     * ***
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
     * 
     */
    @Import(name="modelDeploymentMonitoringJob")
    private @Nullable Output<String> modelDeploymentMonitoringJob;

    /**
     * @return Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
     * 
     */
    public Optional<Output<String>> modelDeploymentMonitoringJob() {
        return Optional.ofNullable(this.modelDeploymentMonitoringJob);
    }

    /**
     * The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,String>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * The region for the resource
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region for the resource
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Output only. Timestamp when this Endpoint was last updated.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Output only. Timestamp when this Endpoint was last updated.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private AiEndpointState() {}

    private AiEndpointState(AiEndpointState $) {
        this.createTime = $.createTime;
        this.deployedModels = $.deployedModels;
        this.description = $.description;
        this.displayName = $.displayName;
        this.effectiveLabels = $.effectiveLabels;
        this.encryptionSpec = $.encryptionSpec;
        this.etag = $.etag;
        this.labels = $.labels;
        this.location = $.location;
        this.modelDeploymentMonitoringJob = $.modelDeploymentMonitoringJob;
        this.name = $.name;
        this.network = $.network;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.region = $.region;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiEndpointState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiEndpointState $;

        public Builder() {
            $ = new AiEndpointState();
        }

        public Builder(AiEndpointState defaults) {
            $ = new AiEndpointState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime (Output)
         * Output only. Timestamp when the DeployedModel was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime (Output)
         * Output only. Timestamp when the DeployedModel was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param deployedModels Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder deployedModels(@Nullable Output<List<AiEndpointDeployedModelArgs>> deployedModels) {
            $.deployedModels = deployedModels;
            return this;
        }

        /**
         * @param deployedModels Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder deployedModels(List<AiEndpointDeployedModelArgs> deployedModels) {
            return deployedModels(Output.of(deployedModels));
        }

        /**
         * @param deployedModels Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder deployedModels(AiEndpointDeployedModelArgs... deployedModels) {
            return deployedModels(List.of(deployedModels));
        }

        /**
         * @param description The description of the Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,String>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param encryptionSpec Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionSpec(@Nullable Output<AiEndpointEncryptionSpecArgs> encryptionSpec) {
            $.encryptionSpec = encryptionSpec;
            return this;
        }

        /**
         * @param encryptionSpec Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionSpec(AiEndpointEncryptionSpecArgs encryptionSpec) {
            return encryptionSpec(Output.of(encryptionSpec));
        }

        /**
         * @param etag Used to perform consistent read-modify-write updates. If not set, a blind &#34;overwrite&#34; update happens.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag Used to perform consistent read-modify-write updates. If not set, a blind &#34;overwrite&#34; update happens.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param labels The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The location for the resource
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location for the resource
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param modelDeploymentMonitoringJob Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
         * 
         * @return builder
         * 
         */
        public Builder modelDeploymentMonitoringJob(@Nullable Output<String> modelDeploymentMonitoringJob) {
            $.modelDeploymentMonitoringJob = modelDeploymentMonitoringJob;
            return this;
        }

        /**
         * @param modelDeploymentMonitoringJob Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
         * 
         * @return builder
         * 
         */
        public Builder modelDeploymentMonitoringJob(String modelDeploymentMonitoringJob) {
            return modelDeploymentMonitoringJob(Output.of(modelDeploymentMonitoringJob));
        }

        /**
         * @param name The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param network The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,String>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param region The region for the resource
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region for the resource
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param updateTime Output only. Timestamp when this Endpoint was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Output only. Timestamp when this Endpoint was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public AiEndpointState build() {
            return $;
        }
    }

}
