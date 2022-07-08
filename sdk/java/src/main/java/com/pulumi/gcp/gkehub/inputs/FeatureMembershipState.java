// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.gkehub.inputs.FeatureMembershipConfigmanagementArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureMembershipState extends com.pulumi.resources.ResourceArgs {

    public static final FeatureMembershipState Empty = new FeatureMembershipState();

    /**
     * Config Management-specific spec. Structure is documented below.
     * 
     */
    @Import(name="configmanagement")
    private @Nullable Output<FeatureMembershipConfigmanagementArgs> configmanagement;

    /**
     * @return Config Management-specific spec. Structure is documented below.
     * 
     */
    public Optional<Output<FeatureMembershipConfigmanagementArgs>> configmanagement() {
        return Optional.ofNullable(this.configmanagement);
    }

    /**
     * The name of the feature
     * 
     */
    @Import(name="feature")
    private @Nullable Output<String> feature;

    /**
     * @return The name of the feature
     * 
     */
    public Optional<Output<String>> feature() {
        return Optional.ofNullable(this.feature);
    }

    /**
     * The location of the feature
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of the feature
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The name of the membership
     * 
     */
    @Import(name="membership")
    private @Nullable Output<String> membership;

    /**
     * @return The name of the membership
     * 
     */
    public Optional<Output<String>> membership() {
        return Optional.ofNullable(this.membership);
    }

    /**
     * The project of the feature
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project of the feature
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private FeatureMembershipState() {}

    private FeatureMembershipState(FeatureMembershipState $) {
        this.configmanagement = $.configmanagement;
        this.feature = $.feature;
        this.location = $.location;
        this.membership = $.membership;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureMembershipState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureMembershipState $;

        public Builder() {
            $ = new FeatureMembershipState();
        }

        public Builder(FeatureMembershipState defaults) {
            $ = new FeatureMembershipState(Objects.requireNonNull(defaults));
        }

        /**
         * @param configmanagement Config Management-specific spec. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder configmanagement(@Nullable Output<FeatureMembershipConfigmanagementArgs> configmanagement) {
            $.configmanagement = configmanagement;
            return this;
        }

        /**
         * @param configmanagement Config Management-specific spec. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder configmanagement(FeatureMembershipConfigmanagementArgs configmanagement) {
            return configmanagement(Output.of(configmanagement));
        }

        /**
         * @param feature The name of the feature
         * 
         * @return builder
         * 
         */
        public Builder feature(@Nullable Output<String> feature) {
            $.feature = feature;
            return this;
        }

        /**
         * @param feature The name of the feature
         * 
         * @return builder
         * 
         */
        public Builder feature(String feature) {
            return feature(Output.of(feature));
        }

        /**
         * @param location The location of the feature
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location of the feature
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param membership The name of the membership
         * 
         * @return builder
         * 
         */
        public Builder membership(@Nullable Output<String> membership) {
            $.membership = membership;
            return this;
        }

        /**
         * @param membership The name of the membership
         * 
         * @return builder
         * 
         */
        public Builder membership(String membership) {
            return membership(Output.of(membership));
        }

        /**
         * @param project The project of the feature
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project of the feature
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public FeatureMembershipState build() {
            return $;
        }
    }

}
