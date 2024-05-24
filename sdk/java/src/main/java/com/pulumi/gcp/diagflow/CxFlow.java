// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.diagflow.CxFlowArgs;
import com.pulumi.gcp.diagflow.inputs.CxFlowState;
import com.pulumi.gcp.diagflow.outputs.CxFlowAdvancedSettings;
import com.pulumi.gcp.diagflow.outputs.CxFlowEventHandler;
import com.pulumi.gcp.diagflow.outputs.CxFlowNluSettings;
import com.pulumi.gcp.diagflow.outputs.CxFlowTransitionRoute;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Flows represents the conversation flows when you build your chatbot agent.
 * 
 * To get more information about Flow, see:
 * 
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 * 
 * ## Example Usage
 * 
 * ### Dialogflowcx Flow Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.diagflow.CxAgent;
 * import com.pulumi.gcp.diagflow.CxAgentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxAgentSpeechToTextSettingsArgs;
 * import com.pulumi.gcp.diagflow.CxFlow;
 * import com.pulumi.gcp.diagflow.CxFlowArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowNluSettingsArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowEventHandlerArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowEventHandlerTriggerFulfillmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var agent = new CxAgent("agent", CxAgentArgs.builder()
 *             .displayName("dialogflowcx-agent")
 *             .location("global")
 *             .defaultLanguageCode("en")
 *             .supportedLanguageCodes(            
 *                 "fr",
 *                 "de",
 *                 "es")
 *             .timeZone("America/New_York")
 *             .description("Example description.")
 *             .avatarUri("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png")
 *             .enableStackdriverLogging(true)
 *             .enableSpellCorrection(true)
 *             .speechToTextSettings(CxAgentSpeechToTextSettingsArgs.builder()
 *                 .enableSpeechAdaptation(true)
 *                 .build())
 *             .build());
 * 
 *         var basicFlow = new CxFlow("basicFlow", CxFlowArgs.builder()
 *             .parent(agent.id())
 *             .displayName("MyFlow")
 *             .description("Test Flow")
 *             .nluSettings(CxFlowNluSettingsArgs.builder()
 *                 .classificationThreshold(0.3)
 *                 .modelType("MODEL_TYPE_STANDARD")
 *                 .build())
 *             .eventHandlers(            
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("custom-event")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(false)
 *                         .messages(CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                             .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("I didn't get that. Can you say it again?")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("sys.no-match-default")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(false)
 *                         .messages(CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                             .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("Sorry, could you say that again?")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("sys.no-input-default")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(false)
 *                         .messages(CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                             .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("One more time?")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dialogflowcx Flow Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.diagflow.CxAgent;
 * import com.pulumi.gcp.diagflow.CxAgentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxAgentSpeechToTextSettingsArgs;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.diagflow.CxFlow;
 * import com.pulumi.gcp.diagflow.CxFlowArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowNluSettingsArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowEventHandlerArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowEventHandlerTriggerFulfillmentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowTransitionRouteArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowTransitionRouteTriggerFulfillmentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowAdvancedSettingsArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowAdvancedSettingsAudioExportGcsDestinationArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxFlowAdvancedSettingsDtmfSettingsArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var agent = new CxAgent("agent", CxAgentArgs.builder()
 *             .displayName("dialogflowcx-agent")
 *             .location("global")
 *             .defaultLanguageCode("en")
 *             .supportedLanguageCodes(            
 *                 "fr",
 *                 "de",
 *                 "es")
 *             .timeZone("America/New_York")
 *             .description("Example description.")
 *             .avatarUri("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png")
 *             .enableStackdriverLogging(true)
 *             .enableSpellCorrection(true)
 *             .speechToTextSettings(CxAgentSpeechToTextSettingsArgs.builder()
 *                 .enableSpeechAdaptation(true)
 *                 .build())
 *             .build());
 * 
 *         var bucket = new Bucket("bucket", BucketArgs.builder()
 *             .name("dialogflowcx-bucket")
 *             .location("US")
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         var basicFlow = new CxFlow("basicFlow", CxFlowArgs.builder()
 *             .parent(agent.id())
 *             .displayName("MyFlow")
 *             .description("Test Flow")
 *             .nluSettings(CxFlowNluSettingsArgs.builder()
 *                 .classificationThreshold(0.3)
 *                 .modelType("MODEL_TYPE_STANDARD")
 *                 .build())
 *             .eventHandlers(            
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("custom-event")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(false)
 *                         .messages(CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                             .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("I didn't get that. Can you say it again?")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("sys.no-match-default")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(false)
 *                         .messages(CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                             .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("Sorry, could you say that again?")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("sys.no-input-default")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(false)
 *                         .messages(CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                             .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("One more time?")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxFlowEventHandlerArgs.builder()
 *                     .event("another-event")
 *                     .triggerFulfillment(CxFlowEventHandlerTriggerFulfillmentArgs.builder()
 *                         .returnPartialResponses(true)
 *                         .messages(                        
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .channel("some-channel")
 *                                 .text(CxFlowEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                                     .texts("Some text")
 *                                     .build())
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .payload("""
 *           {"some-key": "some-value", "other-key": ["other-value"]}
 *                                 """)
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .conversationSuccess(CxFlowEventHandlerTriggerFulfillmentMessageConversationSuccessArgs.builder()
 *                                     .metadata("""
 *             {"some-metadata-key": "some-value", "other-metadata-key": 1234}
 *                                     """)
 *                                     .build())
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .outputAudioText(CxFlowEventHandlerTriggerFulfillmentMessageOutputAudioTextArgs.builder()
 *                                     .text("some output text")
 *                                     .build())
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .outputAudioText(CxFlowEventHandlerTriggerFulfillmentMessageOutputAudioTextArgs.builder()
 *                                     .ssml("""
 *             <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
 *                                     """)
 *                                     .build())
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .liveAgentHandoff(CxFlowEventHandlerTriggerFulfillmentMessageLiveAgentHandoffArgs.builder()
 *                                     .metadata("""
 *             {"some-metadata-key": "some-value", "other-metadata-key": 1234}
 *                                     """)
 *                                     .build())
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .playAudio(CxFlowEventHandlerTriggerFulfillmentMessagePlayAudioArgs.builder()
 *                                     .audioUri("http://example.com/some-audio-file.mp3")
 *                                     .build())
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                                 .telephonyTransferCall(CxFlowEventHandlerTriggerFulfillmentMessageTelephonyTransferCallArgs.builder()
 *                                     .phoneNumber("1-234-567-8901")
 *                                     .build())
 *                                 .build())
 *                         .setParameterActions(                        
 *                             CxFlowEventHandlerTriggerFulfillmentSetParameterActionArgs.builder()
 *                                 .parameter("some-param")
 *                                 .value("123.45")
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentSetParameterActionArgs.builder()
 *                                 .parameter("another-param")
 *                                 .value(serializeJson(
 *                                     "abc"))
 *                                 .build(),
 *                             CxFlowEventHandlerTriggerFulfillmentSetParameterActionArgs.builder()
 *                                 .parameter("other-param")
 *                                 .value(serializeJson(
 *                                     jsonArray("foo")))
 *                                 .build())
 *                         .conditionalCases(CxFlowEventHandlerTriggerFulfillmentConditionalCaseArgs.builder()
 *                             .cases(serializeJson(
 *                                 jsonArray(
 *                                     jsonObject(
 *                                         jsonProperty("condition", "$sys.func.RAND() < 0.5"),
 *                                         jsonProperty("caseContent", jsonArray(
 *                                             jsonObject(
 *                                                 jsonProperty("message", jsonObject(
 *                                                     jsonProperty("text", jsonObject(
 *                                                         jsonProperty("text", jsonArray("First case"))
 *                                                     ))
 *                                                 ))
 *                                             ), 
 *                                             jsonObject(
 *                                                 jsonProperty("additionalCases", jsonObject(
 *                                                     jsonProperty("cases", jsonArray(jsonObject(
 *                                                         jsonProperty("condition", "$sys.func.RAND() < 0.2"),
 *                                                         jsonProperty("caseContent", jsonArray(jsonObject(
 *                                                             jsonProperty("message", jsonObject(
 *                                                                 jsonProperty("text", jsonObject(
 *                                                                     jsonProperty("text", jsonArray("Nested case"))
 *                                                                 ))
 *                                                             ))
 *                                                         )))
 *                                                     )))
 *                                                 ))
 *                                             )
 *                                         ))
 *                                     ), 
 *                                     jsonObject(
 *                                         jsonProperty("caseContent", jsonArray(jsonObject(
 *                                             jsonProperty("message", jsonObject(
 *                                                 jsonProperty("text", jsonObject(
 *                                                     jsonProperty("text", jsonArray("Final case"))
 *                                                 ))
 *                                             ))
 *                                         )))
 *                                     )
 *                                 )))
 *                             .build())
 *                         .build())
 *                     .build())
 *             .transitionRoutes(CxFlowTransitionRouteArgs.builder()
 *                 .condition("true")
 *                 .triggerFulfillment(CxFlowTransitionRouteTriggerFulfillmentArgs.builder()
 *                     .returnPartialResponses(true)
 *                     .messages(                    
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .channel("some-channel")
 *                             .text(CxFlowTransitionRouteTriggerFulfillmentMessageTextArgs.builder()
 *                                 .texts("Some text")
 *                                 .build())
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .payload("""
 *           {"some-key": "some-value", "other-key": ["other-value"]}
 *                             """)
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .conversationSuccess(CxFlowTransitionRouteTriggerFulfillmentMessageConversationSuccessArgs.builder()
 *                                 .metadata("""
 *             {"some-metadata-key": "some-value", "other-metadata-key": 1234}
 *                                 """)
 *                                 .build())
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .outputAudioText(CxFlowTransitionRouteTriggerFulfillmentMessageOutputAudioTextArgs.builder()
 *                                 .text("some output text")
 *                                 .build())
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .outputAudioText(CxFlowTransitionRouteTriggerFulfillmentMessageOutputAudioTextArgs.builder()
 *                                 .ssml("""
 *             <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
 *                                 """)
 *                                 .build())
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .liveAgentHandoff(CxFlowTransitionRouteTriggerFulfillmentMessageLiveAgentHandoffArgs.builder()
 *                                 .metadata("""
 *             {"some-metadata-key": "some-value", "other-metadata-key": 1234}
 *                                 """)
 *                                 .build())
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .playAudio(CxFlowTransitionRouteTriggerFulfillmentMessagePlayAudioArgs.builder()
 *                                 .audioUri("http://example.com/some-audio-file.mp3")
 *                                 .build())
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                             .telephonyTransferCall(CxFlowTransitionRouteTriggerFulfillmentMessageTelephonyTransferCallArgs.builder()
 *                                 .phoneNumber("1-234-567-8901")
 *                                 .build())
 *                             .build())
 *                     .setParameterActions(                    
 *                         CxFlowTransitionRouteTriggerFulfillmentSetParameterActionArgs.builder()
 *                             .parameter("some-param")
 *                             .value("123.45")
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentSetParameterActionArgs.builder()
 *                             .parameter("another-param")
 *                             .value(serializeJson(
 *                                 "abc"))
 *                             .build(),
 *                         CxFlowTransitionRouteTriggerFulfillmentSetParameterActionArgs.builder()
 *                             .parameter("other-param")
 *                             .value(serializeJson(
 *                                 jsonArray("foo")))
 *                             .build())
 *                     .conditionalCases(CxFlowTransitionRouteTriggerFulfillmentConditionalCaseArgs.builder()
 *                         .cases(serializeJson(
 *                             jsonArray(
 *                                 jsonObject(
 *                                     jsonProperty("condition", "$sys.func.RAND() < 0.5"),
 *                                     jsonProperty("caseContent", jsonArray(
 *                                         jsonObject(
 *                                             jsonProperty("message", jsonObject(
 *                                                 jsonProperty("text", jsonObject(
 *                                                     jsonProperty("text", jsonArray("First case"))
 *                                                 ))
 *                                             ))
 *                                         ), 
 *                                         jsonObject(
 *                                             jsonProperty("additionalCases", jsonObject(
 *                                                 jsonProperty("cases", jsonArray(jsonObject(
 *                                                     jsonProperty("condition", "$sys.func.RAND() < 0.2"),
 *                                                     jsonProperty("caseContent", jsonArray(jsonObject(
 *                                                         jsonProperty("message", jsonObject(
 *                                                             jsonProperty("text", jsonObject(
 *                                                                 jsonProperty("text", jsonArray("Nested case"))
 *                                                             ))
 *                                                         ))
 *                                                     )))
 *                                                 )))
 *                                             ))
 *                                         )
 *                                     ))
 *                                 ), 
 *                                 jsonObject(
 *                                     jsonProperty("caseContent", jsonArray(jsonObject(
 *                                         jsonProperty("message", jsonObject(
 *                                             jsonProperty("text", jsonObject(
 *                                                 jsonProperty("text", jsonArray("Final case"))
 *                                             ))
 *                                         ))
 *                                     )))
 *                                 )
 *                             )))
 *                         .build())
 *                     .build())
 *                 .targetFlow(agent.startFlow())
 *                 .build())
 *             .advancedSettings(CxFlowAdvancedSettingsArgs.builder()
 *                 .audioExportGcsDestination(CxFlowAdvancedSettingsAudioExportGcsDestinationArgs.builder()
 *                     .uri(bucket.url().applyValue(url -> String.format("%s/prefix-", url)))
 *                     .build())
 *                 .dtmfSettings(CxFlowAdvancedSettingsDtmfSettingsArgs.builder()
 *                     .enabled(true)
 *                     .maxDigits(1)
 *                     .finishDigit("#")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Flow can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/flows/{{name}}`
 * 
 * * `{{parent}}/{{name}}`
 * 
 * When using the `pulumi import` command, Flow can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:diagflow/cxFlow:CxFlow default {{parent}}/flows/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:diagflow/cxFlow:CxFlow default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:diagflow/cxFlow:CxFlow")
public class CxFlow extends com.pulumi.resources.CustomResource {
    /**
     * Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at the higher level.
     * Hierarchy: Agent-&gt;Flow-&gt;Page-&gt;Fulfillment/Parameter.
     * Structure is documented below.
     * 
     */
    @Export(name="advancedSettings", refs={CxFlowAdvancedSettings.class}, tree="[0]")
    private Output</* @Nullable */ CxFlowAdvancedSettings> advancedSettings;

    /**
     * @return Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at the higher level.
     * Hierarchy: Agent-&gt;Flow-&gt;Page-&gt;Fulfillment/Parameter.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CxFlowAdvancedSettings>> advancedSettings() {
        return Codegen.optional(this.advancedSettings);
    }
    /**
     * The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The human-readable name of the flow.
     * 
     * ***
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The human-readable name of the flow.
     * 
     * ***
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * A flow&#39;s event handlers serve two purposes:
     * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
     * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
     * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     * Structure is documented below.
     * 
     */
    @Export(name="eventHandlers", refs={List.class,CxFlowEventHandler.class}, tree="[0,1]")
    private Output<List<CxFlowEventHandler>> eventHandlers;

    /**
     * @return A flow&#39;s event handlers serve two purposes:
     * They are responsible for handling events (e.g. no match, webhook errors) in the flow.
     * They are inherited by every page&#39;s [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
     * Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     * Structure is documented below.
     * 
     */
    public Output<List<CxFlowEventHandler>> eventHandlers() {
        return this.eventHandlers;
    }
    /**
     * Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent. When you create an agent, the Default Start Flow is created automatically.
     * The Default Start Flow cannot be deleted; deleting the `gcp.diagflow.CxFlow` resource does nothing to the underlying GCP resources.
     * 
     * &gt; Avoid having multiple `gcp.diagflow.CxFlow` resources linked to the same agent with `is_default_start_flow = true` because they will compete to control a single Default Start Flow resource in GCP.
     * 
     */
    @Export(name="isDefaultStartFlow", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isDefaultStartFlow;

    /**
     * @return Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent. When you create an agent, the Default Start Flow is created automatically.
     * The Default Start Flow cannot be deleted; deleting the `gcp.diagflow.CxFlow` resource does nothing to the underlying GCP resources.
     * 
     * &gt; Avoid having multiple `gcp.diagflow.CxFlow` resources linked to the same agent with `is_default_start_flow = true` because they will compete to control a single Default Start Flow resource in GCP.
     * 
     */
    public Output<Optional<Boolean>> isDefaultStartFlow() {
        return Codegen.optional(this.isDefaultStartFlow);
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
    @Export(name="languageCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> languageCode;

    /**
     * @return The language of the following fields in flow:
     * Flow.event_handlers.trigger_fulfillment.messages
     * Flow.event_handlers.trigger_fulfillment.conditional_cases
     * Flow.transition_routes.trigger_fulfillment.messages
     * Flow.transition_routes.trigger_fulfillment.conditional_cases
     * If not specified, the agent&#39;s default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
     * 
     */
    public Output<Optional<String>> languageCode() {
        return Codegen.optional(this.languageCode);
    }
    /**
     * The unique identifier of the flow.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique identifier of the flow.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * NLU related settings of the flow.
     * Structure is documented below.
     * 
     */
    @Export(name="nluSettings", refs={CxFlowNluSettings.class}, tree="[0]")
    private Output</* @Nullable */ CxFlowNluSettings> nluSettings;

    /**
     * @return NLU related settings of the flow.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CxFlowNluSettings>> nluSettings() {
        return Codegen.optional(this.nluSettings);
    }
    /**
     * The agent to create a flow for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parent;

    /**
     * @return The agent to create a flow for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    public Output<Optional<String>> parent() {
        return Codegen.optional(this.parent);
    }
    /**
     * A flow&#39;s transition route group serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
     * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
     * 
     */
    @Export(name="transitionRouteGroups", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> transitionRouteGroups;

    /**
     * @return A flow&#39;s transition route group serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
     * Format:projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/transitionRouteGroups/&lt;TransitionRouteGroup ID&gt;.
     * 
     */
    public Output<Optional<List<String>>> transitionRouteGroups() {
        return Codegen.optional(this.transitionRouteGroups);
    }
    /**
     * A flow&#39;s transition routes serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
     * TransitionRoutes are evalauted in the following order:
     * TransitionRoutes with intent specified.
     * TransitionRoutes with only condition specified.
     * TransitionRoutes with intent specified are inherited by pages in the flow.
     * Structure is documented below.
     * 
     */
    @Export(name="transitionRoutes", refs={List.class,CxFlowTransitionRoute.class}, tree="[0,1]")
    private Output</* @Nullable */ List<CxFlowTransitionRoute>> transitionRoutes;

    /**
     * @return A flow&#39;s transition routes serve two purposes:
     * They are responsible for matching the user&#39;s first utterances in the flow.
     * They are inherited by every page&#39;s [transition routes][Page.transition_routes] and can support use cases such as the user saying &#34;help&#34; or &#34;can I talk to a human?&#34;, which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
     * TransitionRoutes are evalauted in the following order:
     * TransitionRoutes with intent specified.
     * TransitionRoutes with only condition specified.
     * TransitionRoutes with intent specified are inherited by pages in the flow.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<CxFlowTransitionRoute>>> transitionRoutes() {
        return Codegen.optional(this.transitionRoutes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CxFlow(String name) {
        this(name, CxFlowArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CxFlow(String name, CxFlowArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CxFlow(String name, CxFlowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/cxFlow:CxFlow", name, args == null ? CxFlowArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CxFlow(String name, Output<String> id, @Nullable CxFlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/cxFlow:CxFlow", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CxFlow get(String name, Output<String> id, @Nullable CxFlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CxFlow(name, id, state, options);
    }
}
