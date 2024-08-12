// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.diagflow.CxTestCaseArgs;
import com.pulumi.gcp.diagflow.inputs.CxTestCaseState;
import com.pulumi.gcp.diagflow.outputs.CxTestCaseLastTestResult;
import com.pulumi.gcp.diagflow.outputs.CxTestCaseTestCaseConversationTurn;
import com.pulumi.gcp.diagflow.outputs.CxTestCaseTestConfig;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * You can use the built-in test feature to uncover bugs and prevent regressions. A test execution verifies that agent responses have not changed for end-user inputs defined in the test case.
 * 
 * To get more information about TestCase, see:
 * 
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.testCases)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 * 
 * ## Example Usage
 * 
 * ### Dialogflowcx Test Case Full
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
 * import com.pulumi.gcp.diagflow.CxIntent;
 * import com.pulumi.gcp.diagflow.CxIntentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxIntentTrainingPhraseArgs;
 * import com.pulumi.gcp.diagflow.CxPage;
 * import com.pulumi.gcp.diagflow.CxPageArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxPageTransitionRouteArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxPageTransitionRouteTriggerFulfillmentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxPageEventHandlerArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxPageEventHandlerTriggerFulfillmentArgs;
 * import com.pulumi.gcp.diagflow.CxTestCase;
 * import com.pulumi.gcp.diagflow.CxTestCaseArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestConfigArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnUserInputArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnUserInputInputArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnUserInputInputTextArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnVirtualAgentOutputArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnVirtualAgentOutputTriggeredIntentArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnVirtualAgentOutputCurrentPageArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnUserInputInputEventArgs;
 * import com.pulumi.gcp.diagflow.inputs.CxTestCaseTestCaseConversationTurnUserInputInputDtmfArgs;
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
 *             .avatarUri("https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo.png")
 *             .enableStackdriverLogging(true)
 *             .enableSpellCorrection(true)
 *             .speechToTextSettings(CxAgentSpeechToTextSettingsArgs.builder()
 *                 .enableSpeechAdaptation(true)
 *                 .build())
 *             .build());
 * 
 *         var intent = new CxIntent("intent", CxIntentArgs.builder()
 *             .parent(agent.id())
 *             .displayName("MyIntent")
 *             .priority(1)
 *             .trainingPhrases(CxIntentTrainingPhraseArgs.builder()
 *                 .parts(CxIntentTrainingPhrasePartArgs.builder()
 *                     .text("training phrase")
 *                     .build())
 *                 .repeatCount(1)
 *                 .build())
 *             .build());
 * 
 *         var page = new CxPage("page", CxPageArgs.builder()
 *             .parent(agent.startFlow())
 *             .displayName("MyPage")
 *             .transitionRoutes(CxPageTransitionRouteArgs.builder()
 *                 .intent(intent.id())
 *                 .triggerFulfillment(CxPageTransitionRouteTriggerFulfillmentArgs.builder()
 *                     .messages(CxPageTransitionRouteTriggerFulfillmentMessageArgs.builder()
 *                         .text(CxPageTransitionRouteTriggerFulfillmentMessageTextArgs.builder()
 *                             .texts("Training phrase response")
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .eventHandlers(CxPageEventHandlerArgs.builder()
 *                 .event("some-event")
 *                 .triggerFulfillment(CxPageEventHandlerTriggerFulfillmentArgs.builder()
 *                     .messages(CxPageEventHandlerTriggerFulfillmentMessageArgs.builder()
 *                         .text(CxPageEventHandlerTriggerFulfillmentMessageTextArgs.builder()
 *                             .texts("Handling some event")
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var basicTestCase = new CxTestCase("basicTestCase", CxTestCaseArgs.builder()
 *             .parent(agent.id())
 *             .displayName("MyTestCase")
 *             .tags("#tag1")
 *             .notes("demonstrates a simple training phrase response")
 *             .testConfig(CxTestCaseTestConfigArgs.builder()
 *                 .trackingParameters("some_param")
 *                 .page(page.id())
 *                 .build())
 *             .testCaseConversationTurns(            
 *                 CxTestCaseTestCaseConversationTurnArgs.builder()
 *                     .userInput(CxTestCaseTestCaseConversationTurnUserInputArgs.builder()
 *                         .input(CxTestCaseTestCaseConversationTurnUserInputInputArgs.builder()
 *                             .languageCode("en")
 *                             .text(CxTestCaseTestCaseConversationTurnUserInputInputTextArgs.builder()
 *                                 .text("training phrase")
 *                                 .build())
 *                             .build())
 *                         .injectedParameters(serializeJson(
 *                             jsonObject(
 *                                 jsonProperty("some_param", "1")
 *                             )))
 *                         .isWebhookEnabled(true)
 *                         .enableSentimentAnalysis(true)
 *                         .build())
 *                     .virtualAgentOutput(CxTestCaseTestCaseConversationTurnVirtualAgentOutputArgs.builder()
 *                         .sessionParameters(serializeJson(
 *                             jsonObject(
 *                                 jsonProperty("some_param", "1")
 *                             )))
 *                         .triggeredIntent(CxTestCaseTestCaseConversationTurnVirtualAgentOutputTriggeredIntentArgs.builder()
 *                             .name(intent.id())
 *                             .build())
 *                         .currentPage(CxTestCaseTestCaseConversationTurnVirtualAgentOutputCurrentPageArgs.builder()
 *                             .name(page.id())
 *                             .build())
 *                         .textResponses(CxTestCaseTestCaseConversationTurnVirtualAgentOutputTextResponseArgs.builder()
 *                             .texts("Training phrase response")
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxTestCaseTestCaseConversationTurnArgs.builder()
 *                     .userInput(CxTestCaseTestCaseConversationTurnUserInputArgs.builder()
 *                         .input(CxTestCaseTestCaseConversationTurnUserInputInputArgs.builder()
 *                             .event(CxTestCaseTestCaseConversationTurnUserInputInputEventArgs.builder()
 *                                 .event("some-event")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .virtualAgentOutput(CxTestCaseTestCaseConversationTurnVirtualAgentOutputArgs.builder()
 *                         .currentPage(CxTestCaseTestCaseConversationTurnVirtualAgentOutputCurrentPageArgs.builder()
 *                             .name(page.id())
 *                             .build())
 *                         .textResponses(CxTestCaseTestCaseConversationTurnVirtualAgentOutputTextResponseArgs.builder()
 *                             .texts("Handling some event")
 *                             .build())
 *                         .build())
 *                     .build(),
 *                 CxTestCaseTestCaseConversationTurnArgs.builder()
 *                     .userInput(CxTestCaseTestCaseConversationTurnUserInputArgs.builder()
 *                         .input(CxTestCaseTestCaseConversationTurnUserInputInputArgs.builder()
 *                             .dtmf(CxTestCaseTestCaseConversationTurnUserInputInputDtmfArgs.builder()
 *                                 .digits("12")
 *                                 .finishDigit("3")
 *                                 .build())
 *                             .build())
 *                         .build())
 *                     .virtualAgentOutput(CxTestCaseTestCaseConversationTurnVirtualAgentOutputArgs.builder()
 *                         .textResponses(CxTestCaseTestCaseConversationTurnVirtualAgentOutputTextResponseArgs.builder()
 *                             .texts("I didn't get that. Can you say it again?")
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
 * 
 * ## Import
 * 
 * TestCase can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/testCases/{{name}}`
 * 
 * When using the `pulumi import` command, TestCase can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:diagflow/cxTestCase:CxTestCase default {{parent}}/testCases/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:diagflow/cxTestCase:CxTestCase")
public class CxTestCase extends com.pulumi.resources.CustomResource {
    /**
     * When the test was created. A timestamp in RFC3339 text format.
     * 
     */
    @Export(name="creationTime", refs={String.class}, tree="[0]")
    private Output<String> creationTime;

    /**
     * @return When the test was created. A timestamp in RFC3339 text format.
     * 
     */
    public Output<String> creationTime() {
        return this.creationTime;
    }
    /**
     * The human-readable name of the test case, unique within the agent. Limit of 200 characters.
     * 
     * ***
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The human-readable name of the test case, unique within the agent. Limit of 200 characters.
     * 
     * ***
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The latest test result.
     * Structure is documented below.
     * 
     */
    @Export(name="lastTestResults", refs={List.class,CxTestCaseLastTestResult.class}, tree="[0,1]")
    private Output<List<CxTestCaseLastTestResult>> lastTestResults;

    /**
     * @return The latest test result.
     * Structure is documented below.
     * 
     */
    public Output<List<CxTestCaseLastTestResult>> lastTestResults() {
        return this.lastTestResults;
    }
    /**
     * The unique identifier of the page.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/pages/&lt;Page ID&gt;.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique identifier of the page.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/flows/&lt;Flow ID&gt;/pages/&lt;Page ID&gt;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Additional freeform notes about the test case. Limit of 400 characters.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notes;

    /**
     * @return Additional freeform notes about the test case. Limit of 400 characters.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * The agent to create the test case for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parent;

    /**
     * @return The agent to create the test case for.
     * Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
     * 
     */
    public Output<Optional<String>> parent() {
        return Codegen.optional(this.parent);
    }
    /**
     * Tags are short descriptions that users may apply to test cases for organizational and filtering purposes.
     * Each tag should start with &#34;#&#34; and has a limit of 30 characters
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags are short descriptions that users may apply to test cases for organizational and filtering purposes.
     * Each tag should start with &#34;#&#34; and has a limit of 30 characters
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The conversation turns uttered when the test case was created, in chronological order. These include the canonical set of agent utterances that should occur when the agent is working properly.
     * Structure is documented below.
     * 
     */
    @Export(name="testCaseConversationTurns", refs={List.class,CxTestCaseTestCaseConversationTurn.class}, tree="[0,1]")
    private Output</* @Nullable */ List<CxTestCaseTestCaseConversationTurn>> testCaseConversationTurns;

    /**
     * @return The conversation turns uttered when the test case was created, in chronological order. These include the canonical set of agent utterances that should occur when the agent is working properly.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<CxTestCaseTestCaseConversationTurn>>> testCaseConversationTurns() {
        return Codegen.optional(this.testCaseConversationTurns);
    }
    /**
     * Config for the test case.
     * Structure is documented below.
     * 
     */
    @Export(name="testConfig", refs={CxTestCaseTestConfig.class}, tree="[0]")
    private Output</* @Nullable */ CxTestCaseTestConfig> testConfig;

    /**
     * @return Config for the test case.
     * Structure is documented below.
     * 
     */
    public Output<Optional<CxTestCaseTestConfig>> testConfig() {
        return Codegen.optional(this.testConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CxTestCase(java.lang.String name) {
        this(name, CxTestCaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CxTestCase(java.lang.String name, CxTestCaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CxTestCase(java.lang.String name, CxTestCaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/cxTestCase:CxTestCase", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CxTestCase(java.lang.String name, Output<java.lang.String> id, @Nullable CxTestCaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/cxTestCase:CxTestCase", name, state, makeResourceOptions(options, id), false);
    }

    private static CxTestCaseArgs makeArgs(CxTestCaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CxTestCaseArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static CxTestCase get(java.lang.String name, Output<java.lang.String> id, @Nullable CxTestCaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CxTestCase(name, id, state, options);
    }
}
