// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.diagflow.outputs.CxPageTransitionRouteTriggerFulfillment;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CxPageTransitionRoute {
    /**
     * @return The condition to evaluate against form parameters or session parameters.
     * At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
     * 
     */
    private final @Nullable String condition;
    /**
     * @return The unique identifier of an Intent.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/intents/&lt;Intent ID&gt;. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
     * 
     */
    private final @Nullable String intent;
    /**
     * @return -
     * The unique identifier of this event handler.
     * 
     */
    private final @Nullable String name;
    /**
     * @return The target flow to transition to.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;.
     * 
     */
    private final @Nullable String targetFlow;
    /**
     * @return The target page to transition to.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/pages/&lt;Page ID&gt;.
     * 
     */
    private final @Nullable String targetPage;
    /**
     * @return The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
     * Structure is documented below.
     * 
     */
    private final @Nullable CxPageTransitionRouteTriggerFulfillment triggerFulfillment;

    @CustomType.Constructor
    private CxPageTransitionRoute(
        @CustomType.Parameter("condition") @Nullable String condition,
        @CustomType.Parameter("intent") @Nullable String intent,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("targetFlow") @Nullable String targetFlow,
        @CustomType.Parameter("targetPage") @Nullable String targetPage,
        @CustomType.Parameter("triggerFulfillment") @Nullable CxPageTransitionRouteTriggerFulfillment triggerFulfillment) {
        this.condition = condition;
        this.intent = intent;
        this.name = name;
        this.targetFlow = targetFlow;
        this.targetPage = targetPage;
        this.triggerFulfillment = triggerFulfillment;
    }

    /**
     * @return The condition to evaluate against form parameters or session parameters.
     * At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
     * 
     */
    public Optional<String> condition() {
        return Optional.ofNullable(this.condition);
    }
    /**
     * @return The unique identifier of an Intent.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/intents/&lt;Intent ID&gt;. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
     * 
     */
    public Optional<String> intent() {
        return Optional.ofNullable(this.intent);
    }
    /**
     * @return -
     * The unique identifier of this event handler.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The target flow to transition to.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;.
     * 
     */
    public Optional<String> targetFlow() {
        return Optional.ofNullable(this.targetFlow);
    }
    /**
     * @return The target page to transition to.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/pages/&lt;Page ID&gt;.
     * 
     */
    public Optional<String> targetPage() {
        return Optional.ofNullable(this.targetPage);
    }
    /**
     * @return The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
     * Structure is documented below.
     * 
     */
    public Optional<CxPageTransitionRouteTriggerFulfillment> triggerFulfillment() {
        return Optional.ofNullable(this.triggerFulfillment);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CxPageTransitionRoute defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String condition;
        private @Nullable String intent;
        private @Nullable String name;
        private @Nullable String targetFlow;
        private @Nullable String targetPage;
        private @Nullable CxPageTransitionRouteTriggerFulfillment triggerFulfillment;

        public Builder() {
    	      // Empty
        }

        public Builder(CxPageTransitionRoute defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.condition = defaults.condition;
    	      this.intent = defaults.intent;
    	      this.name = defaults.name;
    	      this.targetFlow = defaults.targetFlow;
    	      this.targetPage = defaults.targetPage;
    	      this.triggerFulfillment = defaults.triggerFulfillment;
        }

        public Builder condition(@Nullable String condition) {
            this.condition = condition;
            return this;
        }
        public Builder intent(@Nullable String intent) {
            this.intent = intent;
            return this;
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder targetFlow(@Nullable String targetFlow) {
            this.targetFlow = targetFlow;
            return this;
        }
        public Builder targetPage(@Nullable String targetPage) {
            this.targetPage = targetPage;
            return this;
        }
        public Builder triggerFulfillment(@Nullable CxPageTransitionRouteTriggerFulfillment triggerFulfillment) {
            this.triggerFulfillment = triggerFulfillment;
            return this;
        }        public CxPageTransitionRoute build() {
            return new CxPageTransitionRoute(condition, intent, name, targetFlow, targetPage, triggerFulfillment);
        }
    }
}
