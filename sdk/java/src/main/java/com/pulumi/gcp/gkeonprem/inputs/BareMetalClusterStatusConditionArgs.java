// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BareMetalClusterStatusConditionArgs extends com.pulumi.resources.ResourceArgs {

    public static final BareMetalClusterStatusConditionArgs Empty = new BareMetalClusterStatusConditionArgs();

    /**
     * (Output)
     * Last time the condition transit from one status to another.
     * 
     */
    @Import(name="lastTransitionTime")
    private @Nullable Output<String> lastTransitionTime;

    /**
     * @return (Output)
     * Last time the condition transit from one status to another.
     * 
     */
    public Optional<Output<String>> lastTransitionTime() {
        return Optional.ofNullable(this.lastTransitionTime);
    }

    /**
     * Human-readable message indicating details about last transition.
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return Human-readable message indicating details about last transition.
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * (Output)
     * A human-readable message of the check failure.
     * 
     */
    @Import(name="reason")
    private @Nullable Output<String> reason;

    /**
     * @return (Output)
     * A human-readable message of the check failure.
     * 
     */
    public Optional<Output<String>> reason() {
        return Optional.ofNullable(this.reason);
    }

    /**
     * (Output)
     * The lifecycle state of the condition.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return (Output)
     * The lifecycle state of the condition.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * Type of the condition.
     * (e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of the condition.
     * (e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private BareMetalClusterStatusConditionArgs() {}

    private BareMetalClusterStatusConditionArgs(BareMetalClusterStatusConditionArgs $) {
        this.lastTransitionTime = $.lastTransitionTime;
        this.message = $.message;
        this.reason = $.reason;
        this.state = $.state;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BareMetalClusterStatusConditionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BareMetalClusterStatusConditionArgs $;

        public Builder() {
            $ = new BareMetalClusterStatusConditionArgs();
        }

        public Builder(BareMetalClusterStatusConditionArgs defaults) {
            $ = new BareMetalClusterStatusConditionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param lastTransitionTime (Output)
         * Last time the condition transit from one status to another.
         * 
         * @return builder
         * 
         */
        public Builder lastTransitionTime(@Nullable Output<String> lastTransitionTime) {
            $.lastTransitionTime = lastTransitionTime;
            return this;
        }

        /**
         * @param lastTransitionTime (Output)
         * Last time the condition transit from one status to another.
         * 
         * @return builder
         * 
         */
        public Builder lastTransitionTime(String lastTransitionTime) {
            return lastTransitionTime(Output.of(lastTransitionTime));
        }

        /**
         * @param message Human-readable message indicating details about last transition.
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message Human-readable message indicating details about last transition.
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param reason (Output)
         * A human-readable message of the check failure.
         * 
         * @return builder
         * 
         */
        public Builder reason(@Nullable Output<String> reason) {
            $.reason = reason;
            return this;
        }

        /**
         * @param reason (Output)
         * A human-readable message of the check failure.
         * 
         * @return builder
         * 
         */
        public Builder reason(String reason) {
            return reason(Output.of(reason));
        }

        /**
         * @param state (Output)
         * The lifecycle state of the condition.
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
         * The lifecycle state of the condition.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param type Type of the condition.
         * (e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the condition.
         * (e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public BareMetalClusterStatusConditionArgs build() {
            return $;
        }
    }

}
