// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs extends com.pulumi.resources.ResourceArgs {

    public static final BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs Empty = new BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs();

    /**
     * Specifies the nodes operating system (default: LINUX).
     * Possible values are: `EFFECT_UNSPECIFIED`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.
     * 
     */
    @Import(name="effect")
    private @Nullable Output<String> effect;

    /**
     * @return Specifies the nodes operating system (default: LINUX).
     * Possible values are: `EFFECT_UNSPECIFIED`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.
     * 
     */
    public Optional<Output<String>> effect() {
        return Optional.ofNullable(this.effect);
    }

    /**
     * Key associated with the effect.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return Key associated with the effect.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Value associated with the effect.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return Value associated with the effect.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs() {}

    private BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs $) {
        this.effect = $.effect;
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs $;

        public Builder() {
            $ = new BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs();
        }

        public Builder(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs defaults) {
            $ = new BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param effect Specifies the nodes operating system (default: LINUX).
         * Possible values are: `EFFECT_UNSPECIFIED`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.
         * 
         * @return builder
         * 
         */
        public Builder effect(@Nullable Output<String> effect) {
            $.effect = effect;
            return this;
        }

        /**
         * @param effect Specifies the nodes operating system (default: LINUX).
         * Possible values are: `EFFECT_UNSPECIFIED`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.
         * 
         * @return builder
         * 
         */
        public Builder effect(String effect) {
            return effect(Output.of(effect));
        }

        /**
         * @param key Key associated with the effect.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Key associated with the effect.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value Value associated with the effect.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value associated with the effect.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs build() {
            return $;
        }
    }

}
