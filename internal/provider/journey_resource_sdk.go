// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"epilot-journey/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *JourneyResourceModel) ToSDKType() *shared.JourneyCreationRequest {
	brandID := r.BrandID.ValueString()
	createdBy := new(string)
	if !r.CreatedBy.IsUnknown() && !r.CreatedBy.IsNull() {
		*createdBy = r.CreatedBy.ValueString()
	} else {
		createdBy = nil
	}
	var design *shared.JourneyCreationRequestDesign
	if r.Design != nil {
		logoURL := new(string)
		if !r.Design.LogoURL.IsUnknown() && !r.Design.LogoURL.IsNull() {
			*logoURL = r.Design.LogoURL.ValueString()
		} else {
			logoURL = nil
		}
		theme := make(map[string]interface{})
		for themeKey, themeValue := range r.Design.Theme {
			var themeInst interface{}
			_ = json.Unmarshal([]byte(themeValue.ValueString()), &themeInst)
			theme[themeKey] = themeInst
		}
		design = &shared.JourneyCreationRequestDesign{
			LogoURL: logoURL,
			Theme:   theme,
		}
	}
	journeyID := new(string)
	if !r.JourneyID.IsUnknown() && !r.JourneyID.IsNull() {
		*journeyID = r.JourneyID.ValueString()
	} else {
		journeyID = nil
	}
	logics := make([]shared.JourneyCreationRequestLogics, 0)
	for _, logicsItem := range r.Logics {
		actions := make([]string, 0)
		for _, actionsItem := range logicsItem.Actions {
			actions = append(actions, actionsItem.ValueString())
		}
		autoGeneratedID := new(string)
		if !logicsItem.AutoGeneratedID.IsUnknown() && !logicsItem.AutoGeneratedID.IsNull() {
			*autoGeneratedID = logicsItem.AutoGeneratedID.ValueString()
		} else {
			autoGeneratedID = nil
		}
		conditions := make([]string, 0)
		for _, conditionsItem := range logicsItem.Conditions {
			conditions = append(conditions, conditionsItem.ValueString())
		}
		logics = append(logics, shared.JourneyCreationRequestLogics{
			Actions:         actions,
			AutoGeneratedID: autoGeneratedID,
			Conditions:      conditions,
		})
	}
	name := r.Name.ValueString()
	organizationID := r.OrganizationID.ValueString()
	rules := make([]shared.JourneyCreationRequestRules, 0)
	for _, rulesItem := range r.Rules {
		source := rulesItem.Source.ValueString()
		sourceType := shared.JourneyCreationRequestRulesSourceTypeEnum(rulesItem.SourceType.ValueString())
		target := rulesItem.Target.ValueString()
		type1 := shared.JourneyCreationRequestRulesTypeEnum(rulesItem.Type.ValueString())
		rules = append(rules, shared.JourneyCreationRequestRules{
			Source:     source,
			SourceType: sourceType,
			Target:     target,
			Type:       type1,
		})
	}
	var settings *shared.JourneyCreationRequestSettings
	if r.Settings != nil {
		description := new(string)
		if !r.Settings.Description.IsUnknown() && !r.Settings.Description.IsNull() {
			*description = r.Settings.Description.ValueString()
		} else {
			description = nil
		}
		designID := r.Settings.DesignID.ValueString()
		var embedOptions *shared.JourneyCreationRequestSettingsEmbedOptions
		if r.Settings.EmbedOptions != nil {
			var button *shared.JourneyCreationRequestSettingsEmbedOptionsButton
			if r.Settings.EmbedOptions.Button != nil {
				align := new(shared.JourneyCreationRequestSettingsEmbedOptionsButtonAlignEnum)
				if !r.Settings.EmbedOptions.Button.Align.IsUnknown() && !r.Settings.EmbedOptions.Button.Align.IsNull() {
					*align = shared.JourneyCreationRequestSettingsEmbedOptionsButtonAlignEnum(r.Settings.EmbedOptions.Button.Align.ValueString())
				} else {
					align = nil
				}
				text := new(string)
				if !r.Settings.EmbedOptions.Button.Text.IsUnknown() && !r.Settings.EmbedOptions.Button.Text.IsNull() {
					*text = r.Settings.EmbedOptions.Button.Text.ValueString()
				} else {
					text = nil
				}
				button = &shared.JourneyCreationRequestSettingsEmbedOptionsButton{
					Align: align,
					Text:  text,
				}
			}
			lang := new(shared.JourneyCreationRequestSettingsEmbedOptionsLangEnum)
			if !r.Settings.EmbedOptions.Lang.IsUnknown() && !r.Settings.EmbedOptions.Lang.IsNull() {
				*lang = shared.JourneyCreationRequestSettingsEmbedOptionsLangEnum(r.Settings.EmbedOptions.Lang.ValueString())
			} else {
				lang = nil
			}
			mode := new(shared.JourneyCreationRequestSettingsEmbedOptionsModeEnum)
			if !r.Settings.EmbedOptions.Mode.IsUnknown() && !r.Settings.EmbedOptions.Mode.IsNull() {
				*mode = shared.JourneyCreationRequestSettingsEmbedOptionsModeEnum(r.Settings.EmbedOptions.Mode.ValueString())
			} else {
				mode = nil
			}
			scrollToTop := new(bool)
			if !r.Settings.EmbedOptions.ScrollToTop.IsUnknown() && !r.Settings.EmbedOptions.ScrollToTop.IsNull() {
				*scrollToTop = r.Settings.EmbedOptions.ScrollToTop.ValueBool()
			} else {
				scrollToTop = nil
			}
			topBar := new(bool)
			if !r.Settings.EmbedOptions.TopBar.IsUnknown() && !r.Settings.EmbedOptions.TopBar.IsNull() {
				*topBar = r.Settings.EmbedOptions.TopBar.ValueBool()
			} else {
				topBar = nil
			}
			width := new(string)
			if !r.Settings.EmbedOptions.Width.IsUnknown() && !r.Settings.EmbedOptions.Width.IsNull() {
				*width = r.Settings.EmbedOptions.Width.ValueString()
			} else {
				width = nil
			}
			embedOptions = &shared.JourneyCreationRequestSettingsEmbedOptions{
				Button:      button,
				Lang:        lang,
				Mode:        mode,
				ScrollToTop: scrollToTop,
				TopBar:      topBar,
				Width:       width,
			}
		}
		entityID := new(string)
		if !r.Settings.EntityID.IsUnknown() && !r.Settings.EntityID.IsNull() {
			*entityID = r.Settings.EntityID.ValueString()
		} else {
			entityID = nil
		}
		entityTags := make([]string, 0)
		for _, entityTagsItem := range r.Settings.EntityTags {
			entityTags = append(entityTags, entityTagsItem.ValueString())
		}
		mappingsAutomationID := new(string)
		if !r.Settings.MappingsAutomationID.IsUnknown() && !r.Settings.MappingsAutomationID.IsNull() {
			*mappingsAutomationID = r.Settings.MappingsAutomationID.ValueString()
		} else {
			mappingsAutomationID = nil
		}
		organizationSettings := make(map[string]bool)
		for organizationSettingsKey, organizationSettingsValue := range r.Settings.OrganizationSettings {
			organizationSettingsInst := organizationSettingsValue.ValueBool()
			organizationSettings[organizationSettingsKey] = organizationSettingsInst
		}
		publicToken := new(string)
		if !r.Settings.PublicToken.IsUnknown() && !r.Settings.PublicToken.IsNull() {
			*publicToken = r.Settings.PublicToken.ValueString()
		} else {
			publicToken = nil
		}
		runtimeEntities := make([]shared.JourneyCreationRequestSettingsRuntimeEntitiesEnum, 0)
		for _, runtimeEntitiesItem := range r.Settings.RuntimeEntities {
			runtimeEntities = append(runtimeEntities, shared.JourneyCreationRequestSettingsRuntimeEntitiesEnum(runtimeEntitiesItem.ValueString()))
		}
		safeModeAutomation := new(bool)
		if !r.Settings.SafeModeAutomation.IsUnknown() && !r.Settings.SafeModeAutomation.IsNull() {
			*safeModeAutomation = r.Settings.SafeModeAutomation.ValueBool()
		} else {
			safeModeAutomation = nil
		}
		targetedCustomer := new(string)
		if !r.Settings.TargetedCustomer.IsUnknown() && !r.Settings.TargetedCustomer.IsNull() {
			*targetedCustomer = r.Settings.TargetedCustomer.ValueString()
		} else {
			targetedCustomer = nil
		}
		templateID := new(string)
		if !r.Settings.TemplateID.IsUnknown() && !r.Settings.TemplateID.IsNull() {
			*templateID = r.Settings.TemplateID.ValueString()
		} else {
			templateID = nil
		}
		settings = &shared.JourneyCreationRequestSettings{
			Description:          description,
			DesignID:             designID,
			EmbedOptions:         embedOptions,
			EntityID:             entityID,
			EntityTags:           entityTags,
			MappingsAutomationID: mappingsAutomationID,
			OrganizationSettings: organizationSettings,
			PublicToken:          publicToken,
			RuntimeEntities:      runtimeEntities,
			SafeModeAutomation:   safeModeAutomation,
			TargetedCustomer:     targetedCustomer,
			TemplateID:           templateID,
		}
	}
	steps := make([]shared.JourneyCreationRequestSteps, 0)
	for _, stepsItem := range r.Steps {
		hideNextButton := new(bool)
		if !stepsItem.HideNextButton.IsUnknown() && !stepsItem.HideNextButton.IsNull() {
			*hideNextButton = stepsItem.HideNextButton.ValueBool()
		} else {
			hideNextButton = nil
		}
		name1 := stepsItem.Name.ValueString()
		var schema interface{}
		_ = json.Unmarshal([]byte(stepsItem.Schema.ValueString()), &schema)
		showStepName := new(bool)
		if !stepsItem.ShowStepName.IsUnknown() && !stepsItem.ShowStepName.IsNull() {
			*showStepName = stepsItem.ShowStepName.ValueBool()
		} else {
			showStepName = nil
		}
		showStepSubtitle := new(bool)
		if !stepsItem.ShowStepSubtitle.IsUnknown() && !stepsItem.ShowStepSubtitle.IsNull() {
			*showStepSubtitle = stepsItem.ShowStepSubtitle.ValueBool()
		} else {
			showStepSubtitle = nil
		}
		showStepper := new(bool)
		if !stepsItem.ShowStepper.IsUnknown() && !stepsItem.ShowStepper.IsNull() {
			*showStepper = stepsItem.ShowStepper.ValueBool()
		} else {
			showStepper = nil
		}
		showStepperLabels := new(bool)
		if !stepsItem.ShowStepperLabels.IsUnknown() && !stepsItem.ShowStepperLabels.IsNull() {
			*showStepperLabels = stepsItem.ShowStepperLabels.ValueBool()
		} else {
			showStepperLabels = nil
		}
		stepID := new(string)
		if !stepsItem.StepID.IsUnknown() && !stepsItem.StepID.IsNull() {
			*stepID = stepsItem.StepID.ValueString()
		} else {
			stepID = nil
		}
		subTitle := new(string)
		if !stepsItem.SubTitle.IsUnknown() && !stepsItem.SubTitle.IsNull() {
			*subTitle = stepsItem.SubTitle.ValueString()
		} else {
			subTitle = nil
		}
		title := new(string)
		if !stepsItem.Title.IsUnknown() && !stepsItem.Title.IsNull() {
			*title = stepsItem.Title.ValueString()
		} else {
			title = nil
		}
		var uischema interface{}
		_ = json.Unmarshal([]byte(stepsItem.Uischema.ValueString()), &uischema)
		steps = append(steps, shared.JourneyCreationRequestSteps{
			HideNextButton:    hideNextButton,
			Name:              name1,
			Schema:            schema,
			ShowStepName:      showStepName,
			ShowStepSubtitle:  showStepSubtitle,
			ShowStepper:       showStepper,
			ShowStepperLabels: showStepperLabels,
			StepID:            stepID,
			SubTitle:          subTitle,
			Title:             title,
			Uischema:          uischema,
		})
	}
	out := shared.JourneyCreationRequest{
		BrandID:        brandID,
		CreatedBy:      createdBy,
		Design:         design,
		JourneyID:      journeyID,
		Logics:         logics,
		Name:           name,
		OrganizationID: organizationID,
		Rules:          rules,
		Settings:       settings,
		Steps:          steps,
	}
	return &out

}

