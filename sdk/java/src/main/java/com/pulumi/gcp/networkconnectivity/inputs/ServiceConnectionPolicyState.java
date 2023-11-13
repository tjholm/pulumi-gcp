// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkconnectivity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networkconnectivity.inputs.ServiceConnectionPolicyPscConfigArgs;
import com.pulumi.gcp.networkconnectivity.inputs.ServiceConnectionPolicyPscConnectionArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceConnectionPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceConnectionPolicyState Empty = new ServiceConnectionPolicyState();

    /**
     * The timestamp when the resource was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The timestamp when the resource was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Free-text description of the resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Free-text description of the resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
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
     * The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The type of underlying resources used to create the connection.
     * 
     */
    @Import(name="infrastructure")
    private @Nullable Output<String> infrastructure;

    /**
     * @return The type of underlying resources used to create the connection.
     * 
     */
    public Optional<Output<String>> infrastructure() {
        return Optional.ofNullable(this.infrastructure);
    }

    /**
     * User-defined labels.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return User-defined labels.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location of the ServiceConnectionPolicy.
     * 
     * ***
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of the ServiceConnectionPolicy.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
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
     * Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
     * Structure is documented below.
     * 
     */
    @Import(name="pscConfig")
    private @Nullable Output<ServiceConnectionPolicyPscConfigArgs> pscConfig;

    /**
     * @return Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ServiceConnectionPolicyPscConfigArgs>> pscConfig() {
        return Optional.ofNullable(this.pscConfig);
    }

    /**
     * Information about each Private Service Connect connection.
     * Structure is documented below.
     * 
     */
    @Import(name="pscConnections")
    private @Nullable Output<List<ServiceConnectionPolicyPscConnectionArgs>> pscConnections;

    /**
     * @return Information about each Private Service Connect connection.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<ServiceConnectionPolicyPscConnectionArgs>>> pscConnections() {
        return Optional.ofNullable(this.pscConnections);
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
     * The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.
     * It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
     * 
     */
    @Import(name="serviceClass")
    private @Nullable Output<String> serviceClass;

    /**
     * @return The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.
     * It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
     * 
     */
    public Optional<Output<String>> serviceClass() {
        return Optional.ofNullable(this.serviceClass);
    }

    /**
     * The timestamp when the resource was updated.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The timestamp when the resource was updated.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private ServiceConnectionPolicyState() {}

    private ServiceConnectionPolicyState(ServiceConnectionPolicyState $) {
        this.createTime = $.createTime;
        this.description = $.description;
        this.effectiveLabels = $.effectiveLabels;
        this.etag = $.etag;
        this.infrastructure = $.infrastructure;
        this.labels = $.labels;
        this.location = $.location;
        this.name = $.name;
        this.network = $.network;
        this.project = $.project;
        this.pscConfig = $.pscConfig;
        this.pscConnections = $.pscConnections;
        this.pulumiLabels = $.pulumiLabels;
        this.serviceClass = $.serviceClass;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceConnectionPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceConnectionPolicyState $;

        public Builder() {
            $ = new ServiceConnectionPolicyState();
        }

        public Builder(ServiceConnectionPolicyState defaults) {
            $ = new ServiceConnectionPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime The timestamp when the resource was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The timestamp when the resource was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description Free-text description of the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Free-text description of the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param etag The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param infrastructure The type of underlying resources used to create the connection.
         * 
         * @return builder
         * 
         */
        public Builder infrastructure(@Nullable Output<String> infrastructure) {
            $.infrastructure = infrastructure;
            return this;
        }

        /**
         * @param infrastructure The type of underlying resources used to create the connection.
         * 
         * @return builder
         * 
         */
        public Builder infrastructure(String infrastructure) {
            return infrastructure(Output.of(infrastructure));
        }

        /**
         * @param labels User-defined labels.
         * 
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
         * @param labels User-defined labels.
         * 
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
         * @param location The location of the ServiceConnectionPolicy.
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
         * @param location The location of the ServiceConnectionPolicy.
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
         * @param name The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param network The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
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
         * @param pscConfig Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConfig(@Nullable Output<ServiceConnectionPolicyPscConfigArgs> pscConfig) {
            $.pscConfig = pscConfig;
            return this;
        }

        /**
         * @param pscConfig Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConfig(ServiceConnectionPolicyPscConfigArgs pscConfig) {
            return pscConfig(Output.of(pscConfig));
        }

        /**
         * @param pscConnections Information about each Private Service Connect connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConnections(@Nullable Output<List<ServiceConnectionPolicyPscConnectionArgs>> pscConnections) {
            $.pscConnections = pscConnections;
            return this;
        }

        /**
         * @param pscConnections Information about each Private Service Connect connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConnections(List<ServiceConnectionPolicyPscConnectionArgs> pscConnections) {
            return pscConnections(Output.of(pscConnections));
        }

        /**
         * @param pscConnections Information about each Private Service Connect connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConnections(ServiceConnectionPolicyPscConnectionArgs... pscConnections) {
            return pscConnections(List.of(pscConnections));
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
         * @param serviceClass The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.
         * It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
         * 
         * @return builder
         * 
         */
        public Builder serviceClass(@Nullable Output<String> serviceClass) {
            $.serviceClass = serviceClass;
            return this;
        }

        /**
         * @param serviceClass The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.
         * It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
         * 
         * @return builder
         * 
         */
        public Builder serviceClass(String serviceClass) {
            return serviceClass(Output.of(serviceClass));
        }

        /**
         * @param updateTime The timestamp when the resource was updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The timestamp when the resource was updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public ServiceConnectionPolicyState build() {
            return $;
        }
    }

}
