// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureResourceStateArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeatureResourceStateArgs Empty = new FeatureResourceStateArgs();

    /**
     * (Output)
     * Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.
     * 
     */
    @Import(name="hasResources")
    private @Nullable Output<Boolean> hasResources;

    /**
     * @return (Output)
     * Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.
     * 
     */
    public Optional<Output<Boolean>> hasResources() {
        return Optional.ofNullable(this.hasResources);
    }

    /**
     * (Output)
     * Output only. The &#34;running state&#34; of the Feature in this Hub.
     * Structure is documented below.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return (Output)
     * Output only. The &#34;running state&#34; of the Feature in this Hub.
     * Structure is documented below.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private FeatureResourceStateArgs() {}

    private FeatureResourceStateArgs(FeatureResourceStateArgs $) {
        this.hasResources = $.hasResources;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureResourceStateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureResourceStateArgs $;

        public Builder() {
            $ = new FeatureResourceStateArgs();
        }

        public Builder(FeatureResourceStateArgs defaults) {
            $ = new FeatureResourceStateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hasResources (Output)
         * Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.
         * 
         * @return builder
         * 
         */
        public Builder hasResources(@Nullable Output<Boolean> hasResources) {
            $.hasResources = hasResources;
            return this;
        }

        /**
         * @param hasResources (Output)
         * Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.
         * 
         * @return builder
         * 
         */
        public Builder hasResources(Boolean hasResources) {
            return hasResources(Output.of(hasResources));
        }

        /**
         * @param state (Output)
         * Output only. The &#34;running state&#34; of the Feature in this Hub.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state (Output)
         * Output only. The &#34;running state&#34; of the Feature in this Hub.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public FeatureResourceStateArgs build() {
            return $;
        }
    }

}