func (r *JourneyResourceModel) RefreshFromSDKType(resp *shared.JourneyResponse) {
	r.BrandID = types.StringValue(resp.BrandID)
	if resp.CreatedBy != nil {
		r.CreatedBy = types.StringValue(*resp.CreatedBy)
	} else {
		r.CreatedBy = types.StringNull()
	}
	if resp.Design == nil {
		r.Design = nil
	} else {
		r.Design = &JourneyCreationRequestDesign{}
		if resp.Design.LogoURL != nil {
			r.Design.LogoURL = types.StringValue(*resp.Design.LogoURL)
		} else {
			r.Design.LogoURL = types.StringNull()
		}
		r.Design.Theme = make(map[string]types.String)
		for key, value := range resp.Design.Theme {
			result, _ := json.Marshal(value)
			r.Design.Theme[key] = types.StringValue(string(result))
		}
	}
	if resp.JourneyID != nil {
		r.JourneyID = types.StringValue(*resp.JourneyID)
	} else {
		r.JourneyID = types.StringNull()
	}
	r.Logics = nil
	for _, logicsItem := range resp.Logics {
		var logics1 JourneyCreationRequestLogics
		logics1.Actions = nil
		for _, v := range logicsItem.Actions {
			logics1.Actions = append(logics1.Actions, types.StringValue(v))
		}
		if logicsItem.AutoGeneratedID != nil {
			logics1.AutoGeneratedID = types.StringValue(*logicsItem.AutoGeneratedID)
		} else {
			logics1.AutoGeneratedID = types.StringNull()
		}
		logics1.Conditions = nil
		for _, v := range logicsItem.Conditions {
			logics1.Conditions = append(logics1.Conditions, types.StringValue(v))
		}
		r.Logics = append(r.Logics, logics1)
	}
	r.Name = types.StringValue(resp.Name)
	r.OrganizationID = types.StringValue(resp.OrganizationID)
	r.Rules = nil
	for _, rulesItem := range resp.Rules {
		var rules1 JourneyCreationRequestRules
		rules1.Source = types.StringValue(rulesItem.Source)
		rules1.SourceType = types.StringValue(string(rulesItem.SourceType))
		rules1.Target = types.StringValue(rulesItem.Target)
		rules1.Type = types.StringValue(string(rulesItem.Type))
		r.Rules = append(r.Rules, rules1)
	}
	if resp.Settings == nil {
		r.Settings = nil
	} else {
		r.Settings = &JourneyCreationRequestSettings{}
		if resp.Settings.Description != nil {
			r.Settings.Description = types.StringValue(*resp.Settings.Description)
		} else {
			r.Settings.Description = types.StringNull()
		}
		r.Settings.DesignID = types.StringValue(resp.Settings.DesignID)
		if resp.Settings.EmbedOptions == nil {
			r.Settings.EmbedOptions = nil
		} else {
			r.Settings.EmbedOptions = &JourneyCreationRequestSettingsEmbedOptions{}
			if resp.Settings.EmbedOptions.Button == nil {
				r.Settings.EmbedOptions.Button = nil
			} else {
				r.Settings.EmbedOptions.Button = &JourneyCreationRequestSettingsEmbedOptionsButton{}
				if resp.Settings.EmbedOptions.Button.Align != nil {
					r.Settings.EmbedOptions.Button.Align = types.StringValue(string(*resp.Settings.EmbedOptions.Button.Align))
				} else {
					r.Settings.EmbedOptions.Button.Align = types.StringNull()
				}
				if resp.Settings.EmbedOptions.Button.Text != nil {
					r.Settings.EmbedOptions.Button.Text = types.StringValue(*resp.Settings.EmbedOptions.Button.Text)
				} else {
					r.Settings.EmbedOptions.Button.Text = types.StringNull()
				}
			}
			if resp.Settings.EmbedOptions.Lang != nil {
				r.Settings.EmbedOptions.Lang = types.StringValue(string(*resp.Settings.EmbedOptions.Lang))
			} else {
				r.Settings.EmbedOptions.Lang = types.StringNull()
			}
			if resp.Settings.EmbedOptions.Mode != nil {
				r.Settings.EmbedOptions.Mode = types.StringValue(string(*resp.Settings.EmbedOptions.Mode))
			} else {
				r.Settings.EmbedOptions.Mode = types.StringNull()
			}
			if resp.Settings.EmbedOptions.ScrollToTop != nil {
				r.Settings.EmbedOptions.ScrollToTop = types.BoolValue(*resp.Settings.EmbedOptions.ScrollToTop)
			} else {
				r.Settings.EmbedOptions.ScrollToTop = types.BoolNull()
			}
			if resp.Settings.EmbedOptions.TopBar != nil {
				r.Settings.EmbedOptions.TopBar = types.BoolValue(*resp.Settings.EmbedOptions.TopBar)
			} else {
				r.Settings.EmbedOptions.TopBar = types.BoolNull()
			}
			if resp.Settings.EmbedOptions.Width != nil {
				r.Settings.EmbedOptions.Width = types.StringValue(*resp.Settings.EmbedOptions.Width)
			} else {
				r.Settings.EmbedOptions.Width = types.StringNull()
			}
		}
		if resp.Settings.EntityID != nil {
			r.Settings.EntityID = types.StringValue(*resp.Settings.EntityID)
		} else {
			r.Settings.EntityID = types.StringNull()
		}
		r.Settings.EntityTags = nil
		for _, v := range resp.Settings.EntityTags {
			r.Settings.EntityTags = append(r.Settings.EntityTags, types.StringValue(v))
		}
		if resp.Settings.MappingsAutomationID != nil {
			r.Settings.MappingsAutomationID = types.StringValue(*resp.Settings.MappingsAutomationID)
		} else {
			r.Settings.MappingsAutomationID = types.StringNull()
		}
		r.Settings.OrganizationSettings = make(map[string]types.Bool)
		for key1, value1 := range resp.Settings.OrganizationSettings {
			r.Settings.OrganizationSettings[key1] = types.BoolValue(value1)
		}
		if resp.Settings.PublicToken != nil {
			r.Settings.PublicToken = types.StringValue(*resp.Settings.PublicToken)
		} else {
			r.Settings.PublicToken = types.StringNull()
		}
		r.Settings.RuntimeEntities = nil
		for _, v := range resp.Settings.RuntimeEntities {
			r.Settings.RuntimeEntities = append(r.Settings.RuntimeEntities, types.StringValue(string(v)))
		}
		if resp.Settings.SafeModeAutomation != nil {
			r.Settings.SafeModeAutomation = types.BoolValue(*resp.Settings.SafeModeAutomation)
		} else {
			r.Settings.SafeModeAutomation = types.BoolNull()
		}
		if resp.Settings.TargetedCustomer != nil {
			r.Settings.TargetedCustomer = types.StringValue(*resp.Settings.TargetedCustomer)
		} else {
			r.Settings.TargetedCustomer = types.StringNull()
		}
		if resp.Settings.TemplateID != nil {
			r.Settings.TemplateID = types.StringValue(*resp.Settings.TemplateID)
		} else {
			r.Settings.TemplateID = types.StringNull()
		}
	}
	r.Steps = nil
	for _, stepsItem := range resp.Steps {
		var steps1 JourneyCreationRequestSteps
		if stepsItem.HideNextButton != nil {
			steps1.HideNextButton = types.BoolValue(*stepsItem.HideNextButton)
		} else {
			steps1.HideNextButton = types.BoolNull()
		}
		steps1.Name = types.StringValue(stepsItem.Name)
		schemaResult, _ := json.Marshal(stepsItem.Schema)
		steps1.Schema = types.StringValue(string(schemaResult))
		if stepsItem.ShowStepName != nil {
			steps1.ShowStepName = types.BoolValue(*stepsItem.ShowStepName)
		} else {
			steps1.ShowStepName = types.BoolNull()
		}
		if stepsItem.ShowStepSubtitle != nil {
			steps1.ShowStepSubtitle = types.BoolValue(*stepsItem.ShowStepSubtitle)
		} else {
			steps1.ShowStepSubtitle = types.BoolNull()
		}
		if stepsItem.ShowStepper != nil {
			steps1.ShowStepper = types.BoolValue(*stepsItem.ShowStepper)
		} else {
			steps1.ShowStepper = types.BoolNull()
		}
		if stepsItem.ShowStepperLabels != nil {
			steps1.ShowStepperLabels = types.BoolValue(*stepsItem.ShowStepperLabels)
		} else {
			steps1.ShowStepperLabels = types.BoolNull()
		}
		if stepsItem.StepID != nil {
			steps1.StepID = types.StringValue(*stepsItem.StepID)
		} else {
			steps1.StepID = types.StringNull()
		}
		if stepsItem.SubTitle != nil {
			steps1.SubTitle = types.StringValue(*stepsItem.SubTitle)
		} else {
			steps1.SubTitle = types.StringNull()
		}
		if stepsItem.Title != nil {
			steps1.Title = types.StringValue(*stepsItem.Title)
		} else {
			steps1.Title = types.StringNull()
		}
		uischemaResult, _ := json.Marshal(stepsItem.Uischema)
		steps1.Uischema = types.StringValue(string(uischemaResult))
		r.Steps = append(r.Steps, steps1)
	}

}
