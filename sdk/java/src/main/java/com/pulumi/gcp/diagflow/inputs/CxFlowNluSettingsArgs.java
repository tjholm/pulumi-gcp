// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CxFlowNluSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final CxFlowNluSettingsArgs Empty = new CxFlowNluSettingsArgs();

    /**
     * To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
     * If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
     * 
     */
    @Import(name="classificationThreshold")
    private @Nullable Output<Double> classificationThreshold;

    /**
     * @return To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
     * If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
     * 
     */
    public Optional<Output<Double>> classificationThreshold() {
        return Optional.ofNullable(this.classificationThreshold);
    }

    /**
     * Indicates NLU model training mode.
     * * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
     * * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train.
     *   Possible values are `MODEL_TRAINING_MODE_AUTOMATIC` and `MODEL_TRAINING_MODE_MANUAL`.
     * 
     */
    @Import(name="modelTrainingMode")
    private @Nullable Output<String> modelTrainingMode;

    /**
     * @return Indicates NLU model training mode.
     * * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
     * * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train.
     *   Possible values are `MODEL_TRAINING_MODE_AUTOMATIC` and `MODEL_TRAINING_MODE_MANUAL`.
     * 
     */
    public Optional<Output<String>> modelTrainingMode() {
        return Optional.ofNullable(this.modelTrainingMode);
    }

    /**
     * Indicates the type of NLU model.
     * * MODEL_TYPE_STANDARD: Use standard NLU model.
     * * MODEL_TYPE_ADVANCED: Use advanced NLU model.
     *   Possible values are `MODEL_TYPE_STANDARD` and `MODEL_TYPE_ADVANCED`.
     * 
     */
    @Import(name="modelType")
    private @Nullable Output<String> modelType;

    /**
     * @return Indicates the type of NLU model.
     * * MODEL_TYPE_STANDARD: Use standard NLU model.
     * * MODEL_TYPE_ADVANCED: Use advanced NLU model.
     *   Possible values are `MODEL_TYPE_STANDARD` and `MODEL_TYPE_ADVANCED`.
     * 
     */
    public Optional<Output<String>> modelType() {
        return Optional.ofNullable(this.modelType);
    }

    private CxFlowNluSettingsArgs() {}

    private CxFlowNluSettingsArgs(CxFlowNluSettingsArgs $) {
        this.classificationThreshold = $.classificationThreshold;
        this.modelTrainingMode = $.modelTrainingMode;
        this.modelType = $.modelType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CxFlowNluSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CxFlowNluSettingsArgs $;

        public Builder() {
            $ = new CxFlowNluSettingsArgs();
        }

        public Builder(CxFlowNluSettingsArgs defaults) {
            $ = new CxFlowNluSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param classificationThreshold To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
         * If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
         * 
         * @return builder
         * 
         */
        public Builder classificationThreshold(@Nullable Output<Double> classificationThreshold) {
            $.classificationThreshold = classificationThreshold;
            return this;
        }

        /**
         * @param classificationThreshold To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
         * If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
         * 
         * @return builder
         * 
         */
        public Builder classificationThreshold(Double classificationThreshold) {
            return classificationThreshold(Output.of(classificationThreshold));
        }

        /**
         * @param modelTrainingMode Indicates NLU model training mode.
         * * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
         * * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train.
         *   Possible values are `MODEL_TRAINING_MODE_AUTOMATIC` and `MODEL_TRAINING_MODE_MANUAL`.
         * 
         * @return builder
         * 
         */
        public Builder modelTrainingMode(@Nullable Output<String> modelTrainingMode) {
            $.modelTrainingMode = modelTrainingMode;
            return this;
        }

        /**
         * @param modelTrainingMode Indicates NLU model training mode.
         * * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
         * * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train.
         *   Possible values are `MODEL_TRAINING_MODE_AUTOMATIC` and `MODEL_TRAINING_MODE_MANUAL`.
         * 
         * @return builder
         * 
         */
        public Builder modelTrainingMode(String modelTrainingMode) {
            return modelTrainingMode(Output.of(modelTrainingMode));
        }

        /**
         * @param modelType Indicates the type of NLU model.
         * * MODEL_TYPE_STANDARD: Use standard NLU model.
         * * MODEL_TYPE_ADVANCED: Use advanced NLU model.
         *   Possible values are `MODEL_TYPE_STANDARD` and `MODEL_TYPE_ADVANCED`.
         * 
         * @return builder
         * 
         */
        public Builder modelType(@Nullable Output<String> modelType) {
            $.modelType = modelType;
            return this;
        }

        /**
         * @param modelType Indicates the type of NLU model.
         * * MODEL_TYPE_STANDARD: Use standard NLU model.
         * * MODEL_TYPE_ADVANCED: Use advanced NLU model.
         *   Possible values are `MODEL_TYPE_STANDARD` and `MODEL_TYPE_ADVANCED`.
         * 
         * @return builder
         * 
         */
        public Builder modelType(String modelType) {
            return modelType(Output.of(modelType));
        }

        public CxFlowNluSettingsArgs build() {
            return $;
        }
    }

}
