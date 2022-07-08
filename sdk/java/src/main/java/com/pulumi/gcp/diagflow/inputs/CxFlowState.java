// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.diagflow.inputs.CxFlowEventHandlerArgs;
import com.pulumi.gcp.diagflow.inputs.CxFlowNluSettingsArgs;
import com.pulumi.gcp.diagflow.inputs.CxFlowTransitionRouteArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CxFlowState extends com.pulumi.resources.ResourceArgs {

    public static final CxFlowState Empty = new CxFlowState();

    /**
     * The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The human-readable name of the flow.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The human-readable name of the flow.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * A flow&#39;s event handlers serve two purposes:
     * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
     * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
     * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     * Structure is documented below.
     * 
     */
    @Import(name="eventHandlers")
    private @Nullable Output<List<CxFlowEventHandlerArgs>> eventHandlers;

    /**
     * @return A flow&#39;s event handlers serve two purposes:
     * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
     * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
     * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<CxFlowEventHandlerArgs>>> eventHandlers() {
        return Optional.ofNullable(this.eventHandlers);
    }

    /**
     * The language of the following fields in flow:
     * Flow.event_handlers.trigger_fulfillment.messages
     * Flow.event_handlers.trigger_fulfillment.conditional_cases
     * Flow.transition_routes.trigger_fulfillment.messages
     * Flow.transition_routes.trigger_fulfillment.conditional_cases
     * If not specified, the agent&#39;s default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     * 
     */
    @Import(name="languageCode")
    private @Nullable Output<String> languageCode;

    /**
     * @return The language of the following fields in flow:
     * Flow.event_handlers.trigger_fulfillment.messages
     * Flow.event_handlers.trigger_fulfillment.conditional_cases
     * Flow.transition_routes.trigger_fulfillment.messages
     * Flow.transition_routes.trigger_fulfillment.conditional_cases
     * If not specified, the agent&#39;s default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     * 
     */
    public Optional<Output<String>> languageCode() {
        return Optional.ofNullable(this.languageCode);
    }

    /**
     * - 
     * The unique identifier of this event handler.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return -
     * The unique identifier of this event handler.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * NLU related settings of the flow.
     * Structure is documented below.
     * 
     */
    @Import(name="nluSettings")
    private @Nullable Output<CxFlowNluSettingsArgs> nluSettings;

    /**
     * @return NLU related settings of the flow.
     * Structure is documented below.
     * 
     */
    public Optional<Output<CxFlowNluSettingsArgs>> nluSettings() {
        return Optional.ofNullable(this.nluSettings);
    }

    /**
     * The agent to create a flow for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    @Import(name="parent")
    private @Nullable Output<String> parent;

    /**
     * @return The agent to create a flow for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    public Optional<Output<String>> parent() {
        return Optional.ofNullable(this.parent);
    }

    /**
     * A flow&#39;s transition route group serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
     * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
     * 
     */
    @Import(name="transitionRouteGroups")
    private @Nullable Output<List<String>> transitionRouteGroups;

    /**
     * @return A flow&#39;s transition route group serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
     * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
     * 
     */
    public Optional<Output<List<String>>> transitionRouteGroups() {
        return Optional.ofNullable(this.transitionRouteGroups);
    }

    /**
     * A flow&#39;s transition routes serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
     * 
     */
    @Import(name="transitionRoutes")
    private @Nullable Output<List<CxFlowTransitionRouteArgs>> transitionRoutes;

    /**
     * @return A flow&#39;s transition routes serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
     * 
     */
    public Optional<Output<List<CxFlowTransitionRouteArgs>>> transitionRoutes() {
        return Optional.ofNullable(this.transitionRoutes);
    }

    private CxFlowState() {}

    private CxFlowState(CxFlowState $) {
        this.description = $.description;
        this.displayName = $.displayName;
        this.eventHandlers = $.eventHandlers;
        this.languageCode = $.languageCode;
        this.name = $.name;
        this.nluSettings = $.nluSettings;
        this.parent = $.parent;
        this.transitionRouteGroups = $.transitionRouteGroups;
        this.transitionRoutes = $.transitionRoutes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CxFlowState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CxFlowState $;

        public Builder() {
            $ = new CxFlowState();
        }

        public Builder(CxFlowState defaults) {
            $ = new CxFlowState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName The human-readable name of the flow.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The human-readable name of the flow.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param eventHandlers A flow&#39;s event handlers serve two purposes:
         * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
         * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
         * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventHandlers(@Nullable Output<List<CxFlowEventHandlerArgs>> eventHandlers) {
            $.eventHandlers = eventHandlers;
            return this;
        }

        /**
         * @param eventHandlers A flow&#39;s event handlers serve two purposes:
         * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
         * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
         * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventHandlers(List<CxFlowEventHandlerArgs> eventHandlers) {
            return eventHandlers(Output.of(eventHandlers));
        }

        /**
         * @param eventHandlers A flow&#39;s event handlers serve two purposes:
         * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
         * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
         * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventHandlers(CxFlowEventHandlerArgs... eventHandlers) {
            return eventHandlers(List.of(eventHandlers));
        }

        /**
         * @param languageCode The language of the following fields in flow:
         * Flow.event_handlers.trigger_fulfillment.messages
         * Flow.event_handlers.trigger_fulfillment.conditional_cases
         * Flow.transition_routes.trigger_fulfillment.messages
         * Flow.transition_routes.trigger_fulfillment.conditional_cases
         * If not specified, the agent&#39;s default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(@Nullable Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode The language of the following fields in flow:
         * Flow.event_handlers.trigger_fulfillment.messages
         * Flow.event_handlers.trigger_fulfillment.conditional_cases
         * Flow.transition_routes.trigger_fulfillment.messages
         * Flow.transition_routes.trigger_fulfillment.conditional_cases
         * If not specified, the agent&#39;s default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        /**
         * @param name -
         * The unique identifier of this event handler.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name -
         * The unique identifier of this event handler.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nluSettings NLU related settings of the flow.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder nluSettings(@Nullable Output<CxFlowNluSettingsArgs> nluSettings) {
            $.nluSettings = nluSettings;
            return this;
        }

        /**
         * @param nluSettings NLU related settings of the flow.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder nluSettings(CxFlowNluSettingsArgs nluSettings) {
            return nluSettings(Output.of(nluSettings));
        }

        /**
         * @param parent The agent to create a flow for.
         * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
         * 
         * @return builder
         * 
         */
        public Builder parent(@Nullable Output<String> parent) {
            $.parent = parent;
            return this;
        }

        /**
         * @param parent The agent to create a flow for.
         * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
         * 
         * @return builder
         * 
         */
        public Builder parent(String parent) {
            return parent(Output.of(parent));
        }

        /**
         * @param transitionRouteGroups A flow&#39;s transition route group serve two purposes:
         * They are responsible for matching the user&#39;s first utterances in the flow.
         * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
         * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
         * 
         * @return builder
         * 
         */
        public Builder transitionRouteGroups(@Nullable Output<List<String>> transitionRouteGroups) {
            $.transitionRouteGroups = transitionRouteGroups;
            return this;
        }

        /**
         * @param transitionRouteGroups A flow&#39;s transition route group serve two purposes:
         * They are responsible for matching the user&#39;s first utterances in the flow.
         * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
         * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
         * 
         * @return builder
         * 
         */
        public Builder transitionRouteGroups(List<String> transitionRouteGroups) {
            return transitionRouteGroups(Output.of(transitionRouteGroups));
        }

        /**
         * @param transitionRouteGroups A flow&#39;s transition route group serve two purposes:
         * They are responsible for matching the user&#39;s first utterances in the flow.
         * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
         * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
         * 
         * @return builder
         * 
         */
        public Builder transitionRouteGroups(String... transitionRouteGroups) {
            return transitionRouteGroups(List.of(transitionRouteGroups));
        }

        /**
         * @param transitionRoutes A flow&#39;s transition routes serve two purposes:
         * They are responsible for matching the user&#39;s first utterances in the flow.
         * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
         * 
         * @return builder
         * 
         */
        public Builder transitionRoutes(@Nullable Output<List<CxFlowTransitionRouteArgs>> transitionRoutes) {
            $.transitionRoutes = transitionRoutes;
            return this;
        }

        /**
         * @param transitionRoutes A flow&#39;s transition routes serve two purposes:
         * They are responsible for matching the user&#39;s first utterances in the flow.
         * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
         * 
         * @return builder
         * 
         */
        public Builder transitionRoutes(List<CxFlowTransitionRouteArgs> transitionRoutes) {
            return transitionRoutes(Output.of(transitionRoutes));
        }

        /**
         * @param transitionRoutes A flow&#39;s transition routes serve two purposes:
         * They are responsible for matching the user&#39;s first utterances in the flow.
         * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
         * 
         * @return builder
         * 
         */
        public Builder transitionRoutes(CxFlowTransitionRouteArgs... transitionRoutes) {
            return transitionRoutes(List.of(transitionRoutes));
        }

        public CxFlowState build() {
            return $;
        }
    }

}
