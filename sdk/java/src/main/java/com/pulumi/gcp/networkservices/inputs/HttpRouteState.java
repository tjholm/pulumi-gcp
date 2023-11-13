// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkservices.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networkservices.inputs.HttpRouteRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HttpRouteState extends com.pulumi.resources.ResourceArgs {

    public static final HttpRouteState Empty = new HttpRouteState();

    /**
     * Time the HttpRoute was created in UTC.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Time the HttpRoute was created in UTC.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * A free-text description of the resource. Max length 1024 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A free-text description of the resource. Max length 1024 characters.
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
     * Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
     * Each gateway reference should match the pattern: projects/*{@literal /}locations/global/gateways/&lt;gateway_name&gt;
     * 
     */
    @Import(name="gateways")
    private @Nullable Output<List<String>> gateways;

    /**
     * @return Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
     * Each gateway reference should match the pattern: projects/*{@literal /}locations/global/gateways/&lt;gateway_name&gt;
     * 
     */
    public Optional<Output<List<String>>> gateways() {
        return Optional.ofNullable(this.gateways);
    }

    /**
     * Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
     * 
     */
    @Import(name="hostnames")
    private @Nullable Output<List<String>> hostnames;

    /**
     * @return Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
     * 
     */
    public Optional<Output<List<String>>> hostnames() {
        return Optional.ofNullable(this.hostnames);
    }

    /**
     * Set of label tags associated with the HttpRoute resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the HttpRoute resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
     * Each mesh reference should match the pattern: projects/*{@literal /}locations/global/meshes/&lt;mesh_name&gt;.
     * The attached Mesh should be of a type SIDECAR.
     * 
     */
    @Import(name="meshes")
    private @Nullable Output<List<String>> meshes;

    /**
     * @return Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
     * Each mesh reference should match the pattern: projects/*{@literal /}locations/global/meshes/&lt;mesh_name&gt;.
     * The attached Mesh should be of a type SIDECAR.
     * 
     */
    public Optional<Output<List<String>>> meshes() {
        return Optional.ofNullable(this.meshes);
    }

    /**
     * Name of the HttpRoute resource.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the HttpRoute resource.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * Rules that define how traffic is routed and handled.
     * Structure is documented below.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<HttpRouteRuleArgs>> rules;

    /**
     * @return Rules that define how traffic is routed and handled.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<HttpRouteRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * Server-defined URL of this resource.
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return Server-defined URL of this resource.
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * Time the HttpRoute was updated in UTC.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Time the HttpRoute was updated in UTC.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private HttpRouteState() {}

    private HttpRouteState(HttpRouteState $) {
        this.createTime = $.createTime;
        this.description = $.description;
        this.effectiveLabels = $.effectiveLabels;
        this.gateways = $.gateways;
        this.hostnames = $.hostnames;
        this.labels = $.labels;
        this.meshes = $.meshes;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.rules = $.rules;
        this.selfLink = $.selfLink;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HttpRouteState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HttpRouteState $;

        public Builder() {
            $ = new HttpRouteState();
        }

        public Builder(HttpRouteState defaults) {
            $ = new HttpRouteState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime Time the HttpRoute was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Time the HttpRoute was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
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
         * @param gateways Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
         * Each gateway reference should match the pattern: projects/*{@literal /}locations/global/gateways/&lt;gateway_name&gt;
         * 
         * @return builder
         * 
         */
        public Builder gateways(@Nullable Output<List<String>> gateways) {
            $.gateways = gateways;
            return this;
        }

        /**
         * @param gateways Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
         * Each gateway reference should match the pattern: projects/*{@literal /}locations/global/gateways/&lt;gateway_name&gt;
         * 
         * @return builder
         * 
         */
        public Builder gateways(List<String> gateways) {
            return gateways(Output.of(gateways));
        }

        /**
         * @param gateways Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
         * Each gateway reference should match the pattern: projects/*{@literal /}locations/global/gateways/&lt;gateway_name&gt;
         * 
         * @return builder
         * 
         */
        public Builder gateways(String... gateways) {
            return gateways(List.of(gateways));
        }

        /**
         * @param hostnames Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
         * 
         * @return builder
         * 
         */
        public Builder hostnames(@Nullable Output<List<String>> hostnames) {
            $.hostnames = hostnames;
            return this;
        }

        /**
         * @param hostnames Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
         * 
         * @return builder
         * 
         */
        public Builder hostnames(List<String> hostnames) {
            return hostnames(Output.of(hostnames));
        }

        /**
         * @param hostnames Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
         * 
         * @return builder
         * 
         */
        public Builder hostnames(String... hostnames) {
            return hostnames(List.of(hostnames));
        }

        /**
         * @param labels Set of label tags associated with the HttpRoute resource.
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
         * @param labels Set of label tags associated with the HttpRoute resource.
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
         * @param meshes Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
         * Each mesh reference should match the pattern: projects/*{@literal /}locations/global/meshes/&lt;mesh_name&gt;.
         * The attached Mesh should be of a type SIDECAR.
         * 
         * @return builder
         * 
         */
        public Builder meshes(@Nullable Output<List<String>> meshes) {
            $.meshes = meshes;
            return this;
        }

        /**
         * @param meshes Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
         * Each mesh reference should match the pattern: projects/*{@literal /}locations/global/meshes/&lt;mesh_name&gt;.
         * The attached Mesh should be of a type SIDECAR.
         * 
         * @return builder
         * 
         */
        public Builder meshes(List<String> meshes) {
            return meshes(Output.of(meshes));
        }

        /**
         * @param meshes Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
         * Each mesh reference should match the pattern: projects/*{@literal /}locations/global/meshes/&lt;mesh_name&gt;.
         * The attached Mesh should be of a type SIDECAR.
         * 
         * @return builder
         * 
         */
        public Builder meshes(String... meshes) {
            return meshes(List.of(meshes));
        }

        /**
         * @param name Name of the HttpRoute resource.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the HttpRoute resource.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param rules Rules that define how traffic is routed and handled.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<HttpRouteRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules Rules that define how traffic is routed and handled.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<HttpRouteRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules Rules that define how traffic is routed and handled.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(HttpRouteRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param selfLink Server-defined URL of this resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink Server-defined URL of this resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param updateTime Time the HttpRoute was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Time the HttpRoute was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public HttpRouteState build() {
            return $;
        }
    }

}
