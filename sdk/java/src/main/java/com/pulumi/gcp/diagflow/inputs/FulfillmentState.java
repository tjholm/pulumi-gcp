// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.diagflow.inputs.FulfillmentFeatureArgs;
import com.pulumi.gcp.diagflow.inputs.FulfillmentGenericWebServiceArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FulfillmentState extends com.pulumi.resources.ResourceArgs {

    public static final FulfillmentState Empty = new FulfillmentState();

    /**
     * The human-readable name of the fulfillment, unique within the agent.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The human-readable name of the fulfillment, unique within the agent.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Whether fulfillment is enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether fulfillment is enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     * 
     */
    @Import(name="features")
    private @Nullable Output<List<FulfillmentFeatureArgs>> features;

    /**
     * @return The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<FulfillmentFeatureArgs>>> features() {
        return Optional.ofNullable(this.features);
    }

    /**
     * Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     * 
     */
    @Import(name="genericWebService")
    private @Nullable Output<FulfillmentGenericWebServiceArgs> genericWebService;

    /**
     * @return Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     * 
     */
    public Optional<Output<FulfillmentGenericWebServiceArgs>> genericWebService() {
        return Optional.ofNullable(this.genericWebService);
    }

    /**
     * The unique identifier of the fulfillment. Format: projects/&lt;Project ID&gt;/agent/fulfillment - projects/&lt;Project
     * ID&gt;/locations/&lt;Location ID&gt;/agent/fulfillment
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The unique identifier of the fulfillment. Format: projects/&lt;Project ID&gt;/agent/fulfillment - projects/&lt;Project
     * ID&gt;/locations/&lt;Location ID&gt;/agent/fulfillment
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

    private FulfillmentState() {}

    private FulfillmentState(FulfillmentState $) {
        this.displayName = $.displayName;
        this.enabled = $.enabled;
        this.features = $.features;
        this.genericWebService = $.genericWebService;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FulfillmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FulfillmentState $;

        public Builder() {
            $ = new FulfillmentState();
        }

        public Builder(FulfillmentState defaults) {
            $ = new FulfillmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName The human-readable name of the fulfillment, unique within the agent.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The human-readable name of the fulfillment, unique within the agent.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param enabled Whether fulfillment is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether fulfillment is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param features The field defines whether the fulfillment is enabled for certain features.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder features(@Nullable Output<List<FulfillmentFeatureArgs>> features) {
            $.features = features;
            return this;
        }

        /**
         * @param features The field defines whether the fulfillment is enabled for certain features.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder features(List<FulfillmentFeatureArgs> features) {
            return features(Output.of(features));
        }

        /**
         * @param features The field defines whether the fulfillment is enabled for certain features.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder features(FulfillmentFeatureArgs... features) {
            return features(List.of(features));
        }

        /**
         * @param genericWebService Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder genericWebService(@Nullable Output<FulfillmentGenericWebServiceArgs> genericWebService) {
            $.genericWebService = genericWebService;
            return this;
        }

        /**
         * @param genericWebService Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder genericWebService(FulfillmentGenericWebServiceArgs genericWebService) {
            return genericWebService(Output.of(genericWebService));
        }

        /**
         * @param name The unique identifier of the fulfillment. Format: projects/&lt;Project ID&gt;/agent/fulfillment - projects/&lt;Project
         * ID&gt;/locations/&lt;Location ID&gt;/agent/fulfillment
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The unique identifier of the fulfillment. Format: projects/&lt;Project ID&gt;/agent/fulfillment - projects/&lt;Project
         * ID&gt;/locations/&lt;Location ID&gt;/agent/fulfillment
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

        public FulfillmentState build() {
            return $;
        }
    }

}
