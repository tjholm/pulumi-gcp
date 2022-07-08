// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctionsv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionEventTriggerArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionEventTriggerArgs Empty = new FunctionEventTriggerArgs();

    /**
     * Required. The type of event to observe.
     * 
     */
    @Import(name="eventType")
    private @Nullable Output<String> eventType;

    /**
     * @return Required. The type of event to observe.
     * 
     */
    public Optional<Output<String>> eventType() {
        return Optional.ofNullable(this.eventType);
    }

    /**
     * The name of a Pub/Sub topic in the same project that will be used
     * as the transport topic for the event delivery.
     * 
     */
    @Import(name="pubsubTopic")
    private @Nullable Output<String> pubsubTopic;

    /**
     * @return The name of a Pub/Sub topic in the same project that will be used
     * as the transport topic for the event delivery.
     * 
     */
    public Optional<Output<String>> pubsubTopic() {
        return Optional.ofNullable(this.pubsubTopic);
    }

    /**
     * Describes the retry policy in case of function&#39;s execution failure.
     * Retried execution is charged as any other execution.
     * Possible values are `RETRY_POLICY_UNSPECIFIED`, `RETRY_POLICY_DO_NOT_RETRY`, and `RETRY_POLICY_RETRY`.
     * 
     */
    @Import(name="retryPolicy")
    private @Nullable Output<String> retryPolicy;

    /**
     * @return Describes the retry policy in case of function&#39;s execution failure.
     * Retried execution is charged as any other execution.
     * Possible values are `RETRY_POLICY_UNSPECIFIED`, `RETRY_POLICY_DO_NOT_RETRY`, and `RETRY_POLICY_RETRY`.
     * 
     */
    public Optional<Output<String>> retryPolicy() {
        return Optional.ofNullable(this.retryPolicy);
    }

    /**
     * The email of the service account for this function.
     * 
     */
    @Import(name="serviceAccountEmail")
    private @Nullable Output<String> serviceAccountEmail;

    /**
     * @return The email of the service account for this function.
     * 
     */
    public Optional<Output<String>> serviceAccountEmail() {
        return Optional.ofNullable(this.serviceAccountEmail);
    }

    /**
     * - 
     * The resource name of the Eventarc trigger.
     * 
     */
    @Import(name="trigger")
    private @Nullable Output<String> trigger;

    /**
     * @return -
     * The resource name of the Eventarc trigger.
     * 
     */
    public Optional<Output<String>> trigger() {
        return Optional.ofNullable(this.trigger);
    }

    /**
     * The region that the trigger will be in. The trigger will only receive
     * events originating in this region. It can be the same
     * region as the function, a different region or multi-region, or the global
     * region. If not provided, defaults to the same region as the function.
     * 
     */
    @Import(name="triggerRegion")
    private @Nullable Output<String> triggerRegion;

    /**
     * @return The region that the trigger will be in. The trigger will only receive
     * events originating in this region. It can be the same
     * region as the function, a different region or multi-region, or the global
     * region. If not provided, defaults to the same region as the function.
     * 
     */
    public Optional<Output<String>> triggerRegion() {
        return Optional.ofNullable(this.triggerRegion);
    }

    private FunctionEventTriggerArgs() {}

    private FunctionEventTriggerArgs(FunctionEventTriggerArgs $) {
        this.eventType = $.eventType;
        this.pubsubTopic = $.pubsubTopic;
        this.retryPolicy = $.retryPolicy;
        this.serviceAccountEmail = $.serviceAccountEmail;
        this.trigger = $.trigger;
        this.triggerRegion = $.triggerRegion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionEventTriggerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionEventTriggerArgs $;

        public Builder() {
            $ = new FunctionEventTriggerArgs();
        }

        public Builder(FunctionEventTriggerArgs defaults) {
            $ = new FunctionEventTriggerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param eventType Required. The type of event to observe.
         * 
         * @return builder
         * 
         */
        public Builder eventType(@Nullable Output<String> eventType) {
            $.eventType = eventType;
            return this;
        }

        /**
         * @param eventType Required. The type of event to observe.
         * 
         * @return builder
         * 
         */
        public Builder eventType(String eventType) {
            return eventType(Output.of(eventType));
        }

        /**
         * @param pubsubTopic The name of a Pub/Sub topic in the same project that will be used
         * as the transport topic for the event delivery.
         * 
         * @return builder
         * 
         */
        public Builder pubsubTopic(@Nullable Output<String> pubsubTopic) {
            $.pubsubTopic = pubsubTopic;
            return this;
        }

        /**
         * @param pubsubTopic The name of a Pub/Sub topic in the same project that will be used
         * as the transport topic for the event delivery.
         * 
         * @return builder
         * 
         */
        public Builder pubsubTopic(String pubsubTopic) {
            return pubsubTopic(Output.of(pubsubTopic));
        }

        /**
         * @param retryPolicy Describes the retry policy in case of function&#39;s execution failure.
         * Retried execution is charged as any other execution.
         * Possible values are `RETRY_POLICY_UNSPECIFIED`, `RETRY_POLICY_DO_NOT_RETRY`, and `RETRY_POLICY_RETRY`.
         * 
         * @return builder
         * 
         */
        public Builder retryPolicy(@Nullable Output<String> retryPolicy) {
            $.retryPolicy = retryPolicy;
            return this;
        }

        /**
         * @param retryPolicy Describes the retry policy in case of function&#39;s execution failure.
         * Retried execution is charged as any other execution.
         * Possible values are `RETRY_POLICY_UNSPECIFIED`, `RETRY_POLICY_DO_NOT_RETRY`, and `RETRY_POLICY_RETRY`.
         * 
         * @return builder
         * 
         */
        public Builder retryPolicy(String retryPolicy) {
            return retryPolicy(Output.of(retryPolicy));
        }

        /**
         * @param serviceAccountEmail The email of the service account for this function.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountEmail(@Nullable Output<String> serviceAccountEmail) {
            $.serviceAccountEmail = serviceAccountEmail;
            return this;
        }

        /**
         * @param serviceAccountEmail The email of the service account for this function.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountEmail(String serviceAccountEmail) {
            return serviceAccountEmail(Output.of(serviceAccountEmail));
        }

        /**
         * @param trigger -
         * The resource name of the Eventarc trigger.
         * 
         * @return builder
         * 
         */
        public Builder trigger(@Nullable Output<String> trigger) {
            $.trigger = trigger;
            return this;
        }

        /**
         * @param trigger -
         * The resource name of the Eventarc trigger.
         * 
         * @return builder
         * 
         */
        public Builder trigger(String trigger) {
            return trigger(Output.of(trigger));
        }

        /**
         * @param triggerRegion The region that the trigger will be in. The trigger will only receive
         * events originating in this region. It can be the same
         * region as the function, a different region or multi-region, or the global
         * region. If not provided, defaults to the same region as the function.
         * 
         * @return builder
         * 
         */
        public Builder triggerRegion(@Nullable Output<String> triggerRegion) {
            $.triggerRegion = triggerRegion;
            return this;
        }

        /**
         * @param triggerRegion The region that the trigger will be in. The trigger will only receive
         * events originating in this region. It can be the same
         * region as the function, a different region or multi-region, or the global
         * region. If not provided, defaults to the same region as the function.
         * 
         * @return builder
         * 
         */
        public Builder triggerRegion(String triggerRegion) {
            return triggerRegion(Output.of(triggerRegion));
        }

        public FunctionEventTriggerArgs build() {
            return $;
        }
    }

}
