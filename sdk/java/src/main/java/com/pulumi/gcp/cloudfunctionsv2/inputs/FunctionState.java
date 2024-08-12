// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctionsv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigArgs;
import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionEventTriggerArgs;
import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionServiceConfigArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionState extends com.pulumi.resources.ResourceArgs {

    public static final FunctionState Empty = new FunctionState();

    /**
     * Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     * 
     */
    @Import(name="buildConfig")
    private @Nullable Output<FunctionBuildConfigArgs> buildConfig;

    /**
     * @return Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     * 
     */
    public Optional<Output<FunctionBuildConfigArgs>> buildConfig() {
        return Optional.ofNullable(this.buildConfig);
    }

    /**
     * User-provided description of a function.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return User-provided description of a function.
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
     * The environment the function is hosted on.
     * 
     */
    @Import(name="environment")
    private @Nullable Output<String> environment;

    /**
     * @return The environment the function is hosted on.
     * 
     */
    public Optional<Output<String>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     * 
     */
    @Import(name="eventTrigger")
    private @Nullable Output<FunctionEventTriggerArgs> eventTrigger;

    /**
     * @return An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     * 
     */
    public Optional<Output<FunctionEventTriggerArgs>> eventTrigger() {
        return Optional.ofNullable(this.eventTrigger);
    }

    /**
     * Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
     * It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
     * 
     */
    @Import(name="kmsKeyName")
    private @Nullable Output<String> kmsKeyName;

    /**
     * @return Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
     * It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
     * 
     */
    public Optional<Output<String>> kmsKeyName() {
        return Optional.ofNullable(this.kmsKeyName);
    }

    /**
     * A set of key/value label pairs associated with this Cloud Function.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs associated with this Cloud Function.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location of this cloud function.
     * 
     * ***
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of this cloud function.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
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
     * Describes the Service being deployed.
     * Structure is documented below.
     * 
     */
    @Import(name="serviceConfig")
    private @Nullable Output<FunctionServiceConfigArgs> serviceConfig;

    /**
     * @return Describes the Service being deployed.
     * Structure is documented below.
     * 
     */
    public Optional<Output<FunctionServiceConfigArgs>> serviceConfig() {
        return Optional.ofNullable(this.serviceConfig);
    }

    /**
     * Describes the current state of the function.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return Describes the current state of the function.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * The last update timestamp of a Cloud Function.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The last update timestamp of a Cloud Function.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    /**
     * Output only. The deployed url for the function.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Output only. The deployed url for the function.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private FunctionState() {}

    private FunctionState(FunctionState $) {
        this.buildConfig = $.buildConfig;
        this.description = $.description;
        this.effectiveLabels = $.effectiveLabels;
        this.environment = $.environment;
        this.eventTrigger = $.eventTrigger;
        this.kmsKeyName = $.kmsKeyName;
        this.labels = $.labels;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.serviceConfig = $.serviceConfig;
        this.state = $.state;
        this.updateTime = $.updateTime;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionState $;

        public Builder() {
            $ = new FunctionState();
        }

        public Builder(FunctionState defaults) {
            $ = new FunctionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param buildConfig Describes the Build step of the function that builds a container
         * from the given source.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder buildConfig(@Nullable Output<FunctionBuildConfigArgs> buildConfig) {
            $.buildConfig = buildConfig;
            return this;
        }

        /**
         * @param buildConfig Describes the Build step of the function that builds a container
         * from the given source.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder buildConfig(FunctionBuildConfigArgs buildConfig) {
            return buildConfig(Output.of(buildConfig));
        }

        /**
         * @param description User-provided description of a function.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description User-provided description of a function.
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
         * @param environment The environment the function is hosted on.
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment The environment the function is hosted on.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param eventTrigger An Eventarc trigger managed by Google Cloud Functions that fires events in
         * response to a condition in another service.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventTrigger(@Nullable Output<FunctionEventTriggerArgs> eventTrigger) {
            $.eventTrigger = eventTrigger;
            return this;
        }

        /**
         * @param eventTrigger An Eventarc trigger managed by Google Cloud Functions that fires events in
         * response to a condition in another service.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventTrigger(FunctionEventTriggerArgs eventTrigger) {
            return eventTrigger(Output.of(eventTrigger));
        }

        /**
         * @param kmsKeyName Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
         * It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyName(@Nullable Output<String> kmsKeyName) {
            $.kmsKeyName = kmsKeyName;
            return this;
        }

        /**
         * @param kmsKeyName Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
         * It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyName(String kmsKeyName) {
            return kmsKeyName(Output.of(kmsKeyName));
        }

        /**
         * @param labels A set of key/value label pairs associated with this Cloud Function.
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
         * @param labels A set of key/value label pairs associated with this Cloud Function.
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
         * @param location The location of this cloud function.
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
         * @param location The location of this cloud function.
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
         * @param name A user-defined name of the function. Function names must
         * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A user-defined name of the function. Function names must
         * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
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
         * @param serviceConfig Describes the Service being deployed.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serviceConfig(@Nullable Output<FunctionServiceConfigArgs> serviceConfig) {
            $.serviceConfig = serviceConfig;
            return this;
        }

        /**
         * @param serviceConfig Describes the Service being deployed.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serviceConfig(FunctionServiceConfigArgs serviceConfig) {
            return serviceConfig(Output.of(serviceConfig));
        }

        /**
         * @param state Describes the current state of the function.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Describes the current state of the function.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param updateTime The last update timestamp of a Cloud Function.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The last update timestamp of a Cloud Function.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        /**
         * @param url Output only. The deployed url for the function.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Output only. The deployed url for the function.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public FunctionState build() {
            return $;
        }
    }

}
